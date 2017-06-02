// Package tparse implements data structures and simple functions to parse toml like syntax
package tparse

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Entries Type is used for containing the key value pairs
type Entries map[string]string

// Dict Type is used for containing the whole toml like data
type Dict map[string]Entries

// Find returns value for the given string from the entries data structure.
// If not found it returns an empty string and an error.
func (e Entries) Find(key string) (string, error) {
	ret, ok := e[key]
	if !ok {
		return "", errors.New("No such entry found")
	}
	return ret, nil
}

// NewDict returns the pointer to a Dict instance after initializing it.
func NewDict() *Dict {
	a := make(Dict, 0)
	return &a
}

// Find returns Entries if found and an error if not for a given heading.
func (d Dict) Find(heading string) (Entries, error) {
	ret, ok := d[heading]
	if !ok {
		return nil, errors.New("No such section found")
	}
	return ret, nil
}

// Parse is used to parse the content and store it in the Dict variable.
func (d Dict) Parse(contents string) error {
	lines := strings.Split(contents, "\n")
	head := regexp.MustCompile("^\\s*\\[.*\\]\\s*$")
	kval := regexp.MustCompile("^[^\\[].*=.*")
	empty := regexp.MustCompile("^\\s*$")
	currentHeader := "root"
	for i, val := range lines {
		switch true {
		case kval.MatchString(val):
			key, value := getKeyValPair(val)
			d[currentHeader][key] = value
		case head.MatchString(val):
			currentHeader = getHeader(val)
			d[currentHeader] = make(map[string]string, 0)
		case empty.MatchString(val):
			continue
		default:
			return errors.New(fmt.Sprintf("Illegal Character in line %d", i))
		}
	}
	return nil
}

// getHeader is used to get the heading from the toml section header.
func getHeader(h string) string {
	head := regexp.MustCompile("^\\s*\\[(?P<heading>.*)\\]\\s*$")
	sxn := head.SubexpNames()
	matches := head.FindStringSubmatch(h)
	for i, val := range matches {
		if sxn[i] != "heading" {
			continue
		}
		return strings.TrimSpace(val)
	}
	return ""
}

// getKeyValPair is used to get key and value from the passed toml expression
func getKeyValPair(kv string) (string, string) {
	arr := strings.Split(kv, "=")
	key := arr[0]
	value := ""
	for j, val := range arr {
		if j == 0 {
			continue
		}
		value = value + val
		if j != len(arr)-1 {
			value += "="
		}
	}
	return strings.TrimSpace(key), strings.TrimSpace(value)
}
