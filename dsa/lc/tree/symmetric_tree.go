/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func isMirror(left, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    if left == nil || right == nil {
        return false
    }
    if left.Val == right.Val {
        checkOuter := isMirror(left.Left, right.Right)
        checkInner := isMirror(left.Right, right.Left)
        if checkOuter && checkInner {
            return true
        }
    }
    return false
}
func isSymmetric(root *TreeNode) bool {
    return isMirror(root.Left, root.Right)
}