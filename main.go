package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ajaykujh/gofun/controllers"
)

const (
	first = iota + 1.5 // iota starts from 0 to 1,2,3,4,.......: you can also specify
	//like iota + 0.5 and it will keep that for further evaluation
	second
)

const (
	third = iota //iota would start from 0 here as new const block
)

func main() {

	var i int
	i = 42
	name := "Aaa"
	var f float32 = 3.14
	ptr := &f
	c := complex(3, 4)
	const pi = 3.1415
	fmt.Println(i, f, name, c, ptr, *ptr, pi, first, second, third)

	//Arrays
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr, arr[2])

	//composite declaration
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	// slices: using make command
	var arr3 = make([]string, 5, 10)

	arr3 = append(arr3, "str1", "str2")
	fmt.Println(arr3)                 // append always appends after the max size, you are supposed to use arr[0] = 1 to assign from 0 to 4
	fmt.Println(len(arr3), cap(arr3)) //prints length and capacity

	//array itself is slice
	arr4 := []string{"str1", "str2", "str3"} //note that you have not sized the array, [] is empty. The compiler will treat this as slice
	//and append works. Initially it would assign capacity 3 to the array, every time append is triggered, the capacity would double
	fmt.Println(arr4)

	arr4 = append(arr4, "str4", "str5")
	fmt.Println(arr4)
	fmt.Println(len(arr4), cap(arr4)) //prints length and capacity

	arr4 = append(arr4, "str6", "str7")
	fmt.Println(len(arr4), cap(arr4)) //prints length and capacity: watch the capacity is doubled to the current capacity
	// everytime the length limit is hit

	//create a slice of slice
	arr5 := arr4[2:5]
	fmt.Println(arr5) //pure python list like
	arr5 = append(arr5, "str-appended")
	fmt.Println(arr5) //This is nothing but reference to the original array, so str-appended got added to the orihginal abive one :)
	fmt.Println(len(arr5), cap(arr5))

	fmt.Println(arr4) // will show that str-appended actually replaced arr6

	//appending a slice to a slice
	arr6 := []string{"str10", "str11", "str13"}
	arr4 = append(arr4, arr6...)
	fmt.Println(arr4)

	//loops
	courses := []string{"A", "B", "C"}
	for _, i := range courses {
		fmt.Println(i)
	}

	testMap := map[string]int{"A": 1, "B": 2, "C": 3}

	for k, v := range testMap {
		fmt.Println(k, v)
	}

	//Update map
	testMap["A"] = 100

	//delete map
	delete(testMap, "C")

	fmt.Println(testMap)

	//goroutines
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Print func1")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("Print func2")
	}()

	waitGrp.Wait()

	//Sample WebServer
	controllers.RegisterControllers()
	http.ListenAndServe(":3040", nil)
}
