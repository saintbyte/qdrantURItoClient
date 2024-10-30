package main

import (
	"github.com/saintbyte/qdrantURItoClient"
	"log/slog"
	"os"
)

func main() {
	var qdrant_url string = os.Getenv("QDRANT_URL")
	if len(os.Args) > 1 {
		qdrant_url = os.Args[1]
	} else if qdrant_url == "" {
		qdrant_url = "qdrant://api_key_sfdfdsff@" +
			"qy-blue-block-65767118.eu-central-1.aws.neon.tech?UseTLS=True"
	}
	slog.Info("qdrant_url: ", qdrant_url)
	cl, err := qdrantURItoClient.UriToClient(qdrant_url)
	if err != nil {
		slog.Error("dsn error: ", err)
		return
	}
	slog.Info("cl: ", cl)
}
