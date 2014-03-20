// Copyright (C) 2013 Space Monkey, Inc.

package utils

import (
	"encoding/binary"
	"unsafe"
)

func crc_internal(crc uint32, data uintptr, length int32, out *uint32)

var (
	InitialCRC = CRC(0, nil)
)

func CRC(crc uint32, data []byte) (new_crc uint32) {
	if len(data) > 0 {
		crc_internal(crc, uintptr(unsafe.Pointer(&data[0])), int32(len(data)),
			&new_crc)
	} else {
		crc_internal(crc, 0, 0, &new_crc)
	}
	return new_crc
}

func CRCToBytes(crc uint32) (rv [4]byte) {
	binary.BigEndian.PutUint32(rv[:], crc)
	return rv
}

func CRCFromBytes(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}
