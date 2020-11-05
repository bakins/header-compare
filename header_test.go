package header

import (
	"net/http"
	"strings"
	"testing"
)

const contentTypeHeader = "Content-Type"

func BenchmarkGetCommon(b *testing.B) {
	header := make(http.Header)
	header.Set(contentTypeHeader, "application/json")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = header.Get(contentTypeHeader)
	}
}

func BenchmarkGetCommonLowercase(b *testing.B) {
	header := make(http.Header)
	header.Set(contentTypeHeader, "application/json")

	key := strings.ToLower(contentTypeHeader)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = header.Get(key)
	}
}

func BenchmarkGetCommonDirect(b *testing.B) {
	header := make(http.Header)
	header.Set(contentTypeHeader, "application/json")

	key := http.CanonicalHeaderKey(contentTypeHeader)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Get(header, key)
	}
}

const uncommonHeader = "X-Uncommon-Header"

func BenchmarkGetUncommon(b *testing.B) {
	header := make(http.Header)
	header.Set(uncommonHeader, "application/json")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = header.Get(uncommonHeader)
	}
}

func BenchmarkGetUncommonLowercase(b *testing.B) {
	header := make(http.Header)
	header.Set(uncommonHeader, "application/json")

	key := strings.ToLower(uncommonHeader)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = header.Get(key)
	}
}

func BenchmarkGetUncommonDirect(b *testing.B) {
	header := make(http.Header)
	header.Set(uncommonHeader, "application/json")

	key := http.CanonicalHeaderKey(uncommonHeader)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Get(header, key)
	}
}
