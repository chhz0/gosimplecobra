package main

import (
	"fmt"

	"github.com/chhz0/gosimplecobra"
	"github.com/spf13/pflag"
)

type RootOption struct {
	AppName string       `mapstructure:"app"`
	Server  string       `mapstructure:"server"`
	Print   *PrintOption `mapstructure:"printOption"`
}

func newRootOption() *RootOption {
	return &RootOption{}
}

// LocalFlags implements gosimplecobra.Flags.
func (r *RootOption) LocalFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	fs = pflag.NewFlagSet("RootL", pflag.ExitOnError)
	fs.StringVarP(&r.Server, "server", "s", "", "server address")

	return
}

// PersistentFlags implements gosimplecobra.Flags.
func (r *RootOption) PersistentFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	fs = pflag.NewFlagSet("RootP", pflag.ExitOnError)
	fs.StringVarP(&r.AppName, "app", "a", "go-simplecobra", "app name for the application")

	return
}

var _ gosimplecobra.Flags = (*RootOption)(nil)

type PrintOption struct {
	print string `mapstructure:"print"`
	from  string `mapstructure:"from"`
}

func newPrintCmd() gosimplecobra.SimpleCommander {
	return &PrintOption{}
}

// SimpleCommands implements gosimplecobra.SimpleCommand.
func (p *PrintOption) SimpleCommands() []gosimplecobra.SimpleCommander {
	return nil
}

// LocalFlagsAndRequired implements gosimplecobra.SimpleCommand.
func (p *PrintOption) LocalFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	fs = pflag.NewFlagSet("PrintL", pflag.ExitOnError)
	fs.StringVarP(&p.print, "print", "p", "", "print anything to the screen")
	fs.StringVarP(&p.from, "from", "f", "", "print from where")

	required = []string{"print"}

	return
}

// PersistentFlagsAndRequired implements gosimplecobra.SimpleCommand.
func (p *PrintOption) PersistentFlagsAndRequired() (fs *pflag.FlagSet, required []string) {

	return
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
func (p *PrintOption) PreRun(args []string) error {
	fmt.Println("print cmd pre run")

	return nil
}

// Run implements gosimplecobra.Commander.
func (p *PrintOption) Run(args []string) error {
	fmt.Println("print cmd run")

	return nil
}

// Use implements gosimplecobra.Commander.
func (p *PrintOption) Use() string {
	return "print"
}

var _ gosimplecobra.SimpleCommander = (*PrintOption)(nil)

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
func (e *EchoOption) PreRun(args []string) error {
	fmt.Println("echo cmd pre run")
	return nil
}

// Run implements gosimplecobra.Commander.
func (e *EchoOption) Run(args []string) error {
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
func (t *TimesOption) PreRun(args []string) error {
	fmt.Println("times cmd pre run")
	return nil
}

// Run implements gosimplecobra.Commander.
func (t *TimesOption) Run(args []string) error {
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
