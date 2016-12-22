package main

import (
	"fmt"
	"log"
)

func main() {
	var weight float64
	var height float64
	//fmt.Println("Enter your weight(kg) and height(m): ")
	//fmt.Scanln(&weight, &height)
	fmt.Print("Enter your weight(kg): ")
	fmt.Scan(&weight)
	fmt.Print("Enter your height(m): ")
	fmt.Scan(&height)
	fmt.Printf("Your BMI is : %f\n", calculateBMI(weight, height))
}

func calculateBMI(weight float64, height float64) float64 {
	log.Println(weight)
	log.Println(height)
	BMI := weight / (height * height)
	return BMI
}
