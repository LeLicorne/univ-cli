package filesystemcli

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var CpCmd = &cobra.Command{
	Use:   "cp <file> <destination>",
	Short: "Copie le fichier spécifié vers la destination fournie",
	Long:  `Copie le fichier spécifié vers la destination fournie.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		source := args[0]
		destination := args[1]

		srcFile, err := os.Open(source)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors de l'ouverture du fichier source: %v\n", err)
			return
		}
		defer srcFile.Close()

		destFile, err := os.Create(destination)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors de la création du fichier de destination: %v\n", err)
			return
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors de la copie du fichier: %v\n", err)
			return
		}

		fmt.Printf("Fichier copié avec succès de '%s' vers '%s'\n", source, destination)
	},
}
