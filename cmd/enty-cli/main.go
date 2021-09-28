package main

import (
	"entysquare/enty-cli/service"
	"entysquare/enty-cli/storage"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"sort"
)

func main() {
	db, err := storage.NewDatabase()
	if err != nil {
		fmt.Println("err:", err)
		panic("db failed init")
	}
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
					filerName := c.Args().Get(0)
					amount := c.Args().Get(1)
					println("add order " + amount + " nanofil for" + filerName)
					err := service.InsertOrder(db, filerName, amount)
					if err != nil {
						println("fail to add order")
						return err
					}
					println("add order success")
					return nil
				},
			},
			{
				Name:    "withdraw",
				Aliases: []string{"w"},
				Usage:   "withdraw fil ",
				Action: func(c *cli.Context) error {
					if c.NArg() <= 0 {
						println("amount required")
						return nil
					}
					filerName := c.Args().Get(0)
					amount := c.Args().Get(1)
					err := service.Withdraw(db, filerName, amount)
					if err != nil {
						println("fail to withdraw")
						return err
					}
					println("withdraw success")
					return nil
				},
			},
			{
				Name:    "deposit",
				Aliases: []string{"d"},
				Usage:   "deposit fil",
				Action: func(c *cli.Context) error {
					if c.NArg() <= 0 {
						println("amount required")
						return nil
					}
					filerName := c.Args().Get(0)
					amount := c.Args().Get(1)
					err := service.Deposit(db, filerName, amount)
					if err != nil {
						println("fail to deposit")
						return err
					}
					println("deposit success")
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list income",
				Action: func(c *cli.Context) error {
					err := service.QueryIncomeList(db)
					if err != nil {
						println("fail to get income list")
						return err
					}
					return nil
				},
			},
		},
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	err = app.Run(os.Args)
	if err != nil {
		fmt.Println("err", err)
	}
}
