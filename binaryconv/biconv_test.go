//go:build !race

package binaryconv

import (
	"bytes"
	"encoding/binary"
	"testing"
	"unsafe"
)

func makeTestData() []byte {
	var h HolderWrapper
	h.Bachpana.Mauj = 123456789
	h.Bachpana.Masti = 42
	h.Budhapa.Kshama = 0x01020304
	h.Budhapa.Tirth = 0x05060708
	h.Budhapa.Yagna = 1234
	h.Budhapa.Aarohi = 5678
	h.Budhapa.Krodh = 17
	h.Budhapa.Paap = 3
	copy(h.Payload[:], []byte("hello world"))
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &h)
	return buf.Bytes()
}

func BenchmarkReadHolderWrapper_BinaryRead(b *testing.B) {
	data := makeTestData()
	const HolderWrapperSize = int(unsafe.Sizeof(HolderWrapper{}))

	buf := bytes.NewReader(data[:HolderWrapperSize])

	for b.Loop() {
		ReadHolderWrapper_BinaryRead(data, buf)
	}
}

func BenchmarkReadHolderWrapper_UnsafeCast(b *testing.B) {
	data := makeTestData()
	for b.Loop() {
		ReadHolderWrapper_UnsafeCast(data)
	}
}

func BenchmarkReadHolderWrapper_Manual(b *testing.B) {
	data := makeTestData()

	for b.Loop() {
		ReadHolderWrapper_Manual(data)
	}
}
