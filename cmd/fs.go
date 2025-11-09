package cmd

import (
	"fmt"

	filesystemcli "example.com/univ-cli/cmd/filesystem-cli"
	"github.com/spf13/cobra"
)

var fsCmd = &cobra.Command{
	Use:     "fs",
	Aliases: []string{"filesystem"},
	Short:   "Permet de naviguer et manipuler le système de fichiers local.",
	Long: `La commande 'fs' permet aux utilisateurs de naviguer et de manipuler
			le système de fichiers local. Elle offre des fonctionnalités telles que la
			navigation dans les répertoires, la création, la suppression et la modification
			de fichiers et de dossiers, ainsi que la gestion des permissions et des
			attributs des fichiers. Cette commande est essentielle pour interagir avec
			le système de fichiers de manière efficace et sécurisée.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Utilisez une des sous-commandes: ls, cat, cp, mkdir")
	},
}

func init() {
	rootCmd.AddCommand(fsCmd)
	fsCmd.AddCommand(filesystemcli.LsCmd)
	fsCmd.AddCommand(filesystemcli.CatCmd)
	fsCmd.AddCommand(filesystemcli.CpCmd)
	fsCmd.AddCommand(filesystemcli.MkdirCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
