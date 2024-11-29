
package main

type StructName struct {
        x int8
};


func (s StructName) method(x int8) {
        println("method", x);
}

func main () {
        var s StructName;
        s.method(5);
}
