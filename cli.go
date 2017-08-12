package main

import "github.com/urfave/cli"
import "os"

func startCli() {
	var message string
	var keyfile string
	app := cli.NewApp()
	app.Name = "cipher"
	app.Usage = "Liam Naddell's terrible encryption program"
	app.Copyright = "©️ Liam Naddell github.com/liamnaddell/cipher"
	app.Commands = []cli.Command{
		{
			//we need the choice of reading the message from a file or not
			//and we need the option of the keyfile
			Name:    "encode",
			Aliases: []string{"e", "-e"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "m, mfile",
					Usage:       "specify a file where the message is contained",
					Destination: &message,
				},
				cli.StringFlag{
					Name:        "k, keyfile",
					Usage:       "specify a file where the key is contained, otherwise be ready to break out the copy paste button",
					Destination: &keyfile,
				},
			},
			Category: "Text Options",
			Usage:    "encode a string",
			Action: func(c *cli.Context) error {
				return newEncode(message, keyfile)
			},
		},
		{
			//we need a keyfile option
			Name:    "genkey",
			Aliases: []string{"g", "-g"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "k, keyfile",
					Usage:       "specify a file, otherwise your tty will not be happy",
					Destination: &keyfile,
				},
			},
			Usage: "generate a key and store it in a file or print to stdout",
			Action: func(c *cli.Context) error {
				return newKey(keyfile)
			},
		},
		{
			//we need a file option for the message and a file option for the keyfile
			Name:    "decode",
			Aliases: []string{"d", "-d"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "m, mfile",
					Usage:       "specify a file where the message is stored",
					Destination: &message,
				},
				cli.StringFlag{
					Name:        "k, keyfile",
					Usage:       "specify a file where the key is stored, mandatory",
					Destination: &keyfile,
				},
			},
			Category: "Text Options",
			Usage:    "decode a encrypted string",
			Action: func(c *cli.Context) error {
				return newDecode(message, keyfile)
			},
		},
	}

	app.Run(os.Args)
}
