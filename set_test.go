package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T)  {
	s1 := New("1","2","3","4")
	s1.Add("4","5","6")
	fmt.Println(s1)
}
                                   
func TestHas(t *testing.T)  {
	s1 := New("1","2","3","4")
	fmt.Println(s1.Has("8") )
}

func TestUnion(t *testing.T) {
	s1 := New("5","4","2","1","3")
	s2 := New("3","5","7","8","9")
	fmt.Println(s1.Union(s2))
}

func TestIntersect(t *testing.T) {
	s1 := New("5","4","2","1","3")
	s2 := New("3","5","7","8","9")
	fmt.Println(s1.Intersect(s2))
}

func TestMinus(t *testing.T) {
	s1 := New("5","4","2","1","3")
	s2 := New("3","5","7","8","9")
	fmt.Println(s2.Minus(s1))
}

func TestComplement(t *testing.T) {
	s1 := New("5","4","2","1","3")
	s2 := New("5","4","2","1","3")
	fmt.Println(s2.Complement(s1))
}