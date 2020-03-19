package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tangxusc/envoy-authz/pkg/config"
	"github.com/tangxusc/envoy-authz/pkg/server"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	newCommand := NewCommand(ctx)
	HandlerNotify(cancel)

	_ = newCommand.Execute()
	cancel()
}

func NewCommand(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:   "start",
		Short: "start authz",
		RunE: func(cmd *cobra.Command, args []string) error {
			rand.Seed(time.Now().Unix())
			config.InitLog()

			fmt.Println(config.Instance.Debug, config.Instance.Allow)
			go func() {
				http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
					writer.Write([]byte("headers:\n"))
					for key, values := range request.Header {
						writer.Write([]byte(fmt.Sprintf("%v:%+v \n", key, values)))
					}
					//writer.Header().Add("Access-Control-Allow-Methods", "GET,POST,DELETE,PATCH,OPTIONS,PUT")
					//writer.Header().Add("Access-Control-Allow-Origin", "*")
					//writer.Header().Add("Access-Control-Max-Age", "3600")
					//writer.Write([]byte("OK"))
				})
				http.ListenAndServe(":8080", nil)
			}()
			err := server.StartServer()
			if err != nil {
				return err
			}
			<-ctx.Done()
			return nil
		},
	}
	logrus.SetFormatter(&logrus.TextFormatter{})
	config.BindParameter(command)

	return command
}

func HandlerNotify(cancel context.CancelFunc) {
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, os.Kill)
		<-signals
		cancel()
	}()
}
