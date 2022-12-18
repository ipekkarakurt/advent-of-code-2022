package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name      string
	size      uint64
	children  []*node
	parent    *node
	isDir     bool
	isVisited bool
}

func main() {
	f, _ := os.Open("/Users/ipekkarakurt/Documents/advent-of-code-2022/day-7/input.txt")
	scanner := bufio.NewScanner(f)

	root := &node{"/", 0, make([]*node, 0), nil, true, true}
	currentPtr := root

	scanner.Scan()
	line := scanner.Text()
	notEOF := true

	for notEOF {
		if string(line[0]) == "$" {
			// New command
			cmd := line[2:]
			switch cmd[:2] {
			case "cd":
				param := line[5:]
				if param == "/" {
					currentPtr = root
				} else if param == ".." {
					currentPtr = currentPtr.parent
				} else {
					// Point to param dir
					for _, trNode := range currentPtr.children {
						if trNode.name == param {
							currentPtr = trNode
							break
						}
					}
				}
				scanner.Scan()
				line = scanner.Text()
				if line == "" {
					notEOF = false
					break
				}
			case "ls":
				for {
					scanner.Scan()
					line = scanner.Text()
					if line == "" {
						notEOF = false
						break
					}
					listed := strings.Split(line, " ")

					if string(line[0]) == "$" {
						// Encountered new command, break
						if currentPtr.isDir && !currentPtr.isVisited {
							currentPtr.isVisited = true
							parentNode := currentPtr.parent
							for parentNode != nil {
								parentNode.size += currentPtr.size
								parentNode = parentNode.parent
							}
						}
						break
					}
					if string(listed[0]) == "dir" {
						// Listing directory
						thisNode := &node{listed[1], 0, nil, currentPtr, true, false}
						currentPtr.children = append(currentPtr.children, thisNode)
					} else {
						// Listing file
						size, _ := strconv.Atoi(listed[0])
						thisNode := &node{listed[1], uint64(size), nil, currentPtr, false, true}
						currentPtr.children = append(currentPtr.children, thisNode)
						currentPtr.size += thisNode.size
					}
				}
			}
		}
	}

	fmt.Println("First part: ", dfsSum(root, 0))
	toFree := 30000000 - (70000000 - root.size)
	minToFree := uint64(10000000000000)
	fmt.Println("Second part: ", dfsMin(root, toFree, minToFree))
	//dfsPrint(root, 0)
}

func dfsSum(n *node, sum uint64) uint64 {
	if n.isDir && n.size < 100000 {
		sum += n.size
	}
	for _, child := range n.children {
		sum += dfsSum(child, 0)
	}

	return sum
}

func dfsMin(n *node, toFree uint64, minToFree uint64) uint64 {
	if toFree < n.size && minToFree > n.size {
		minToFree = n.size
	}

	for _, child := range n.children {
		minToFree = dfsMin(child, toFree, minToFree)
	}

	return minToFree
}

func dfsPrint(n *node, level int) {
	fmt.Println(n.name, n.size)
	for _, child := range n.children {
		p := child.parent
		for p != nil {
			p = p.parent
			fmt.Print("-")
		}
		dfsPrint(child, level)
	}
}

func getDirectorySizes(dir *node) []int {
	var sizes []int
	sizes = append(sizes, int(dir.size))
	for _, dir := range dir.children {
		sizes = append(sizes, getDirectorySizes(dir)...)
	}

	return sizes
}
