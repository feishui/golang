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
input:
3 3
1 3 4
2 2 3
1 2 4

correct output:
60

*/
// Complete the surfaceArea function below.
func surfaceArea(A [][]int32) int32 {
    rs, cs := int32(len(A)), int32(len(A[0]))

    surface := int32(0)
    for r := int32(0); r < rs; r++ {
        surface = surface + A[r][0]
        for c := int32(1); c < cs; c++ {
            higher := A[r][c] - A[r][c-1]
            if higher > 0 {
                surface = surface + higher
            } else {
                surface = surface - higher
            }
        }
        surface = surface + A[r][cs-1]
    }
    for c := int32(0); c < cs; c++ {
        surface = surface + A[0][c]
        for r := int32(1); r < rs; r++ {
            higher := A[r][c] - A[r-1][c]
            if higher > 0 {
                surface = surface + higher
            } else {
                surface = surface - higher
            }
        } 
        surface = surface + A[rs-1][c]
    }

    return surface + rs * cs * 2
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    HW := strings.Split(readLine(reader), " ")

    HTemp, err := strconv.ParseInt(HW[0], 10, 64)
    checkError(err)
    H := int32(HTemp)

    WTemp, err := strconv.ParseInt(HW[1], 10, 64)
    checkError(err)
    W := int32(WTemp)

    var A [][]int32
    for i := 0; i < int(H); i++ {
        ARowTemp := strings.Split(readLine(reader), " ")

        var ARow []int32
        for _, ARowItem := range ARowTemp {
            AItemTemp, err := strconv.ParseInt(ARowItem, 10, 64)
            checkError(err)
            AItem := int32(AItemTemp)
            ARow = append(ARow, AItem)
        }

        if len(ARow) != int(int(W)) {
            panic("Bad input")
        }

        A = append(A, ARow)
    }

    result := surfaceArea(A)

    fmt.Fprintf(writer, "%d\n", result)

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
