package main

import (
	"fmt"
	"os"
	"strconv"
)

type packet struct {
	version, typeId, value int64
	subPackets             []packet
}

func (p packet) isLiteral() bool {
	return p.typeId == 4
}

func (p packet) sumVersion() int64 {
	var total int64
	for _, subPacket := range p.subPackets {
		total += subPacket.sumVersion()
	}
	return p.version + total
}

func (p packet) compute() int64 {
	var total int64
	if p.typeId == 0 { //sum
		if len(p.subPackets) == 1 {
			total = p.subPackets[0].compute()
		} else {
			for _, subPacket := range p.subPackets {
				total += subPacket.compute()
			}
		}
	} else if p.typeId == 1 { // product
		if len(p.subPackets) == 1 {
			total = p.subPackets[0].compute()
		} else {
			total = 1
			for _, subPacket := range p.subPackets {
				total *= subPacket.compute()
			}
		}
	} else if p.typeId == 2 { // minimum
		for i, subPacket := range p.subPackets {
			value := subPacket.compute()
			if i == 0 {
				total = value
			} else {
				if total > value {
					total = value
				}
			}
		}
	} else if p.typeId == 3 { // maximum
		for i, subPacket := range p.subPackets {
			value := subPacket.compute()
			if i == 0 {
				total = value
			} else {
				if total < value {
					total = value
				}
			}
		}
	} else if p.typeId == 4 { // literal
		total = p.value
	} else if p.typeId == 5 { //greater than
		if len(p.subPackets) != 2 {
			panic("subpacket of type 5 has not 2 packets")
		}
		if p.subPackets[0].compute() > p.subPackets[1].compute() {
			total = 1
		} else {
			total = 0
		}
	} else if p.typeId == 6 { // less than
		if len(p.subPackets) != 2 {
			panic("subpacket of type 6 has not 2 packets")
		}
		if p.subPackets[0].compute() < p.subPackets[1].compute() {
			total = 1
		} else {
			total = 0
		}
	} else if p.typeId == 7 { // equal to
		if len(p.subPackets) != 2 {
			panic("subpacket of type 7 has not 2 packets")
		}
		if p.subPackets[0].compute() == p.subPackets[1].compute() {
			total = 1
		} else {
			total = 0
		}
	}
	return total
}

func hexToBinary(msg string) string {
	r := ""
	for _, c := range msg {
		i, _ := strconv.ParseUint(string(c), 16, 32)
		r += fmt.Sprintf("%04b", i)
	}
	return r
}

func strBinaryToInt(s string) int64 {
	r, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return r
}

func decodeLiteral(msg string) (int64, string) {
	offset := 0
	blocks := ""
	for {
		block := msg[offset : offset+5]
		blocks += block[1:]
		offset += 5
		if block[0] == '0' {
			break
		}
	}
	res := strBinaryToInt(blocks)
	return res, msg[offset:]
}

func isLiteral(version string) bool {
	return version == "100"
}

func decodeOperator0(msg string) ([]packet, string) {
	subPacketLength := strBinaryToInt(msg[:15])
	subPacket := msg[15 : 15+subPacketLength]
	packets := make([]packet, 0)
	for {
		p, remainder := decodePacket(subPacket)
		packets = append(packets, p)
		subPacket = remainder
		if remainder == "" {
			break
		}
	}

	return packets, msg[15+subPacketLength:]
}

func decodeOperator1(msg string) ([]packet, string) {
	subPacketCount := strBinaryToInt(msg[:11])
	remainder := msg[11:]
	packets := make([]packet, subPacketCount)
	for i := 0; i < int(subPacketCount); i++ {
		p, rest := decodePacket(remainder)
		remainder = rest
		packets[i] = p
	}

	return packets, remainder
}

func decodePacket(msg string) (packet, string) {
	p := packet{version: strBinaryToInt(msg[:3]), typeId: strBinaryToInt(msg[3:6])}
	if p.isLiteral() { // literal
		value, remainder := decodeLiteral(msg[6:])
		p.value = value
		return p, remainder
	} else { // operator
		if msg[6] == '0' {
			subPackets, remainder := decodeOperator0(msg[7:])
			p.subPackets = subPackets
			return p, remainder
		} else {
			subPackets, remainder := decodeOperator1(msg[7:])
			p.subPackets = subPackets
			return p, remainder
		}
	}
}

func main() {
	contents, _ := os.ReadFile(os.Args[1])
	messageHex := string(contents)
	messageBin := hexToBinary(messageHex)
	p, _ := decodePacket(messageBin)
	fmt.Println(p.compute())
}
