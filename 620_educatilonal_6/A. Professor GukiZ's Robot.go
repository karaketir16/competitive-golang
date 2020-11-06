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

func abs(a int64) int64 {
	if a > 0{
		return a
	}
	return 0 - a
}

func main(){
	rs := bufio.NewReader(os.Stdin)
	ws := bufio.NewWriter(os.Stdout)

	strs := make(chan *string, 100)

	go inputter(rs, strs)

	x1, _ := strconv.ParseInt(*<-strs, 10, 64)
	y1, _ := strconv.ParseInt(*<-strs, 10, 64)
	x2, _ := strconv.ParseInt(*<-strs, 10, 64)
	y2, _ := strconv.ParseInt(*<-strs, 10, 64)

	f1 := abs(x1 - x2)
	f2 := abs(y1 - y2)
	
	if f1 > f2 {
		ws.WriteString(fmt.Sprint(f1))
	} else {
		ws.WriteString(fmt.Sprint(f2))
	}

	ws.Flush()
	return
}
