package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"

	"golang.zabbix.com/sdk/plugin"
	"golang.zabbix.com/sdk/plugin/container"
)

type Plugin struct {
	plugin.Base
}

type dirDiscovery struct {
	Name string `json:"name"`
}

var impl Plugin

func (p *Plugin) Export(key string, params []string, ctx plugin.ContextProvider) (result interface{}, err error) {
	p.Infof("received request to handle %s key with %d parameters", key, len(params))

	var lld []dirDiscovery

	dirSearch := strings.Split(params[0], "|")
	patterns := strings.Split(params[1], "|")
	patterns_array := strings.Join(patterns, "|")

	for _, dir := range dirSearch {
		fmt.Println(dir)
		filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				checkRegex := validadePath(patterns_array, path)
				if checkRegex {
					lld = append(lld, dirDiscovery{Name: path})
				} else {
					fmt.Println("Erro na regex", path)
				}
			}
			return nil
		})
	}
	b, err := json.Marshal(&lld)
	if err != nil {
		fmt.Println(err)
	}
	return string(b), nil

}

func init() {
	plugin.RegisterMetrics(&impl, "DirSearch", "dir.search", "Returns a json with the list of directories.")
}
func main() {
	h, err := container.NewHandler(impl.Name())
	if err != nil {
		panic(fmt.Sprintf("failed to create plugin handler %s", err.Error()))
	}
	impl.Logger = &h

	err = h.Execute()
	if err != nil {
		panic(fmt.Sprintf("failed to execute plugin handler %s", err.Error()))
	}

}

func validadePath(pattern, path string) bool {
	patterns := strings.Split(pattern, "|")
	// fmt.Println(patterns)
	regexes := make([]*regexp.Regexp, 0)
	for _, pattern := range patterns {
		regexes = append(regexes, regexp.MustCompile(pattern))
	}
	for _, re := range regexes {
		validate := re.MatchString(path)
		if validate {
			return validate
		}
	}
	return false
}
