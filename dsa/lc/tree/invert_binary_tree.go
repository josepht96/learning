/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func invertTree(root *TreeNode) *TreeNode {
    //navigate to node, switch child nodes
    if root == nil {
        return nil
    }
    tempLeftNode := root.Left
    //switch nodes
    root.Left = root.Right
    root.Right = tempLeftNode
    
    _ = invertTree(root.Left)
    _ = invertTree(root.Right)
    
    return root
}