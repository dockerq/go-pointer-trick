package main

import "fmt"

type R struct {
	ID  int
	CPU float32
	MEM float32
}

func main() {
	rr := GetR()
	var rs []*R
	for k, r := range *rr {
		fmt.Printf("%dth r, id: %d, cpu: %f, mem: %f\n", k, r.ID, r.CPU, r.MEM)
		rs = append(rs, &r)
	}
	fmt.Println(rs)
	iter(&rs)
}

func NewR(i int, c, m float32) *R {
	return &R{
		ID:  i,
		CPU: c,
		MEM: m,
	}
}

func iter(rs *[]*R) {
	for _, r := range *rs {
		fmt.Printf("id: %d, cpu: %f, mem: %f\n", r.ID, r.CPU, r.MEM)
	}
}

func GetR() *[]R {
	rr := &[]R{
		*NewR(0, 4.0, 16000.0),
		*NewR(1, 2.0, 8000.0),
		*NewR(2, 1.0, 4000.0),
	}
	return rr
}
