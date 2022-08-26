package cmd

import (
	"context"
	"flag"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang/app/database"
	"golang/app/students"
	"golang/app/users"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var serveHTTPCmd = &cobra.Command{
	Use:   "serveHTTP",
	Short: "http",
	Long:  "Activate http",
	Run:   RunServeHTTP,
}

func init() {
	rootCmd.AddCommand(serveHTTPCmd)
}

func RunServeHTTP(*cobra.Command, []string) {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	//Init Database
	studentsDB := database.CreateConnection("mysql", viper.GetString("MYSQL_CONNECTION"))
	//Create New Session
	dbSes := studentsDB.NewSession(nil)

	//Init Validator
	v := validator.New()

	// Add your routes as needed
	router := mux.NewRouter().StrictSlash(false)
	students.InitStudents(router, dbSes, v)
	users.InitUsers(router, dbSes, v)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Error(err)
		}

	}()
	logrus.Info("Running On Port 8080")
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	logrus.Info("shutting down")
	os.Exit(0)
}
