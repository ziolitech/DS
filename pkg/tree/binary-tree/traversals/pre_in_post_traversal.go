package traversals

import (
	"errors"
	"fmt"

	"github.com/craftizmv/DS/pkg/stack"
	binary_tree "github.com/craftizmv/DS/pkg/tree/binary-tree"
)

type Pair[T any] struct {
	Element *binary_tree.GNode[T]
	Rank    int
}

// PreInPostTraversal - performs the traversal using rank
func PreInPostTraversal[T any](root *binary_tree.GNode[T]) error {
	if root == nil {
		return errors.New("wrong input")
	}

	// 0. init results
	pre, inO, post := make([]T, 0), make([]T, 0), make([]T, 0)

	//1. seed position
	// Here stack T = *Pair[T], meaning every element is of pair type.
	st := stack.NewStack[*Pair[T]]()
	st.Push(&Pair[T]{
		Element: root,
		Rank:    1,
	})

	//2. actual code.
	for !st.IsEmpty() {
		//0. Pop
		topPair, e := st.Pop()
		if e != nil {
			return e
		}

		element, rank := topPair.Element, topPair.Rank

		if rank == 1 {
			// add to result
			pre = append(pre, element.Data)

			// repush with update rank.
			st.Push(&Pair[T]{
				Element: element,
				Rank:    rank + 1,
			})

			// check the presence of left
			if element.Left != nil {
				st.Push(&Pair[T]{
					Element: element.Left,
					Rank:    1,
				})
			}
		}

		if rank == 2 {
			// add to result
			inO = append(inO, element.Data)

			// repush with update rank.
			st.Push(&Pair[T]{
				Element: element,
				Rank:    rank + 1,
			})

			// check the presence of left
			if element.Right != nil {
				st.Push(&Pair[T]{
					Element: element.Right,
					Rank:    1,
				})
			}
		}

		if rank == 3 {
			// add to result
			post = append(post, element.Data)
		}

	}

	for i := range pre {
		fmt.Println("Pre Data \n at index : ", i, " Data : ", pre[i])
		fmt.Println("InO Data \n at index : ", i, " Data : ", inO[i])
		fmt.Println("Ppst Data \n at index : ", i, " Data : ", post[i])
	}

	return nil
}
