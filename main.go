package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// new comment
func main() {

	var math_operation_sum, math_operation_sub, math_operation_mult, math_operation_div, argument1, argument2, roman_indicator uint
	var result int
	var input_text string
	var sliced_text []string
	var sliced_text_part1, sliced_text_part2 string
	map_arabic := map[string]uint{"1": 1, "10": 10, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	map_roman_to_arabic := map[string]uint{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	map_arabic_to_roman := map[uint]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X"}
	map_arabic_to_roman_dec := map[uint]string{1: "X", 2: "XX", 3: "XXX", 4: "XL", 5: "L", 6: "LX", 7: "LXX", 8: "LXXX", 9: "XC", 10: "C"}

	for {
		math_operation_sum, math_operation_sub, math_operation_mult, math_operation_div, roman_indicator = 0, 0, 0, 0, 0

		fmt.Println("Введите арифметическое выражение:")

		input_text, _ = bufio.NewReader(os.Stdin).ReadString('\n')

		cleared_input_text := strings.ReplaceAll(input_text, " ", "")
		cleared_input_text = cleared_input_text[:len(cleared_input_text)-1]

		separator_check := strings.Contains(input_text, "+")
		if separator_check == true {
			math_operation_sum = 1
			sliced_text = strings.Split(cleared_input_text, "+")

		}

		separator_check = strings.Contains(input_text, "-")
		if separator_check == true {
			math_operation_sub = 1
			sliced_text = strings.Split(cleared_input_text, "-")
		}

		separator_check = strings.Contains(input_text, "*")
		if separator_check == true {
			math_operation_mult = 1
			sliced_text = strings.Split(cleared_input_text, "*")
		}

		separator_check = strings.Contains(input_text, "/")
		if separator_check == true {
			math_operation_div = 1
			sliced_text = strings.Split(cleared_input_text, "/")
		}

		if math_operation_sum+math_operation_sub+math_operation_mult+math_operation_div != 1 {
			fmt.Println("Ошибка, выражение должно содержать одно арифмитеческое действие с положительными числами")
			os.Exit(0)
		}
		if len(sliced_text) != 2 {
			fmt.Println("Ошибка, выражение должно содержать два аргумента")
			os.Exit(0)
		}
		sliced_text_part1, sliced_text_part2 = sliced_text[0], sliced_text[1]

		argument1 = map_arabic[sliced_text_part1]
		argument2 = map_arabic[sliced_text_part2[:len(sliced_text_part2)-1]]

		if argument1*argument2 == 0 {
			argument1 = map_roman_to_arabic[sliced_text_part1]
			argument2 = map_roman_to_arabic[sliced_text_part2[:len(sliced_text_part2)-1]]

			roman_indicator = 1
		}
		if argument1*argument2 == 0 {
			fmt.Println("Ошибка аргументов выражения")
			os.Exit(0)
		}

		if math_operation_div == 1 {
			result = int(argument1 / argument2)
		} else {
			result = int((argument1+argument2)*math_operation_sum + (argument1-argument2)*math_operation_sub + (argument1*argument2)*math_operation_mult)

		}
		if roman_indicator == 1 {
			if result < 1 {
				fmt.Println("Недопустимые значения аргументов, результат меньше I")
			} else {
				if result < 10 {
					fmt.Println("Результат:", map_arabic_to_roman[uint(result)])
				}

				if result > 10 {
					fmt.Println("Результат:", map_arabic_to_roman_dec[uint(result/10)]+map_arabic_to_roman[uint((result-10*int(result/10)))])
				}
			}

		} else {
			fmt.Println("Результат:", result)
		}
	}

}
