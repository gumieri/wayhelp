package main

import (
	"regexp"
)

// Categories are all groups of key bidings
var Categories = []*Category{
	&Category{
		Label: "Sway",
		Matchings: []*Matching{
			&Matching{
				String:  "reload",
				Replace: "Reload Sway configurations",
			},
			&Matching{
				String:  "exit",
				Replace: "Exit Sway session",
			},
			&Matching{
				Pattern: regexp.MustCompile(`.*swaylock.*`),
				Replace: "Lock screen",
			},
			&Matching{
				Pattern: regexp.MustCompile(`.*(rofi|dmenu).*`),
				Replace: "Menu Applications",
			},
			&Matching{
				Pattern: regexp.MustCompile(`.*(xterm|gnome-terminal|termite|alacritty|kitty).*`),
				Replace: "Open terminal",
			},
		},
	},
	&Category{
		Label: "Workspace",
		Matchings: []*Matching{
			&Matching{
				Pattern: regexp.MustCompile(`^workspace`),
				Replace: "Switch to Workspace",
			},
		},
	},
	&Category{
		Label: "Resizing",
		Matchings: []*Matching{
			&Matching{
				String:  `mode "default"`,
				Replace: `Leave "resize mode"`,
			},
			&Matching{
				String:  `mode "resize"`,
				Replace: `Activate "resize mode"`,
			},
			&Matching{
				Pattern: regexp.MustCompile(`^resize`),
				Replace: " ",
			},
		},
	},
	&Category{
		Label: "Containers",
		Matchings: []*Matching{
			&Matching{
				Pattern: regexp.MustCompile(`^move (up|down|left|right)`),
				Replace: "Move container $1",
			},
			&Matching{Pattern: regexp.MustCompile(`^move container to workspace .*`)},
			&Matching{
				Pattern: regexp.MustCompile(`^focus (up|down|left|right)`),
				Replace: "Focus the container $1",
			},
			&Matching{
				String:  "floating toggle",
				Replace: "Toggle focused container between floating and tiling",
			},
			&Matching{
				String:  "focus mode_toggle",
				Replace: "Toggle focus between floating and tiling containers",
			},
			&Matching{
				String:  "focus parent",
				Replace: "Focus the parent containers",
			},
			&Matching{
				String:  "kill",
				Replace: "Close the focused container",
			},
		},
	},
	&Category{
		Label: "Layout",
		Matchings: []*Matching{
			&Matching{
				String:  "fullscreen",
				Replace: "Toggle the with fullscreen layout",
			},
			&Matching{
				Pattern: regexp.MustCompile(`^layout toggle (.*)`),
				Replace: "Toggle with $1 layout",
			},
			&Matching{
				Pattern: regexp.MustCompile(`^layout (.*)`),
				Replace: "Change to $1 layout",
			},
		},
	},
	&Category{
		Label: "Split",
		Matchings: []*Matching{
			&Matching{
				String:  "splith",
				Replace: "Split the tile horizontally",
			},
			&Matching{
				String:  "splitv",
				Replace: "Split the tile vertically",
			},
		},
	},
	&Category{
		Label: "Scratchpad",
		Matchings: []*Matching{
			&Matching{
				Pattern: regexp.MustCompile(`^move( to)? scratchpad`),
				Replace: "Move the focused container to the scratchpad",
			},
			&Matching{
				String:  "scratchpad show",
				Replace: "Show as window the first container sent to scratchpad",
			},
		},
	},
}
