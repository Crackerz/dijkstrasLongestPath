package main

import "fmt"

const MAXUINT64 uint64 = (1<<64)-1 //Max value of int64

type Dam struct {
	Matrix [][]int
}

func (r *Dam) dijkstra() []int {
	//roots:=[]int{0,1,2,3,4,5,6,7,8,9,10,11,12,13}
	roots:=r.getRootNodes()
	nodeCount:=len(roots)
	for i:=0;i<nodeCount;i++ {
		visited:=make([]bool,len(r.Matrix))
		dist:=make([]uint64,len(r.Matrix))
		//Initialize all values to max value of int
		for i:=0;i<len(r.Matrix);i++ {
			dist[i]= MAXUINT64
		}
		u:=roots[i]
		dist[u] = 0
		fmt.Println("Root:", u)
		for {
			fmt.Println("Selected Node ",u)
			//Get u's children
			child:=r.getChildren(u)
			fmt.Println("Node ",u,"'s children:",child)
			childCount:=len(child)
			for j:=0;j<childCount;j++ {
				if(dist[child[j]]>dist[u]+uint64(r.Matrix[child[j]][u])) {
					dist[child[j]]=dist[u]+uint64(r.Matrix[child[j]][u])
					fmt.Println(child[j],"=",dist[u],"+",r.Matrix[child[j]][u],"=",dist[child[j]])
				}
			}
			visited[u]=true
			u=r.getSmallestValue(dist,visited)
			if u==-1 {
				break
			}
		}
		fmt.Println(dist)
	}
	return nil
}

func (r *Dam) getSmallestValue(dist []uint64, visited []bool) int {
	values:=len(dist)
	minValue:=MAXUINT64
	index:=(-1)
	for i:=0;i<values;i++ {
		if !visited[i]&&dist[i]!=MAXUINT64 {
			if(index==-1||minValue>dist[i]) {
				minValue = dist[i]
				index = i
			}
		}
	}
	return index
}

func (r *Dam) getChildren(u int) []int{
	values:=len(r.Matrix)
	children:=make([]int, 0, values)
	for i:=0;i<values;i++ {
		if r.Matrix[i][u]!=0 {
			//We have found a child
			children=children[:len(children)+1]
			children[len(children)-1]=i
		}
	}
	return children
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
		}
	}
	return roots
}
