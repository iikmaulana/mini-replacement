package lib

import (
	"math"
	"math/rand"
	"time"
)

func RandomCharacter(length int) string {
	tmpStr := "0123456789"
	e, r := len(tmpStr), ""

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		c := int(math.Floor(rand.Float64() * float64(e)))
		r += string([]rune(tmpStr)[c])
	}
	return r
}
