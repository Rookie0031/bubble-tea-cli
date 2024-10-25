package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

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
		case "enter":
			m.selected = m.choices[m.cursor]
			if m.step == 0 {
				if m.cursor == 0 {
					return m, m.accessDevEks
				} else if m.cursor == 1 {
					m.step = 1
					m.choices = []string{"로컬 경로 입력", "원격 경로 입력", "파일 복사 실행"}
					m.cursor = 0
				}
			} else if m.step == 1 {
				if m.cursor == 0 {
					return m, m.promptForInput("로컬 경로를 입력하세요: ")
				} else if m.cursor == 1 {
					return m, m.promptForInput("원격 경로를 입력하세요: ")
				} else if m.cursor == 2 {
					return m, m.copyFile
				}
			}
		}
	case string:
		if m.step == 1 {
			if m.localPath == "" {
				m.localPath = msg
			} else {
				m.remotePath = msg
			}
		}
	}
	return m, nil
}
