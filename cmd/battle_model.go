package cmd

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type gameState int

const (
	stateSelectPokemon gameState = iota
	stateBattle
)

type battleModel struct {
	state            gameState
	availablePokemon []Pokemon
	player           Pokemon
	opponent         Pokemon
	selectedAttack   int
	message          string
	isPlayerTurn     bool
	pokemonChoice    int
	battleIsOver     bool
}

func initialBattleModel() battleModel {
	rand.Seed(time.Now().UnixNano())
	return battleModel{
		state:            stateSelectPokemon,
		availablePokemon: createPokemon(),
		selectedAttack:   0,
		isPlayerTurn:     true,
		pokemonChoice:    0,
		battleIsOver:     false,
	}
}

func (m battleModel) Init() tea.Cmd {
	return nil
}

type attackMsg struct{}

func performAttack() tea.Msg {
	time.Sleep(1500 * time.Millisecond)
	return attackMsg{}
}

func (m battleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		if key == "ctrl+c" || key == "q" {
			return m, tea.Quit
		}

		if key == "r" && m.battleIsOver {
			return initialBattleModel(), nil
		}

		if m.state == stateSelectPokemon {
			return m.handlePokemonSelection(key)
		}

		if m.state == stateBattle && !m.battleIsOver {
			return m.handleBattle(key)
		}

	case attackMsg:
		return m.opponentAttack()
	}

	return m, nil
}

func (m battleModel) handlePokemonSelection(key string) (tea.Model, tea.Cmd) {
	if key == "left" {
		if m.pokemonChoice > 0 {
			m.pokemonChoice--
		}
	}

	if key == "right" {
		if m.pokemonChoice < len(m.availablePokemon)-1 {
			m.pokemonChoice++
		}
	}

	if key == "enter" || key == " " {
		m.player = m.availablePokemon[m.pokemonChoice]
		m.opponent = m.availablePokemon[rand.Intn(len(m.availablePokemon))]
		m.state = stateBattle
		m.message = fmt.Sprintf("Un %s sauvage apparaît!", m.opponent.Name)
	}

	return m, nil
}

func (m battleModel) handleBattle(key string) (tea.Model, tea.Cmd) {
	if !m.isPlayerTurn {
		return m, nil
	}

	if key == "up" {
		if m.selectedAttack >= 2 {
			m.selectedAttack = m.selectedAttack - 2
		}
	}

	if key == "down" {
		if m.selectedAttack < 2 {
			m.selectedAttack = m.selectedAttack + 2
		}
	}

	if key == "left" {
		if m.selectedAttack == 1 || m.selectedAttack == 3 {
			m.selectedAttack--
		}
	}

	if key == "right" {
		if m.selectedAttack == 0 || m.selectedAttack == 2 {
			m.selectedAttack++
		}
	}

	if key == "enter" || key == " " {
		return m.playerAttack()
	}

	return m, nil
}

func (m battleModel) playerAttack() (tea.Model, tea.Cmd) {
	attack := m.player.Attacks[m.selectedAttack]
	m.message = fmt.Sprintf("%s utilise %s!", m.player.Name, attack.Name)

	reduceHP(&m.opponent, attack.Damage)

	if isFainted(m.opponent) {
		m.battleIsOver = true
		m.message = fmt.Sprintf("🎉 %s est K.O.! Vous avez gagné! 🎉", m.opponent.Name)
		return m, nil
	}

	m.isPlayerTurn = false
	return m, performAttack
}

func (m battleModel) opponentAttack() (tea.Model, tea.Cmd) {
	randomIndex := rand.Intn(len(m.opponent.Attacks))
	attack := m.opponent.Attacks[randomIndex]
	m.message = fmt.Sprintf("%s utilise %s!", m.opponent.Name, attack.Name)

	reduceHP(&m.player, attack.Damage)

	if isFainted(m.player) {
		m.battleIsOver = true
		m.message = fmt.Sprintf("💀 %s est K.O.! Vous avez perdu! 💀", m.player.Name)
		return m, nil
	}

	m.isPlayerTurn = true
	return m, nil
}

func (m battleModel) View() string {
	if m.state == stateSelectPokemon {
		return m.viewSelectPokemon()
	}

	if m.battleIsOver {
		return m.viewGameOver()
	}

	return m.viewBattle()
}

func (m battleModel) viewSelectPokemon() string {
	var s strings.Builder

	s.WriteString(battleTitleStyle.Render("⚔️  SÉLECTION DE POKÉMON  ⚔️"))
	s.WriteString("\n\n")
	s.WriteString("Choisissez votre Pokémon:\n\n")

	var cards []string
	for i, p := range m.availablePokemon {
		cards = append(cards, m.renderPokemonCard(p, i == m.pokemonChoice))
	}

	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, cards...))
	s.WriteString("\n\n")
	s.WriteString(helpStyle.Render("↑/↓ ou ←/→: naviguer • enter: sélectionner • q: quitter"))

	return s.String()
}

