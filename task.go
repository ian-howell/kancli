package main

import (
	"github.com/charmbracelet/bubbles/list"
)

type Status int

const (
	todo Status = iota
	inProgress
	done
)

type Task struct {
	title       string
	description string
}

var _ list.Item = (*Task)(nil)

func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}
