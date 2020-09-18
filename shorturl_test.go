package shorturl

import (
	"strings"
	"testing"
)

func TestShortUrl(t *testing.T) {
	total := 1000
	generated := make(map[string]bool)
	for i := 0; i < total; i++ {
		result, err := CreateShortURL("https://longurltoshortne.com", "http://shorturl:3000")
		if err != nil {
			t.Error(err)
		}
		shortid := strings.Split(result, "/")[3]
		if generated[shortid] {
			t.Errorf("shortid already created")
		}
		generated[shortid] = true
	}
}
