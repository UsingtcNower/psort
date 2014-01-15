package main

import (
	"fmt"
	"psort/psort"
)

func main() {
	var MAXLEN int = 100000
	a := make(psort.IntSlice, MAXLEN)
	for i:=0;i<MAXLEN;i++ {
		a[i]=MAXLEN-i
	}
	psort.Psort(a)
	fmt.Println(a)
}
