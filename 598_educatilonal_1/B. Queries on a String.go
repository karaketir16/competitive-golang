package main

import (
	// "fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func readOneLineInts(inputreader *bufio.Reader, n int) (res chan int64) {
	res = make(chan int64)
	go func() {
		input, _ :=	inputreader.ReadString('\n')
		ints := strings.Split(strings.TrimSpace(input), " ")
		for i := range ints{
			x, _ := strconv.ParseInt(ints[i],10,64)
			res <- x
		}
		close(res)
	}()

	return 
}

func gcd(a, b int64) int64 {
	if b==0 {
		return a;
	} else {
		return gcd(b, a%b);
	}
}

func rotate(nums []byte, k int64) {
    if k < 0 || len(nums) == 0 {
        return
	}
	
	ln := int64(len(nums))

	k = k%ln
	k = ln - k

	var d,i,j int64;
	var temp byte
	d = -1
	for i = int64(0); i<gcd(ln,k); i++ {
		j=i;
		temp = nums[i];
		for {
			d=(j+k)%ln;
			if(d==i) {
				break;
			}
			nums[j] = nums[d];
			j=d;
		}
		nums[j]=temp;
	}
}
func main(){
	rs := bufio.NewReader(os.Stdin)
	ws := bufio.NewWriter(os.Stdout)

	input, _ := rs.ReadString('\n')
	mission := []byte(strings.TrimSpace(input))

	input, _ = rs.ReadString('\n')
	m, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)


	
	for i := int64(0); i < m; i++ {
		ch := readOneLineInts(rs, 3);
		l := <-ch; r := <-ch; k := <-ch;
		l--; r--;
		rotate(mission[l:r+1], k)
	}

	ws.WriteString(string(mission))

	ws.Flush()
}


