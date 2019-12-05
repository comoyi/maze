package main

import (
	"flag"
	"fmt"
)

func main() {

	var mazeFilePath string
	flag.StringVar(&mazeFilePath, "file", "", "maze map file path")
	flag.Parse()
	fmt.Println(mazeFilePath)
	var mazeMap [][]int
	var err error
	//mazeFilePath = "./mazes/maze-1"
	mazeMap, err = ReadMaze(mazeFilePath)
	if err != nil {
		fmt.Println("Read maze file failed!")
		return
	}
	fmt.Println("mazeFile: ", mazeFilePath, "mazeMap: ", mazeMap)

	PrintMazeMap(mazeMap)
}
