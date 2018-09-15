package cmd

import (
	"fmt"
	"github.com/douglas-vaz/kube-ecr-helper/auth"
	"github.com/spf13/cobra"
	"log"
)

var SecretName string
var Email string

var createSecretCmd = &cobra.Command{
	Use: "get-apply",

	Long: `Outputs the kubectl command to create/replace K8s secret with refreshed ECR credentials.
Usage: kube-ecr-helper get-apply`,

	Run: func(cmd *cobra.Command, args []string) {
		user, err := auth.Login()
		check(err)

		token, err := user.GetToken()
		check(err)

		fmt.Println(buildCommand(token))
	},
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func buildCommand(t auth.Token) string {
	return fmt.Sprintf("kubectl create secret docker-registry %s --docker-server=%s --docker-username=%s --docker-password=%s --docker-email=%s --dry-run -o yaml",
		SecretName, t.Server, t.Username, t.Password, Email)
}

func init() {
	createSecretCmd.Flags().StringVarP(&SecretName, "secret", "s", "docker-credentials", "Name of the secret resource")
	createSecretCmd.Flags().StringVarP(&Email, "email", "e", "", "Email (optional)")

	createSecretCmd.MarkFlagRequired("secret")
	createSecretCmd.MarkFlagRequired("email")

	rootCmd.AddCommand(createSecretCmd)
}
