/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }
    
    leftMin := 1 + minDepth(root.Left)
    rightMin := 1 + minDepth(root.Right)
	//if leftMin or rightMin is 1 (returned zero, we know there wasn't a leafnode)
    if leftMin == 1 {
        leftMin = rightMin
    }
    if rightMin == 1 {
        rightMin = leftMin
    }
    
    if leftMin < rightMin {
        return leftMin
    } else {
        return rightMin
    }
}