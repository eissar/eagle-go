package eagle

import (
	"fmt"

	"github.com/eissar/eagle-go/endpoints"
)

type TagItem struct {
	Name       string   `json:"name"`
	ImageCount int      `json:"imageCount"`
	Groups     []string `json:"groups"` // seems to be an array but exclusive?
	Pinyin     string   `json:"pinyin"`
}

// not implemented
func TagAll(baseUrl string) {
	// ep := endpoints.TagAll
	// uri := baseUrl + ep.Path
	//
	// // no params
	//
	// err := Request(ep.Method, uri, nil, nil, &resp)
}

func TagList(baseUrl string) ([]*TagItem, error) {
	ep := endpoints.TagAll
	uri := baseUrl + ep.Path

	// no params
	var resp struct {
		EagleResponse
		Data []*TagItem `json:"data"`
	}

	err := Request(ep.Method, uri, nil, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("list: err=%w", err)
	}

	if resp.Status != "success" {
		return nil, fmt.Errorf("list: err=%w", ErrStatusErr)
	}

	return resp.Data, nil
}

// not implemented
func TagGroups(baseUrl string) {
}

// not implemented
func TagListRecent(baseUrl string) {
}
