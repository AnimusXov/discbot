package config

import (
	"encoding/json"
	"errors"
	"github.com/dghubble/oauth1"
	"log"
	"net/http"
	"os"
)

var (
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
	twitConfig        *twitConfigStruct
)

type twitConfigStruct struct {
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"AccessTokenSecret"`
}

func ReadTwitConfigConfig() error {
	file, err := os.ReadFile("./" + "twitterconfig" + ".json")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = json.Unmarshal(file, &twitConfig)
	if err != nil {
		log.Fatal(err)
		return err
	}
	ConsumerKey = twitConfig.ConsumerKey
	ConsumerSecret = twitConfig.ConsumerSecret
	AccessToken = twitConfig.AccessToken
	AccessTokenSecret = twitConfig.AccessTokenSecret
	return nil
}

func CreateTwitterClient() *http.Client {
	if _, err := os.Stat("./twitterconfig.json"); errors.Is(err, os.ErrNotExist) {
		//Create oauth client with consumer keys and access token
		config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
		token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
		return config.Client(oauth1.NoContext, token)
	}
	err := ReadTwitConfigConfig()
	if err != nil {
		return nil
	}
	config := oauth1.NewConfig(twitConfig.ConsumerKey, twitConfig.ConsumerSecret)
	token := oauth1.NewToken(twitConfig.AccessToken, twitConfig.AccessTokenSecret)

	return config.Client(oauth1.NoContext, token)
}
