package main

import "fmt"

func main() {
	examScore := 90
	absenceScore := 80

	isExamPass := examScore > 80
	isAbsencePass := absenceScore > 80

	isGraduate := isExamPass && isAbsencePass

	fmt.Println("Graduate: ", isGraduate) // false
}
