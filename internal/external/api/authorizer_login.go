package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Pos-Tech-Challenge-48/delivery-api/internal/entities"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type AuthorizerAPI struct {
	cognitoClient *cognitoidentityprovider.CognitoIdentityProvider
	userPoolID    string
	clientID      string
}

func NewAuthorizer() *AuthorizerAPI {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	userPoolID := os.Getenv("AWS_USER_POOL_ID")
	clientID := os.Getenv("AWS_CLIENT_ID")

	cognitoClient := cognitoidentityprovider.New(sess)
	return &AuthorizerAPI{
		cognitoClient: cognitoClient,
		userPoolID:    userPoolID,
		clientID:      clientID,
	}
}

func (a *AuthorizerAPI) Execute(ctx context.Context, login *entities.Login) (string, error) {
	params := &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:        aws.String(a.userPoolID),
		Username:          aws.String(login.Email),
		TemporaryPassword: aws.String("Temp@123"),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(login.Email),
			},
		},
	}

	_, err := a.cognitoClient.AdminCreateUser(params)
	if err != nil {
		log.Println("Err to create a user:", err)
		return "", err
	}

	authParams := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow:       aws.String("ADMIN_USER_PASSWORD_AUTH"),
		UserPoolId:     aws.String(a.userPoolID),
		ClientId:       aws.String(a.clientID),
		AuthParameters: map[string]*string{"USERNAME": aws.String(login.Email), "PASSWORD": aws.String("Temp@123")},
	}

	authResp, err := a.cognitoClient.AdminInitiateAuth(authParams)
	if err != nil {
		log.Println("Err to auth the user:", err)
		return "", err
	}

	if *authResp.ChallengeName != "NEW_PASSWORD_REQUIRED" {
		log.Println("Erro to create the challenge user:", err)
		return "", err
	}

	challengeResponses := map[string]*string{
		"USERNAME":     aws.String(login.Email),
		"PASSWORD":     aws.String(login.Password),
		"NEW_PASSWORD": aws.String(login.Password),
	}

	respondToAuthChallengeInput := &cognitoidentityprovider.AdminRespondToAuthChallengeInput{
		UserPoolId:         aws.String(a.userPoolID),
		ClientId:           aws.String(a.clientID),
		ChallengeName:      aws.String("NEW_PASSWORD_REQUIRED"),
		ChallengeResponses: challengeResponses,
		Session:            authResp.Session,
	}

	respondOutput, err := a.cognitoClient.AdminRespondToAuthChallenge(respondToAuthChallengeInput)
	if err != nil {
		fmt.Println("Err to create new password:", err)
		return "", err
	}

	return aws.StringValue(respondOutput.AuthenticationResult.AccessToken), nil
}
