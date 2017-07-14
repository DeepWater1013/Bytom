package commands

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	tmflags "github.com/node_p2p/node_p2p/commands/flags"
	cfg "github.com/node_p2p/config"
	"github.com/tendermint/tmlibs/log"
)

var (
	config = cfg.DefaultConfig()
	logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "main")
)

func init() {
	RootCmd.PersistentFlags().String("log_level", config.LogLevel, "Log level")
}

var RootCmd = &cobra.Command{
	Use:   "node_p2p",
	Short: "node_p2p in Go",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := viper.Unmarshal(config)
		if err != nil {
			return err
		}
		config.SetRoot(config.RootDir)
		cfg.EnsureRoot(config.RootDir)
		logger, err = tmflags.ParseLogLevel(config.LogLevel, logger)
		if err != nil {
			return err
		}
		return nil
	},
}
