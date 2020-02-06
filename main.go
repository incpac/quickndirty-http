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

var sslCertPath string
var sslPrivateKeyPath string

var logfile string

var Version string

func serve(cmd *cobra.Command, args []string) {

	fs := http.FileServer(http.Dir(contentPath))
	http.Handle("/", fs)

	if sslCertPath == "" {
		log.Printf("Listening on http://" + bindAddress + "...\n")
		http.ListenAndServe(bindAddress, logger(http.DefaultServeMux))
	} else {
		if sslPrivateKeyPath == "" {
			log.Fatalf("Missing SSL private key")
		} else {
			log.Printf("Listening on https://" + bindAddress + "...\n")
			http.ListenAndServeTLS(bindAddress, sslCertPath, sslPrivateKeyPath, logger(http.DefaultServeMux))
		}
	}
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
				if logfile != "" {
					f, _ := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
					defer f.Close()
					log.SetOutput(f)
				}

				serve(cmd, args)
			}
		},
	}

	command.Flags().StringVarP(&contentPath, "content", "c", "./", "path to the content to serve")
	command.Flags().StringVarP(&bindAddress, "bind", "b", "127.0.0.1:3000", "the address to bind the server to")
	command.Flags().BoolVarP(&showVersion, "version", "v", false, "display the version")
	command.Flags().StringVarP(&sslCertPath, "sslcert", "", "", "SSL certificate for HTTPS traffic")
	command.Flags().StringVarP(&sslPrivateKeyPath, "sslkey", "", "", "Private key for the SSL certificate")
	command.Flags().StringVarP(&logfile, "logfile", "", "", "File to write logs to")

	if err := command.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
