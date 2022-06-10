/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

 func connect(root *Node) *Node {
	//check left
    //
    if root == nil {
        return nil
    }
    
    if root.Left != nil {
        root.Left.Next = root.Right
    }
    
    if root.Right != nil {
        if root.Next == nil {
            root.Right.Next = nil
        } else {
            root.Right.Next = root.Next.Left
        }
    }
    
    if root.Left != nil {
        _ = connect(root.Left)
    }
    if root.Right != nil {
        _ = connect(root.Right)
    }
    
    return root
}