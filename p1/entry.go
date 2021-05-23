package main

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/klyvechen/go-practice-2/p0"
	"math"
)
var x int = 2
func b() {
	fmt.Println(x)
}


func a() {
	var x int = 3
	fmt.Println(x)
}

type Grade int

const (
	aa Grade = iota
	bb
	cc
	dd
)


type MyInt int

func (m MyInt) Double() int { 
	return int(m*2)
}

func (m *MyInt) Double2() int { 
	return int(*m*2)
}

type Pointer struct {
	x int
	y int
}

func (p *Pointer) toOrigin() float64 {
	return math.Sqrt(math.Pow(float64(p.x), 2) + math.Pow(float64(p.y), 2))
}

func (p Pointer) toOrigin2() float64 {
	return math.Sqrt(math.Pow(float64(p.x), 2) + math.Pow(float64(p.y), 2))
}

func main() {

	data.PrintX()

	var mi MyInt = 3
	fmt.Print("mi ")
	fmt.Println(mi.Double());
	fmt.Print("mi ")
	fmt.Println(MyInt(6).Double())

	// var mi2 *MyInt

	// *mi2 = 4

	fmt.Print("mi ref ")
	fmt.Println(mi.Double2())

	p := Pointer{x: 1, y: 2}


	fmt.Print("p ")
	fmt.Println(p.toOrigin())


	fmt.Print("p ref ")
	fmt.Println(p.toOrigin2())




	type FloatType float64

	var i int
	var k int = 3
	var f FloatType = 2.7
	var x int = 100
	var y = 100

	var ip *int = &k

	fmt.Println(i)
	fmt.Println(*ip)

	ip = &x

	i = *ip

	s := strconv.Itoa(i+x+y+int(f))
	fmt.Printf(s)
	fmt.Println()
	fmt.Println(i)
	fmt.Println(*ip)
	a()
	b()

	fmt.Println(bb)
	fmt.Println(aa)

	i = 0

	for {
		fmt.Println("hi")
		if i == 10 {
			break
		}
		i++
	}
	i= 0

	switch i {
	case 0:
		fmt.Printf("Is %d\n", 0)
	case 1:
		fmt.Printf("Is %d\n", 1)
	default:
		fmt.Printf("Is %d\n", 2)		
	}

	var ss string

	fmt.Printf("The answer is: ")
	for i = 0; i < 0; i++ {
		fmt.Scan(&ss)
		fmt.Printf("The read in word is %s\n", ss)
	}

	var xtemp int
	x1 := 0
	x2 := 1
	for x:=0; x<5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)

	ary := [...]string{"a", "b", "c", "d"}
	sl1 := ary[0:1]
	fmt.Printf("%d, %d\n", len(sl1), cap(sl1))

	sl1 = append(sl1, "e")
	fmt.Print(sl1);
	fmt.Print(ary);
	for i = 0; i< 6; i++ {
		sl1 = append(sl1, "e2")
	}
	fmt.Println(sl1);
	fmt.Println(ary);

	var idMap map[string]int
	idMap = make(map[string]int)
	idMap["myfirstId"] = 0
	idMap["mySecondId"] = 1
	fmt.Println(idMap["mySecondId"])
	
	idMap2 := map[string]int{"test": 3}
	// idMap = make(map[string]int)
	idMap2["myfirstId"] = 0
	idMap2["mySecondId"] = 1
	fmt.Println(idMap2["test"])

	for key, val := range idMap2 {
		fmt.Println(idMap2[key], val)
	}

	x9 := []int {4, 8, 5}
  	y1 := -1
  	for _, elt := range x9 {
		  fmt.Printf("%d \n", elt)
    if elt > y1 {
      y1 = elt
    }
  }
  fmt.Print(y)
  cnt, err := http.Get("http://www.uci.edu")
  if err != nil {
	fmt.Print(cnt.Body)
	}
}
