package main

import (
	"fmt"
)

func (m model) View() string {
	s := "개발자 도구:\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\n위/아래 화살표로 이동, 엔터로 선택, q로 종료\n"

	if m.localPath != "" {
		s += fmt.Sprintf("\n로컬 경로: %s\n", m.localPath)
	}
	if m.remotePath != "" {
		s += fmt.Sprintf("원격 경로: %s\n", m.remotePath)
	}

	return s
}
