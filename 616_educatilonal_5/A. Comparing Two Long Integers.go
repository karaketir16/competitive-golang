package main

import (
	// "fmt"
	"bufio"
	"os"
	"strings"
	// "strconv"
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

func main(){
	rs := bufio.NewReader(os.Stdin)
	ws := bufio.NewWriter(os.Stdout)

	strs := make(chan *string, 100)

	go inputter(rs, strs)


	number1 := strings.TrimLeft(*<-strs, "0")
	number2 := strings.TrimLeft(*<-strs, "0")

	if len(number1) + len(number2) == 0{
		ws.WriteString("=")
	} else if len(number1) > len(number2){
		ws.WriteString(">")
	} else if len(number1) < len(number2){
		ws.WriteString("<")
	} else {
		if number1 > number2{
			ws.WriteString(">")
		} else if number1 < number2{
			ws.WriteString("<") 
		} else{
			ws.WriteString("=")
		}
	}
	ws.Flush()
	return
}
