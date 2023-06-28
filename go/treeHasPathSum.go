// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	targetSum = targetSum - root.Val
	if hasPathSum(root.Left, targetSum) {
		return true
	}

	if hasPathSum(root.Right, targetSum) {
		return true
	}

	return false

}