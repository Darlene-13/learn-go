package main

// The length and size of an array slice may change as long as it fits within the limits of the  underlying array.
// The capacity of the slice accessible by the built-in function (cap), reports the maximum length a slice may assume.

// Append Function to append data to a slice if the data exceeds the capacity the slice is re-allocated and the resulting slice is returned
// Slices and arrays are stored in contiguous memory.

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l+len(data) > cap(slice) { //re allocate.
		//Allocate double what's needed
		newSlice := make([]byte, (l+len(data))*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0 : l+len(data)]
	copy(slice[l+len(data):], slice)
	return slice
}

func main() {

}
