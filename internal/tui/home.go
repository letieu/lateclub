package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HomeState struct {
}

func (m model) HomeInit() tea.Cmd {
	return nil
}

func (m model) HomeUpdate(msg tea.Msg) (model, tea.Cmd) {
	return m, nil
}

func (m model) HomeView() string {
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		m.HeaderView(),
		"",
		m.ChatBox(),
	)

	return m.renderer.NewStyle().PaddingLeft(2).PaddingRight(2).PaddingTop(1).PaddingBottom(1).Render(content)
}

func (m model) HeaderView() string {
	logo := m.theme.TextAccent().Bold(true).Render("Lateclub") + m.theme.Base().Background(m.theme.Brand()).Render(" ")
	sub := m.theme.TextBody().Render("What's up guy")

	row := lipgloss.JoinVertical(
		lipgloss.Left,
		logo,
		sub,
	)

	return row
}

func (m model) ChatBox() string {
	label := m.theme.TextAccent().Background(m.theme.Border()).Render(" Chat ")
	a1 := m.theme.TextBody().Render("[ tieule ]")
	msg1 := m.theme.TextAccent().Render("Hello everyone")
	item1 := lipgloss.JoinHorizontal(lipgloss.Top, a1, " ", msg1)

	a2 := m.theme.TextBody().Render("[ trant ]")
	msg2 := m.theme.TextAccent().Render("this app is nice")
	item2 := lipgloss.JoinHorizontal(lipgloss.Top, a2, " ", msg2)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		label,
		item1,
		item2,
	)
}
