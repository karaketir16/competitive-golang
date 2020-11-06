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

	str := strings.Split(*<-strs, ":")
	a , _ := strconv.ParseInt(str[0], 10, 64)
	b , _ := strconv.ParseInt(str[1], 10, 64)
	c , _ := strconv.ParseInt(*<-strs, 10, 64)

	b += c

	a += b/60
	b %= 60
	a %= 24

	ws.WriteString(fmt.Sprintf("%02d:%02d", a, b))
	ws.Flush()
}


