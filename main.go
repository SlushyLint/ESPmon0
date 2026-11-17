package main

import (
	"fmt"
	"time"

	"tinygo.org/x/espradio"
)

const (
	// Wi-Fi channels are listened to one at a time; change this for your lab AP.
	channel       = uint8(6)
	sampleWindow  = time.Second
	retryInterval = time.Second
)

func main() {
	// give a serial monitor time to connect after the board resets.
	time.Sleep(time.Second)
	fmt.Println("ESPmon0,boot")

	// enable initializes the radio; Start brings up the Wi-Fi driver.
	if err := espradio.Enable(espradio.Config{}); err != nil { // enable radio
		fmt.Println("ESPmon0,error,enable,", err)
		return

	}

	fmt.Printf("ESPmon0,ready,channel=%d\n", channel)
	if err := espradio.Start(); err != nil { // start radio
		fmt.Println("ESPmon0,error,start,", err)
		return
	}

	// this tool's job, temporarily, is to list APs and some extra information about them
	// this  means that we will ignore `espradio.BeginMonitor` until we fetch more information about said access points

	//	if err := espradio.BeginMonitor(channel); err != nil {
	//		fmt.Printf("ESPmon,err,monitor,", err)
	//		return
	//	}
	fmt.Printf("\nESPmon,scanning,channel=%d\n", channel)

	//scan APs and print their results
	var i int
	for {
		time.Sleep(750 * time.Millisecond)
		ap, e := espradio.Scan()
		if e != nil {
			fmt.Printf("scan err: %v\n", e)
			return
		}
		// RSSI is signal strength.
		// we will also display that beside the AP, it could be potentially useful information
		// to who idk.
		// but sure

		for _, ap := range ap {
			fmt.Printf("ESSID:\t%-28q\tBSSID:\t%02X", ap.SSID, ap.BSSID[0])
			for _, i := range ap.BSSID[1:] {
				fmt.Printf(":%02X", i)
			}
			fmt.Printf("\tRSSI:\t%4d dBm\x1b[K\n", ap.RSSI)

		}
		if i++; i == 15 {
			i = 0
			fmt.Printf("\x1b[2J\x1b[H")

		}
	}
}
