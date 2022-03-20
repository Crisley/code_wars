package desafios

import "fmt"

func PrinterError(s string) string {
	noErrors := 0
	errors := 0

	for _, v := range s {
		if v > 'm' {
			errors++
		}
		noErrors++
	}

	result := fmt.Sprint(errors) + "/" + fmt.Sprint(noErrors)

	return result

}
