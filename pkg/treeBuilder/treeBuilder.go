package treebuilder

import "container/heap"

type HuffmanNode struct {
	Character rune
	Frequency int
	Left      *HuffmanNode
	Right     *HuffmanNode
}

type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Frequency < pq[j].Frequency
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*HuffmanNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func BuildHuffmanTree(freqMap map[rune]int) *HuffmanNode {
	var pq PriorityQueue
	heap.Init(&pq)

	for char, freq := range freqMap {
		heap.Push(&pq, &HuffmanNode{Character: char, Frequency: freq})
	}

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(*HuffmanNode)
		right := heap.Pop(&pq).(*HuffmanNode)

		merged := &HuffmanNode{
			Frequency: left.Frequency + right.Frequency,
			Left:      left,
			Right:     right,
		}

		heap.Push(&pq, merged)
	}

	return heap.Pop(&pq).(*HuffmanNode)
}

func GenerateHuffmanCodes(root *HuffmanNode) map[rune]string {
	encoding := make(map[rune]string)
	generateCodesHelper(root, "", encoding)
	return encoding
}

func generateCodesHelper(node *HuffmanNode, code string, encoding map[rune]string) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		encoding[node.Character] = code
		return
	}

	generateCodesHelper(node.Left, code+"0", encoding)
	generateCodesHelper(node.Right, code+"1", encoding)
}
