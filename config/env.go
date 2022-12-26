package config

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
)

type env struct {
	Address      string
	RemoteAPI    string
	BdConnection string
}

var Env env

func CheckFlagEnv() {
	var address string
	var remoteAPI string
	var dbConnection string

	_ = godotenv.Load()

	if os.Getenv("RUN_ADDRESS") != "" {
		address = os.Getenv("RUN_ADDRESS")
	} else {
		address = "localhost:8080"
	}

	if os.Getenv("ACCRUAL_SYSTEM_ADDRESS") != "" {
		remoteAPI = os.Getenv("ACCRUAL_SYSTEM_ADDRESS")

	} else {
		remoteAPI = ""
	}

	if os.Getenv("DATABASE_URI") != "" {
		dbConnection = os.Getenv("DATABASE_URI")

	} else {
		dbConnection = ""
	}

	var flagAddress = flag.String("a", "", "Server name")
	var flagRemoteAPI = flag.String("r", "", "Remote api")
	var flagDSN = flag.String("d", "", "Base dsn connection")

	flag.Parse()

	if *flagAddress != "" {
		address = *flagAddress
	}

	if *flagRemoteAPI != "" {
		remoteAPI = *flagRemoteAPI
	}

	if *flagDSN != "" {

		dbConnection = *flagDSN
	}

	Env = env{
		Address:      address,
		RemoteAPI:    remoteAPI,
		BdConnection: dbConnection,
	}
}
