package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	sway "github.com/gumieri/go-sway"
)

// OtherCategory if the key matches nothing
var OtherCategory = &Category{
	Label: "Other",
}

func main() {
	sc, err := sway.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	c := parseConfig(sc)
	for key, v := range c.configurations {
		categorized := false

		for _, category := range Categories {
			if found, matching := category.Match(v); found {
				categorized = true

				category.AddAction(v, key, matching)
			}
		}

		if !categorized {
			OtherCategory.AddAction(v, key, nil)
		}
	}

	for _, category := range append(Categories, OtherCategory) {
		if len(category.Actions) == 0 {
			continue
		}

		fmt.Println("#", category.Label, "actions")

		sort.Slice(category.Actions, func(i, j int) bool {
			return category.Actions[i].Command < category.Actions[j].Command
		})

		for _, action := range category.Actions {

			sort.Slice(action.Keys, func(i, j int) bool {
				iLen := len(action.Keys[i])
				jLen := len(action.Keys[j])
				if iLen < jLen {
					return true
				} else if iLen > jLen {
					return false
				} else {
					return action.Keys[i] < action.Keys[j]
				}
			})

			fmt.Println("  ", action.Description()+":", strings.Join(action.Keys, " / "))
		}
	}
}
