package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
We define  to be a permutation of the first  natural numbers in the range . Let  denote the value at position in permutation  using -based indexing.

 is considered to be an absolute permutation if  holds true for every .

Given  and , print the lexicographically smallest absolute permutation . If no absolute permutation exists, print -1.

For example, let  giving us an array . If we use  based indexing, create a permutation where every . If , we could rearrange them to :

pos[i]	i	|Difference|
3	1	2
4	2	2
1	3	2
2	4	2
Function Description

Complete the absolutePermutation function in the editor below. It should return an integer that represents the smallest lexicographically smallest permutation, or  if there is none.

absolutePermutation has the following parameter(s):

n: the upper bound of natural numbers to consider, inclusive
k: the integer difference between each element and its index
Input Format

The first line contains an integer , the number of test cases.
Each of the next  lines contains  space-separated integers,  and .

Constraints


1 <= t <= 10
1 <= n <= 10**5
0 <= k < n

8
6 0
n=6, k=0
result=[0 1 2 3 4 5]
6 1
n=6, k=1
result=[2 1 4 3 6 5]
6 2
n=6, k=2
result=[]
6 3
n=6, k=3
result=[4 5 6 1 2 3]
6 4
n=6, k=4
result=[]
6 5
n=6, k=5
result=[]
*/
// Complete the absolutePermutation function below.
func absolutePermutation(n int32, k int32) []int32 {
	size := int(n)
	result := make([]int32, size)
	for i := 0; i < size; i++ {
		result[i] = int32(i + 1)
	}
	if k == 0 {
		return result
	}
	if n == 1 {
		return result[0:0]
	}
	i, j := int32(0), int32(0)
	for ; i <= n-k*2; i = i + k*2 {
		for j = i; j < i+k; j++ {
			result[j], result[j+k] = result[j+k], result[j]
		}
	}
	if j+k >= n {
		return result
	}
	return result[0:0]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nk := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nk[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		kTemp, err := strconv.ParseInt(nk[1], 10, 64)
		checkError(err)
		k := int32(kTemp)

		fmt.Printf("n=%d, k=%d\n", n, k)
		result := absolutePermutation(n, k)
		fmt.Printf("result=%v\n", result)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		if len(result) == 0 {
			fmt.Fprintf(writer, "-1")
		}
		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
