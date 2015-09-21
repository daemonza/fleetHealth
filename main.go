package main

// TODO :
// * add graceful shutdown

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
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
				unitDown := color.YellowString(unit.Name)
				unitState := color.RedString(unit.SystemdSubState)

				fmt.Println(unitDown + " : " + unitState)
			}
		}
	}

}

func main() {
	for {
		// Connect to fleet api and get unit list state
		healthCheck()

		interval, err := strconv.Atoi(os.Getenv("FH_CHECK_INTERVAL"))
		// if FH_CHECK_INTERVAL is not set default to 5 secs
		if err != nil {
			interval = 5
		}

		time.Sleep(time.Second * time.Duration(interval))
	}
}
