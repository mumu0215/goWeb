package main

import (
	"strconv"
	"testing"
)

func CreateMap(m map[int]string,i int)  {
	m[i]=strconv.Itoa(i)
}

func BenchmarkMap(b *testing.B)  {
	a:=make(map[int]string,100000)
	for i:=0;i<100000;i++{
		CreateMap(a,i)
	}
}
func BenchmarkMap1(b *testing.B)  {
	a:=make(map[int]string,100000)
	for i:=0;i<100000;i++{
		CreateMap(a,i)
	}
}