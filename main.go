/*
 * Copyright (C) 2021 The "MysteriumNetwork/asymmetric-udp-echo" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

/* Receives datagram on receiver port, extracts Port and UUID and sends UUID
back to originating address on specified Port.
*/

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	PortFieldSize = 2
	UUIDSize      = 16
	PacketSize    = PortFieldSize + UUIDSize
)

var version = "undefined"

var (
	bindReceiver = flag.String("bind-receiver", "0.0.0.0:4589", "socket address for request datagrams")
	bindSender   = flag.String("bind-sender", "0.0.0.0:4590", "socket address for response datagrams")
	showVersion  = flag.Bool("version", false, "show program version and exit")
)

func run() int {
	flag.Parse()
	if *showVersion {
		fmt.Println(version)
		return 0
	}

	rxAddr, err := net.ResolveUDPAddr("udp", *bindReceiver)
	if err != nil {
		log.Fatalf("Can't resolve receiver bind address: %v", err)
	}

	txAddr, err := net.ResolveUDPAddr("udp", *bindSender)
	if err != nil {
		log.Fatalf("Can't resolve sender bind address: %v", err)
	}

	rxSocket, err := net.ListenUDP("udp", rxAddr)
	if err != nil {
		log.Fatalf("Can't bind receiver socket: %v", err)
	}
	defer rxSocket.Close()

	txSocket, err := net.ListenUDP("udp", txAddr)
	if err != nil {
		log.Fatalf("Can't bind receiver socket: %v", err)
	}
	defer txSocket.Close()

	buf := make([]byte, PacketSize)
	for {
		n, peerAddress, err := rxSocket.ReadFromUDP(buf)
		if err != nil {
			if n == 0 {
				log.Fatalf("UDP receive failed: %v", err)
			}
			continue
		}

		if n != PacketSize {
			continue
		}

		dstPort := binary.BigEndian.Uint16(buf)
		uuid := buf[PortFieldSize:]
		peerAddress.Port = int(dstPort)

		_, err = txSocket.WriteToUDP(uuid, peerAddress)
		if err != nil {
			log.Printf("UDP response send failed: %v", err)
		}
	}
	return 0
}

func main() {
	os.Exit(run())
}
