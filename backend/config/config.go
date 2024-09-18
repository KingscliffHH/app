package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
)

type Auth0Config struct {
	Domain       string
	ClientID     string
	ClientSecret string
	Audience     string
}

type Env struct {
	MONGODB_URI  string
	MONGODB_NAME string
	Auth0        struct {
		App Auth0Config
		Api Auth0Config
	}
	ListenAddr string
	S3Bucket   string
}

func LoadConfig() (Env, error) {
	listenAddr := flag.String("listen-addr", ":8000", "server listen address")
	flag.Parse()

	godotenv.Load()

	return Env{
		MONGODB_URI:  os.Getenv("MONGODB_URI"),
		MONGODB_NAME: os.Getenv("MONGODB_NAME"),
		Auth0: struct {
			App Auth0Config
			Api Auth0Config
		}{
			App: Auth0Config{
				Domain:   os.Getenv("AUTH0_APP_DOMAIN"),
				ClientID: os.Getenv("AUTH0_APP_CLIENT_ID"),
				Audience: os.Getenv("AUTH0_APP_AUDIENCE"),
			},
			Api: Auth0Config{
				Domain:       os.Getenv("AUTH0_API_DOMAIN"),
				ClientID:     os.Getenv("AUTH0_API_CLIENT_ID"),
				ClientSecret: os.Getenv("AUTH0_API_CLIENT_SECRET"),
				Audience:     os.Getenv("AUTH0_API_AUDIENCE"),
			},
		},
		ListenAddr: *listenAddr,
		S3Bucket:   os.Getenv("S3_BUCKET"),
	}, nil
}
