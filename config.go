package main

import (
	"bufio"
	"strings"

	sway "github.com/gumieri/go-sway"
)

const defaultModKey string = "Mod1"

var modKeys = map[string]string{
	"Mod1": "Alt",
	"Mod2": "Num_Lock",
	"Mod3": "",
	"Mod4": "Super",
	"Mod5": "Mode_Switch",
}

type config struct {
	modifier       string
	variables      map[string]string
	configurations map[string]string
}

func (c config) parseLine(line string) {
	fields := strings.Fields(line)
	if len(fields) == 0 {
		return
	}

	switch fields[0] {
	case "set":
		if strings.Contains(line, "Mod") {
			c.modifier = modKeys[fields[2]]
			c.variables[fields[1]] = modKeys[fields[2]]
			return
		}

		c.variables[fields[1]] = strings.Join(fields[2:], " ")

	case "bindsym":
		var key, value string
		if strings.HasPrefix(fields[1], "--") {
			key = fields[2]
			value = strings.Join(fields[3:], " ")
		} else {
			key = fields[1]
			value = strings.Join(fields[2:], " ")
		}

		arK := strings.Split(key, "+")
		for i := 0; i < cap(arK); i++ {
			if strings.HasPrefix(arK[i], "$") {
				arK[i] = c.variables[arK[i]]
			}
		}

		for k, v := range c.variables {
			value = strings.Replace(value, k, v, -1)
		}

		c.configurations[strings.Join(arK, "+")] = value
	}
}

func parseConfig(sc sway.Config) (c config) {
	c = config{
		modifier:       modKeys[defaultModKey],
		variables:      make(map[string]string),
		configurations: make(map[string]string),
	}

	scanner := bufio.NewScanner(strings.NewReader(sc.Config))
	for scanner.Scan() {
		c.parseLine(scanner.Text())
	}

	return
}
