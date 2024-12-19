package gosimplecobra

import (
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	useCommandFlags = "[command] [flags]"
	useFlagsArgs    = "[flags] [args]"
)

// SimpleCommand 简易命令行接口, 用以处理带有标志的命令行
// 需要实现 Commander 和 Flags 接口
type SimpleCommander interface {
	Commander
	Flags
	SimpleCommands() []SimpleCommander
}

// commander 用以实现 不包含标志 的命令行
type Commander interface {
	Use() string

	ShortAndLong() (string, string)

	PreRun(args []string) error

	Run(args []string) error

	Commanders() []Commander
}

// Flags 实现标志的接口
// 注意：你应该返回各函数返回值的空值，而不是 nil
type Flags interface {
	// PersistentFlags 持久化的标志
	PersistentFlagsAndRequired() (fs *pflag.FlagSet, required []string)

	// LocalFlags 本地的标志
	LocalFlagsAndRequired() (fs *pflag.FlagSet, required []string)
}

type RootCommand struct {
	AppName string
	Short   string
	Long    string

	Version string
	// todo: 添加自定义的help
	Help string
	// todo: 启动默认的config && 绑定到子命令的标志
	EnableConf bool

	FlagSet Flags
	Args    cobra.PositionalArgs

	Initialize []func()
	PreRunFunc func(ctx context.Context, args []string) error
	RunFunc    func(ctx context.Context, args []string) error

	// SimpleCommand 这是一个实现了 SimpleCommander 接口的集合
	SimpleCommander []SimpleCommander
	// Commander 这是一个实现了 Commander 接口的集合
	Commander []Commander

	rootCobra *cobra.Command
}

