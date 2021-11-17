package model

import "strings"

func (i Invoice) Code() string {
	url := i.Links.Self.Href
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
