package utl

import (
	"net/url"
	"path"
	"strings"
)

type Page struct {
	Name    string
	Server  string
	Project string
}

func (p *Page) String() string {
	return path.Join(p.Server, p.Project, p.Name)
}

func ParsePagePath(pagePath, defaultProject, defaultServer string) *Page {
	names := strings.Split(pagePath, "/")
	pageName := names[len(names)-1]

	projectName := defaultProject
	serverName := defaultServer
	if len(names) > 1 {
		projectName = names[len(names)-2]
	}
	if len(names) > 2 {
		serverName = names[len(names)-3]
	}

	return &Page{
		Name:    pageName,
		Project: projectName,
		Server:  serverName,
	}
}

func GenerateBodyQuery(contents string) string {
	values := url.Values{}
	values.Add("body", contents)
	return values.Encode()
}
