package filesystemcli

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var CatCmd = &cobra.Command{
	Use:   "cat <file>",
	Short: "Affiche le contenu du fichier spécifié",
	Long:  `Affiche le contenu du fichier spécifié.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors de la lecture du fichier: %v\n", err)
			return
		}

		fmt.Print(string(content))
	},
}
