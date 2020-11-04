package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	// "sort"
)

//N _
const (
	N = 2*1e5 + 5
)

func inputter(inputReader *bufio.Reader,strs chan *string){
	for {
		input , _ := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		if len(input) > 0 {
			inputs := strings.Split(input, " ")
			for i := range inputs{
				strs <- &inputs[i]
			}
		}
	}
}

func main(){
	rs := bufio.NewReader(os.Stdin)
	ws := bufio.NewWriter(os.Stdout)

	strs := make(chan *string, 100)

	go inputter(rs, strs)

	n , _ := strconv.ParseInt(*<-strs, 10, 64)
	m , _ := strconv.ParseInt(*<-strs, 10, 64)

	a := make([]int, m)

	for i := 0; i < int(n); i++ {
		x, _ := strconv.ParseInt(*<-strs, 10, 32)
		a[int(x) - 1]++
	}
	tot := 0
	for i := 0; i < int(m); i++ {
		for j := i + 1; j < int(m); j++ {
			tot += a[i] * a[j]
		}
	}

	ws.WriteString(fmt.Sprint(tot))

	ws.Flush()
}


