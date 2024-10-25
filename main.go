package main

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	step        int
	localPath   string
	inputFocus  bool
	selectedEKS int
	eksOptions  []string
}

var (
	titleStyle     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF69B4"))
	stepStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF"))
	highlightStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFD700"))
	listStyle      = lipgloss.NewStyle().MarginLeft(4)
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.inputFocus {
			switch msg.String() {
			case "enter":
				m.inputFocus = false
				m.executeStep() // 로컬 경로 입력 후 명령 실행
			case "backspace":
				if len(m.localPath) > 0 {
					m.localPath = m.localPath[:len(m.localPath)-1]
				}
			default:
				m.localPath += msg.String()
			}
		} else {
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "n":
				m.step++
			case "b":
				if m.step > 0 {
					m.step--
				}
			case "up":
				if m.selectedEKS > 0 {
					m.selectedEKS--
				}
			case "down":
				if m.selectedEKS < len(m.eksOptions)-1 {
					m.selectedEKS++
				}
			case "enter":
				if m.step == 0 {
					m.executeStep()
					m.step++            // EKS 선택 후 다음 단계로 이동
					m.inputFocus = true // 로컬 경로 입력 모드로 전환
				}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var view string
	switch m.step {
	case 0:
		view = stepStyle.Render("Step 1: EKS 클러스터 선택\n")
		for i, option := range m.eksOptions {
			cursor := " "
			if i == m.selectedEKS {
				cursor = ">"
			}
			view += listStyle.Render(fmt.Sprintf("%s %s\n", cursor, option))
		}
		view += highlightStyle.Render("Use arrow keys to select, 'enter' to confirm, 'q' to quit.\n")
	case 1:
		selectedCluster := m.eksOptions[m.selectedEKS]
		view = stepStyle.Render(fmt.Sprintf("Step 2: 로컬로 데이터 복사 (현재 EKS: %s)\n", selectedCluster)) +
			highlightStyle.Render("Enter local path: ") + m.localPath + "\n" +
			highlightStyle.Render("Press 'b' for back, 'enter' to confirm, 'q' to quit.\n")
	case 2:
		view = highlightStyle.Render("작업이 완료되었습니다. 'q'를 눌러 종료하세요.\n")
	}
	return titleStyle.Render("Bubble Tea TUI\n") + view
}

func (m model) executeStep() {
	switch m.step {
	case 0:
		selectedCluster := m.eksOptions[m.selectedEKS]
		cmd := exec.Command("aws", "eks", "update-kubeconfig", "--region", "ap-northeast-2", "--name", selectedCluster, "--role-arn", "arn:aws:iam::991518841123:role/developer-allow-describe-dev-eks-role")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	case 1:
		if m.localPath != "" {
			cmd := exec.Command("kubectl", "cp", "ad/efs-bastion-dev-ad:/efs", m.localPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			m.step++ // 모든 작업이 완료된 후 다음 단계로 이동
		}
	}
}

func main() {
	initialModel := model{
		eksOptions: []string{"dev-eks", "prod-eks", "test-eks"},
	}
	p := tea.NewProgram(initialModel)
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
