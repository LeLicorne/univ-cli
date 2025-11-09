package cmd

import (
	"fmt"

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
		fmt.Println("Ahahahaha! La commande battle est en cours de développement. Restez à l'écoute pour plus de fonctionnalités passionnantes!")
	},
}

func init() {
	rootCmd.AddCommand(battleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// battleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// battleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
