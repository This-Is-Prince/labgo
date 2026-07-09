package main

import (
	"fmt"

	dnslookup "github.com/This-Is-Prince/labgo/networking/dns_lookup"
	"github.com/This-Is-Prince/labgo/networking/h2frame_client"
	"github.com/This-Is-Prince/labgo/networking/inspect_tls_handshake"
	"github.com/This-Is-Prince/labgo/networking/layerlab"
	rawtcpserver "github.com/This-Is-Prince/labgo/networking/raw_tcp_server"
	"github.com/This-Is-Prince/labgo/networking/sniff"
	tcpclient "github.com/This-Is-Prince/labgo/networking/tcp_client"
	tlsclient "github.com/This-Is-Prince/labgo/networking/tls_client"
	udpserver "github.com/This-Is-Prince/labgo/networking/udp_server"
)

func main() {
	fmt.Print("Networking examples\n\n")

	tlsclient.RunTlsClient(false)
	rawtcpserver.RunRawTcpServer(false)
	dnslookup.RunDnsLookup(false)
	tcpclient.RunTcpClient(false)
	udpserver.RunUdpServer(false)
	inspect_tls_handshake.RunInspectTlsHandshake(false)
	layerlab.RunLayerLab(true)
	h2frame_client.RunH2FrameClient(false)
	sniff.RunSniff(false)
}
