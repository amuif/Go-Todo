package cli

import (
	tea "charm.land/bubbletea/v2"
	"github.com/amuif/go-cli/internal/todo"
)

type view int

const (
	listView view = iota
	addView
	helpView
)

type Model struct {
	todoService *todo.Service

	// UI
	currentView view
	todos       []todo.Todo
	cursor      int //the selected todo
	helpVisible bool

	// to add new todo
	addInputs struct {
		title       string
		description string
		dueDate     string
		focused     string // which input is focused
	}

	// dimenstions
	width  int
	height int

	err error
}

func NewModel(service *todo.Service) *Model {
	return &Model{
		todoService: service,
		currentView: listView,
		cursor:      0,
		helpVisible: false,
		todos:       []todo.Todo{},
	}
}

func (m Model) Init() tea.Cmd {
	return m.loadTodos()
}
func (m Model) loadTodos() tea.Cmd {
	return func() tea.Msg {
		todos, err := m.todoService.GetAllTodos()
		if err != nil {
			return errMsg{err}
		}
		return todosLoaderMsg{todos}
	}
}

type todosLoaderMsg struct {
	todos []todo.Todo
}
type errMsg struct {
	err error
}

func (e errMsg) Error() string {
	return e.err.Error()
}
