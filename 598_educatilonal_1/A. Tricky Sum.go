package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func calc(n int64, ws *bufio.Writer, before, next chan bool) {
	tot := (n*(n+1))/2
	ch := make(chan int64, 100)

	go func(){
		for i := int64(1); i <= n; i *= 2 {
			ch <- i*2
		}
		close(ch)
	}()

	for x := range ch{
		tot -= x
	} 
	a := <-before
	ws.WriteString(fmt.Sprintln(tot))
	next <- a
}

func main(){
	rs := bufio.NewReader(os.Stdin)
	ws := bufio.NewWriter(os.Stdout)

	input, _ := rs.ReadString('\n')
	t, _ := strconv.ParseUint(strings.TrimSpace(input), 10, 64)

	first := make(chan bool)
	
	oldFirst := first
	
	for i := uint64(0); i < t; i++ {
		input, _ := rs.ReadString('\n')
		n, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
		second := make(chan bool)
		go calc(n, ws,first, second)
		first = second
	}

	oldFirst <- true
	<-first

	ws.Flush()
}


