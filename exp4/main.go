package main

import (
		"fmt"
		"time"
		"math/rand")

func makeGraph(n int) [][]bool{
	R := rand.New(rand.NewSource(time.Now().UnixNano()))
	graph := make([][]bool,n)
	for i := 0 ;i < n; i++ {
		graph[i] = make([]bool,n)
	}
    for i := 0; i < n; i++ {
        graph[i][i] = true
        for j := (i + 1); j < n; j++ {
            graph[i][j] = (R.Intn(0x80000000) > 0x7FFFFFFF/2)
            graph[j][i] = graph[i][j]
        }
    }
    fmt.Println("生成的矩阵图 >> ")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] {
				fmt.Printf("1 ")
			} else {
				fmt.Printf("0 ")
			}
		}
		fmt.Println()
    }
    return graph
}
func do_check(graph [][]bool, n int , vis []bool, now int) {
    if vis[now] {
        return
    }
    vis[now] = true
    for i:= 0; i < n; i++ {
        if graph[now][i] {
            do_check(graph, n, vis, i)
        }
    }
}

func check(graph [][]bool, n int) bool { 
	//连通性检查
    vis := make([]bool,n)
    for i:= 0; i < n; i++ {
    	vis[i] = false
    }
    do_check(graph, n, vis, 0);
    for i:= 0; i < n; i++ {
        if vis[i] != true {
            return false
        }
    }
    return true
}

type Stack struct {
    top int
    node[100] int
}

var s Stack

func DFS(graph [][]bool, n int, x int) {
    s.top++
    s.node[s.top] = x
    for i:= 0; i < n; i++ {
        if graph[i][x] == true {
            graph[i][x] = false
            graph[x][i] = false
            DFS(graph, n, i)
            break
        }
    }
}

func Fleury(graph [][]bool, n int, ans []int, x int) int {
    var t,count int
    count = 0
    s.top = 0
    s.node[s.top] = x
	for s.top >= 0 {
        t =  0
        for i:= 0; i < n; i++ {
            if graph[s.node[s.top]][i] == true {
                t = 1
                break
            }
        }
        if (t== 0) {
            ans[count] = s.node[s.top] + 1
            count++
            s.top--
        } else {
            s.top--
            DFS(graph, n, s.node[s.top + 1])
        }
    }
    fmt.Println()
    return count
}
func main() {
	var n int
	fmt.Printf("请输入节点数 >> ")
	fmt.Scanf("%d",&n)
	if n <= 0 {
		return 
	}
	graph := makeGraph(n)
	if !check(graph,n) {
		fmt.Println("非连通图")
		return 
	}
	
	var start,num int	
	for i:= 0; i < n; i++ {
		degree := 0
		for j:= 0; j < n; j++ {
			if graph[i][j] {
				degree += 1
			} else {
				degree += 0
			}
		}
		if  degree % 2 == 1 {
			start = i
			num++
		}
	}
	if num == 0 || num == 2 {
		ans := make([]int,n*4)
		count := Fleury(graph, n, ans, start)
		if (ans[0] == ans[n - 1]) {
			fmt.Println("该图为欧拉图，欧拉回路为: ")
		} else {
			fmt.Println("该图为半欧拉图，欧拉路为: ")
		}
		var path string = ""
		for i:= 0; i < count; i++ {
			path += string(ans[i] + 0x30) + " "
		}
		fmt.Println(path)
	} else {
		fmt.Println("非欧拉图或半欧拉图")
	}
}
