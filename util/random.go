package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomFileName() string {
	return RandomString(rand.Int())
}

func RandomArchiveSize() float64 {
	return RandomFloat64(0, 100990000)
}
func RandomTotalSize() float64 {
	return RandomFloat64(0, 100000000)
}

func RandomTotalFiles() float64 {
	return RandomFloat64(0, 10)
}

func RandomFilePath() string {
	return RandomString(6)
}
func RandomSize() float64 {
	return RandomFloat64(0, 1000)
}

func RandomMimeType() string {
	types := []string{"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"application/xml",
		"image/jpeg",
		"image/png",
	}
	n := len(types)
	return types[rand.Intn(n)]
}
