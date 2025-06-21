package main

import (
	"fmt"
	"bufio"
	"os"
)

var numbers []uint32
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var N, M, K int
	fmt.Fscanln(reader, &N, &M, &K)
	numbers = make([]uint32, N)
	for i := 0; i < N; i++ {
		fmt.Fscanln(reader, &(numbers[i]))
	}
	allStoreTree := NewStoreTree(0, uint32(N)-1)
	var mode, numberOne, numberTwo uint32
	for i := 0; i < M+K; i++ {
		fmt.Fscanln(reader, &mode, &numberOne, &numberTwo)
		numberOne--
		if mode == 1 {
			numbers[numberOne] = numberTwo
			allStoreTree.ModifyNumber(numberOne)
		} else if mode == 2 {
			numberTwo--
			fmt.Fprintln(writer, allStoreTree.GetPartialMultiply(numberOne, numberTwo))
		} else {
			i--
		}
	}
}

type StoreTree struct {
	StartIndex uint32
	EndIndex   uint32
	Length     uint32
	Value      uint
	Children   []*StoreTree
}

func NewStoreTree(startIndex uint32, endIndex uint32) StoreTree {
	storeTree := StoreTree{startIndex, endIndex, endIndex - startIndex, 0, make([]*StoreTree, 0)}
	storeTree.GetMultiply()
	return storeTree
}
func (st *StoreTree) GetMultiply() uint {
	if st.Value != 0 {
		return st.Value
	}
	if st.Length == 0 {
		st.Value = uint(numbers[st.StartIndex])
		return st.Value
	} else if st.Length == 1 {
		st.Value = uint(numbers[st.StartIndex] * numbers[st.EndIndex]) % 1000000007
		return st.Value
	}
	c1 := NewStoreTree(st.StartIndex, st.StartIndex+(st.Length/2))
	c2 := NewStoreTree(st.StartIndex+(st.Length/2)+1, st.EndIndex)
	st.Value = (c1.GetMultiply() * c2.GetMultiply()) % 1000000007 //(A * B) mod C = (A mod C * B mod C) mod C, 수 하나하나는 1백만 이하이므로 나머지 구해도 같음, 10억 7
	st.Children = append(st.Children, &c1)
	st.Children = append(st.Children, &c2)
	return st.Value
}
func (st *StoreTree) GetPartialMultiply(startIndex uint32, endIndex uint32) uint {
	if st.Length == 0 || st.Length == 1 {
		return st.Value
	}
	if startIndex > st.StartIndex {
		if startIndex > st.Children[0].EndIndex {
			return st.Children[1].GetPartialMultiply(startIndex, endIndex)
		} else {
			if endIndex > st.Children[0].EndIndex {
				return (st.Children[0].GetPartialMultiply(startIndex, st.Children[0].EndIndex) * st.Children[1].GetPartialMultiply(st.Children[1].StartIndex, endIndex))% 1000000007
			}
			return st.Children[0].GetPartialMultiply(startIndex, endIndex)
		}
	}

	//st.StartIndex == startIndex
	if endIndex == st.EndIndex {
		return st.Value
	}
	result := st.Children[0].GetMultiply()
	if endIndex < st.EndIndex {
		result = (result * st.Children[1].GetPartialMultiply(st.Children[1].StartIndex, endIndex)) % 1000000007
	}
	return result
}
func (st *StoreTree) ModifyNumber(index uint32) uint {
	if st.StartIndex > index || st.EndIndex < index {
		return st.Value
	}
	if st.Length == 1 {
		st.Value = uint(numbers[st.StartIndex] * numbers[st.EndIndex]) % 1000000007
		return st.Value
	} else if st.Length == 0 {
		st.Value = uint(numbers[st.StartIndex]) //numbers <= 1000000
		return st.Value
	}
	st.Value = (st.Children[0].ModifyNumber(index) * st.Children[1].ModifyNumber(index)) % 1000000007
	return st.Value
}
func (st *StoreTree) PrintTree(space int) {
	for i := 0; i < space; i++ {
		fmt.Print("  ")
	}
	fmt.Printf("ㄴ[%d, %d]: %d(자식 수: %d)\n", st.StartIndex, st.EndIndex, st.GetMultiply(), len(st.Children))
	for _, child := range st.Children {
		child.PrintTree(space + 1)
	}
}
