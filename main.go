package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func initialModel() model {
	return model{
		choices: []Option{
			{ID: 1, Label: "Create a new task"},
			{ID: 2, Label: "List created todo"},
		},
		cursor: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.selected = m.choices[m.cursor].ID
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	var s strings.Builder
	s.WriteString("Welcome to Go todo\nWhat's the schedule for today?\n\n")

	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		fmt.Fprintf(&s, "%s [%d] %s\n", cursor, choice.ID, choice.Label)
	}

	s.WriteString("\n(Press q to quit)\n")
	return tea.NewView(s.String())
}
func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
