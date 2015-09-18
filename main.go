package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

// FleetState is a struct for the fleet api json
type FleetState struct {
	States []struct {
		Hash               string `json:"hash"`
		MachineID          string `json:"machineID"`
		Name               string `json:"name"`
		SystemdActiveState string `json:"systemdActiveState"`
		SystemdLoadState   string `json:"systemdLoadState"`
		SystemdSubState    string `json:"systemdSubState"`
	} `json:"states"`
}

func healthCheck() {
	fmt.Println("== health check")

	fleetAPI := os.Getenv("FH_FLEET_API")
	if fleetAPI == "" {
		fleetAPI = "127.0.0.1"
	}
	url := "http://" + fleetAPI + "/fleet/v1/state"

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err.Error())
		}

		var data FleetState
		json.Unmarshal(body, &data)

		for _, unit := range data.States {
			if unit.SystemdSubState != "running" {
				fmt.Println(unit.Name + " : " + unit.SystemdSubState)
			}
		}
	}

}

func main() {
	for {
		// Connect to fleet api and get unit list state
		healthCheck()

		// sleep
		interval, err := strconv.Atoi(os.Getenv("FH_CHECK_INTERVAL"))
		// if FH_CHECK_INTERVAL is not set default to 5 secs
		if err != nil {
			interval = 5
		}

		time.Sleep(time.Second * time.Duration(interval))
	}
}
