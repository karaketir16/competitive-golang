package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	// "sort"
	// "math"
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

func min(values []int64) int64{
	mn := values[0]
	for _, v := range values {
			if (v < mn) {
				mn = v
			}
	}
	return mn
}

func main(){
	rs := bufio.NewReader(os.Stdin)
	ws := bufio.NewWriter(os.Stdout)

	strs := make(chan *string, 100)

	go inputter(rs, strs)

	n , _ := strconv.ParseInt(*<-strs, 10, 64)
	m , _ := strconv.ParseInt(*<-strs, 10, 64)

	a := make([]int64, m)
	
	mx := int64(0)
	for i := int64(0); i < n; i++ {
		
		for j := int64(0); j < m; j++ {
			x, _ := strconv.ParseInt(*<-strs, 10, 64)
			a[j] = x
		}

		y := min(a)
		
		if y > mx {
			mx = y
		}
	}

	ws.WriteString(fmt.Sprint(mx))

	ws.Flush()
}


