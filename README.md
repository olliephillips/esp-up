# esp-up
Simple Wifi configuration utility for ESP8266 devices running Espruino.
Returns IP address of device once connected, ideal for going on to connect with the Web IDE using the telnet option.

Quick dirty and experimental.

## Setup
Written in Go. Assuming you have a go environment, use ```go get``` to obtain the package code.

```
go get github.com/olliephillips/esp-up
```

Then build the package to binary and optionally install it.

## About
This is a proof of concept for myself, the objective being to make initial setup of ESP8266 (running Espruino) as simple as possible. It negates the need for the Web IDE in the first instance, and abstracts away the need write, and indeed understand, code, in order to configure the ESP8266.

## Roadmap
It's a very small leap to deploy a working program(sketch) to the ESP8266 using this same approach. This might come in handy.
Also a cross platform tool that could flash ESP8266 with Espruino, setup the Wifi and deploy a javascript program, would be very cool indeed (in my opinion).
