package main

import "strings"

// Action refers to a configured sway command
type Action struct {
	Command  string
	Matching *Matching
	Keys     []string
}

// Description of the action
func (a *Action) Description() (description string) {
	if a.Matching == nil || a.Matching.Replace == "" {
		description = a.Command
	} else if a.Matching.Pattern != nil {
		description = a.Matching.Pattern.ReplaceAllString(a.Command, a.Matching.Replace)
	} else {
		description = a.Matching.Replace
	}

	description = strings.Trim(description, " ")

	return
}
