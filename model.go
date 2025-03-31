package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	focused Status
	lists   []list.Model
}

var _ tea.Model = (*Model)(nil)

func NewModel() Model {
	todoTasks := NewList([]list.Item{
		Task{title: "Task 1", description: "Description 1"},
		Task{title: "Task 2", description: "Description 2"},
	})
	inProgressTasks := NewList([]list.Item{
		Task{title: "Task 3", description: "Description 3"},
		Task{title: "Task 4", description: "Description 4"},
	})
	doneTasks := NewList([]list.Item{
		Task{title: "Task 5", description: "Description 5"},
		Task{title: "Task 6", description: "Description 6"},
	})

	return Model{
		focused: todo,
		lists: []list.Model{
			todo:       todoTasks,
			inProgress: inProgressTasks,
			done:       doneTasks,
		},
	}
}

func NewList(items []list.Item) list.Model {
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.SetShowHelp(false)
	return l
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		for listID, list := range m.lists {
			width := msg.Width / len(m.lists)
			list.SetSize(width, 20)
			m.lists[listID] = list
		}
	case tea.KeyMsg:
		m = m.handleKeyMsg(msg)
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) handleKeyMsg(msg tea.KeyMsg) Model {
	switch msg.String() {
	// case "ctrl+c", "q":
	// 	return m
	case "l":
		m = m.nextFocused()
	case "h":
		m = m.previousFocused()
	}
	return m
}

func (m Model) nextFocused() Model {
	m.focused = Status((int(m.focused) + 1) % len(m.lists))
	return m
}

func (m Model) previousFocused() Model {
	m.focused = Status((int(m.focused) - 1 + len(m.lists)) % len(m.lists))
	return m
}

func (m Model) View() string {
	views := make([]string, 0, len(m.lists))
	for status, list := range m.lists {
		var view string
		if Status(status) == m.focused {
			view = focusedStyle.Render(list.View())
		} else {
			view = columnStyle.Render(list.View())
		}
		// view = list.View()
		views = append(views, view)
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, views...)
}
