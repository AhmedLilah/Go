package main

import "fmt"

func main() {
        var x int8 = 127; 
        const y uint64 = 758903475;
        var z uint8 = 255;
        fmt.Println("Hello, World!");
        fmt.Print(x, y, z, "\n");
        complain();
}

func complain() {
        fmt.Println("WTF is PrintLn() spelled like that?!!");
}
