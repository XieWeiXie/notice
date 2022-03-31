package main

import "fmt"

func main() {
	//BiAn := service.NewIService(service.BiAn)
	//BiAn.Notice("")
	fmt.Println(Sum([]int{1, 2, 3, 4}))
	fmt.Println(Sum2([]int{1, 2, 3, 4}))
	fmt.Println(Sum([]int64{1, 2, 3, 4}))
	fmt.Println(Sum2([]int64{1, 2, 3, 4}))
	fmt.Println(Sum([]float64{1, 2, 3, 4}))
	fmt.Println(Sum2([]float64{1, 2, 3, 4}))
	fmt.Println(Sum([]float32{1, 2, 3, 4}))
	fmt.Println(Sum2([]float32{1, 2, 3, 4}))

	Print([]string{"a", "b", "c"})
	Print([]int{1, 2, 3})
	Print([]float32{1, 2, 3})

	fmt.Println(SumMap(map[int]int{1: 1, 2: 2, 3: 3}))
	fmt.Println(SumMap(map[int]float64{1: 1.2, 2: 2.2, 3: 3.3}))

	a := Test[int]{list: []int{1, 2, 3, 4}}
	b := Test[string]{list: []string{"a", "b", "c"}}
	c := Test[float64]{list: []float64{1.1, 1.2, 1.3}}
	a.Print()
	b.Print()
	c.Print()
}

type Number interface {
	int | float64 | int64 | float32 | string
}

func Sum[N Number](sum []N) N {
	var total N
	for _, i := range sum {
		total += i
	}
	return total
}

func Sum2[m int | int64 | float64 | float32](sum []m) m {
	var total m
	for _, i := range sum {
		total += i
	}
	return total
}

func Print[t any](s []t) {
	// any is equal to interface
	for _, i := range s {
		fmt.Println(i)
	}
}

func SumMap[k int, v int | float64](m map[k]v) v {
	var total v
	for _, i := range m {
		total += i
	}
	return total
}

type Test[T Number] struct {
	list []T
}

func (t Test[T]) Print() {
	for _, i := range t.list {
		fmt.Println(i)
	}
}


