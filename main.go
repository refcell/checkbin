package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "checkbin",
		Usage: "checkbin [global options] binary [command options] [arguments...]",
		// Flags: []cli.Flag{
		//       &cli.StringFlag{
		//           Name:  "algorithm",
		//           Value: "sha256",
		//           Usage: "language for the greeting",
		//       },
		//   },
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() == 0 {
				return errors.New("missing binary name")
			}
			return checksumBinary(ctx)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func checksumBinary(ctx *cli.Context) error {
	bin := ctx.Args().Get(0)
	if bin == "" {
		return errors.New("missing binary name")
	}

	f, err := os.Open(bin)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))

	return nil
}
