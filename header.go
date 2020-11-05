package header

import "net/http"

// based on https://golang.org/pkg/net/textproto/

// Get a header value by using a key that is already in canonical form.
func Get(h http.Header, key string) string {
	if h == nil {
		return ""
	}

	v := h[key]

	if len(v) == 0 {
		return ""
	}

	return v[0]
}

// Set a header value by using a key that is already canonical.
func Set(h http.Header, key string, value string) {
	h[key] = []string{value}
}

// Add a header value by using a key that is already canonical.
func Add(in http.Header, key string, value string) {
	in[key] = append(in[key], value)
}


// Values returns all values associated with the given key.
// Key must already be in canonical form.
func  Values(h http.Header, key string) []string {
	if h == nil {
		return nil
	}

	return h[key]
}
