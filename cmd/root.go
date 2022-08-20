package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:     "golang",
	Short:   "golang service is a revamped service for gold",
	Example: "For running http services. You may run: `go run main.go serveHTTP`. For Cron jobs. Some command example: `go run main.go cron expired`",
}

func init() {
	cobra.OnInitialize(LoadConfig)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

// Execute func.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(LoadConfig)
}

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}
