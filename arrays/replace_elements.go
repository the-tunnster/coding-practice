/*
Leet Code #1299

Given an array arr, replace every element in that array with the greatest element among the
elements to its right, and replace the last element with -1.
After doing so, return the array.
*/

package arrays

func ReplaceElements(arr []int) []int {
	result := make([]int, len(arr))

	max := arr[len(arr)-1]
	arr[len(arr)-1] = -1

	for i := len(arr)-2; i >=0; i--{
		result[i] = max
		if arr[i] > max{
			max = arr[i]
		}
	}

	return (result)
}
