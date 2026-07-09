package sniff

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func RunSniff(isEnabled bool) {
	if !isEnabled {
		return
	}

	fmt.Print("Sniff\n\n")

	iface := "lo0" // macOS loopback
	if len(os.Args) > 1 {
		iface = os.Args[1]
	}

	filter := "tcp port 8080 or tcp port 8443"
	if len(os.Args) > 2 {
		filter = os.Args[2]
	}

	fmt.Println("interface:", iface)
	fmt.Println("BPF filter:", filter)

	handle, err := pcap.OpenLive(iface, 65535, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(filter); err != nil {
		log.Fatal(err)
	}

	source := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range source.Packets() {
		printPacket(packet)
	}
}

func printPacket(packet gopacket.Packet) {
	networkLayer := packet.NetworkLayer()
	transportLayer := packet.TransportLayer()

	if networkLayer == nil || transportLayer == nil {
		return
	}

	srcIP, dstIP := networkLayer.NetworkFlow().Endpoints()

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer == nil {
		return
	}

	tcp, ok := tcpLayer.(*layers.TCP)
	if !ok {
		return
	}

	fmt.Println()
	fmt.Println("==================================================")
	fmt.Println("TCP PACKET")
	fmt.Println("==================================================")

	fmt.Printf("%s:%s -> %s:%s\n", srcIP, tcp.SrcPort, dstIP, tcp.DstPort)
	fmt.Printf("Seq=%d Ack=%d Window=%d\n", tcp.Seq, tcp.Ack, tcp.Window)
	fmt.Printf("Flags: SYN=%v ACK=%v FIN=%v RST=%v PSH=%v\n", tcp.SYN, tcp.ACK, tcp.FIN, tcp.RST, tcp.PSH)
	fmt.Printf("Payload bytes: %d\n", len(tcp.Payload))

	if len(tcp.Payload) > 0 {
		max := len(tcp.Payload)
		if max > 512 {
			max = 512
		}

		fmt.Println("Payload preview:")
		fmt.Println(hex.Dump(tcp.Payload[:max]))

		if len(tcp.Payload) > max {
			fmt.Printf("... truncated, total payload %d bytes\n", len(tcp.Payload))
		}
	}
}
