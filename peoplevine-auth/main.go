package main

import (
	"os"
	"peoplevine-auth/cmdargs"
	"peoplevine-auth/logger"
	"peoplevine-auth/rest"
)

// ********************************************************
func RestServer() (err error) {
	var (
		MainRestServer rest.RestServer
	)
	MainRestServer.RestServerConf = rest.RestServerConf
	err = MainRestServer.Start()
	if err != nil {
		logger.Logger.Errorln("Error on run REST server EXIT...")
		os.Exit(127)
	}
	return err
}

// ********************************************************
func main() {

	// Parse argv
	cmdargs.Parse(&cmdargs.Args)

	// Set debug level
	if cmdargs.Args.Loggig != "" {
		logger.InitLogger(cmdargs.Args.Loggig, "stdout")
	} else {
		logger.InitLogger("DEBUG", "stdout")
	}

	// Set listen addr:port
	if cmdargs.Args.Listen != "" {
		rest.RestServerConf.Listen = cmdargs.Args.Listen
	}

	_ = RestServer()
}
