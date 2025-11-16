package cmd

import "github.com/charmbracelet/lipgloss"

var (
	battleTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FAFAFA")).
				Background(lipgloss.Color("#7D56F4")).
				PaddingLeft(2).
				PaddingRight(2).
				MarginBottom(1)

	pokemonBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 2).
			Width(40)

	opponentBoxStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#FF5F87")).
				Padding(1, 2).
				Width(40)

	attackStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#3C3C3C")).
			Padding(1, 2).
			Margin(0, 1).
			Width(35)

	selectedAttackStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#1a1a1a")).
				Background(lipgloss.Color("#04B575")).
				Bold(true).
				Padding(1, 2).
				Margin(0, 1).
				Width(35)

	hpBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575"))

	hpBarLowStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5F87"))

	messageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFF00")).
			Bold(true).
			Italic(true).
			MarginTop(1).
			MarginBottom(1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(1)

	victoryStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true).
			MarginTop(1)

	defeatStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF5F87")).
			Bold(true).
			MarginTop(1)

	pokemonCardStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#626262")).
				Padding(1, 2).
				Width(35).
				Margin(0, 1)

	selectedPokemonCardStyle = lipgloss.NewStyle().
					Border(lipgloss.ThickBorder()).
					BorderForeground(lipgloss.Color("#04B575")).
					Padding(1, 2).
					Width(35).
					Margin(0, 1)

	cardTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Align(lipgloss.Center)

	cardHPStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true)

	cardAttackStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			PaddingLeft(1)

	cardAttackLabelStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#7D56F4")).
				Bold(true).
				MarginTop(1)
)
