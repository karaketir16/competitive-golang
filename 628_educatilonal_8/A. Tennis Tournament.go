package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	// "sort"
	"math"
)

//N _
const (
	N = 1e6 + 5
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

func tester(a int64) int64 {
	chTest := make(chan int64)
	go func(){
		for i:=int64(1) ; ;i++{
			for j := int64(1); j<=i; j++{
				chTest <- j
			}
		}
	}()
	var x int64
	for i := int64(0); i < a; i++{
		x = <-chTest
	}
	return x
}




func calculator(a int64) int64{
	x := math.Sqrt(float64(2*a))
	y := int64(x)

	for (y*(y+1))/2 >=a {
		y--
	}
	return a - (y*(y+1))/2
}

func main(){
	rs := bufio.NewReader(os.Stdin)
	ws := bufio.NewWriter(os.Stdout)

	strs := make(chan *string, 100)

	go inputter(rs, strs)

	n, _ := strconv.ParseInt(*<-strs, 10, 64)
	b, _ := strconv.ParseInt(*<-strs, 10, 64)
	p, _ := strconv.ParseInt(*<-strs, 10, 64)

	y := n * p
	
	x := int64(0)

	for a := int64(1 << 9); a > 1; a >>= 1 {
		if a & n != 0 {
			x += a * b + (a >> 1)
			n -= (a >> 1)
			a <<= 1
		}
	}

	ws.WriteString(fmt.Sprintf("%d %d", x, y))

	ws.Flush()
	return
}
