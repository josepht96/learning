/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    
    if p == nil || q == nil {
        return false
    }
    
    if p.Val != q.Val {
        return false
    }
    
    checkLeft := isSameTree(p.Left, q.Left)
    checkRight := isSameTree(p.Right, q.Right)
    
    if checkLeft && checkRight {
        return true
    } else {
        return false
    }
}