package main

import "github.com/charmbracelet/lipgloss"

var (
	columnStyle = lipgloss.NewStyle().
			Padding(1, 2)
	focusedStyle = columnStyle.
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)
