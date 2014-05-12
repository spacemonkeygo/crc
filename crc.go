/*
Package crc is a much faster implementation of the standard CRC polynomial.
Instead of operating on each byte like the stdlib version does, this version
operates on each word.

The main logic of this code was taking from zlib.

On my crappy laptop I get:
  BenchmarkCRC_C	   10000	    121849 ns/op	 537.84 MB/s
  BenchmarkCRC_Go	    5000	    317655 ns/op	 206.31 MB/s

The difference is much more pronounced on ARM devices.
*/
package crc

import (
	"encoding/binary"
	"unsafe"
)

func crc_internal(crc uint32, data uintptr, length int32, out *uint32)

const (
	InitialCRC uint32 = 0
)

// Calculates/updates a CRC from a byte array.
func CRC(crc uint32, data []byte) (new_crc uint32) {
	if len(data) > 0 {
		crc_internal(crc, uintptr(unsafe.Pointer(&data[0])), int32(len(data)),
			&new_crc)
	} else {
		crc_internal(crc, 0, 0, &new_crc)
	}
	return new_crc
}

// Given a uint32 crc, returns a 4-byte big endian representation
func CRCToBytes(crc uint32) (rv [4]byte) {
	binary.BigEndian.PutUint32(rv[:], crc)
	return rv
}

// Given a 4-byte big endian CRC representation, turns it into a uint32
func CRCFromBytes(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}
