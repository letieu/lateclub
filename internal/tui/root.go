package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width  int
	height int
	theme  Theme
	state
	renderer *lipgloss.Renderer
}

type state struct {
	splash SplashState
	cursor cursorState
}

func NewModel(renderer *lipgloss.Renderer) *model {
	return &model{
		theme: BasicTheme(renderer, nil),
	}
}

func (m model) Init() tea.Cmd {
	return m.SplashInit()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case CursorTickMsg:
		m, cmd := m.CursorUpdate(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	return m.SplashView()
}
