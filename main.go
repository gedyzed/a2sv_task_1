package main

import (
	"fmt"
)

type Subject struct {
	name string 
	grade float32
}

type Student struct {
	name string
	noOfSubjects int
	subjects []Subject
	averageGrade float32
}
var student Student

func takeInput(){

	fmt.Print("Enter Your Name : ")
	fmt.Scan(&student.name)
	fmt.Print("Enter Number of Subjects : ")
	fmt.Scan(&student.noOfSubjects)

    for i := 0; i < student.noOfSubjects; i++{
	
		fmt.Printf("\n%d. subject name : ", i + 1)
		
		var name string
		fmt.Scan(&name)

		fmt.Printf("%d. subject grade : ", i + 1)
		var grade float32
		fmt.Scan(&grade)

		subject := Subject{name, grade}
		student.subjects = append(student.subjects, subject)

		fmt.Println()
	} 
}

func calculateAverage(subjects []Subject) float32 {
	
	var sum float32
	for _, subject := range subjects {
		sum += subject.grade
	}

	noOfSubjects := student.noOfSubjects

	return (sum / float32(noOfSubjects))

}

func displayResult(student Student){

	fmt.Printf("Name :  %s", student.name)
	fmt.Println("\n======================================")
	fmt.Println("Grades")
	fmt.Println("======================================")
	for _, subject := range student.subjects {

		fmt.Printf("Subject Name :  %s\n", subject.name)
		fmt.Printf("Grade : %f", subject.grade)

		fmt.Println()
	}
	fmt.Println("======================================")
	fmt.Printf("Average Grade : %f\n\n", student.averageGrade)

}


func main(){

	takeInput();
	avg := calculateAverage(student.subjects);
	student.averageGrade = avg
	displayResult(student)
}