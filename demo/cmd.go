package main

import (
	"context"
	"fmt"

	"github.com/chhz0/gosimplecobra"
)

func initCli() *gosimplecobra.Executor {
	return gosimplecobra.NewRootCmd("gosimplecobra",
		gosimplecobra.WithRootShort("gosimplecobra short desc"),
		gosimplecobra.WithRootLong("this is a long description for gosimplecobra"),
		gosimplecobra.WithVersion("0.0.1 Snapshot"),
		gosimplecobra.WithFlagSets(newRootOption()),
		gosimplecobra.WithInitialize(func() {
			fmt.Println("gosimplecobra init func")
		}),
		gosimplecobra.WithPreRunFunc(func(ctx context.Context, args []string) error {
			fmt.Println("gosimplecobra prerun func")
			return nil
		}),
		gosimplecobra.WithRunFunc(func(ctx context.Context, args []string) error {
			fmt.Println("gosimplecobra run func")
			return nil
		}),
		gosimplecobra.WithSimpleCommand([]gosimplecobra.SimpleCommander{
			newPrintCmd(),
		}),
		gosimplecobra.WithCommander([]gosimplecobra.Commander{
			newEchoCmd(),
		}),
	)
}

func main() {
	if err := initCli().Execute(); err != nil {
		panic(err)
	}
}
