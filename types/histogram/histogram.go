package histogram

import (
	"fmt"
	"math"
	"strings"
)

type HistType byte

const (
	Linear HistType = iota
	Exponential
)

type hvals interface {
	int | uint |
		int64 | int32 | int16 | int8 |
		uint64 | uint32 | uint16 | uint8 |
		float64 | float32
}

type Histogram[T hvals] struct {
	htype HistType
	base  T

	Buckets  []uint // slot -> count
	Min, Max T
	Sum      float64
	Count    uint
}

func NewLinear[T hvals](base T, buckets int) *Histogram[T] {
	return &Histogram[T]{
		htype:   Linear,
		base:    base,
		Buckets: make([]uint, buckets),
	}
}

func NewExponential[T hvals](base T, buckets int) *Histogram[T] {
	return &Histogram[T]{
		htype:   Exponential,
		base:    base,
		Buckets: make([]uint, buckets),
	}
}

func (h *Histogram[T]) Reset() {
	for i := range h.Buckets {
		h.Buckets[i] = 0
	}

	h.Min = 0
	h.Max = 0
	h.Sum = 0
	h.Count = 0
}

func (h *Histogram[T]) String() string {
	res := new(strings.Builder)
	switch h.htype {
	case Linear:
		for i := 0; i < len(h.Buckets)-1; i++ {
			v := float64(h.base) * float64(i)
			fmt.Fprintf(res, "[%v, %v) => %d\n", v, v+float64(h.base), h.Buckets[i])
		}
		fmt.Fprintf(res, "[%v, inf) => %d\n", float64(h.base)*float64(len(h.Buckets)-1), h.Buckets[len(h.Buckets)-1])
	case Exponential:
		fmt.Fprintf(res, "[0, %v) => %d\n", float64(h.base), h.Buckets[0])
		for i := 1; i < len(h.Buckets)-1; i++ {
			v := math.Pow(float64(h.base), float64(i))
			fmt.Fprintf(res, "[%v, %v) => %d\n", v, v*float64(h.base), h.Buckets[i])
		}
		fmt.Fprintf(res, "[%v, inf) => %d\n", math.Pow(float64(h.base), float64(len(h.Buckets))-1), h.Buckets[len(h.Buckets)-1])
	}
	return res.String()
}

func (h *Histogram[T]) Median() T {
	var s uint = 0
	c := h.Count / 2
	for i, bv := range h.Buckets {
		s += bv
		if s >= c {
			// found the bucket
			if h.htype == Linear {
				return T(i+1) * h.base
			}
			return T(math.Pow(float64(h.base), float64(i+1)))
		}
	}
	return h.Max
}

func (h *Histogram[T]) Add(v T) {
	if h.Count == 0 {
		h.Max = v
		h.Min = v
	} else {
		if v > h.Max {
			h.Max = v
		} else if v < h.Min {
			h.Min = v
		}
	}

	h.Sum += float64(v)
	h.Count++

	var slot int
	switch h.htype {
	case Linear:
		slot = int(math.Floor(float64(v / T(h.base))))
	case Exponential:
		slot = int(math.Floor(math.Log(float64(v)) / math.Log(float64(h.base))))
	}

	if slot >= len(h.Buckets) {
		h.Buckets[len(h.Buckets)-1]++
	} else if slot < 0 {
		h.Buckets[0]++
	} else {
		h.Buckets[slot]++
	}
}
