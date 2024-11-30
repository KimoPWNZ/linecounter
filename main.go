package main  

import (  
    "bufio"  
    "fmt"  
    "os"  
)  

func countLines(filePath string) (int, error) {  
    file, err := os.Open(filePath)  
    if err != nil {  
        return 0, err  
    }  
    defer file.Close()  

    scanner := bufio.NewScanner(file)  
    lineCount := 0  
    for scanner.Scan() {  
        lineCount++  
    }  

    if err := scanner.Err(); err != nil {  
        return 0, err  
    }  
    return lineCount, nil  
}  

func main() {  
    if len(os.Args) < 2 {  
        fmt.Println("Usage: linecounter <file>")  
        return  
    }  

    lineCount, err := countLines(os.Args[1])  
    if err != nil {  
        fmt.Println("Error:", err)  
        return  
    }  
    fmt.Printf("Line count: %d\n", lineCount)  
}
