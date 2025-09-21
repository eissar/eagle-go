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

// not implemented (this just gets data.tags, data.recent, data.groups
func TagAll(baseUrl string) {
	// ep := endpoints.TagAll
	// uri := baseUrl + ep.Path
	//
	// // no params
	//
	// err := Request(ep.Method, uri, nil, nil, &resp)
}

func TagList(baseUrl string) ([]*TagItem, error) {
	ep := endpoints.TagList
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

// id        : MFUAWYUQRHY2G
// name      : test
// tags      : {}
// editable  : False
// $$hashKey : object:32256

func TagGroups(baseUrl string) ([]*TagsGroup, error) {
	ep := endpoints.TagGroups
	uri := baseUrl + ep.Path

	// no param
	var resp struct {
		EagleResponse
		Data []*TagsGroup `json:"data"`
	}

	err := Request(ep.Method, uri, nil, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("groups: err=%s", err)
	}
	if resp.Status != "success" {
		return nil, fmt.Errorf("groups: err=%s", err)
	}

	return resp.Data, nil
}

func TagListRecent(baseUrl string) ([]*TagItem, error) {
	ep := endpoints.TagListRecent
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
