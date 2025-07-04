package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "azxcvbnmasdfghjklpoiuytrewq"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // min->max
}

// RandomString generates a random string  of length n
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

// RandomOwner generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomOwner generates a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "RUB"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email address
func RandomEmail() string {
	return RandomString(6) + "@" + RandomString(6) + ".com"
}
