package filesystemcli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var MkdirCmd = &cobra.Command{
	Use:   "mkdir <dir>",
	Short: "Crée un nouveau répertoire avec le nom spécifié",
	Long:  `Crée un nouveau répertoire avec le nom spécifié.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirName := args[0]

		err := os.Mkdir(dirName, 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors de la création du répertoire: %v\n", err)
			return
		}

		fmt.Printf("Répertoire '%s' créé avec succès\n", dirName)
	},
}
