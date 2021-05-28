package Service

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
	"os"
)

type TwitterCredentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func tweet() (err error) {
	credentials := TwitterCredentials{
		ConsumerKey:       os.Getenv("TWITTER_CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("TWITTER_CONSUMER_SECRET"),
		AccessToken:       "",
		AccessTokenSecret: "",
	}
	log.Println("Here")
	_, err = getTwitterClient(&credentials)
	if err != nil {
		log.Println(err)
	}
	return
}

func getTwitterClient(cred *TwitterCredentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(cred.ConsumerKey, cred.ConsumerSecret)
	token := oauth1.NewToken(cred.AccessToken, cred.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}
	log.Println("Success")
	return client, nil
}
