package main

// Category is the grouping of a set of key bidings
type Category struct {
	Label     string
	Actions   []*Action
	Matchings []*Matching
}

// AddAction append a key
func (c *Category) AddAction(command string, key string, matching *Matching) {
	for _, action := range c.Actions {
		if command == action.Command {
			action.Keys = append(action.Keys, key)
			return
		}
	}
	c.Actions = append(c.Actions, &Action{
		Command:  command,
		Matching: matching,
		Keys:     []string{key},
	})
}

// Match check if given value match to the category
func (c *Category) Match(v string) (matches bool, matching *Matching) {
	for _, m := range c.Matchings {
		if m.Pattern != nil {
			if m.Pattern.MatchString(v) {
				matches = true
				matching = m
				return
			}
		}
		if m.String == v {
			matches = true
			matching = m
			return
		}
	}

	return
}
