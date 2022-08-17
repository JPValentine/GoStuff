package main

func Sum(numbers []int)int{
	sum := 0
	for _, number := range numbers{
		sum+=number
	}//for
	return sum
}//sum
