package server

import (
	"assignment/routes"
	"flag"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func StartServer() {
	writeToFile := flag.Bool("f", false, "write logs to file")
	loglevel := flag.String("l", "ERROR", "logging verbosity")
	flag.Parse()

	//set log level
	level, err := logrus.ParseLevel(*loglevel)
	if err != nil {
		level = logrus.ErrorLevel
	}
	logrus.SetLevel(level)
	logrus.SetReportCaller(true)

	// create log file if enabled
	if *writeToFile {
		f, err := os.OpenFile("/var/log/assignment.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			panic(err)
		}
		logrus.SetOutput(f)
	}

	//initialize API routes
	r := routes.Routes()

	// start server
	listenAddr := "0.0.0.0:3339"
	logrus.Infof("starting server at: %s", listenAddr)
	logrus.Error(http.ListenAndServe(listenAddr, r))
}
