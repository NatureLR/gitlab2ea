package migrate

import (
	"fmt"

	sdkgitea "code.gitea.io/sdk/gitea"
	"github.com/naturelr/gitlab2ea/pkg/gitea"
	"github.com/naturelr/gitlab2ea/pkg/gitlab"
	"k8s.io/klog/v2"
)

type Migrage struct {
	Gitlab *gitlab.Gitlab
	Gitea  *gitea.Gitea
}

func New(gitlab *gitlab.Gitlab, gitea *gitea.Gitea) *Migrage {
	return &Migrage{
		Gitlab: gitlab,
		Gitea:  gitea,
	}
}

type GroupRepo map[string]map[string]struct{}

func (g GroupRepo) Add(group, repo string) {
	if g[group] == nil {
		g[group] = make(map[string]struct{})
	}
	r := g[group]
	if repo != "" {
		r[repo] = struct{}{}
	}
}

func (g GroupRepo) orgExist(group string) bool {
	_, ok := g[group]
	return ok
}

func (g GroupRepo) RepoExist(group, repo string) bool {
	repos, ok := g[group]
	if !ok {
		return false
	}
	if _, ok := repos[repo]; !ok {
		return false
	}
	return true
}

func (m *Migrage) InitGroupRepo() (GroupRepo, error) {
	groupRepo := make(GroupRepo)

	orgs, err := m.Gitea.ListMyOrgsAll()
	if err != nil {
		return nil, err
	}

	for _, org := range orgs {
		repos, err := m.Gitea.ListOrgReposAll(org.UserName)
		if err != nil {
			return nil, err
		}
		for _, repo := range repos {
			groupRepo.Add(org.UserName, repo.Name)
		}
	}

	users, err := m.Gitea.AdminListUsersAll()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		repos, err := m.Gitea.ListUserReposAll(user.UserName)
		if err != nil {
			return nil, err
		}
		for _, repo := range repos {
			groupRepo.Add(user.UserName, repo.Name)
		}
	}

	return groupRepo, nil
}

func (m *Migrage) Do() error {
	groupRepo, err := m.InitGroupRepo()
	if err != nil {
		return err
	}

	user, _, err := m.Gitea.Client.GetMyUserInfo()
	if err != nil {
		return err
	}

	projects, err := m.Gitlab.ListProject()
	if err != nil {
		fmt.Println(err)
		return err
	}

	var migrateSuccess, migrateAlreadyExist int

	for _, p := range projects {
		if p.Namespace.Name == "Administrator" {
			p.Namespace.Name = user.UserName
		}

		if !groupRepo.orgExist(p.Namespace.Name) {
			klog.Warningf("org %s not exits create", p.Namespace.Name)
			_, _, err := m.Gitea.Client.CreateOrg(sdkgitea.CreateOrgOption{
				Name:                      p.Namespace.Name,
				FullName:                  p.Namespace.Name,
				Visibility:                sdkgitea.VisibleTypeLimited,
				RepoAdminChangeTeamAccess: true,
			})
			if err != nil {
				return err
			}
			groupRepo.Add(p.Namespace.Name, "")
		}

		if groupRepo.RepoExist(p.Namespace.Name, p.Name) {
			klog.Infof("already exists %s/%s", p.Namespace.Name, p.Name)
			migrateAlreadyExist++
			continue
		}

		opt := sdkgitea.MigrateRepoOption{
			Service:      sdkgitea.GitServiceGitlab,
			RepoName:     p.Name,
			RepoOwner:    p.Namespace.Name,
			Description:  p.Description,
			CloneAddr:    p.HTTPURLToRepo,
			AuthToken:    m.Gitlab.Token,
			Wiki:         true,
			Milestones:   true,
			Labels:       true,
			Issues:       true,
			PullRequests: true,
			Releases:     true,
		}

		_, _, err := m.Gitea.Client.MigrateRepo(opt)
		if err != nil {
			return fmt.Errorf("migrate %s/%s err %v", p.Namespace.Name, p.Name, err)
		}
		migrateSuccess++
		klog.Infof("[%d/%d] migrate %s/%s success skip %d ", migrateSuccess, len(projects), p.Namespace.Name, p.Name, migrateAlreadyExist)
	}

	klog.Infof("migrate success [%d/%d] skip %d ", migrateSuccess, len(projects), migrateAlreadyExist)
	return nil
}
