package gitea

import (
	"strconv"

	"code.gitea.io/sdk/gitea"
)

type Gitea struct {
	Addr   string
	Token  string
	Client *gitea.Client
}

func New(addr, token string) (*Gitea, error) {
	c, err := gitea.NewClient(addr, gitea.SetToken(token))
	if err != nil {
		return nil, err
	}
	return &Gitea{
		Addr:   addr,
		Token:  token,
		Client: c,
	}, nil
}

func ListALL(listAPI func(listOpts gitea.ListOptions) (*gitea.Response, error), pageSize int) error {
	for page := 1; ; page++ {
		listOpt := gitea.ListOptions{
			Page:     page,
			PageSize: pageSize,
		}
		resp, err := listAPI(listOpt)
		if err != nil {
			return err
		}
		total, err := strconv.Atoi(resp.Header.Get("X-Total-Count"))
		if err != nil {
			return err
		}

		if pageSize*page >= total {
			break
		}
	}
	return nil
}

func (g *Gitea) ListMyOrgsAll() ([]*gitea.Organization, error) {
	var orgs []*gitea.Organization

	api := func(listOpts gitea.ListOptions) (*gitea.Response, error) {
		o, resp, err := g.Client.ListMyOrgs(gitea.ListOrgsOptions{ListOptions: listOpts})
		if err != nil {
			return nil, err
		}
		orgs = append(orgs, o...)
		return resp, nil
	}
	if err := ListALL(api, 20); err != nil {
		return nil, err
	}

	return orgs, nil
}

func (g *Gitea) ListOrgReposAll(org string) ([]*gitea.Repository, error) {
	var repos []*gitea.Repository

	api := func(listOpts gitea.ListOptions) (*gitea.Response, error) {
		repo, resp, err := g.Client.ListOrgRepos(org, gitea.ListOrgReposOptions{ListOptions: listOpts})
		if err != nil {
			return nil, err
		}
		repos = append(repos, repo...)
		return resp, nil

	}
	if err := ListALL(api, 20); err != nil {
		return nil, err
	}

	return repos, nil
}

func (g *Gitea) AdminListUsersAll() ([]*gitea.User, error) {
	var users []*gitea.User

	api := func(listOpts gitea.ListOptions) (*gitea.Response, error) {
		u, resp, err := g.Client.AdminListUsers(gitea.AdminListUsersOptions{ListOptions: listOpts})
		if err != nil {
			return nil, err
		}
		users = append(users, u...)
		return resp, nil
	}
	if err := ListALL(api, 20); err != nil {
		return nil, err
	}

	return users, nil
}

func (g *Gitea) ListUserReposAll(user string) ([]*gitea.Repository, error) {
	var repos []*gitea.Repository

	api := func(listOpts gitea.ListOptions) (*gitea.Response, error) {
		r, resp, err := g.Client.ListUserRepos(user, gitea.ListReposOptions{})
		if err != nil {
			return nil, err
		}
		repos = append(repos, r...)
		return resp, nil
	}
	if err := ListALL(api, 20); err != nil {
		return nil, err
	}

	return repos, nil
}
