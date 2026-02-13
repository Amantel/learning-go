// Уровень сложности: сложный

// Если нам даны длины трёх отрезков, то можно проверить, можно ли из них составить треугольник. Для этого нужно проверить неравенство треугольника: каждая сторона должна быть меньше суммы двух других.

// Если из отрезков можно построить треугольник, то он получится одного из трёх типов:

// Остроугольный (Acute)

// Прямоугольный (Right)

// Тупоугольный (Obtuse)

// Задан набор констант: Impossible, Acute, Right, Obtuse

// Нужно написать функцию InspectTriangle( a, b, c int) int

// Которая принимает длины трёх отрезков и возвращает одну из заданных констант, в зависимости от того, существует ли треугольник и какого он типа.

// Например InspectTriangle (3, 6, 10) должно вернуть impossible, т.к. не выполняется неравенство треугольника.

// Например InspectTriangle (3, 10, 7) должно вернуть Impossible, т.к. по-прежнему не выполняется неравенство треугольника.

// Например InspectTriangle (5, 6, 7) должно вернуть Acute

// Например InspectTriangle( 3, 5, 4) должно вернуть Right

// Например InspectTriangle (2, 9, 10) должно вернуть Obtuse

// Подсказка: проверить прямоугольность, остроугольность, тупоугольность можно через Теорему Пифагора и её следствие. Если стороны треугольника обозначить a, b, c, причем c – самая большая из сторон, то для прямоугольного треугольника выполняется равенство: aa+bb=c*c

// А для остроугольного и тупоугольного треугольника будут выполняться соответствующие неравенства:

// aa+bb>c*c

// aa+bb<c*c

package main

import (
	"fmt"
	"math"
	"slices"
)

const (
	Impossible = iota
	Acute
	Right
	Obtuse
)

func InspectTriangle(a, b, c int) int {
	arr := []int{a, b, c}
	slices.Sort(arr)
	a = arr[0]
	b = arr[1]
	c = arr[2]

	if c >= a+b || b >= a+c || a >= b+c {
		return Impossible
	}
	quadrAB := math.Pow(float64(a), 2) + math.Pow(float64(b), 2)
	quardC := math.Pow(float64(c), 2)
	// fmt.Println("a", a, math.Pow(float64(a), 2))
	// fmt.Println("b", b, math.Pow(float64(b), 2))
	// fmt.Println("c", c, quardC)
	// fmt.Println(quadrAB, quardC)
	if quadrAB > quardC {
		return Acute
	}
	if quadrAB < quardC {
		return Obtuse
	}
	return Right
}

func main() {
	fmt.Println(">>>>")
	fmt.Println(InspectTriangle(3, 6, 10))
	fmt.Println(InspectTriangle(3, 10, 7))
	fmt.Println(InspectTriangle(5, 6, 7), "Acute - 1")
	fmt.Println(InspectTriangle(3, 5, 4), "Right - 2")
	fmt.Println(InspectTriangle(2, 9, 10), "Obtuse - 3")
	fmt.Println("<<<<")
}
