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
	var mazeMap *MazeMap
	var err error
	mazeMap, err = ReadMaze(mazeFilePath)
	if err != nil {
		fmt.Println("Read maze file failed!")
		return
	}
	fmt.Println("mazeFile: ", mazeFilePath, "mazeMap: ", mazeMap)

	PrintMazeMap(mazeMap)

	Walk(mazeMap)
}
