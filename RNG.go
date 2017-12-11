package Thoth

import (
	"math/rand"
	"time"
)

// https://stackoverflow.com/questions/45030618/generate-a-random-bool-in-go
type Boolgen struct {
	src       rand.Source
	cache     int64
	remaining int
}

func (b *Boolgen) RandBool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}
	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--
	return result
}

func NewBoolgen() *Boolgen {
	return &Boolgen{src: rand.NewSource(time.Now().UnixNano())}
}

func RandFloat() float64 {
	return rand.Float64()
}

func RandFloatBetween(min float64, max float64) float64 {
	r := max - min
	return (rand.Float64() * r) + min
}

func RNG_iRand() int {
	return 0
}

// TODO - Integer specific RNG per thread
func RandPostiveIntUpTo(max int) int {
	return rand.Intn(max)
}

func RandPostiveIntBetween(min int, max int) int {
	return rand.Intn(max - min) + min
}
