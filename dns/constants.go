package dns

const (
	ipV4FragmentLen = 20
	protocolOffset  = 9

	//tcpProtocol byte = 0x06
	// tcp flags: source port (2) + target ports (2) + seq# (4) + ack# (4) + header (4bit) + reserved (4bit)
	//tcpFlagsLen = 13

	// tcp flags from beginning of IPv4 packet
	//tcpFlagsOffset = ipV4FragmentLen + tcpFlagsLen

	// |CWR|ECE|URG|ACK|PSH|RST|SYN|FIN|
	// |-0-|-0-|-0-|-0-|-0-|-0-|-1-|-0-|
	//tcpSYNflag = 0x02

	// |CWR|ECE|URG|ACK|PSH|RST|SYN|FIN|
	// |-0-|-0-|-0-|-0-|-0-|-1-|-0-|-0-|
	// tcpRSTflag = 0x04
	// tcpFlags   = tcpRSTflag | tcpSYNflag

	udpProtocol byte = 0x11

	// udp header: source port (2) + target ports (2) , payload length (2), checksum (2)
	udpHeaderLen    = 8
	udpHeaderOffset = ipV4FragmentLen

	// dns response source port takes one byte
	// 00110101
	udpSourcePort byte = 0x35 //53, byte(0b00110101)

	// udp one byte source port from beginning of IPv4 packet
	udpSourcePortOffset = udpHeaderOffset + 1

	// udp payload segment from beginning of IPv4 packet
	// ip (20) + udp (8)
	//udpPayloadOffset = ipV4FragmentLen + udpHeaderLen
)
