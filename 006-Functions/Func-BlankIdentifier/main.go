package main

import "fmt"

func rect(L, W float64) (float64, float64) {
	var area = L * W
	var perimeter = (L + W) * 2
	return area, perimeter
}
func main() {
	var l float64 = 12.0
	w := 23.0
	area, _ := rect(l, w) // here we used 'blank identifier' to avoid perimeter return value
	fmt.Println("The area and Perimeter is", area)

}
