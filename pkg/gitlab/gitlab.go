package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"k8s.io/klog/v2"
)

type Gitlab struct {
	Addr   string
	Token  string
	Client *http.Client
}

func New(addr, token string) *Gitlab {
	return &Gitlab{
		Addr:  addr,
		Token: token,
		Client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (g *Gitlab) Request(api string, page, prePage int) ([]byte, int, error) {
	url := fmt.Sprintf("%s/%s?per_page=%d&page=%d", g.Addr, api, prePage, page)
	klog.Info("request gitlab url:", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Add("PRIVATE-TOKEN", g.Token)

	resp, err := g.Client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	total, err := strconv.Atoi(resp.Header.Get("x-total"))
	if err != nil {
		return nil, 0, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return data, total, nil

}

func (g *Gitlab) ListProject() ([]Project, error) {
	var ret []Project
	prePage := 20
	for page := 1; ; page++ {
		data, total, err := g.Request("api/v4/projects", page, prePage)
		if err != nil {
			return nil, err
		}

		var objs []Project
		err = json.Unmarshal(data, &objs)
		if err != nil {
			return nil, err
		}

		ret = append(ret, objs...)

		fmt.Printf("total:%d page:%d prePage:%d\n", total, page, prePage)

		if prePage*page >= total {
			break
		}
	}

	return ret, nil
}
