package gosimplecobra

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// DefaultInitConfig 默认的初始化Viper配置
func DefaultInitConfigFunc() {
	// todo: 读取环境变量

	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")

		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read home dir(%s): %v\n", home, err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home + "." + "simplecobra")
		viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", cfgFile, err)
		os.Exit(1)
	}
}
