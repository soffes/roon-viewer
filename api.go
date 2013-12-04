package main

import (
	"fmt"
	"net/http"
	"errors"
	"encoding/json"
)

func Get(blogSubdomain string, postSlug string) (*Post, error) {
	url := fmt.Sprintf("https://roon.io/api/v1/blogs/%s/posts/%s", blogSubdomain, postSlug)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	r := new(Post)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
