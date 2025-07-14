package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SplashState struct {
	data  bool
	delay bool
}

type UserSignedInMsg struct {
}

type DelayCompleteMsg struct{}

func (m model) IsLoadingComplete() bool {
	return m.state.splash.data &&
		m.state.splash.delay
}

func (m model) SplashInit() tea.Cmd {
	disableMouseCmd := func() tea.Msg {
		return tea.DisableMouse()
	}

	return tea.Batch(m.CursorInit(), disableMouseCmd)
}

func (m model) SplashView() string {
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			"",
			"",
			"",
			"",
			m.LogoView(),
			"",
			"",
		),
	)
}
