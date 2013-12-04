package main

import "fmt"

type Author struct {
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Bio        string `json:"bio_html"`
	Website    string
	Twitter    string
	AvatarURL  string
}

func (a Author) Name() string {
	return fmt.Sprintf("%s %s", a.GivenName, a.FamilyName)
}
