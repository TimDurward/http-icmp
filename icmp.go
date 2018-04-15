package main

import (
	"fmt"
	"net"
	"time"

	"github.com/sparrc/go-ping"
)

// Ping represents the Result of an ICMP response
type Ping struct {
	// Address is the address of the host being pinged.
	Address string `json:"address"`

	// IPAddress is the address of the host being pinged.
	IPAddress *net.IPAddr `json:"ip_address"`

	// Results represents the nested structure which includes Packets & Statistics.
	Results Results `json:"results"`
}

// Results represents the data that is returned by ICMP.
type Results struct {
	// Packets represents the overall response for Packet data.
	Packets Packets `json:"packets"`

	// Statistics represents the overall response for statistic calculations.
	Statistics Statistics `json:"statistics"`
}

// Packets represents a received and processed ICMP echo packet.
type Packets struct {
	// Sent is the number of packets sent.
	Sent int `json:"sent"`

	// Received is the number of packets received.
	Received int `json:"received"`

	// Packet Loss is the percentage of packets lost.
	Loss float64 `json:"lost"`
}

// Statistics represent the stats of a finished ICMP response.
type Statistics struct {
	// Minimum is the minimum round-trip time sent via ICMP.
	Minimum time.Duration `json:"minimum"`
	// Maximum is the maximum round-trip time sent via ICMP.
	Maximum time.Duration `json:"maximum"`
	// Average is the average round-trip time sent via ICMP.
	Average time.Duration `json:"average"`
}

func icmp(host string, count int) Ping {
	var icmpResp Ping

	p, err := ping.NewPinger(host)
	if err != nil {
		fmt.Printf("[ Ping Service ] : %s\n", err.Error())
	}

	p.Count = count
	p.Interval = time.Second
	p.Timeout = time.Second * 100000
	p.SetPrivileged(false)

	p.OnFinish = func(stats *ping.Statistics) {

		resp := Ping{
			Address:   stats.Addr,
			IPAddress: stats.IPAddr,
			Results: Results{
				Statistics: Statistics{
					Minimum: stats.MinRtt,
					Maximum: stats.MaxRtt,
					Average: stats.AvgRtt,
				},
				Packets: Packets{
					Sent:     stats.PacketsSent,
					Received: stats.PacketsRecv,
					Loss:     stats.PacketLoss,
				},
			},
		}

		icmpResp = resp
	}

	p.Run()
	return icmpResp

}