func NewRootCmd(appName string, opts ...RootOption) *Executor {
	rootCmd := &RootCommand{
		AppName: appName,
	}

	for _, o := range opts {
		o(rootCmd)
	}

	rootCmd.rootCobra = buildcobra(rootCmd, rootCmd.FlagSet)
	rootCmd.rootCobra.Args = rootCmd.Args

	if rootCmd.Version != "" {
		rootCmd.rootCobra.Version = rootCmd.Version
	}

	if rootCmd.EnableConf {
		rootCmd.loadConfig()
	}

	if rootCmd.Initialize != nil {
		cobra.OnInitialize(rootCmd.Initialize...)
	}

	if len(rootCmd.SimpleCommander) != 0 {
		rootCmd.buildSimpleCommander(rootCmd.rootCobra)
	}

	if len(rootCmd.Commander) != 0 {
		rootCmd.buildCommander(rootCmd.rootCobra)
	}

	return &Executor{exec: rootCmd.rootCobra}
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

func WithVersion(version string) RootOption {
	return func(r *RootCommand) {
		r.Version = version
	}
}

func WithConfig(enable bool) RootOption {
	return func(r *RootCommand) {
		r.EnableConf = enable
	}
}

func WithFlagSets(flagSets Flags) RootOption {
	return func(r *RootCommand) {
		r.FlagSet = flagSets
	}
}

func WithArgs(args cobra.PositionalArgs) RootOption {
	return func(r *RootCommand) {
		r.Args = args
	}
}

func WithSimpleCommand(simpleCommand []SimpleCommander) RootOption {
	return func(r *RootCommand) {
		r.SimpleCommander = simpleCommand
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

func WithPreRunFunc(preRunF func(ctx context.Context, args []string) error) RootOption {
	return func(r *RootCommand) {
		r.PreRunFunc = preRunF
	}
}

func WithRunFunc(runF func(ctx context.Context, args []string) error) RootOption {
	return func(r *RootCommand) {
		r.RunFunc = runF
	}
}

func (rc *RootCommand) loadConfig() {
	rc.rootCobra.Flags().StringVarP(&cfgFile, "config", "c", ".simplecobra.yaml", "config file (default is $HOME/.simplecobra/config.yaml)")
	cobra.OnInitialize(DefaultInitConfigFunc)
}

func (rc *RootCommand) buildSimpleCommander(rootCobra *cobra.Command) {
	for _, simpleCmd := range rc.SimpleCommander {
		simpleCmdBuilder := &commandBuilder{
			SimpleCommander: simpleCmd,
		}
		simpleCmdBuilder.buildCobraInSimpleCommander()
		rootCobra.AddCommand(simpleCmdBuilder.CobraCommand)
	}
}

func (rc *RootCommand) buildCommander(rootCobra *cobra.Command) {

	for _, cmder := range rc.Commander {
		simpleCmd := &commandBuilder{
			Commander: cmder,
		}
		simpleCmd.buildCobraInCommander()
		rootCobra.AddCommand(simpleCmd.CobraCommand)
	}
}

func (rc *RootCommand) Use() string {
	return rc.AppName
}

func (rc *RootCommand) ShortAndLong() (string, string) {
	return rc.Short, rc.Long
}

func (rc *RootCommand) PreRun(args []string) error {
	if rc.PreRunFunc == nil {
		return nil
	}
	return rc.PreRunFunc(context.Background(), args)
}

func (rc *RootCommand) Run(args []string) error {
	if rc.RunFunc == nil {
		return nil
	}

	if rc.EnableConf {
		if err := viper.BindPFlags(rc.rootCobra.PersistentFlags()); err != nil {
			return err
		}

		if err := viper.BindPFlags(rc.rootCobra.Flags()); err != nil {
			return err
		}

		if err := viper.Unmarshal(rc.FlagSet); err != nil {
			return err
		}
		printConfig()
	}

	return rc.RunFunc(context.Background(), args)
}

func (rc *RootCommand) Commanders() []Commander {
	return rc.Commander
}

func (rc *RootCommand) PersistentFlagsAndRequired() (fs *pflag.FlagSet, required []string) {

	return rc.FlagSet.PersistentFlagsAndRequired()
}

func (rc *RootCommand) LocalFlagsAndRequired() (fs *pflag.FlagSet, required []string) {

	return fs, required
}

type commandBuilder struct {
	SimpleCommander SimpleCommander
	Commander       Commander
	FlagSet         Flags

	CobraCommand *cobra.Command
}

func (cb *commandBuilder) buildCobraInSimpleCommander() {
	if cb.CobraCommand != nil {
		return
	}

	cb.CobraCommand = cb.buildcobra()
	if cb.SimpleCommander.SimpleCommands() != nil {
		for _, simpleCmd := range cb.SimpleCommander.SimpleCommands() {
			subBuilder := &commandBuilder{
				Commander: simpleCmd,
				FlagSet:   simpleCmd,
			}
			subBuilder.buildCobraInSimpleCommander()
			cb.CobraCommand.AddCommand(subBuilder.CobraCommand)
		}
	}

}

func (cb *commandBuilder) buildCobraInCommander() {
	if cb.CobraCommand != nil {
		return
	}

	cb.CobraCommand = cb.buildcobra()
	if cb.Commander.Commanders() != nil {
		for _, cmder := range cb.Commander.Commanders() {
			subBuilder := &commandBuilder{
				Commander: cmder,
			}
			subBuilder.buildCobraInCommander()
			cb.CobraCommand.AddCommand(subBuilder.CobraCommand)
		}
	}
}

func (cb *commandBuilder) buildcobra() *cobra.Command {
	if cb.SimpleCommander != nil {
		cb.Commander = cb.SimpleCommander
		cb.FlagSet = cb.SimpleCommander
	}

	return buildcobra(cb.Commander, cb.FlagSet)
}

func buildcobra(cmder Commander, fs Flags) *cobra.Command {
	short, long := cmder.ShortAndLong()
	cobraCmd := &cobra.Command{
		Use:   use(cmder),
		Short: short,
		Long:  long,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return cmder.PreRun(args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmder.Run(args)
		},
		SilenceErrors:              true,
		SilenceUsage:               true,
		SuggestionsMinimumDistance: 2,
	}

	applyFlags(cobraCmd, fs)

	return cobraCmd
}

func applyFlags(cmd *cobra.Command, f Flags) {
	if cmd == nil || f == nil {
		return
	}

	pfs, reqf := f.PersistentFlagsAndRequired()
	cmd.PersistentFlags().AddFlagSet(pfs)

	for _, rf := range reqf {
		_ = cmd.MarkPersistentFlagRequired(rf)
	}

	lfs, reqf := f.LocalFlagsAndRequired()
	cmd.Flags().AddFlagSet(lfs)

	for _, rf := range reqf {
		_ = cmd.MarkFlagRequired(rf)
	}
}

func use(cmder Commander) string {
	var line = useCommandFlags
	if cmder.Commanders() == nil || len(cmder.Commanders()) == 0 {
		line = useFlagsArgs
	}
	return fmt.Sprintf("%s %s", cmder.Use(), line)
}

// printConfig 打印配置
// copy from https://github.com/marmotedu/iam/blob/master/pkg/app/config.go
func printConfig() {
	if keys := viper.AllKeys(); len(keys) > 0 {
		fmt.Printf("%v Configuration items:\n", color.GreenString("==>"))
		table := uitable.New()
		table.Separator = " "
		table.MaxColWidth = 80
		table.RightAlign(0)
		for _, k := range keys {
			table.AddRow(fmt.Sprintf("%s:", k), viper.Get(k))
		}
		fmt.Printf("%v\n", table)
	}
}

type Executor struct {
	exec *cobra.Command
}

func (e *Executor) Execute() error {
	return e.exec.Execute()
}
