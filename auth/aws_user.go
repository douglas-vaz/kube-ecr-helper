package auth

import (
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
	"log"
	"os"
	"strings"
	"time"
)

type AwsUser struct {
	sess *session.Session
}

type Token struct {
	Username  string
	Password  string
	Server    string
	ExpiresAt time.Time
}

func Login() (*AwsUser, error) {
	region := mandatoryEnvVar("AWS_REGION")
	mandatoryEnvVar("AWS_ACCESS_KEY_ID")
	mandatoryEnvVar("AWS_SECRET_ACCESS_KEY")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewEnvCredentials(),
	})

	if err != nil {
		return nil, err
	}

	return &AwsUser{sess: sess}, nil
}

func (user *AwsUser) GetToken() (Token, error) {
	svc := ecr.New(user.sess)
	input := &ecr.GetAuthorizationTokenInput{}

	result, err := svc.GetAuthorizationToken(input)
	if err != nil {
		return Token{}, err
	}

	data := result.AuthorizationData[0]

	byteArray, err := base64.StdEncoding.DecodeString(*data.AuthorizationToken)
	if err != nil {
		return Token{}, err
	}

	bytesAsString := string(byteArray[:])
	decodedString := strings.Split(bytesAsString, ":")

	if len(decodedString) != 2 {
		return Token{}, fmt.Errorf("Invalid token %s\n", bytesAsString)
	}

	token := Token{
		Username:  decodedString[0],
		Password:  decodedString[1],
		ExpiresAt: *data.ExpiresAt,
		Server:    strings.Replace(*data.ProxyEndpoint, "https://", "", 1),
	}

	return token, nil
}

func mandatoryEnvVar(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("%s not set in environment", key)
	}
	return value
}
