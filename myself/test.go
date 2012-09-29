package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"io"
	"flag"
	"bufio"
	"errors"
	"math"
)

var numberFlag = flag.Bool("n", false, "number each line")

type Address struct {
	Type	string
	City	string
	Country	string
}

type VCard struct {
	FirstName	string
	LastName	string
	Addresses	[]*Address
	Remark		string
}

type Page struct {
	Title	string
	Body	[]byte
}

func (p *Page) Save() error{
	fmt.Println(p.Title)
	fmt.Println(p.Body)
	if err := ioutil.WriteFile(p.Title, p.Body, 0644); err != nil {
		fmt.Println(err)
	}
	return nil
}

func LoadFile(p *Page)  {
	if buf, err := ioutil.ReadFile(p.Title); err != nil {
		fmt.Println(err)
	} else {
		p.Body = buf
	}
}

func cat(r *bufio.Reader)  {
	i := 1
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *numberFlag {
			fmt.Fprintf(os.Stdout, "%d %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
	return
}

func badCall() {
	i := 0
	fmt.Println(1/i)
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicking %s\n", e)
		}
	}()
	badCall()
	fmt.Println("send bad call")
}
func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printling in g", i)
	g(i + 1)
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func CovertInt64ToInt(i int64) int {
	if i < math.MinInt32 || math.MaxInt32 < i {
		panic(fmt.Sprintf("Cant't change %v to int", i))
	}
	return int(i)
}

func IntFromInt64(i int64) (r int, err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			r = 0
			err = errors.New("Can't change")
		}
	}()
	return CovertInt64ToInt(i), nil
}

func sendData(done chan bool, data chan int)  {
	for i := 0; i < 10; i++ {
		data <- i * 10
	}
	close(data)
	done <- true
}

func recvData(done chan bool, data chan int)  {
	for i := range data {
		fmt.Println("Recv", i)
	}

	done <- true
}

func buildLazy(fn func(uint64, uint64) uint64, i1, i2 uint64) func() uint64 {
	retChan := make(chan uint64)

	go func() {
		v1, v2 := i1, i2
		for	{
			v1, v2 = v2, fn(v1, v2)
			retChan <- v1
		}
	}()

	return func() uint64 {
		return <-retChan
	}
}

func main() {
	fib := buildLazy(func(i, j uint64) uint64 {
		return i + j
	}, 0, 1)

	for i := 0; i < 100; i++ {
		fmt.Println(fib())
	}
}
