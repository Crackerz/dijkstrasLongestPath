package main

import "fmt"

type Dam struct {
	Matrix [][]int
}

func (r *Dam) dijkstra() []int {
	return nil
}

func (r *Dam) getRootNodes() []int {
	roots:=make([]int,0,len(r.Matrix))
	verticies:=len(r.Matrix)
	for i:=0;i<verticies;i++ {
		hasChild:=false
		//Check each node for children
		for j:=0;j<verticies;j++ {
			if r.Matrix[i][j] != 0 {
				hasChild=true
				break
			}
		}
		if !hasChild {
			roots=roots[:len(roots)+1]
			roots[len(roots)-1]=i
			fmt.Println(len(roots));
		}
	}
	return roots
}
