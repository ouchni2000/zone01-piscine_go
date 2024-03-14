package main

import "fmt"

func Enigma(a ***int, b *int, c *******int, d ****int) {
	n := ***a
	nb := *******c
	***a = *b
	*b = ****d
	*******c = n
	****d = nb
}

func main() {
	x := 9
	y := &x
	z := &y
	a := &z

	w := 1
	b := &w

	u := 8
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j

	k := 5
	l := &k
	m := &l
	n := &m
	d := &n

	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

	Enigma(a, b, c, d)

	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

}
