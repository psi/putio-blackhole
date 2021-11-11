package main

import (
	"context"

	"github.com/putdotio/go-putio"
	"golang.org/x/oauth2"
)

func NewPutIOClient(token string) *putio.Client {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	oauthClient := oauth2.NewClient(context.TODO(), tokenSource)

	return putio.NewClient(oauthClient)
}

func AddTorrentToPutIO(torrentUrl string, client *putio.Client) error {
	_, err := client.Transfers.Add(context.Background(), torrentUrl, 0, "")
	return err
}
