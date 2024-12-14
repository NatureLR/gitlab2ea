package gitlab

import (
	"fmt"
	"testing"
)

var g = New("", "")

func TestGitlab(t *testing.T) {
	prj, err := g.ListProject()
	if err != nil {
		t.Log(err)
	}

	for _, p := range prj {
		fmt.Println(p.Namespace.Name, p.Name, p.Description, p.WebURL)
	}

	fmt.Println(len(prj))
}
