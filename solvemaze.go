// solvemaze.go - Simple program to solve a maze.
// The input file contains a description of the multiway intersections in the maze.
//
// That is, the file describes only those nodes in the maze where you have
// a choice of which way to go. Each such node has a number, and each node
// has letters describing the choices.  In the original maze for which
// this program was written - from the 2024-05-26 issue of the New York
// Times - each node in the maze has only two choices, so the directions
// to go for that maze are "a" and "b".
//
// The input file contains lines, each of which describes a path leading out of a given node.
//
// - "1a 2"  means that if you start at node 1 and take direction a, you get to node 2.
// - "1b x"  means that if you start at node 1 and take direction b, you will reach a dead end.
// - "2a e"  means that if you start at node 2 and take direction a, you will reach the end.
//
// Blank lines, and lines starting with "#", are ignored.
// It's assumed that the first multiway node in the maze is numbered 1.
// The maze must be annotated (outside of this program) as to which nodes have
// numbers, what those numbers are, and which direction out of a node is "a" and "b".
//
// Mark Riordan  01-JUL-2024

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const EXIT_NODE_NUM int = -1

// NextNode represents a direction and destination node from a node.
type NextNode struct {
	Direction string
	NodeNum   int
}

// Node represents a junction in the maze.
type Node struct {
	Name     string
	Children []NextNode // List of directions and destination nodes from this node.
}

// Maze is a slice of nodes.
// Maze[i] is a list of directions that can be taken from node i, and
// the destination node number for each direction.
// If Maze[i] is empty, then this entry should be ignored.
var Maze []Node

// Parse a string of the form "12a" into a node number and a direction.
func parseNodeFrom(nodeFrom string) (nodeNum int, direction string) {
	nodeNumStr := ""
	for i, r := range nodeFrom {
		if unicode.IsDigit(r) {
			nodeNumStr += string(r)
		} else {
			direction = nodeFrom[i:]
			break
		}
	}
	nodeNum, _ = strconv.Atoi(nodeNumStr)
	return nodeNum, direction
}

// Debug print the data structure representing the maze.
func dumpMaze() {
	for i, node := range Maze {
		if len(node.Children) > 0 {
			fmt.Print("Node ", i)
			for _, nextNode := range node.Children {
				fmt.Print("  ", nextNode.Direction, nextNode.NodeNum, ";")
			}
			fmt.Println()
		}
	}
}

// Parse a line from the input file and add the node to the maze.
func parseLine(line string) {
	line = strings.TrimSpace(line)
	// Ignore comments and empty lines.
	if len(line) == 0 {
		return
	}
	if line[0] == '#' {
		return
	}
	tokens := strings.Fields(line)
	if len(tokens) != 2 {
		fmt.Println("Invalid line:", line)
		return
	}
	nodeFrom := tokens[0]
	// nodeFrom now looks like 12a.
	nodeToStr := tokens[1]
	// nodeToStr now looks like 13, or x, or e.
	var nodeTo int
	nodeNum, direction := parseNodeFrom(nodeFrom)
	if "x" == nodeToStr {
		// No need to add a node for a dead end
		return
	} else if "e" == nodeToStr {
		nodeTo = EXIT_NODE_NUM
		// Add a node for the end
	} else {
		nodeTo, _ = strconv.Atoi(nodeToStr)
	}
	// If the node number is greater than the length of the Maze slice, add additional elements
	if len(Maze) < nodeNum+1 {
		// Calculate the number of elements to add
		additionalElementsNeeded := (nodeNum + 4) - len(Maze)
		// Create a slice with the additional elements initialized to their zero value
		additionalElements := make([]Node, additionalElementsNeeded)
		// Append the additional elements to Maze
		Maze = append(Maze, additionalElements...)
	}
	nextNode := NextNode{Direction: direction, NodeNum: nodeTo}
	Maze[nodeNum].Children = append(Maze[nodeNum].Children, nextNode)
}

// Process the input file, adding nodes to the maze.
func processFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parseLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	//dumpMaze()
}

// Solve the maze.
// startNodeNum is the number of the node to start at.
// path is the path taken so far, in the form "1a 2b 3a".
// Returns the path taken to solve the maze, or an empty string if
// the maze cannot be solved.
func solveMaze(startNodeNum int, path string) string {
	for _, nextNode := range Maze[startNodeNum].Children {
		if nextNode.NodeNum == EXIT_NODE_NUM {
			return path + " " + strconv.Itoa(startNodeNum) + nextNode.Direction
		}
		possiblePath := solveMaze(nextNode.NodeNum, path+" "+strconv.Itoa(startNodeNum)+nextNode.Direction)
		if possiblePath != "" {
			return possiblePath
		}
	}
	return ""
}

func main() {
	processFile("maze20240526.txt")
	solvedPath := solveMaze(1, "")
	fmt.Println("Solved path:", solvedPath)
}
