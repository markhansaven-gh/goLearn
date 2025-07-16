package main

func main() {

	// Int
	var a int = 5
	var b int = 6
	println(a + b)

	// String
	var myStr string = "Hihihi!"
	println(myStr)

	// Bool
	var isTrue bool = false
	println(isTrue)

	// Array and for loop
	var arr = []string{"hi", "there", "how", "are", "you", "doing", "today?"}
	for x := 0; x < len(arr); x++ {
		println(arr[x])
	}

}
