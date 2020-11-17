package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	// "strconv"
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

	str := <-strs
	i := len(*str)-1
	last := (*str)[i]

	tot := int64(0)

	if (last - '0') % 4 == 0 {
		tot++
	}

	ctr := int64(0)

	for i := len(*str)-2; i >= 0; i-- {
		ths := (*str)[i] - '0'
		x := ths * 10 + last
		if x % 4 == 0{
			ctr++
		}
		tot += ctr
		last = ths
		if last % 4 == 0 {
			tot++
		}
	}

	ws.WriteString(fmt.Sprint(tot))

	ws.Flush()
}


