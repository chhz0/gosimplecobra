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
		gosimplecobra.WithInitialize(func() {
			fmt.Println("gosimplecobra init func")
		}),
		gosimplecobra.WithPreRunFunc(func(ctx context.Context) error {
			fmt.Println("gosimplecobra prerun func")
			return nil
		}),
		gosimplecobra.WithRunFunc(func(ctx context.Context, args []string) error {
			fmt.Println("gosimplecobra run func")
			return nil
		}),
		gosimplecobra.WithCommander([]gosimplecobra.Commander{
			newPrintCmd(),
			newEchoCmd(),
		}),
	)
}

func main() {
	if err := initCli().Execute(); err != nil {
		panic(err)
	}
}
