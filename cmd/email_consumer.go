package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	nsqApp "golang/app/nsq"
	"golang/app/students/delevery/consumer"
	"golang/app/students/usecase"
)

var EmailConsumer = &cobra.Command{
	Use:   "emailConsumer",
	Short: "consumer",
	Long:  "Activate Email Consumer",
	Run:   RunEmailConsumer,
}

func init() {
	rootCmd.AddCommand(EmailConsumer)
}

func RunEmailConsumer(*cobra.Command, []string) {
	topic := viper.GetString("NSQ_TOPIC")
	channel := viper.GetString("NSQ_CHANNEL")
	NSQConsumer := nsqApp.InitConsumer(topic, channel)

	useca := usecase.NewEmailConsumerUsecase()
	consum := consumer.NewEmailConsumer(useca, NSQConsumer)
	consum.StartEmailConsumer(useca)
}
