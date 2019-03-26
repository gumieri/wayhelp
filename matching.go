package main

import "regexp"

// Matching of a command
type Matching struct {
	Pattern *regexp.Regexp
	String  string
	Replace string
}
