package main

import (
	"testing"
)

func Test_Tree1_Sum(t *testing.T) {
	root := Tree{
		value: 1,
	}

	Test{t}.Expect(sum(root)).ToEqual(1)
}

func Test_Tree2_Sum(t *testing.T) {
	leftLeaf := Tree{
		value: 2,
	}

	rightLeaf := Tree{
		value: 3,
	}

	root := Tree{
		value:     5,
		leftNode:  &leftLeaf,
		rightNode: &rightLeaf,
	}

	Test{t}.Expect(sum(root)).ToEqual(10)
}

func Test_Tree1_Max(t *testing.T) {
	root := Tree{
		value: 1,
	}

	Test{t}.Expect(max(root)).ToEqual(1)
}

func Test_Tree2_Max(t *testing.T) {
	leftLeaf := Tree{
		value: 2,
	}

	root := Tree{
		value:    1,
		leftNode: &leftLeaf,
	}

	Test{t}.Expect(max(root)).ToEqual(2)
}

func Test_Tree3_Max(t *testing.T) {
	leftLeaf := Tree{
		value: 5,
	}

	middleNode := Tree{
		value:    2,
		leftNode: &leftLeaf,
	}

	root := Tree{
		value:    1,
		leftNode: &middleNode,
	}

	Test{t}.Expect(max(root)).ToEqual(5)
}

func Test_Tree4_Max(t *testing.T) {
	leftLeaf := Tree{
		value: 5,
	}

	rightLeaf := Tree{
		value:    20,
		leftNode: &leftLeaf,
	}

	root := Tree{
		value:     1,
		leftNode:  &leftLeaf,
		rightNode: &rightLeaf,
	}

	Test{t}.Expect(max(root)).ToEqual(20)
}

func Test_Tree5_Max(t *testing.T) {
	leaf1 := Tree{
		value: 5,
	}
	leaf2 := Tree{
		value: 10,
	}

	leaf3 := Tree{
		value: 11,
	}
	leaf4 := Tree{
		value: 8,
	}

	middleNode1 := Tree{
		value:     2,
		leftNode:  &leaf1,
		rightNode: &leaf2,
	}

	middleNode2 := Tree{
		value:     3,
		leftNode:  &leaf3,
		rightNode: &leaf4,
	}

	root := Tree{
		value:     1,
		leftNode:  &middleNode1,
		rightNode: &middleNode2,
	}

	Test{t}.Expect(max(root)).ToEqual(11)
}
