package gosimplecobra

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type Commander interface {
	Use() string

	ShortAndLong() (string, string)

	PreRun() error

	Run() error

	Commanders() []Commander
}

type Flags interface {
	// PersistentFlags 持久化的标志
	PersistentFlags(*pflag.FlagSet)

	// LocalFlags 本地的标志
	LocalFlags(*pflag.FlagSet)

	RequiredFlags(*pflag.FlagSet) []string
}

type RootCommand struct {
	AppName string
	Short   string
	Long    string

	Version string
	Help    string

	FlagSet *pflag.FlagSet
	Args    cobra.PositionalArgs

	Initialize []func()
	PreRunFunc func(ctx context.Context) error
	RunFunc    func(ctx context.Context, args []string) error

	Commander []Commander
	RootCobra *cobra.Command

	*DefaultFlags
}

func NewRootCmd(appName string, opts ...RootOption) *Executor {
	rootCmd := &RootCommand{
		AppName: appName,
	}

	for _, o := range opts {
		o(rootCmd)
	}

	rootCmd.RootCobra = &cobra.Command{
		Use:   rootCmd.Use(),
		Short: rootCmd.Short,
		Long:  rootCmd.Long,

		Args: rootCmd.Args,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return rootCmd.PreRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return rootCmd.Run()
		},
	}

	if rootCmd.Initialize != nil {
		cobra.OnInitialize(rootCmd.Initialize...)
	}

	rootCmd.buildCommander()

	return &Executor{exec: rootCmd.RootCobra}
}

type RootOption func(r *RootCommand)

func WithRootShort(short string) RootOption {
	return func(r *RootCommand) {
		r.Short = short
	}
}

func WithRootLong(long string) RootOption {
	return func(r *RootCommand) {
		r.Long = long
	}
}

func WithFlagSets(flagSets *pflag.FlagSet) RootOption {
	return func(r *RootCommand) {
		r.FlagSet = flagSets
	}
}

func WithArgs(args cobra.PositionalArgs) RootOption {
	return func(r *RootCommand) {
		r.Args = args
	}
}

func WithCommander(commanders []Commander) RootOption {
	return func(r *RootCommand) {
		r.Commander = commanders
	}
}

func WithInitialize(initF ...func()) RootOption {
	return func(r *RootCommand) {
		r.Initialize = initF
	}
}

func WithPreRunFunc(preRunF func(ctx context.Context) error) RootOption {
	return func(r *RootCommand) {
		r.PreRunFunc = preRunF
	}
}

func WithRunFunc(runF func(ctx context.Context, args []string) error) RootOption {
	return func(r *RootCommand) {
		r.RunFunc = runF
	}
}

// 构建commandBuilder
func (rc *RootCommand) buildCommander() {
	for _, cmder := range rc.Commander {
		simpleCmd := &commandBuilder{
			Command: cmder,
		}
		simpleCmd.buildCobra()
		rc.RootCobra.AddCommand(simpleCmd.CobraCommand)
	}
}

func (rc *RootCommand) Use() string {
	var use = fmt.Sprintf("%s [flags] [args]", rc.AppName)
	if rc.Commander != nil {
		use = fmt.Sprintf("%s [command] [flags]", rc.AppName)
	}
	return use
}

func (rc *RootCommand) PreRun() error {
	if rc.PreRunFunc == nil {
		return nil
	}
	return rc.PreRunFunc(context.Background())
}

func (rc *RootCommand) Run() error {
	if rc.RunFunc == nil {
		return nil
	}
	return rc.RunFunc(context.Background(), nil)
}

func (rc *RootCommand) Commanders() []Commander {
	return rc.Commander
}

type commandBuilder struct {
	Command Commander

	CobraCommand *cobra.Command
}

func (cb *commandBuilder) buildCobra() {
	if cb.CobraCommand != nil {
		return
	}
	short, long := cb.Command.ShortAndLong()
	cb.CobraCommand = &cobra.Command{
		Use:   cb.Command.Use(),
		Short: short,
		Long:  long,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return cb.Command.PreRun()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cb.Command.Run()
		},
		SilenceErrors:              true,
		SilenceUsage:               true,
		SuggestionsMinimumDistance: 2,
	}

	if cb.Command.Commanders() != nil {
		for _, cmder := range cb.Command.Commanders() {
			subBuilder := &commandBuilder{
				Command: cmder,
			}
			subBuilder.buildCobra()
			cb.CobraCommand.AddCommand(subBuilder.CobraCommand)
		}
	}
}

type Executor struct {
	exec *cobra.Command
}

func (e *Executor) Execute() error {
	return e.exec.Execute()
}
