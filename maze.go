package main

import (
	"fmt"
	"os"
)

func ReadMaze(mazeFilePath string) ([][]int, error) {
	var mazeMap [][]int

	fmt.Println(mazeFilePath)

	var f *os.File

	var err error
	f, err = os.Open(mazeFilePath)
	if err != nil {
		fmt.Println("Open maze file err")
		return mazeMap, err
	}
	var row int
	var col int
	fmt.Fscanf(f, "%d %d", &row, &col)
	fmt.Println("row: ", row)
	fmt.Println("col: ", col)

	mazeMap = make([][]int, row)
	fmt.Println(mazeMap)

	for i := range mazeMap {
		mazeMap[i] = make([]int, col)
		for j := range mazeMap[i] {
			fmt.Fscanf(f, "%d", &mazeMap[i][j])
			fmt.Println(mazeMap)
		}
	}
	fmt.Println(mazeMap)
	return mazeMap, err
}

func PrintMazeMap(mazeMap [][]int) {
	for i := range mazeMap {
		for j := range mazeMap[i] {
			fmt.Print(mazeMap[i][j])
		}
		fmt.Println()
	}
}
