package main

import "github.com/urfave/cli"
import "github.com/liamnaddell/cipher"
import "os"

//import "errors"

func startCli() {
	var message string
	var keyfile string
	app := cli.NewApp()
	app.Name = "cipher"
	app.Version = Version
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
				var err error
				if message == "" {
					if len(c.Args()) < 1 {
						return cli.NewExitError("No message given", 1)
					}
					message = c.Args()[0]
					err := cipher.NewEncode(message, keyfile)
					if err != nil {
						return cli.NewExitError(err, 1)
					}
				} else {
					err = cipher.NewFileEncode(message, keyfile)
					if err != nil {
						return cli.NewExitError(err, 1)
					}
				}
				return nil
			},
		},
		{
			//we need a keyfile option
			Name:    "genkey",
			Aliases: []string{"g", "-g"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "k, keyfile",
					Usage:       "specify a file, or else",
					Destination: &keyfile,
				},
			},
			Usage: "generate a key and store it in a file or print to stdout",
			Action: func(c *cli.Context) error {
				err := cipher.NewKey(keyfile)
				if err != nil {
					return cli.NewExitError(err, 1)
				}
				return nil
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
				if message == "" {
					if len(c.Args()) < 1 {
						return cli.NewExitError("No message to decode given", 1)
					}
					message = c.Args()[0]
					err := cipher.NewDecode(message, keyfile)
					if err != nil {
						return cli.NewExitError(err, 1)
					}
				} else {
					err := cipher.NewFileDecode(message, keyfile)
					if err != nil {
						return cli.NewExitError(err, 1)
					}
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
