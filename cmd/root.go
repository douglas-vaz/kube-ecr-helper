package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kube-ecr",
	Short: "Utility to fetch and refresh private AWScECR credentials for Kubernetes",
	Long:  `Generates a shell command to create a Kubernetes secret to ECR image pulls`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
