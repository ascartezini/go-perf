package main

import (
	"sync"
	"unsafe"
)

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var personPool = sync.Pool{
	// New creates an object when the pool has nothing available to return.
	// New must return an interface{} to make it flexible. You have to cast
	// your type after getting it.
	New: func() interface{} {
		// Pools often contain things like *bytes.Buffer, which are
		// temporary and re-usable.
		return &Person{}
	},
}

func readableStream() <-chan Person {
	c := make(chan Person, 100)

	go func() {
		b := make([]byte, 0, 10)
		for i := 0; i < 1e6; i++ {
			b = append(b, "Andre-"...)
			// b = append(b, strconv.Itoa(i)...)
			person := Person{Id: +i, Name: *(*string)(unsafe.Pointer(&b))}
			// person := personPool.Get().(*Person)
			// person.Id = +i
			// person.Name = fmt.Sprintf("Andre-%d", i)
			// data, err := json.Marshal(&person)

			// if err != nil {
			// 	fmt.Println(err)
			// 	break
			// }

			c <- person
			// personPool.Put(person)
			b = b[:0]
		}
		close(c)
	}()

	return c
}

// func transform(data <-chan string) <-chan string {

// 	c := make(chan string, 100)

// 	go func() {
// 		c <- "id,name\n"

// 		for v := range data {
// 			person := Person{}
// 			json.Unmarshal([]byte(v), &person)
// 			// transformedData := fmt.Sprintf("%d,%s\n", person.Id, strings.ToUpper(person.Name))
// 			transformedData := fmt.Sprintf("%d\n", person.Id)

// 			c <- transformedData
// 		}
// 		close(c)
// 	}()

// 	return c
// }

// func writableStream(data <-chan string, w *bufio.Writer, done chan bool) {
// 	go func() {
// 		for v := range data {
// 			w.WriteString(v)
// 		}
// 		done <- true
// 	}()
// }

// func merge(cs ...<-chan string) <-chan string {
// 	var wg sync.WaitGroup
// 	out := make(chan string, 100)

// 	// Start an output goroutine for each input channel in cs.  output
// 	// copies values from c to out until c is closed, then calls wg.Done.
// 	output := func(c <-chan string) {
// 		for n := range c {
// 			out <- n
// 		}
// 		wg.Done()
// 	}
// 	wg.Add(len(cs))
// 	for _, c := range cs {
// 		go output(c)
// 	}

// 	// Start a goroutine to close out once all the output goroutines are
// 	// done.  This must start after the wg.Add call.
// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()
// 	return out
// }

// func main() {
// 	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
// 	// defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
// 	path, err := os.Getwd()

// 	if err != nil {
// 		fmt.Println("error when getting working directory")
// 	}

// 	b := strings.Builder{}
// 	b.WriteString(path)
// 	b.WriteString("/data.csv")

// 	file, err := os.Create(b.String())

// 	if err != nil {
// 		fmt.Printf("%s", err)
// 	}

// 	defer file.Close()

// 	rand.Seed(time.Now().UnixNano())
// 	start := time.Now()

// 	done := make(chan bool)
// 	c := readableStream()
// 	t1 := transform(c)
// 	t2 := transform(c)

// 	mergedTransform := merge(t1, t2)
// 	file.Sync()
// 	w := bufio.NewWriter(file)

// 	writableStream(mergedTransform, w, done)
// 	w.Flush()

// 	<-done

// 	close(done)

// 	elapsed := time.Since(start)
// 	fmt.Println(elapsed)
// }