func (m battleModel) renderPokemonCard(p Pokemon, isSelected bool) string {
	var content strings.Builder

	content.WriteString(cardTitleStyle.Render(p.Name))
	content.WriteString("\n\n")

	content.WriteString(cardHPStyle.Render(fmt.Sprintf("HP: %d", p.HP)))
	content.WriteString("\n")

	content.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render(strings.Repeat("─", 31)))
	content.WriteString("\n")

	content.WriteString(cardAttackLabelStyle.Render("Attaques:"))
	content.WriteString("\n")

	for _, attack := range p.Attacks {
		attackLine := fmt.Sprintf("• %s", attack.Name)
		damageText := fmt.Sprintf("(%d dmg)", attack.Damage)

		spacing := 31 - len(attackLine) - len(damageText)
		if spacing < 1 {
			spacing = 1
		}

		content.WriteString(cardAttackStyle.Render(attackLine))
		content.WriteString(strings.Repeat(" ", spacing))
		content.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5F87")).Render(damageText))
		content.WriteString("\n")
	}

	style := pokemonCardStyle
	if isSelected {
		style = selectedPokemonCardStyle
	}

	return style.Render(content.String())
}

func (m battleModel) viewBattle() string {
	var s strings.Builder

	s.WriteString(battleTitleStyle.Render("⚔️  BATAILLE EN COURS  ⚔️"))
	s.WriteString("\n\n")

	if m.message != "" {
		s.WriteString(messageStyle.Render(m.message))
		s.WriteString("\n\n")
	}

	playerBox := m.renderPokemonBox(m.player, true)
	opponentBox := m.renderPokemonBox(m.opponent, false)
	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, playerBox, "  ", opponentBox))
	s.WriteString("\n\n")

	if m.isPlayerTurn {
		s.WriteString("Choisissez une attaque:\n\n")

		var attack0, attack1, attack2, attack3 string

		for i, attack := range m.player.Attacks {
			attackText := fmt.Sprintf("%s\nDégâts: %d", attack.Name, attack.Damage)

			var attackBox string
			if i == m.selectedAttack {
				attackBox = selectedAttackStyle.Render(attackText)
			} else {
				attackBox = attackStyle.Render(attackText)
			}

			switch i {
			case 0:
				attack0 = attackBox
			case 1:
				attack1 = attackBox
			case 2:
				attack2 = attackBox
			case 3:
				attack3 = attackBox
			}
		}

		row1 := lipgloss.JoinHorizontal(lipgloss.Top, attack0, attack1)
		row2 := lipgloss.JoinHorizontal(lipgloss.Top, attack2, attack3)

		s.WriteString(row1 + "\n")
		s.WriteString(row2 + "\n\n")
		s.WriteString(helpStyle.Render("↑/↓/←/→: naviguer • enter: attaquer • q: abandonner"))
	} else {
		s.WriteString(helpStyle.Render("Tour de l'adversaire..."))
	}

	return s.String()
}

func (m battleModel) viewGameOver() string {
	s := battleTitleStyle.Render("⚔️  FIN DE LA BATAILLE  ⚔️") + "\n\n"

	if strings.Contains(m.message, "gagné") {
		s += victoryStyle.Render(m.message)
	} else {
		s += defeatStyle.Render(m.message)
	}

	s += "\n\n"
	s += helpStyle.Render("r: rejouer • q: quitter")

	return s
}

func (m battleModel) renderPokemonBox(p Pokemon, isPlayer bool) string {
	var s strings.Builder

	style := pokemonBoxStyle
	if !isPlayer {
		style = opponentBoxStyle
	}

	title := "Votre Pokémon"
	if !isPlayer {
		title = "Adversaire"
	}

	s.WriteString(lipgloss.NewStyle().Bold(true).Render(title))
	s.WriteString("\n\n")
	s.WriteString(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA")).Render(p.Name))
	s.WriteString("\n\n")

	hpPercent := float64(p.HP) / float64(p.MaxHP)
	barWidth := 30
	filledWidth := int(hpPercent * float64(barWidth))

	hpStyle := hpBarStyle
	if hpPercent < 0.3 {
		hpStyle = hpBarLowStyle
	}

	hpBar := hpStyle.Render(strings.Repeat("█", filledWidth)) +
		lipgloss.NewStyle().Foreground(lipgloss.Color("#3C3C3C")).Render(strings.Repeat("░", barWidth-filledWidth))

	s.WriteString(fmt.Sprintf("HP: %d/%d\n", p.HP, p.MaxHP))
	s.WriteString(hpBar)

	return style.Render(s.String())
}
