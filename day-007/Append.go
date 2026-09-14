package main

// It is a built-in function that is used to add elements dynamically to a slice.
// If the underlying array is not large enough append() will create a new underlying array and point the slice to it.
// Append is variadic.

func main() {
	slice := []int{1, 2, 3, 4, 5}
	//Suppose we want to add some other numbers to the slice
	//Take this slice and add something ie 6 to it.
	slice = appendSlice(slice, 6)
	slice = append(slice, 7, 8)
	slice = append(slice, slice...)
}

func appendSlice(slice []int, elems ...int) []int {
	return append(slice, elems...)
}

// Example 2
type cost struct {
	day   int
	value float64
}

func getCostByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		//While the day is greater than the costs by day
		//If I encounter a day that I do not have its cost or anything append 0.0
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}

	return costsByDay
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

// One thing with append is that we append to the same slice and not create a new slice variable

// Java version of create Matrix
//public static [][] createMatrix(int rows, int cols ){
//	int [][] matrix = new int[rows][cols]
//	for(int i = 0; i < rows, i++){
//		for (int j = 0; j < cols, j++) {
//			matrix[i][j] = i * j
//		}
//	}

//	return matrix;
//}

// Major difference between go and java in that example is how we create a matrix
//int [][] matrix = new int[rows][cols] and make([][]int)
