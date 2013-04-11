package main

import(
	"network"
	"time"
)

func main() {
	network.Channelinit()

	network.UDPinit()
	network.TCPinit()
	
	for {
		time.Sleep(time.Millisecond)
	}
}
