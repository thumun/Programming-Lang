package main

func main() {
	// 1.) Why is rrslice less efficient than rslice
		// In the first slice is faster because the slice is initialized up front 
		// so there is need allocate more memory which is a faster process.
		
		// In the second slice, the slice size is changed after each append which 
		// is a slower process. 

	// 2.) What is r3slice said to suck?
		// Because the memory is allocated and the data is appended so the array 
		// is twice the size it is supposed to be. (the allocated data is just 
		// empty)

	// 3.) How would you fix r3slice?
		// by either removing the size initalization or removing the append (using
		// same idea as the first method)
}
