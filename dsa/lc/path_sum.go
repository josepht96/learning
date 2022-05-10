/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func hasPathSum(root *TreeNode, targetSum int) bool {
    //only once you hit a root can you check if its correct
    //if the sum surpasses targetSum return regardless
    //targetSum can become remainder
    if root == nil {
        return false
    }
    if root.Val == targetSum && root.Left == nil && root.Right == nil {
        return true
    }
    
    remainder := targetSum - root.Val
    checkLeft := hasPathSum(root.Left, remainder)
    checkRight := hasPathSum(root.Right, remainder)
    if checkLeft == true {
        return true
    } 
    if checkRight == true {
        return true
    }
    return false
}