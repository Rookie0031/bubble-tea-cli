package main

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	choices    []string
	cursor     int
	selected   string
	step       int
	localPath  string
	remotePath string
}

func initialModel() model {
	return model{
		choices: []string{"dev-eks에 접근", "파일 복사"},
		cursor:  0,
		step:    0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
