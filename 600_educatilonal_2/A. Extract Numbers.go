package main

import (
	// "fmt"
	"bufio"
	"os"
	"strings"
	// "strconv"
)

func isNumber(str *string) bool{
	if len(*str) == 0{
		return false
	}
	if len(*str) == 1 && (*str)[0] == '0'{
		return true
	}
	if (*str)[0] == '0'{
		return false
	}
	for _, x:= range (*str){
		if x < '0' || x > '9'{
			return false
		}
	}
	return true
}

const(
	n = int(1e5 + 5)
)
func filter(input chan *string, intOut chan *string, strOut chan *string){
	for str := range input{
		if isNumber(str){
			intOut <- str
		} else {
			strOut <- str
		}
	}
	close(intOut)
	close(strOut)
}

func main(){
	rs := bufio.NewReader(os.Stdin)
	ws := bufio.NewWriter(os.Stdout)

	input, _ := rs.ReadString('\n')
	input = strings.TrimSpace(input)

	// parsed := strings.FieldsFunc(input, func(a rune) bool {return a == ';' || a == ','})
	parsed1 := strings.Split(input, ",")
	parsed	:= make([]string, 0, n)

	for i := range parsed1{
		if len(parsed1[i]) > 0{
			parsed2 := strings.Split(parsed1[i], ";")
			parsed = append(parsed, parsed2...)
		} else {
			parsed = append(parsed, parsed1[i])
		}
	}


	ints := make(chan *string)
	strs := make(chan *string)
	inputs := make(chan *string)

	go filter(inputs, ints, strs)

	go func (){
		for i := range parsed{
			inputs <- &parsed[i]
		}
		close(inputs)
	}()

	intString := make([]byte, 0, n)
	strString := make([]byte, 0, n)
	
	for {
		select {
		case x, ok := <-ints:
			if !ok {
				ints = nil
			} else {
				intString = append(intString, ',')
				intString = append(intString, []byte(*x)...)
			}

		case x, ok := <-strs:
			if !ok {
				strs = nil
			} else {
				strString = append(strString, ',')
				strString = append(strString, []byte(*x)...)
			}
		}
	
		if ints == nil && strs == nil {
			break
		}
	}

	if len(intString) > 0{
		ws.WriteString("\"" + string(intString[1:]) + "\"\n")
	} else {
		ws.WriteString("-\n")
	}


	if len(strString) > 0{
		ws.WriteString("\"" + string(strString[1:]) + "\"\n")
	} else {
		ws.WriteString("-\n")
	}

	ws.Flush()
}


