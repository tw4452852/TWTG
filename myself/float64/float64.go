package float64

import (
	"fmt"
)

type Sorter interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

func Sort(data Sorter) {
	for i := 0; i < data.Len(); i++ {
		for j := i; j < data.Len() - 1; j++ {
			if data.Less(j + 1, j) {
				data.Swap(j, j + 1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	for i := 0; i < data.Len() - 1; i++ {
		if data.Less(i+1, i) {
			return false
		}
	}
	return true
}

type Float64Array []float64

func NewFlow64Array() Float64Array {
	return make([]float64, 25)
}

func (f Float64Array) Len() int {
	return len(f)
}

func (f Float64Array) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f Float64Array) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Float64Array) List() string {
	s := "{"

	for i := 0; i < f.Len(); i++ {
		if p[i] == 0 {
			continue
		}
		s += fmt.Sprintf("%3.1f", p[i])
	}
	s += "}"

	return s
}

func (f Float64Array) String() string {
	return f.List()
}
