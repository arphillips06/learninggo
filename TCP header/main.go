package main

import "fmt"

func main() {
	//header length (measured in 32-bit words) is 5
	var headerWords uint8 = 5
	//header length in bytes is 20
	headerLen := headerWords * 32 / 8
	//build a slice of 32 bytes to store the TCP header
	b := make([]byte, headerLen)
	// Shift header words bits to the left to fit
	// the Header length field of the TCP header
	s := headerWords << 4
	//OR operation on byte 13 and the store new value
	b[13] = b[13] | s
	fmt.Printf("%08b\n", b[13])
	//assume that this is SYN message
	var tcpSYN uint8 = 1
	//SYN flag is the 2nd bit from the right so
	//we shift it right my 1 positions
	ack := tcpSYN << 1
	//OR operation on byte 14 and store the new value
	b[14] = b[14] | ack
	//print the byte 14 of the TCP header
	fmt.Printf("%08b\n", b[14])
	//only interested if a TCP SYN flag has been set
	tcpSynFlag := (b[14] & 0x02) != 0
	//shift header length right, drop any low-order bits
	parsedHeaderWords := b[13] >> 4
	//prints "TCP Flag is set: true"
	fmt.Printf("TCP Flag is set: %t\n", tcpSynFlag)
	// prints "TCP header words: 5"
	fmt.Printf("TCP header words: %d\n", parsedHeaderWords)
}
