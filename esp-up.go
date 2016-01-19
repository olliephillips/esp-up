// Utility to connect ESP8266 running Espruino to Wifi and save settings using wifi.save()
// Ollie Phillips 2016
// MIT License

package main

import (
	"bufio"
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"github.com/spf13/cobra"
	"log"
	"os"
	//"strings"
	"bytes"
)

var (
	serialPort string
	username   string
	password   string
)

func configureWifi() {

	options := serial.OpenOptions{
		PortName:        serialPort,
		BaudRate:        115200,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 1,
	}

	s, err := serial.Open(options)
	if err != nil {
		log.Println("No device connected at", serialPort)
		os.Exit(0)
	}
	// Clean up
	defer s.Close()

	// Send script
	var configScript string = `
var wifi = require("Wifi");	
wifi.stopAP();
wifi.connect("%s", {password: "%s"}, function(){
	console.log("Connected");
	console.log(wifi.getIP().ip);
});
wifi.save();
`
	send := fmt.Sprintf(configScript, username, password)
	_, err = s.Write([]byte(send))
	if err != nil {
		log.Println("Unable to write to connected device")
		os.Exit(0)
	}

	// Handle terminal output
	reader := bufio.NewReader(s)
	readCount := 0
	go func() {
		// Read buffer to terminal
		for {
			response, _ := reader.ReadBytes('\n')
			responseString := stripBytes(response)
			if readCount == 1 {
				log.Println("IP address:", responseString)
				os.Exit(0)
			}
			if responseString == "Connected" {
				log.Println("Successfully connected to Wifi")
				readCount++
			}
		}
	}()
	select {}
}

func stripBytes(buf []byte) string {
	output := bytes.Replace(buf, []byte{62}, []byte{}, -1)
	output = bytes.Replace(output,[]byte{8}, []byte{}, -1)
	output = bytes.Replace(output, []byte{13}, []byte{}, -1)
	output = bytes.Replace(output, []byte{10}, []byte{}, -1)
	return string(output)
}

func main() {
	cmd := &cobra.Command{
		Use:   "esp-up",
		Short: "Connect your Espruino ESP8266 to Wifi instantly",
		Run: func(cmd *cobra.Command, args []string) {
			configureWifi()
		},
	}
	cmd.Flags().StringVarP(&serialPort, "serialport", "s", "/dev/tty.SLAB_USBtoUART", "Serial port to communicate over")
	cmd.Flags().StringVarP(&username, "username", "u", "ssid", "SSID of the Wifi network to connect to")
	cmd.Flags().StringVarP(&password, "password", "p", "password", "Password of the Wifi network to connect to")
	cmd.Execute()
}
