package main

import (
	"fmt"
	"math"
)

type Hitung interface {
	Luas() float64
	Keliling() float64
}

type Lingkaran struct {
	diameter float64
}

func (l *Lingkaran) Keliling() float64 {
	return l.diameter / 2
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * math.Pow(l.Keliling(), 2)
}

type Persegi struct {
	sisi float64
}

func (p Persegi) Luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p Persegi) Keliling() float64 {
	return p.sisi * 4
}

func GetLuas(h Hitung) float64 {
	return h.Luas()
}

func main() {
	var bangunDatar Hitung

	bangunDatar = &Lingkaran{diameter: 10}
	fmt.Println("keliling", bangunDatar.Keliling())
	fmt.Println("luas", bangunDatar.Luas())

	bangunDatar = Persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.Luas())
	fmt.Println("keliling  :", bangunDatar.Keliling())
	fmt.Println(GetLuas(&Lingkaran{diameter: 23}))

}
