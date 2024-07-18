package main

import (
	"fmt"
	"sort"

	"github.com/Masterminds/semver/v3"
)

type Language struct {
	Language string
	Version  string
}

var items = []Language{
	{"cpp", "2.7"}, {"cpp", "2020"},
	{"golang", "1.21"}, {"golang", "1.20"}, {"golang", "1.2"}, {"golang", "1.19"}, {"golang", "1.12.13"}, {"golang", "1.17"}, {"golang", "1.8"},
	{"lua", "5"},
}

func main() {
	m := make(map[string][]string)
	for _, one := range items {
		m[one.Language] = append(m[one.Language], one.Version)
	}
	for l, vs := range m {
		fmt.Println(l)
		fmt.Printf("original:\t%s\n", m[l])
		m[l] = sortVersions(vs)
		fmt.Printf("sorted:\t\t%s\n", m[l])
	}
}

func sortVersions(vs []string) []string {
	var versions []*semver.Version
	for _, v := range vs {
		versions = append(versions, semver.MustParse(v))
	}
	// newer version comes first
	sort.Sort(sort.Reverse(semver.Collection(versions)))
	var ret []string
	for _, v := range versions {
		ret = append(ret, v.Original())
	}
	return ret
}
