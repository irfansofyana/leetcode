/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import "fmt"

type QueueElement struct {
    Depth int
    Node *TreeNode
}
type Queue struct {
    Elements []QueueElement
    CurrIdx int
}
func (q *Queue) Push(el QueueElement) {
    q.Elements = append(q.Elements, el)
}
func (q *Queue) Front() QueueElement {
    front := q.Elements[q.CurrIdx]
    q.CurrIdx++
    return front
}
func (q *Queue) IsEmpty() bool {
    return q.CurrIdx == len(q.Elements)
}

func CreateAQueue() Queue {
    el := make([]QueueElement, 0)
    return Queue {
        CurrIdx: 0,
        Elements: el,
    }
}

func isEvenOddTree(root *TreeNode) bool {
    q := CreateAQueue()
    q.Push(QueueElement{
        Depth: 0,
        Node: root,
    })

    mep := make(map[int][]int)
    maxDepth := 0
    for !q.IsEmpty() {
        f := q.Front()
        depth := f.Depth
        if depth > maxDepth {
            maxDepth = depth
        }
        node := f.Node
        if len(mep[depth]) == 0 {
            mep[depth] = make([]int, 0)
        }
        mep[depth] = append(mep[depth], node.Val)
        if (node.Left != nil) {
            q.Push(QueueElement{
                Depth: depth+1,
                Node: node.Left,
            })
        }
        if (node.Right != nil) {
            q.Push(QueueElement{
                Depth: depth+1,
                Node: node.Right,
            })
        }
    }

    for i := 0; i <= maxDepth; i++ {
        for j := 0; j < len(mep[i]); j++ {
            if i % 2 == 0 && mep[i][j] % 2 == 0 {
                return false
            } 
            if i % 2 == 1 && mep[i][j] % 2 == 1 {
                return false
            }
            if i % 2 == 0 && j > 0 && mep[i][j] <= mep[i][j-1] {
                return false
            }
            if i % 2 == 1 && j > 0 && mep[i][j] >= mep[i][j-1] {
                return false
            }
        }
    }

    return true
}
