package sols

import (
	"dsa/lib"
	"fmt"
)

/*
Constraints:
1 <= n <= 2000
0 <= prereqs.length <= n * (n - 1)
prereqs[i].length == 2
0 <= a, b < n
a != b
All the pairs [a, b] are distinct.
Example 1:
Input: n = 2, prereqs = [[1,0]]

Output: [0,1]

Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].

Example 2 :
Input: n = 4, prereqs = [[1,0],[2,0],[3,1],[3,2]]

Output: [0,2,1,3]

Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finish course 0.

Two correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

Example 3 :
Input: n = 1, prereqs = []

Output: [0]
*/

/*
Solution-explanation

This is DAG and if there are no inwards or the indegree is 0 to the course then it can be started at first. If there isn't any course with indegree -> 0 means it is cyclic and courses can't be finished

1st condition: Atleast 1 course indegree should be 0

As mentioned we can finish in any order only the condition is to meet the prerequisties every time


Now we can solve this by Kanh's algo (BFS) or DFS by finding the levels (constructing the max heights from any root)

BFS:

-> get all in-degrees and connections like start,end

Now traverse and form the order
	In each traversal, decrease the neighbour in-degree and add to ans (if updated in-degree is 0 or push to the queue)

*/

// q1 takes no of courses and prerequisties list, where courses are numbers and returns the order of courses to finish
func Q1(n int, prereqs [][]int) ([]int, error) {
	return dfs(n, prereqs)
}

func bfs(n int, prereqs [][]int) ([]int, error) {

	ans := make([]int, 0, n)
	inDegrees := make(map[int]int, n)
	list := make(map[int][]int, n)

	for _, edge := range prereqs {
		if len(edge) != 2 {
			return nil, fmt.Errorf("size of prequisties element should be 2")
		}
		inDegrees[edge[0]]++
		if _, exists := list[edge[1]]; !exists {
			list[edge[1]] = make([]int, 0)
		}
		list[edge[1]] = append(list[edge[1]], edge[0])
	}

	for i := 0; i < n; i++ {
		if _, exists := inDegrees[i]; !exists {
			inDegrees[i] = 0
		}
	}

	if len(inDegrees) != n {
		return nil, fmt.Errorf("size ")
	}

	nodes := make([]interface{}, 0, len(prereqs)+n)
	// push root nodes first
	for node, inDegree := range inDegrees {
		if inDegree == 0 {
			nodes = append(nodes, node)
		}
	}

	for node, inDegree := range inDegrees {
		if inDegree != 0 {
			nodes = append(nodes, node)
		}
	}

	queue, err := lib.NewQueue(nodes)
	if err != nil {
		return nil, err
	}

	for !queue.IsEmpty() {

		ele, err := queue.Front()
		if err != nil {
			return nil, err
		}
		queue.Pop()

		if inDegrees[ele.(int)] == -1 {
			continue
		} else if inDegrees[ele.(int)] > 0 {
			queue.Push(ele)
			continue
		}

		// mark it as done
		inDegrees[ele.(int)] = -1
		ans = append(ans, ele.(int))

		if dep, exists := list[ele.(int)]; exists {
			for _, neighbour := range dep {
				inDegrees[neighbour]--
				queue.Push(neighbour)
			}
		}

	}

	return ans, nil
}

// Idea is to create the forest and max level from any root and finish the courses from desc to asc order of level

func dfs(n int, prereqs [][]int) ([]int, error) {

	ans := make([]int, 0, n)

	roots := make([]int, 0, n)
	level := make(map[int]int)
	inDegrees := make(map[int]int, n)
	list := make(map[int][]int, n)

	for _, edge := range prereqs {
		if len(edge) != 2 {
			return nil, fmt.Errorf("size of prequisties element should be 2")
		}
		inDegrees[edge[0]]++
		if _, exists := list[edge[1]]; !exists {
			list[edge[1]] = make([]int, 0)
		}
		list[edge[1]] = append(list[edge[1]], edge[0])
	}

	for i := 0; i < n; i++ {
		if _, exists := inDegrees[i]; !exists {
			inDegrees[i] = 0
		}
	}

	if len(inDegrees) != n {
		return nil, fmt.Errorf("size ")
	}

	// push root nodes first
	for node, inDegree := range inDegrees {
		if inDegree == 0 {
			roots = append(roots, node)
			level[node] = 0
		}
	}

	for _, root := range roots {
		st, err := lib.NewStack([]interface{}{root})
		if err != nil {
			return nil, err
		}

		for !st.IsEmpty() {
			ele, err := st.Top()
			if err != nil {
				return nil, err
			}
			st.Pop()

			if dep, exists := list[ele.(int)]; exists {
				for _, neighbour := range dep {
					level[neighbour] = lib.Max(level[neighbour], level[ele.(int)]+1)
					st.Push(neighbour)
				}
			}
		}
	}

	levelArr := make(map[int][]int, 0)
	for k, v := range level {
		if _, exists := levelArr[v]; !exists {
			levelArr[v] = make([]int, 0)
		}
		levelArr[v] = append(levelArr[v], k)
	}

	for i := 0; i < n; i++ {
		if _, exists := levelArr[i]; !exists {
			continue
		}
		ans = append(ans, levelArr[i]...)
	}
	return ans, nil
}
