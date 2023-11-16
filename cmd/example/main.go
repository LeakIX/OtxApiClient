package main

import (
	"github.com/LeakIX/OtxApiClient"
	"log"
	"os"
)

func main() {
	client, err := OtxApiClient.NewClient(os.Getenv("OTX_KEY"))
	if err != nil {
		panic(err)
	}
	// Get current user
	user, err := client.GetUserService().GetMe()
	if err != nil {
		panic(err)
	}
	log.Println(user.Username)

	// Get Domain IOCs
	var domainId OtxApiClient.DomainGeneral
	err = client.GetIndicatorsService().GetIndicators(&domainId, "leakix.net", "general")
	if err != nil {
		panic(err)
	}
	for _, pulse := range domainId.PulseInfo.Pulses {
		log.Println(pulse.Name)
	}

	// Add IOC to Pulse
	err = client.GetPulsesService().AddIndicator("65565252c2a155709ed4fb53", OtxApiClient.PulseIndicator{
		Type:      OtxApiClient.PulseHostnameIndicator,
		Indicator: os.Args[1],
	})
	if err != nil {
		panic(err)
	}

}
