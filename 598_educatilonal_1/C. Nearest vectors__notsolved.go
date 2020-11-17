package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
	"math"
	"math/cmplx"
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

type vec struct{
	index int
	c complex128
}


func main(){
	rs := bufio.NewReader(os.Stdin)
	ws := bufio.NewWriter(os.Stdout)

	strs := make(chan *string, 100)

	go inputter(rs, strs)

	n , _ := strconv.ParseInt(*<-strs, 10, 64)

	arr := make([]vec, 0, n)

	for a :=0 ; a < int(n); a++{
		x , _ := strconv.ParseFloat(*<-strs, 64)
		y , _ := strconv.ParseFloat(*<-strs, 64)
		arr = append(arr, vec{index: a + 1 , c : complex(x, y)})
	}

	sort.SliceStable(arr, func(i, j int)bool{
		return cmplx.Phase(arr[i].c) < cmplx.Phase(arr[j].c)
	})

	// ws.WriteString(fmt.Sprintf("%v", arr) + "\n")

	// for i := range arr{
		// ws.WriteString(fmt.Sprintf("%v, %v | ", arr[i].c, cmplx.Phase(arr[i].c) / math.Pi ))
	// }

	// ws.WriteString("\n")
	minI := -1
	minV := math.Pi * 3

	for i := range arr{
		tmp := math.Abs(cmplx.Phase(arr[(i + 1) % int(n)].c) - cmplx.Phase(arr[i].c))
		if tmp >= math.Pi * 2 - tmp{
			tmp = math.Pi * 2 - tmp
		}
		if tmp < minV{
			minV = tmp
			minI = i
			// ws.WriteString(fmt.Sprintf("%d %v | ",i, tmp / math.Pi))
		} else {
			// ws.WriteString(fmt.Sprintf("%d %v | ",i, tmp / math.Pi))
		}
	}
	// ws.WriteString("\n")
	
	ws.WriteString(fmt.Sprintf("%d %d", arr[minI].index, arr[(minI + 1) % int(n)].index))
	ws.Flush()
}


