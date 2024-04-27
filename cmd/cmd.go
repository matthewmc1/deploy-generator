package cmd

import (
	"fmt"
	"validator/service-definition/internal"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

func BaseCommand() *cobra.Command {
	intro := figure.NewFigure("Deployment", "basic", true).ColorString()

	deploydef := &cobra.Command{
		Use:   "deploy",
		Short: "Deployment Definition Generator",
		Long:  "CLI for translating a small deployment file, including some of the current best practices for security and replication to avoid users making issues on creation of K8s deployments",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(intro)
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Prints the application version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(internal.Version())
		},
	}
	deploydef.AddCommand(versionCmd)

	create := &cobra.Command{
		Use:   "create",
		Short: "Generates a compliant k8s manifest",
		Run: func(cmd *cobra.Command, args []string) {
			internal.Generator(args[0])
		},
	}
	deploydef.AddCommand(create)

	return deploydef
}
