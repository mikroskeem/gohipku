package main

import (
	"fmt"
	"net"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/mikroskeem/gohipku/pkg"
)

func main() {
	app := &cli.App{
		Name:  "hipku",
		Usage: "Encode & decode IP address to/from haiku",
		Commands: []*cli.Command{
			{
				Name:   "encode",
				Usage:  "Encode IP to haiku",
				Action: encode,
			},
			{
				Name:   "decode",
				Usage:  "Decode IP from haiku",
				Action: decode,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func encode(cctx *cli.Context) (err error) {
	if cctx.NArg() < 1 {
		return fmt.Errorf("Expected at least one argument")
	}

	rawIP := cctx.Args().Get(0)
	ip := net.ParseIP(rawIP)

	haiku, ok := gohipku.Encode(ip)
	if !ok {
		return fmt.Errorf("Invalid IP? %s", ip)
	}

	fmt.Println(haiku)

	return
}

func decode(cctx *cli.Context) (err error) {
	if cctx.NArg() < 1 {
		return fmt.Errorf("Expected at least one argument")
	}

	haiku := cctx.Args().Get(0)

	ip, err := gohipku.Decode(haiku)
	if err != nil {
		return fmt.Errorf("Invalid hipku? %w", err)
	}

	fmt.Println(ip)

	return
}
