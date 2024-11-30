package main  

import (  
    "os"  
    "testing"  
)  

func TestCountLines(t *testing.T) {  
    // Создаем временный файл с тестовым содержимым  
    tempFile, err := os.CreateTemp("", "testfile.txt")  
    if err != nil {  
        t.Fatal(err)  
    }  
    defer os.Remove(tempFile.Name()) // чистим после теста  

    content := []byte("Line 1\nLine 2\nLine 3\n")  
    if _, err := tempFile.Write(content); err != nil {  
        t.Fatal(err)  
    }  
    tempFile.Close()  

    // Проверка количества строк  
    lineCount, err := countLines(tempFile.Name())  
    if err != nil {  
        t.Fatal(err)  
    }  

    if lineCount != 3 {  
        t.Errorf("expected 3 lines, got %d", lineCount)  
    }  
}
