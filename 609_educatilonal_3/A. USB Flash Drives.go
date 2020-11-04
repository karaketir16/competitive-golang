package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
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

	a := make([]int, 0, N)

	for i := 0; i < int(n); i++ {
		x, _ := strconv.ParseInt(*<-strs, 10, 32)
		a = append(a, int(x))
	}

	sort.SliceStable(a, func(i,j int)bool{return a[i] > a[j]})

	tot := int64(0)

	for i := 0; i < int(n); i++ {
		tot += int64(a[i])
		if tot >= m{
			ws.WriteString(fmt.Sprint(i + 1))
			break
		}
	}
	
	ws.Flush()
}


