package main

import "os"
import "fmt"
import "flag"
import "bufio"
import "strconv"
import "sort"



func b_search(array []int, num int, start int, end int) int{
mid := start
if end >= start{
	mid = start + ( end - start ) / 2
	if array[mid] == num {
		return mid
	}
	if array[mid] > num {
		//if mid-1 >= start {
		return b_search(array,num, start, mid-1)
		//}
	} else {
		//if mid+1 <= end {
		return b_search(array,num, mid+1, end)
		//}
	}
}
/*
if mid > end {
	mid = end
} else if mid < start {
	mid = start
}
*/
return mid
}

func check_and_update(top []int,num int,N int) int {
var i int
if num == top[N-1] {
	return 0
}
if num > top[N-1] {
	for i = 1; i < N; i++ {
		top[i-1] = top[i]
	}
	top[N-1] = num
return 0
}
if num <= top[0] {
	return 0
}
if num > top[0] && num < top[N-1] {
	pos := 0
	pos = b_search(top, num, 0 , N-1)
	if top[pos] == num {
		return 0
	}
	if pos == 0{
		top[pos] = num
	} else {
		for i = 0; i < pos-1; i++ {
			top[i] = top[i+1]
		}	
		top[pos -1] = num
	}
	
}
return 0
}

func main(){

file_name := flag.String("input", "", "Input filename (Compulsory)")
top_N := flag.Int("N",5,"The count of highest numbers (Optional)")

/*
var svar string
flag.StringVar(&svar, "svar", "bar", "a string var")
*/
flag.Parse()

if *file_name == "" || *top_N == -1{
	flag.Usage()
	os.Exit(1)
}

N := *top_N;

//fmt.Println("File name", *file_name)
//fmt.Println("N ", N)

//top to hold N numbers
top :=  make([]int,N)

inpFile, err := os.Open(*file_name);
if err != nil {
	fmt.Println(err)
}
defer inpFile.Close()
scanner := bufio.NewScanner(inpFile)

//fmt.Println("len(top) ", len(top))

fill := 0
numLines := 0
var num int
last_pos := 0
for scanner.Scan() {
	//fmt.Println(scanner.Text())
	//num,_ = strconv.ParseInt(scanner.Text(), 10, 32)
	num,_ = strconv.Atoi(scanner.Text())
	//fmt.Println(num)

	numLines++
	if numLines <= N {
		pos := 0
		if numLines > 1 {
			pos = b_search(top,num,0, fill - 1)
		//	pos = b_search(top,num,0, fill)
				//fmt.Println("fill = " , fill)
				//fmt.Println("pos = " , pos)
				//fmt.Println("numl = " , numLines)
/*
			if pos > N-1 {
				fmt.Println("Oh yes");
				//fmt.Println(top[N-1], num)
				pos = N-1
			}

			if pos < 0 {
				fmt.Println("Oh yes 2");
				pos = 0
			}
*/
			if top[pos] == num {
				continue;
			}
		}
		last_pos = fill
		fill++ 
		top[last_pos] = num
		if numLines == N {
			sort.Ints(top)
		}
		
	} else {
		check_and_update(top,num,N)
	}
}
fmt.Println("Top N -----")
for fill = N-1; fill >= 0; fill-- {
	fmt.Println(top[fill]);
}
}

