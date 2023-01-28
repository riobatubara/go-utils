package uuid

import (
	"crypto/rand"
	"io"
	"net"
	"sync/atomic"
	"time"
)

var hardwareAddr []byte
var clockSeq uint32

const (
	VariantNCSCompat = 0
	VariantIETF      = 2
	VariantMicrosoft = 6
	VariantFuture    = 7
)

func init() {
	if interfaces, err := net.Interfaces(); err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagLoopback == 0 && len(i.HardwareAddr) > 0 {
				hardwareAddr = i.HardwareAddr
				break
			}
		}
	}
	if hardwareAddr == nil {
		// If failed to obtain the MAC address of the current computer,
		// it will use a randomly generated 6 byte sequence instead and set
		// the multicast bit as recommended in RFC 4122.
		hardwareAddr = make([]byte, 6)
		_, err := io.ReadFull(rand.Reader, hardwareAddr)
		if err != nil {
			panic(err)
		}
		hardwareAddr[0] = hardwareAddr[0] | 0x01
	}

	// initialize the clock sequence with a random number
	var clockSeqRand [2]byte
	io.ReadFull(rand.Reader, clockSeqRand[:])
	clockSeq = uint32(clockSeqRand[1])<<8 | uint32(clockSeqRand[0])
}

func (this UUIDBuilder) Timebased() UUID {
	timeutc := time.Now().In(time.UTC)
	t := uint64(timeutc.Unix()-*this.Timestamp)*10000000 + uint64(timeutc.Nanosecond()/100)

	var u UUID
	u[0], u[1], u[2], u[3] = byte(t>>24), byte(t>>16), byte(t>>8), byte(t)
	u[4], u[5] = byte(t>>40), byte(t>>32)
	u[6], u[7] = byte(t>>56)&0x0F, byte(t>>48)

	clock := atomic.AddUint32(&clockSeq, 1)
	u[8] = byte(clock >> 8)
	u[9] = byte(clock)

	copy(u[10:], hardwareAddr)

	u[6] |= 0x10 // set version to 1 (time based uuid)
	u[8] &= 0x3F // clear variant
	u[8] |= 0x80 // set to IETF variant
	return u
}
