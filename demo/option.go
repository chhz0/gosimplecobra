package main

import (
	"fmt"

	"github.com/chhz0/gosimplecobra"
)

type PrintOption struct {
	print string
	from  string
}

func newPrintCmd() gosimplecobra.Commander {
	return &PrintOption{}
}

// Long implements gosimplecobra.Commander.
func (p *PrintOption) ShortAndLong() (string, string) {
	return "Print anything to the screen",
		`print is for printing anything back to the screen.
For many years people have printed back to the screen.`
}

// Commanders implements gosimplecobra.Commander.
func (p *PrintOption) Commanders() []gosimplecobra.Commander {
	return nil
}

// PreRun implements gosimplecobra.Commander.
func (p *PrintOption) PreRun() error {
	fmt.Println("print cmd pre run")

	return nil
}

// Run implements gosimplecobra.Commander.
func (p *PrintOption) Run() error {
	fmt.Println("print cmd run")

	return nil
}

// Use implements gosimplecobra.Commander.
func (p *PrintOption) Use() string {
	return "print"
}

var _ gosimplecobra.Commander = (*PrintOption)(nil)

type EchoOption struct {
}

func newEchoCmd() gosimplecobra.Commander {
	return &EchoOption{}
}

// ShortAndLong implements gosimplecobra.Commander.
func (e *EchoOption) ShortAndLong() (string, string) {
	return "Echo anything to the screen",
		`echo is for echoing anything back to the screen.
For many years people have echoed back to the screen.`
}

// Commanders implements gosimplecobra.Commander.
func (e *EchoOption) Commanders() []gosimplecobra.Commander {
	return []gosimplecobra.Commander{
		newTimesCmd(),
	}
}

// PreRun implements gosimplecobra.Commander.
func (e *EchoOption) PreRun() error {
	fmt.Println("echo cmd pre run")
	return nil
}

// Run implements gosimplecobra.Commander.
func (e *EchoOption) Run() error {
	fmt.Println("echo cmd run")
	return nil
}

// Use implements gosimplecobra.Commander.
func (e *EchoOption) Use() string {
	return "echo"
}

var _ gosimplecobra.Commander = (*EchoOption)(nil)

type TimesOption struct {
}

func newTimesCmd() gosimplecobra.Commander {
	return &TimesOption{}
}

// Commanders implements gosimplecobra.Commander.
func (t *TimesOption) Commanders() []gosimplecobra.Commander {
	return nil
}

// PreRun implements gosimplecobra.Commander.
func (t *TimesOption) PreRun() error {
	fmt.Println("times cmd pre run")
	return nil
}

// Run implements gosimplecobra.Commander.
func (t *TimesOption) Run() error {
	fmt.Println("times cmd run")
	return nil
}

// ShortAndLong implements gosimplecobra.Commander.
func (t *TimesOption) ShortAndLong() (string, string) {
	return "Echo anything to the screen more times",
		`echo things multiple times back to the user by providing
a count and a string.`
}

// Use implements gosimplecobra.Commander.
func (t *TimesOption) Use() string {
	return "times"
}

var _ gosimplecobra.Commander = (*TimesOption)(nil)
