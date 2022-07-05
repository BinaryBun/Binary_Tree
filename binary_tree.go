/*
    __init__: mindTree{value: <int>}    // initializing a binary tree with an input element
    Insert: mindTree.Insert(<int>)      // insertion into a binary tree with <int> input data
    findMax: mindTree.findMax() <int>   // search for the largest element in the tree with output data of type <int>
    findMin: mindTree.findMin() <int>   // search for the smallest element in the tree with <int> type output data
*/
package main

import ("errors")

type mindTree struct {
  value int
  left *mindTree
  right *mindTree
}

func (m *mindTree) Insert (val int) error {
  if m == nil {
		return errors.New("Tree is nil!!")
	}

  if m.value > val {
    if m.left == nil {
      m.left = &mindTree{value: val}
      return nil
    }
    return m.left.Insert(val)
  }

  if m.value < val {
    if m.right == nil {
      m.right = &mindTree{value: val}
      return nil
    }
    return m.right.Insert(val)
  }
  return nil
}

func (m *mindTree) findMin () int {
  if m.left == nil {
    return m.value
  }
  return m.left.findMin()
}

func (m *mindTree) findMax () int {
  if m.right == nil {
    return m.value
  }
  return m.right.findMax()
}

func (t *mindTree) printInorder() {
	if t == nil {
		return
	}

	t.left.printInorder()
	fmt.Print(t.value, " ")
	t.right.printInorder()
}
