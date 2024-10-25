package main

import (
	"fmt"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) accessDevEks() tea.Msg {
	cmd := exec.Command("aws", "eks", "update-kubeconfig",
		"--region", "ap-northeast-2",
		"--name", "dev-eks",
		"--role-arn", "arn:aws:iam::991518841123:role/developer-allow-describe-dev-eks-role")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("오류 발생: %v\n%s", err, output)
	}
	return "dev-eks에 성공적으로 접근했습니다."
}

func (m model) promptForInput(prompt string) tea.Cmd {
	return func() tea.Msg {
		fmt.Print(prompt)
		var input string
		fmt.Scanln(&input)
		return input
	}
}

func (m model) copyFile() tea.Msg {
	cmd := exec.Command("kubectl", "cp", m.remotePath, m.localPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("파일 복사 중 오류 발생: %v\n%s", err, output)
	}
	return "파일이 성공적으로 복사되었습니다."
}
