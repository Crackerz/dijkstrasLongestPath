package main

import "fmt"

const MAXUINT64 uint64 = (1<<64)-1 //Max value of int64

type Dam struct {
	Matrix [][]int
}

func (r *Dam) printMatrix() {
	length:=len(r.Matrix)
	maxVal:=r.getMaxVal()
	digits:=getDigits(maxVal)
	indexDigits:=getDigits(len(r.Matrix)-1)
	if(indexDigits>digits) {digits=indexDigits} //make sure col widths allow for indexs
	printIndex(length,indexDigits,digits)
	for i:=0;i<length;i++ {
		printNumber(i,indexDigits)
		fmt.Print(" [")
		for j:=0;j<length;j++ {
			if j!=0 {fmt.Print(",")}
			printNumber(r.Matrix[i][j],digits)
		}
		fmt.Print("] ")
		printNumber(i,indexDigits)
		fmt.Print("\n")
	}
	printIndex(length,indexDigits,digits)
}

func printIndex(cols,w,colWidth int) {
	for i:=0;i<w+2;i++ {
		fmt.Print(" ")
	}
	for i:=0;i<cols;i++ {
		digitWidth:=getDigits(i)
		Lpad:=(colWidth-digitWidth)/2
		Rpad:=(colWidth-digitWidth)-Lpad
		for j:=0;j<Lpad;j++ {
			fmt.Print(" ")
		}
		printNumber(i, digitWidth+Rpad)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}

func printNumber(val, digits int) {
	padding:=digits-getDigits(val)
	for i:=0;i<padding;i++ {
		fmt.Print(" ")
	}
	fmt.Print(val)
}
func (r *Dam) getMaxVal() int {
	maxVal:=-1
	for i:=0;i<len(r.Matrix);i++ {
		for j:=0;j<len(r.Matrix[i]);j++ {
			if(r.Matrix[i][j]>maxVal) {maxVal=r.Matrix[i][j]}
		}
	}
	return maxVal
}

func getDigits(val int) int {
	if(val==0) {return 1}
	digits:=0
	for val>0 {
		digits++
		val/=10
	}
	return digits
}

//Dijkstra's shortest path algorithm. I have implemented it
//using a breadth first approach. The breadth first approach
//enables the algorithm to be used on maps with negative
//weights for edges and for finding longest paths. This is
//a slower approach but allows the algorithm to be applied
//in more scenarios. The boolean indicates whether or not
//Dijkstra is attempting to find a shortest or longest path.
func (r *Dam) Dijkstra() []int {
	fmt.Println("Input")
	r.printMatrix();
	roots:=r.getRootNodes()
	nodeCount:=len(roots)
	var result []int
	maxDist:=uint64(0)
	fmt.Println("\n\nPossible Routes: \n")
	for i:=0;i<nodeCount;i++ {
		visited:=make([]bool,len(r.Matrix))
		dist:=make([]uint64,len(r.Matrix))
		nextCheck:=NewQueue(roots[i])
		//Initialize all values to max value of int
		for i:=0;i<len(r.Matrix);i++ {
			dist[i]= MAXUINT64 //Max value is reserved for Infinity or -1 depending on search
		}
		u,_:=nextCheck.Pop()
		dist[u] = 0
		visited[u] = true
		fmt.Println("Root:", u)
		for {
			//Get u's children
			child:=r.getChildren(u)
			childCount:=len(child)
			for j:=0;j<childCount;j++ {
				nextCheck.Push(child[j])
				//if(dist[child[j]]>dist[u]+uint64(r.Matrix[child[j]][u]) { //shortest path
				if(dist[child[j]]<dist[u]+uint64(r.Matrix[child[j]][u])||dist[child[j]]==MAXUINT64) { //longest path
					dist[child[j]]=dist[u]+uint64(r.Matrix[child[j]][u])
				}
			}
			visited[u]=true
			u=getNextValue(nextCheck,visited)
			if u==-1 {
				break
			}
		}
		cp:=r.getCP(dist,visited)
		fmt.Println("Local Critical Path:",cp," Distance:",dist[cp[len(cp)-1]],"\n")
		if result==nil||dist[cp[len(cp)-1]]>maxDist {
			result=cp
		}
	}

	fmt.Println("\nOveral Critical Path: ",result);
	return result
}

func getNextValue(nextCheck Queue, visited []bool) int {
	val,err:=-1,error(nil)
	for {
		val,err=nextCheck.Pop()
		if err!=nil {
			return -1
		}
		if !visited[val] {
			break
		}
	}
	return val
}

func (r *Dam) getChildren(u int) []int {
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

func (r *Dam) getParents(u int) []int {
	values:=len(r.Matrix)
	parents:=make([]int, 0, values)
	for i:=0;i<values;i++ {
		if r.Matrix[u][i]!=0 {
			//We have found a child
			parents=parents[:len(parents)+1]
			parents[len(parents)-1]=i
		}
	}
	return parents
}

//Find the longest set of nodes that can be traversed to get to
//a node in a map.
//Params:
//dist = distance traversed to reach node.
//visited = if the node has been visited.
//Returns:
//[]int the node ids along the critical path.
func (r *Dam) getCP(dist []uint64, visited []bool) []int {
	u:=getLargestValue(dist,visited)
	if u==-1 {return nil}
	vals:=len(dist)
	path:=make([]int,vals)
	pathLen:=1
	path[vals-pathLen] = u
	pathLen++
	parents:=r.getParents(u)
	for parents!=nil {
		u=getLargestParent(parents,dist,visited)
		if u==-1 {break}
		path[vals-pathLen]=u
		pathLen++
		parents=r.getParents(u)
	}
	return path[vals-pathLen+1:]
}

func getLargestParent(index []int, dist []uint64, visited []bool) int {
	values:=len(index)
	result:=-1
	maxVal:=uint64(0)
	for i:=0;i<values;i++ {
		if visited[index[i]]&&(result==-1||maxVal<dist[index[i]]) {
			result=index[i]
			maxVal=dist[index[i]]
		}
	}
	return result
}

//Returns index of largest value
func getLargestValue(dist []uint64, visited []bool) int {
	nodes:=len(dist)
	index:=-1
	maxVal:=uint64(0)
	for i:=0;i<nodes;i++ {
		if visited[i]&&(index==-1||maxVal<dist[i]) {
			maxVal=dist[i]
			index=i
		}
	}
	return index
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
