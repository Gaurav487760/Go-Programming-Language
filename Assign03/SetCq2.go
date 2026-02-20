package main
import "fmt"

type Employee struct {
	eno    int
	ename  string
	salary float64
}

func main() {
	var n int
	fmt.Print("Enter number of employees: ")
	fmt.Scan(&n)

	employees := make([]Employee, n)

	for i := 0; i < n; i++ {
		fmt.Println("\nEnter details for employee", i+1)
		fmt.Println("Enter Number of employee")
		fmt.Scan(&employees[i].eno)
		fmt.Println("Enter Name of employee")
		fmt.Scan(&employees[i].ename)
		fmt.Println("Enter Salary of employee")
		fmt.Scan(&employees[i].salary)
	}

	maxSalary := employees[0].salary

	for i := 1; i < n; i++ {
		if employees[i].salary > maxSalary {
			maxSalary = employees[i].salary
		}
	}
	fmt.Println("\nEmployees with Maximum Salary:")
	for i := 0; i < n; i++ {
		if employees[i].salary == maxSalary {
			fmt.Println("Employee No:", employees[i].eno)
			fmt.Println("Name:", employees[i].ename)
			fmt.Println("Salary:", employees[i].salary)
			fmt.Println()
		}
	}
}