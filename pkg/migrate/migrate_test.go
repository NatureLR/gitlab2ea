package migrate

import (
	"testing"

	"github.com/naturelr/gitlab2ea/pkg/gitea"
	"github.com/naturelr/gitlab2ea/pkg/gitlab"
)

func TestMigrate(t *testing.T) {
	gitlab := gitlab.New("", "")
	gea, err := gitea.New("", "")
	if err != nil {
		t.Log(err)
		return
	}

	m := New(gitlab, gea)

	if err := m.Do(); err != nil {
		t.Log(err)
	}
}
