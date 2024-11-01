package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnoprstuvxyz"

var rng *rand.Rand

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))

}

// RandomInt generaters a ramdom integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min-1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(500, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{TL, EURO, USD}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}

// RandomEmail generates
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
