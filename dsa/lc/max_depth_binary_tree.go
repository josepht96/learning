/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func maxDepth(root *TreeNode) int {
    if (root == nil) {
        return 0;
    }
    leftMax := (1 + maxDepth(root.Left))
    rightMax := (1 + maxDepth(root.Right))
    if leftMax > rightMax {
        return leftMax
    } else {
        return rightMax
    }
}