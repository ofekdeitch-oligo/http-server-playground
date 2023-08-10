package main

type Tree struct {
	value     int
	leftNode  *Tree
	rightNode *Tree
}

func sum(tree Tree) int {
	leftNodeSum := 0
	rightNodeSum := 0

	if tree.leftNode != nil {
		leftNodeSum = sum(*tree.leftNode)
	}

	if tree.rightNode != nil {
		rightNodeSum = sum(*tree.rightNode)
	}

	return tree.value + leftNodeSum + rightNodeSum
}

func max(tree Tree) int {
	maxValue := tree.value

	if tree.leftNode != nil {
		leftMax := max(*tree.leftNode)

		if maxValue < leftMax {
			maxValue = leftMax
		}
	}

	if tree.rightNode != nil {
		rightMax := max(*tree.rightNode)

		if maxValue < rightMax {
			maxValue = rightMax
		}
	}

	return maxValue
}
