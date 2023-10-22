package dns

import (
	"golang.org/x/net/dns/dnsmessage"
)

// as in a number of other places, the reality of DNS does not match RFC 1035.
// nobody knows what multiple questions really mean - if one of them fails to
// answer correctly, how would that be signalled etc?
// so in practice, multiple questions don't work.
// One of the later RFCs might clarify this.
// even if theoretically multiple questions are valid according to the spec,
// servers don't actually support them. There will be at most one question here.
//
// implementation implications:
// we can skip loop counting over response's question/answer section
// var qc uint32
// for {
// 	err := p.SkipQuestion()
// 	if err != nil {
// 		break
// 	}
// 	qc++
// }
// return qc

func IsDNS(packet []byte) bool {
	// confirm protocol & source port
	// assuming BigIndian order & single byte port number
	return packet[protocolOffset] == udpProtocol &&
		packet[udpSourcePortOffset] == udpSourcePort
}

func IsBlockedDNSResponse(packet []byte) bool {
	var p dnsmessage.Parser
	// we check by Rcode
	if hh, err := p.Start(packet[:]); err != nil ||
		hh.RCode != dnsmessage.RCodeNameError {
		return false
	}
	return true
}
