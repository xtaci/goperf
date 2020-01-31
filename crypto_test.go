package main

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"hash/crc32"
	"io"
	"testing"
)

func BenchmarkCryptoRead(b *testing.B) {
	buf := make([]byte, 1<<20)
	b.SetBytes(1 << 20)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		io.ReadFull(rand.Reader, buf)
	}
}
func BenchmarkCRC32(b *testing.B) {
	content := make([]byte, 1024)
	b.SetBytes(int64(len(content)))
	for i := 0; i < b.N; i++ {
		crc32.ChecksumIEEE(content)
	}
}

func BenchmarkMD5(b *testing.B) {
	var data [md5.Size]byte
	b.SetBytes(md5.Size)

	for i := 0; i < b.N; i++ {
		data = md5.Sum(data[:])
	}
}
func BenchmarkSHA1(b *testing.B) {
	var data [sha1.Size]byte
	b.SetBytes(sha1.Size)

	for i := 0; i < b.N; i++ {
		data = sha1.Sum(data[:])
	}
}
