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
	p , _ := strconv.ParseInt(*<-strs, 10, 64)
	q , _ := strconv.ParseInt(*<-strs, 10, 64)
	str := <-strs

	var x,y int64

	for i := int64(0); i * p <= n; i ++{
		if (n - i * p)%q == 0{
			x = i
			y = (n - i * p)/q 
		}
	}
	
	if x + y == 0{
		ws.WriteString("-1")
		ws.Flush()
		return
	} 
	
	ws.WriteString(fmt.Sprintln(x + y))


	for i := int64(0); i < x; i++{
		ws.WriteString((*str)[i * p: (i + 1) * p ] + "\n")
	}
	
	for i := int64(0); i < y; i++{
		ws.WriteString((*str)[x*p + i*q : x*p + (i + 1)*q ] + "\n")
	}

	ws.Flush()
}


