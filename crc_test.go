package crc

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"testing"
)

func gocrc(data []byte) (out [4]byte) {
	binary.BigEndian.PutUint32(out[:], crc32.ChecksumIEEE(data))
	return out
}

func TestCRCAgainstStdlib(t *testing.T) {
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			buf := make([]byte, 4096+i)
			_, err := rand.Read(buf)
			if err != nil {
				t.Fatal(err)
			}
			if CRCToBytes(CRC(InitialCRC, buf[j:])) != gocrc(buf[j:]) {
				t.Fatal("crc doesn't match")
			}
		}
	}
}

func TestCRC(t *testing.T) {
	cases := []struct {
		input string
		hex   string
	}{
		{"hello world", "0d4a1185"},
		{"look ma i can checksum", "38c6c764"},
		{"", "00000000"},
	}
	for i, test := range cases {
		out := fmt.Sprintf("%x", CRCToBytes(CRC(InitialCRC, []byte(test.input))))
		if out != test.hex {
			t.Errorf("%d: %s != %s", i, out, test.hex)
		}
	}
}

func BenchmarkCRC_Zlib(b *testing.B) {
	buf := make([]byte, 65536)
	_, err := rand.Read(buf)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.SetBytes(int64(len(buf)))
	for i := 0; i < b.N; i++ {
		CRCToBytes(CRC(InitialCRC, buf))
	}
}

func BenchmarkCRC_Stdlib(b *testing.B) {
	buf := make([]byte, 65536)
	_, err := rand.Read(buf)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	b.SetBytes(int64(len(buf)))
	for i := 0; i < b.N; i++ {
		gocrc(buf)
	}
}
