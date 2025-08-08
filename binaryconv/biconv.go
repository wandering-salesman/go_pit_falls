package binaryconv

import (
	"encoding/binary"
	"errors"
	"io"
	"unsafe"
)

const PayloadSize = 16 * 1024 // 4kb - 16KB
type Chakra struct {
	Kshama uint32
	Tirth  uint32
	Yagna  uint16
	Aarohi uint16
	Krodh  uint8
	Paap   uint8
	_      [2]byte // padding to align next field to 4 bytes
}

type Bachpan struct {
	Mauj  uint64
	Masti uint32
	_     [4]byte // padding to align Chakra to 8-byte boundary
}

type HolderWrapper struct {
	Bachpana Bachpan
	Budhapa  Chakra
	Payload  [PayloadSize]byte
}

func ReadHolderWrapper_BinaryRead(data []byte, buf io.Reader) (HolderWrapper, error) {
	var h HolderWrapper
	if len(data) < int(unsafe.Sizeof(h)) {
		return h, errors.New("insufficient data for HolderWrapper")
	}
	err := binary.Read(buf, binary.LittleEndian, &h)
	return h, err
}

func ReadHolderWrapper_UnsafeCast(data []byte) (HolderWrapper, error) {
	if len(data) < int(unsafe.Sizeof(HolderWrapper{})) {
		return HolderWrapper{}, errors.New("data too short for HolderWrapper")
	}
	return *(*HolderWrapper)(unsafe.Pointer(&data[0])), nil
}

func ReadHolderWrapper_Manual(data []byte) (HolderWrapper, error) {
	var h HolderWrapper
	if len(data) < int(unsafe.Sizeof(h)) {
		return h, errors.New("data too short")
	}

	// MetaHeader
	h.Bachpana.Mauj = binary.LittleEndian.Uint64(data[0:8])
	h.Bachpana.Masti = binary.LittleEndian.Uint32(data[8:12])
	// NetHeader
	h.Budhapa.Kshama = binary.LittleEndian.Uint32(data[16:20])
	h.Budhapa.Tirth = binary.LittleEndian.Uint32(data[20:24])
	h.Budhapa.Yagna = binary.LittleEndian.Uint16(data[24:26])
	h.Budhapa.Aarohi = binary.LittleEndian.Uint16(data[26:28])
	h.Budhapa.Krodh = data[28]
	h.Budhapa.Paap = data[29]
	// skip 2 bytes padding

	// Data
	copy(h.Payload[:], data[32:32+len(h.Payload)]) // since slice won't be directly assignable to array

	return h, nil
}
