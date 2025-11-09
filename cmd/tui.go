/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	titleStyle = lipgloss.NewStyle().
			MarginLeft(2).
			Foreground(lipgloss.Color("205")).
			Bold(true)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4)

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("170"))
)

type item struct {
	title string
	value string
}

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.title)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + s[0])
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			if m.choice != "" && !m.quitting {
				m.choice = ""
				return m, nil
			}
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i.value)
			}

			switch m.choice {
			case "datetime":
				return m, nil
			case "message":
				return m, nil
			case "quit":
				m.quitting = true
				return m, tea.Quit
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" && !m.quitting {
		var output string
		switch m.choice {
		case "datetime":
			output = fmt.Sprintf("\n🕐 Date et heure: %s\n\nAppuyez sur 'q' pour revenir au menu", time.Now().Format("02/01/2006 15:04:05"))
		case "message":
			output = "\n👋 Bienvenue dans univ-cli !\n\nCette application vous permet de naviguer et manipuler le système de fichiers.\n\nAppuyez sur 'q' pour revenir au menu"
		}
		return output
	}

	if m.quitting {
		return "Au revoir! 👋\n"
	}

	return "\n" + titleStyle.Render("Menu TUI") + "\n\n" + m.list.View()
}

func initialModel() model {
	items := []list.Item{
		item{"Afficher la date et l'heure actuelles", "datetime"},
		item{"Afficher un message de bienvenue", "message"},
		item{"Quitter l'application", "quit"},
	}

	const defaultWidth = 80

	l := list.New(items, itemDelegate{}, defaultWidth, 14)
	l.Title = "Sélectionnez une option"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle

	return model{list: l}
}

// tuiCmd represents the tui command
var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Lance une interface utilisateur textuelle (TUI)",
	Long: `Lance une interface utilisateur textuelle (TUI) simple qui affiche un menu avec les options suivantes:
- Afficher la date et l'heure actuelles.
- Afficher un message de bienvenue.
- Quitter l'application.`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(initialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Erreur lors du lancement du TUI: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tuiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tuiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
