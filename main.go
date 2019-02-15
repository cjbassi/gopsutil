package main

import (
    "fmt"

    "github.com/shirou/gopsutil/host"
)

// // type Errors []error

// func (errors []error) Error() string {
// 	s := "["
// 	for _, err := range errors {
// 		if err != nil {
// 			s += err.Error() + " "
// 		}
// 	}
// 	if len(s) > 1 {
// 		s = s[:len(s)-1]
// 	}
// 	return s + "]"
// }

func main() {
    // var errors []error
    // errors = append(errors, fmt.Errorf("asdf"))
    // fmt.Println(errors)
    temperatures, err := host.SensorsTemperatures()
    fmt.Println(temperatures)
    fmt.Println(err)
}
