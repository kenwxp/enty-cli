package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"sort"
)

func main() {
	//var localhost bool
	//db, err := storage.NewDatabase(localhost)
	//if err != nil {
	//	fmt.Println("err:", err)
	//	panic("db failed init")
	//}
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add order",
				Action: func(c *cli.Context) error {
					if c.NArg() <= 0 {
						println("amount required")
						return nil
					}
					amount := c.Args().Get(0)
					println("add order " + amount + " nanofil")
					return nil
				},
			},
			{
				Name:    "withdraw",
				Aliases: []string{"w"},
				Usage:   "withdraw fil (nanofil)",
				Action: func(c *cli.Context) error {
					if c.NArg() <= 0 {
						println("amount required")
						return nil
					}
					amount := c.Args().Get(0)
					println("withdraw  " + amount + " nanofil")
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list income",
				Action: func(c *cli.Context) error {
					println("list income ")
					return nil
				},
			},
		},
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("err", err)
	}
}
