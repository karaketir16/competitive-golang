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
	N = 1e6 + 5
)

var values []byte 

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

func digits(ch1 chan int64, ch2 chan byte){
	for x := range ch1{
		a := byte(0)
		for x > 0 {
			a += values[x%10]
			x/=10
		}
		ch2 <- a
	}
	close(ch2)
}

func main(){
	rs := bufio.NewReader(os.Stdin)
	ws := bufio.NewWriter(os.Stdout)

	strs := make(chan *string, 100)

	go inputter(rs, strs)

	a , _ := strconv.ParseInt(*<-strs, 10, 64)
	b , _ := strconv.ParseInt(*<-strs, 10, 64)

	values = []byte{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}

	ch1 := make(chan int64, N)
	ch2 := make(chan byte, N)

	go digits(ch1, ch2)

	go func(){
		for i := a; i <= b; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	total := int64(0)
	for x := range ch2 {
		total += int64(x)
	}

	ws.WriteString(fmt.Sprint(total))
	ws.Flush()
}


