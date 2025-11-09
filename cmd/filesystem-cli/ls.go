package filesystemcli

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var LsCmd = &cobra.Command{
	Use:   "ls [path]",
	Short: "Liste les fichiers et dossiers dans le répertoire spécifié",
	Long: `Liste les fichiers et dossiers dans le répertoire spécifié 
(ou le répertoire courant si aucun chemin n'est fourni).`,
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		files, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors de la lecture du répertoire: %v\n", err)
			return
		}

		for _, file := range files {
			if file.IsDir() {
				fmt.Printf("%s/\n", file.Name())
			} else {
				fmt.Println(file.Name())
			}
		}
	},
}
