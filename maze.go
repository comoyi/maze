package main

import (
	"fmt"
	"os"
)

type MazeMap struct {
	Name  string
	Level int
	Map   [][]int
}

type Point struct {
	x int
	y int
}

func NewPoint(x int, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

type Direction struct {
	x int
	y int
}

var directions = [4]*Direction{
	{x: -1, y: 0},
	{x: 0, y: 1},
	{x: 1, y: 0},
	{x: 0, y: -1},
}

var steps [][]int

func ReadMaze(mazeFilePath string) (*MazeMap, error) {
	var mazeMap *MazeMap

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

	//mazeMap = make([][]int, row)
	mazeMap = &MazeMap{
		Map: make([][]int, row),
	}
	fmt.Println(mazeMap)

	for i := range mazeMap.Map {
		mazeMap.Map[i] = make([]int, col)
		for j := range mazeMap.Map[i] {
			fmt.Fscanf(f, "%d", &mazeMap.Map[i][j])
			fmt.Println(mazeMap)
		}
	}
	fmt.Println(mazeMap)
	return mazeMap, err
}

func PrintMazeMap(mazeMap *MazeMap) {
	for i := range mazeMap.Map {
		for j := range mazeMap.Map[i] {
			fmt.Print(mazeMap.Map[i][j])
		}
		fmt.Println()
	}
}

func Walk(mazeMap *MazeMap) {
	fmt.Println(mazeMap)
	steps = make([][]int, len(mazeMap.Map))
	for i := range steps {
		steps[i] = make([]int, len(mazeMap.Map[0]))
	}
	var (
		startPoint   *Point
		endPoint     *Point
		direction    *Direction
		currentPoint *Point
		allowPoints  []*Point
		nextPoint    *Point
		queue        []*Point
	)
	startPoint = NewPoint(0, 0)
	endPoint = NewPoint(len(mazeMap.Map)-1, len(mazeMap.Map)-1)
	fmt.Println(startPoint)
	fmt.Println(endPoint)

	currentPoint = startPoint
	queue = append(queue, currentPoint)
	sequence := 0
	for {
		sequence++
		fmt.Println("queue len: ", len(queue))
		if len(queue) == 0 {
			PrintSteps(steps, mazeMap)
			fmt.Println("No way!")
			return
		}
		currentPoint := queue[0]
		queue = queue[1:]
		fmt.Println("queue:", queue)
		if currentPoint.x == endPoint.x && currentPoint.y == endPoint.y {
			PrintSteps(steps, mazeMap)
			PrintShortestWay(steps, mazeMap, startPoint, endPoint)
			fmt.Println("Reach!")
			return
		}

		allowPoints = make([]*Point, 0)
		for _, direction = range directions {
			nextPoint = NextPoint(currentPoint, direction)
			fmt.Println("next point:", nextPoint)
			if IsPointValid(nextPoint, mazeMap, startPoint) {
				allowPoints = append(allowPoints, nextPoint)
			}
		}
		//for _, point := range allowPoints {
		//	fmt.Println("allow points:", point)
		//}
		for _, point := range allowPoints {
			fmt.Println("add to queue: ", point)
			queue = append(queue, point)
			steps[point.x][point.y] = sequence
			fmt.Println("point:", point, "sequence:", sequence)
		}
	}
}

func NextPoint(currentPoint *Point, direction *Direction) *Point {
	var point *Point
	point = &Point{
		x: currentPoint.x + direction.x,
		y: currentPoint.y + direction.y,
	}
	return point
}

func IsPointValid(point *Point, mazeMap *MazeMap, startPoint *Point) bool {
	if point.x < 0 {
		return false
	}
	if point.x > len(mazeMap.Map)-1 {
		return false
	}

	if point.y < 0 {
		return false
	}
	if point.y > len(mazeMap.Map[0])-1 {
		return false
	}

	if point.x == startPoint.x && point.y == startPoint.y {
		return false
	}

	if steps[point.x][point.y] >= 1 {
		return false
	}

	// wall
	if mazeMap.Map[point.x][point.y] == 1 {
		return false
	}

	return true
}

func PrintSteps(steps [][]int, mazeMap *MazeMap) {
	for i := range steps {
		for j := range steps[i] {
			if mazeMap.Map[i][j] == 1 {
				fmt.Printf("%5s", "x")
			} else {
				fmt.Printf("%5d", steps[i][j])
			}
		}
		fmt.Println()
	}
}

func PrintShortestWay(steps [][]int, mazeMap *MazeMap, startPoint *Point, endPoint *Point) {
	var maxStep int
	maxStep = steps[endPoint.x][endPoint.y]
	fmt.Println(maxStep)

	var log []*Point
	log = append(log, endPoint)

	var direction *Direction
	var currentPoint *Point
	var nextPoint *Point
	var choosePoint *Point
	currentPoint = endPoint
	for {

		for _, direction = range directions {
			nextPoint = NextPoint(currentPoint, direction)
			fmt.Println("next point:", nextPoint)
			if nextPoint.x < 0 {
				fmt.Println("11")
				continue
			}
			if nextPoint.x > len(steps)-1 {
				fmt.Println("22")
				continue
			}
			if nextPoint.y < 0 {
				fmt.Println("33")
				continue
			}
			//fmt.Println(nextPoint.y,len(steps[0])-1)
			if nextPoint.y > len(steps[0])-1 {
				fmt.Println("44")
				continue
			}
			if choosePoint == nil && steps[nextPoint.x][nextPoint.y] != 0 {
				fmt.Println("a")
				choosePoint = nextPoint
			}
			fmt.Println("v:", steps[nextPoint.x][nextPoint.y])
			if steps[nextPoint.x][nextPoint.y] != 0 && steps[nextPoint.x][nextPoint.y] < steps[choosePoint.x][choosePoint.y] {
				fmt.Println("b")
				choosePoint = nextPoint
			}
		}
		fmt.Println("choose point:", choosePoint)
		log = append(log, choosePoint)
		currentPoint = choosePoint
		if steps[choosePoint.x][choosePoint.y] == 1{
			break
		}
	}

	for i := len(log) - 1; i >= 0; i-- {
		fmt.Println(log[i])
	}
}
