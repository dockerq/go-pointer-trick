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
	fmt.Printf("rr type is %T\n", rr)
	fmt.Println("rr is", rr)
	fmt.Println("*rr is", *rr)
	// fmt.Printf("*rs type is %T", *rs)
	for k, r := range *rr {
		fmt.Printf("%dth r, id: %d, cpu: %f, mem: %f\n", k, r.ID, r.CPU, r.MEM)
		// fmt.Println("rs is ", rs)
		// fmt.Println("r is ", r)
		// rs = append(rs, &(*rr)[k])
		rs = append(rs, &r)
		fmt.Println(rs)
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
