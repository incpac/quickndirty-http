package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var contentPath string
var bindAddress string

var Version string

func serve(cmd *cobra.Command, args []string) {

	fs := http.FileServer(http.Dir(contentPath))
	http.Handle("/", fs)

	log.Printf("Listening on " + bindAddress + "...\n")
	http.ListenAndServe(bindAddress, logger(http.DefaultServeMux))
}

func logger(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf(
			"%s %s %s\n",
			r.RemoteAddr,
			r.Method,
			r.URL,
		)

		handler.ServeHTTP(w, r)
	})
}

func main() {

	var showVersion bool

	command := &cobra.Command{
		Use:   "qndhttp",
		Short: "A quick and dirty webserver",
		Long:  `A simple webserver for quickly serving static files.`,
		Run: func(cmd *cobra.Command, args []string) {
			if showVersion {
				fmt.Printf("%s\n", Version)
			} else {
				serve(cmd, args)
			}
		},
	}

	command.Flags().StringVarP(&contentPath, "content", "c", "./", "path to the content to serve")
	command.Flags().StringVarP(&bindAddress, "bind", "b", "127.0.0.1:3000", "the address to bind the server to")
	command.Flags().BoolVarP(&showVersion, "version", "v", false, "display the version")

	if err := command.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
