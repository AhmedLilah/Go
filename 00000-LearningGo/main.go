package main

import (
	"fmt"
	// "io"
	"log"
	"math"
	"time"
	// "net/http"
	// "regexp"
	"sync"
)

type Image struct{}


// to create  a var outside a funciton it must be like that
var (	x string  = "x";
	w string = "w";
)

// struct 
type Vertex struct {
	x, y int
}

type Cartesian struct {
	x, y int
}

// Interfaces
type Printable interface {
	print()
}

func vertexFunction(v Vertex) (result float64) {
	result = math.Sqrt(float64(v.x * v.x + v.y * v.y))
	return
}

func (v Vertex) vertexMethod () (result float64) {
	result = math.Sqrt(float64(v.x * v.x + v.y * v.y))
	return
}

func (v Vertex) print () {
	fmt.Print("Vertex:\n\t", v.x, v.y, "\n\n") 
}

func (c Cartesian) print () {
	fmt.Print("Cartesian:\n\t", c.x, c.y, "\n\n") 
}

func handleError(param int) {
	fmt.Println("Deferred: ", param)
}

// functions as values
func transform(slice [] float64, fn func(float64) float64) {
	for i , _ := range slice {
		slice[i] = fn(slice[i])
	}
}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// Generics
type List[T any] struct {
	next *List[T]
	val  T
}



// Goroutines 
func routine(s string, c chan int){
	time.Sleep(time.Second)
	fmt.Println(s)
	c <- 5
	c <- 6
}


// URL fetcher

// URL tracker using a mutex
var (
	mu      sync.Mutex
	visited = make(map[string]bool)
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	if depth <= 0 {
		return
	}

	// Protect access to the visited map
	mu.Lock()
	if visited[url] {
		mu.Unlock()
		return
	}

	visited[url] = true
	mu.Unlock()


	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	
	// Spawn goroutines for the next level
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			Crawl(u, depth-1, fetcher)
		}(u)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}


// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// func (f *fakeFetcher) Fetch(url string) (body string, urls []string, err error) {
// 	// Gets the response form the url
// 	resp, err := http.Get(url)

// 	if err != nil {
// 		return
// 	}
// 	
// 	defer resp.Body.Close()

// 	// Reads the body from the response
// 	bodyint, err := io.ReadAll(resp.Body)

// 	if err != nil {
// 		return
// 	}

// 	body = string(bodyint)

// 	// Define a regex pattern to match URLs
// 	urlRegex := `https?://[^\s"'>]+`
// 	re := regexp.MustCompile(urlRegex)

// 	// Find all matching URLs
// 	urls = re.FindAllString(body, -1)

// 	return
// }




