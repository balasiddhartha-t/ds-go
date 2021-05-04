//program to find minimum element in a sorted and rotated array

package main

import "fmt"

func MinSortedArray(incomingArray []int, start int, end int) int{

	mid := start+ (start+end/2)
	
	if (incomingArray[mid] > incomingArray[mid+1]){
		return incomingArray[mid];
	}
	if (mid > start && incomingArray[mid] < incomingArray[mid - 1]){
		return incomingArray[mid];
	}

	if (incomingArray[end] > incomingArray[mid]){
		return MinSortedArray(incomingArray, start, mid-1);

	}

	MinSortedArray(incomingArray, mid+1, end);
	return 0;

} 


func main() {
	testArray:= []int{4,5,1,2,3}
	testArrayLength := len(testArray)
	answer := MinSortedArray(testArray, 0, testArrayLength);
	fmt.Println("The Answer is:", answer);

	testArray2:= []int{4,5,6,7,2,3}
	testArrayLength2 := len(testArray2)
	answer2 := MinSortedArray(testArray2, 0, testArrayLength2);
	fmt.Println("The Second Answer is:", answer2);
}