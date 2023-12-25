package algo

import "fmt"

func CheckCard(numbers [16]uint8) bool {
	var sum uint8 = 0
	fmt.Println(numbers)
	for i := 0; i < 16; i++ {
		digit := numbers[i]
		if i%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	fmt.Println(sum)
	return (sum % 10) == 0
}
