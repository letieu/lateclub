package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SplashState struct {
	complete bool
}

type UserSignedInMsg struct {
}

type DelayCompleteMsg struct{}

func (m model) IsLoadingComplete() bool {
	return m.state.splash.complete
}

func (m model) SplashInit() tea.Cmd {
	// 2 second delay
	delayCmd := tea.Tick(
		500_000_000, // 1 second in nanoseconds
		func(t time.Time) tea.Msg {
			return DelayCompleteMsg{}
		},
	)

	disableMouseCmd := func() tea.Msg {
		return tea.DisableMouse()
	}

	return tea.Batch(m.CursorInit(), disableMouseCmd, delayCmd)
}

func (m model) SplashUpdate(msg tea.Msg) (model, tea.Cmd) {
	switch msg.(type) {
	case DelayCompleteMsg:
		m.state.splash.complete = true
	}
	return m, nil
}

func (m model) SplashView() string {
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			m.LogoView(),
		),
	)
}
