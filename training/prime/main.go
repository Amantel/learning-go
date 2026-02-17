// Уровень сложности: сложный

// Вспомним о простых числах: простым числом считается такое натуральное число, которое делится нацело только на два числа: на единицу и на само себя.

// Отрицательные числа, 0 и 1 не являются простыми.

// Простыми числами являются 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, …

// Обратите внимание, что есть простые числа, разность которых равна 2: например 17 и 19, 29 и 31.

// Такие простые числа называются простыми числами близнецами.

// Есть гипотеза, что их бесконечно много, но точно мы этого не знаем.

// Задача:

// Написать функцию IsPrimeTwin(n int) bool

// Функция должна принимать число и возвращать true, если число является простым близнецом, и false во всех остальных случаях . Например: IsPrimeTwin(10) должно вернуть false, т.к. 10 – не простое.

// Например: IsPrimeTwin(23) должно вернуть false, т.к. число 23 – простое, но его «соседи» 21 и 25 – не простые.

// Например: IsPrimeTwin(29) должно вернуть true, т.к. 29 – простое и его «старший сосед» 31 – тоже простое.

// Например: IsPrimeTwin(19) должно вернуть true, т.е. число 19 – простое и его «младший сосед» 17 – тоже простое

package main

import (
	"fmt"
)

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}

	}

	return true
}

func IsPrimeTwin(n int) bool {
	if n < 3 {
		return false
	}

	if !IsPrime(n) {
		fmt.Println(n, "is NOT prime")
		return false
	}
	fmt.Println(n, "is prime")

	prev := IsPrime(n - 2)
	fmt.Println(n-2, "prev is Prime", prev)
	next := IsPrime(n + 2)
	fmt.Println(n+2, "next is Prime", next)
	if prev || next {
		fmt.Println("!!!!! THIS IS PRIME TWIN")
	}
	fmt.Println("--------")

	return prev || next
}

func main() {
	fmt.Println(">>>>")
	fmt.Println(IsPrimeTwin(-1))
	fmt.Println(IsPrimeTwin(0))
	fmt.Println(IsPrimeTwin(1))
	fmt.Println(IsPrimeTwin(2))
	fmt.Println(IsPrimeTwin(3))
	fmt.Println(IsPrimeTwin(4))
	fmt.Println(IsPrimeTwin(5))
	fmt.Println(IsPrimeTwin(6))
	fmt.Println(IsPrimeTwin(7))
	fmt.Println(IsPrimeTwin(23))
	fmt.Println(IsPrimeTwin(29))
	fmt.Println(IsPrimeTwin(19))
	fmt.Println("<<<<")
}
