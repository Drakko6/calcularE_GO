package main

import "fmt"

func main(){

	var n int
	fmt.Scan(&n)
	var e float64 = 0

    for i:=0; i<n; i++{
		var j float64 = float64(i)
		
        var factorial float64 = 1 
        for j > 0 {
            factorial = factorial * j
			j = j- 1
		}
        e = e + 1/factorial
	
	}
	fmt.Println(e)

}