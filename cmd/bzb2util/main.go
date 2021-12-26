package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/pkg/errors"

	"github.com/corntoole/bzb2util/backblaze"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {

	// set logging flags
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	keyID := os.Getenv("KEY_ID")
	applicationKey := os.Getenv("APPLICATION_KEY")
	bucketName := os.Getenv("BUCKET_NAME")
	endpoint := os.Getenv("ENDPOINT")
	region := os.Getenv("REGION")

	// new Backblaze B2 client
	b2Client, err := backblaze.NewB2Client(endpoint, region, keyID, applicationKey, "", bucketName)
	if err != nil {
		log.Fatalln("[fatal][app] failed to get B2 client", err)
	}
	var _ = b2Client

	app := &cli.App{
		Name: "bzb2util",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l", "ls"},
				Usage:   "List contents of a bucket",
				Action: func(c *cli.Context) error {
					results, err := b2Client.List()
					if err != nil {
						return err
					}
					if len(results) > 0 {
						log.Println("printing results")
					}
					for _, result := range results {
						fmt.Println(result)
					}
					return nil
				},
			},
			{
				Name:    "upload",
				Aliases: []string{"u"},
				Usage:   "upload a file or directory",
				Action: func(c *cli.Context) error {
					if filename := c.String("file"); filename != "" {
						remoteFilepath := path.Base(filename)
						if p := c.String("output-file"); p != "" {
							remoteFilepath = p
						}
						log.Printf("uploading %s to %s\n", filename, remoteFilepath)
						return b2Client.Upload(remoteFilepath, filename)
					}
					return errors.New("invalid filename")
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Aliases:  []string{"f"},
						Usage:    "Upload file to bucket",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "output-file",
						Aliases: []string{"o"},
						Usage:   "Filepath for remote file",
					},
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
