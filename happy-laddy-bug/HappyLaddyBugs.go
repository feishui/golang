package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the happyLadybugs function below.
func happyLadybugs(b string) string {
    var exist [26]uint8
    underscore := false

    for i := 0; i < len(b); i++ {
        switch {
        case b[i] == '_': 
            underscore = true
        case exist[b[i] - 'A'] <= 1:
            exist[b[i] - 'A']++
        }
    }
    if !underscore {
        cnt := 1
        prv := b[0]
        for i := 1; i < len(b); i++ {
            if b[i] != prv {
                if cnt < 2 {
                    return "NO"
                }
                prv, cnt = b[i], 0
            }
            cnt++
        }
        if cnt == 1 {
            return "NO"
        }
        return "YES"
    }
    for _, e := range exist {
        if e == 1 {
            return "NO"
        }
    }
    return "YES"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    gTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    g := int32(gTemp)

    for gItr := 0; gItr < int(g); gItr++ {
        // nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        _, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        // n := int32(nTemp)

        b := readLine(reader)

        result := happyLadybugs(b)

        fmt.Fprintf(writer, "%s\n", result)
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
