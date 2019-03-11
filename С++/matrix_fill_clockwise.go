// Заполнение матрицы по часовой стрелке
//
// Используются конечные автоматы
// Состояния: право, вниз, лево, вверх
//
// Время работы алгоритма оценивается как: O(n)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var err error
	var size int

	fmt.Print("Размерность матрицы: ")

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		size, err = strconv.Atoi(input.Text())
		if err != nil {
			fmt.Fprint(os.Stderr, "Некорректное значение, введите еще раз: ")
			continue
		}
		break
	}

	now := time.Now()

	matrix := initMatrix(size)
	fillMatrix(matrix, size)
	printMatrix(matrix, size)

	fmt.Printf("Общее время: %.4fс\n", time.Since(now).Seconds())
}

// инициализация матрицы
func initMatrix(size int) [][]int {
	// создание строк
	matrix := make([][]int, size)
	for k := range matrix {
		// создание столбцов
		matrix[k] = make([]int, size)
	}

	return matrix
}

// заполнение матрицы
func fillMatrix(matrix [][]int, size int) {
	var row, col int
	var dir = "right"
	var step, count = size, size

	for i := 1; i <= size*size; i++ {
		switch dir {
		case "right":
			{
				matrix[row][col] = i
				col++
				step--
				if step == 0 {
					col--
					row++
					count--
					step = count
					dir = "down"
				}
			}
		case "down":
			{
				matrix[row][col] = i
				row++
				step--
				if step == 0 {
					row--
					col--
					step = count
					dir = "left"
				}
			}
		case "left":
			{
				matrix[row][col] = i
				col--
				step--
				if step == 0 {
					col++
					row--
					count--
					step = count
					dir = "up"
				}
			}
		case "up":
			{
				matrix[row][col] = i
				row--
				step--
				if step == 0 {
					row++
					col++
					step = count
					dir = "right"
				}
			}
		}
	}
}

// вывод матрицы
func printMatrix(matrix [][]int, size int) {
	padCount := len(strconv.Itoa(size * size))
	repeatCount := (padCount+3)*size - 1
	fmt.Println("┌" + strRepeat("─", repeatCount) + "┐")
	for r := 0; r < size; r++ {
		fmt.Print("│ ")
		for c := 0; c < size; c++ {
			fmt.Printf("%0"+strconv.Itoa(padCount)+"d │ ", matrix[r][c])
		}
		fmt.Println()
		if r+1 < size {
			fmt.Println("│" + strRepeat("─", repeatCount) + "│ ")
		}
	}
	fmt.Println("└" + strRepeat("─", repeatCount) + "┘")
}

// повторение строки
func strRepeat(s string, n int) string {
	rep := make([]string, n)
	for i := 0; i < n; i++ {
		rep[i] = s
	}

	return strings.Join(rep, "")
}