func main() {
	// to create a variable inside a function it could be both ways
	y := "y"
	var z = "z"
	fmt.Println(x, w)
	fmt.Println(y, z)
	fmt.Println()


	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)


	// For loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("Sum :", sum)
	fmt.Println()

	True := true
	times := 0
	for True == true {
		times += 1

		if times == 5 {
			True = false
		}
	} 

	fmt.Println("times :", times)

	for {
		fmt.Println(times)
		times += 1
		if times == 10 {
			break
		}
	}
	fmt.Println()


	// switch case
	switchVar := 5

	switch switchVar {
	case 1:
		fmt.Println("one")
	case int(2.5*2):
		fmt.Println("Switching on :", switchVar)
	default:
		fmt.Println("default case.")
	}
	fmt.Println()


	// Defer 
	deferParam :=5
	defer handleError(deferParam)
	deferParam = 10

	ptr := new(int)
	*ptr = 5

	defer func() {
		if ptr != nil {
			fmt.Println("Memory got freed")
			ptr = nil
		} else {
			fmt.Println("Pointer is already invalid")
		}
	}()
	fmt.Println()

	
	// Structs	
	vertex := Vertex {1, 2}
	vertexPtr := &vertex

	vertexPtr.x = 5
	vertexPtr.y = 10

	fmt.Printf("vertex : {%d, %d}\n", vertexPtr.x, vertexPtr.y)
	fmt.Println()


	// Arrays
	var arrayOfInts [3] int 
	arrayOfInts[0] = 1
	arrayOfInts[1] = 10
	arrayOfInts[2] = 100

	arrayOfStrings := [2]string{"Hello", "World"}
	
	fmt.Println("Array of ints:\n\t", arrayOfInts[0], arrayOfInts[1])
	fmt.Println()
	fmt.Printf("Array of strings:\n\t%s, %s!\n", arrayOfStrings[0], arrayOfStrings[1])
	fmt.Println()


	// Slices
	var sliceOfInts []int = arrayOfInts[0:2]
	sliceOfStrings := arrayOfStrings[0:1]

	fmt.Printf("Slices:\n\t%v %v\n", sliceOfInts, sliceOfStrings)
	fmt.Println()

	sliceOfInts[0] = 100

	fmt.Printf("Slices Modifying the Array:\n\t%v %v\n", sliceOfInts, arrayOfInts)
	fmt.Println()

	// Resizing slices
	slice := []int{1, 2, 3}

	fmt.Printf("Slices resizing:\n\t%v\n", slice)
	fmt.Println()

	for i:=0; i<5; i++ {
		slice = append(slice, i) 
	}

	fmt.Printf("Slices resizing:\n\t%v\n", slice)
	fmt.Println()

	slice = slice[:0]

	fmt.Printf("Slices resizing:\n\t%v\n", slice)
	fmt.Println()

	slice = slice[:cap(slice)]

	fmt.Printf("Slices resizing:\n\t%v\n", slice)
	fmt.Println()


	// Range based for loop
	for i, v := range slice {
		fmt.Printf("%d\t%d\n", i, v)
	}
	fmt.Println()


	// Mps
	var mapSS map[string]string
	mapSS = make(map[string]string)
	mapSS["name"] = "Ahmed S. Lilah"
	mapSS["age"]  = "25"

	for i, v := range mapSS {
		fmt.Printf("Key: %s, Value: %s\n", i, v)
	}
	fmt.Println()


	// functions as values
	sliceOfFloats := []float64{0, 1, 2, 3, 4, 5}
	fmt.Println(sliceOfFloats)
	fmt.Println()

	double := func(number float64) (result float64) {
		result = number * 2;
		return;
	}

	transform(sliceOfFloats, double)
	fmt.Println(sliceOfFloats)
	fmt.Println()


	// Functions vs Methods
	fmt.Printf("Vertex function : %f\n", vertexFunction(vertex))
	fmt.Printf("Vertex method   : %f\n", vertex.vertexMethod())
	fmt.Println()


	// Interfaces
	var printable Printable
	printable = vertex

	printable.print()


	// Type assertion 
	v, ok := printable.(Vertex)

	if ok {
		fmt.Printf("The interface %T hold the type %T and the return vlues are %v and %t\n\n", printable, vertex, v, ok)
	}


	// Type switches
	switch v := printable.(type) {
	case Cartesian:
		fmt.Printf("This inteface holds a value of type %T.\n\n", v)
	case Vertex:
		fmt.Printf("This inteface holds a value of type %T.\n\n", v)
	default:
		fmt.Printf("This inteface holds a value of unknown type.\n\n")
	} 
	
	
	// Type parameters
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Printf("Index of %v in %v is %v\n\n", 15, si, Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Printf("Index of %v in %v is %v\n\n", "hello", ss, Index(ss, "hello"))


	// Generics
	list := new(List[float64])
	list.val = 5.0
	list.next = new(List[float64])
	list.next.val = 6.0

	fmt.Printf("Head :  %p, value : %v, next : %p, next vaue : %v, next next : %p\n\n", list, list.val, list.next, list.next.val, list.next.next)


	// Goroutine
	c := make(chan int) 
	go routine("Routine result\n", c)
	fmt.Println("Main thread\n")
	arc := <- c *  <- c
	fmt.Println("the value calculated after the satrt of the routine", arc)
	fmt.Println()

	Crawl("https://golang.org/", 10, fetcher)
}
