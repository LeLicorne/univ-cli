package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var battleCmd = &cobra.Command{
	Use:     "battle",
	Aliases: []string{"b"},
	Short:   "Gérer les batailles",
	Long: `La commande 'battle' permet de gérer les batailles dans le jeu.
Elle offre des fonctionnalités telles que le démarrage d'une bataille, l'affichage des statistiques
des personnages et la gestion des objets.`,

	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(initialBattleModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Erreur: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(battleCmd)
}
