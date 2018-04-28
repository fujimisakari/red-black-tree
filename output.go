package main

import (
	"fmt"
	"strings"
)

const (
	LeafLength = 10
	LEFT       = "left"
	RIGHT      = "right"
)

func printTree(root int, loopCnt int, parentVal int, n *Node, balance string, valueStore [][]int) {
	loopCnt++
	if loopCnt > 1 {
		// ツリーの空白を作るときに計算をするデータをvalueStoreに格納する
		if balance == LEFT {
			valueStore = append(valueStore, []int{n.value, parentVal})
		} else {
			valueStore = append(valueStore, []int{parentVal, n.value})
		}
	}
	if n.isNilNode() {
		return
	}

	printTree(root, loopCnt, n.value, n.right, RIGHT, valueStore)
	output(root, loopCnt, n, balance, valueStore)
	printTree(root, loopCnt, n.value, n.left, LEFT, valueStore)
}

func output(root int, loopCnt int, n *Node, balance string, valueStore [][]int) {
	if n.value == root {
		fmt.Println(fmt.Sprintf("%s", n.valueToStrng()))
		return
	}

	// valueStoreからスペース文字列を計算する
	var fillSpace string
	for _, vs := range valueStore[:len(valueStore)-1] {
		v := n.value
		if vs[0] < v && v < vs[1] {
			fillSpace += fmt.Sprintf("│%s", strings.Repeat(" ", LeafLength-1))
		} else {
			fillSpace += fmt.Sprintf("%s", strings.Repeat(" ", LeafLength))
		}
	}

	// 出力
	lineCnt := LeafLength - (len(fmt.Sprint(n.value)) + 1)
	if balance == LEFT {
		fmt.Println(fmt.Sprintf("%s└%s %s", fillSpace, strings.Repeat("─", lineCnt), n.valueToStrng()))
	} else {
		fmt.Println(fmt.Sprintf("%s┌%s %s", fillSpace, strings.Repeat("─", lineCnt), n.valueToStrng()))
	}
}
