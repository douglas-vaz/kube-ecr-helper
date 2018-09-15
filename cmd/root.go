package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kube-ecr-helper",
	Short: "Utility to fetch and refresh private AWScECR credentials for Kubernetes",
	Long: `Generates a kubectl command to create a Kubernetes secret to pull private ECR image.
Prerequisite: AWS_REGION, AWS_ACCESS_KEY_ID, and AWS_SECRET_ACCESS_KEY are set as environment variables`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
