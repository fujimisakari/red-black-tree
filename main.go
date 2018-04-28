package main

import (
	"flag"
	"strconv"
)

func addTree(rootNode *Node, v int) *Node {
	// 最初のノードを追加した場合はルートノードを作る
	if rootNode == nil {
		return newRootNode(v)
	}

	// 追加するまでのノード経路を取得(同じノードだった場合は終了)
	path := searchPath(rootNode, v)
	if path.noList() {
		return rootNode
	}

	// 新しいノードを追加
	parentNode := path.pop()
	chaildNode := newNode(v)
	if parentNode.value > v {
		parentNode.left = chaildNode
	} else {
		parentNode.right = chaildNode
	}

	// 親ノードが赤ならばツリーのバランスを整える
	if parentNode.isRed() {
		rootNode = rebalance(rootNode, chaildNode, parentNode, path)
	}
	return rootNode
}

func searchPath(root *Node, v int) *Path {
	n := root
	path := &Path{}
	for {
		if n == nilNode {
			break
		}

		path.list = append(path.list, n)
		switch {
		case n.value > v:
			n = n.left
		case n.value < v:
			n = n.right
		default:
			path = &Path{}
			n = nilNode
		}
	}
	return path
}

func rebalance(rootNode *Node, cn *Node, pn *Node, path *Path) *Node {
	var (
		grandNode     *Node
		parentNodeBro *Node
		parentNode    *Node = pn
		chaildNode    *Node = cn
	)
	for {
		// 祖父ノードから親ノード兄弟を取得
		grandNode = path.pop()
		if grandNode == nil {
			break
		}
		if grandNode.left == parentNode {
			parentNodeBro = grandNode.right
		} else {
			parentNodeBro = grandNode.left
		}

		// 親ノード兄弟が黒か赤でバランス処理を整える処理は大きく異なる
		if parentNodeBro.isBlack() {
			// 黒ならば回転処理させて終了
			rootNode = rotateTree(rootNode, chaildNode, parentNode, grandNode, path)
			break
		} else {
			// 赤ならば回転できないので、親ノードたちを黒に塗り替えて上のツリーを修正していく
			parentNode.withBlack()
			parentNodeBro.withBlack()

			// 祖父ノードがルートなら終了
			if grandNode == rootNode {
				break
			}

			// 祖父ノードを赤に書き換えて上のツリー修正を続行する
			grandNode.withRed()
			chaildNode = grandNode
			parentNode = path.pop()
		}
	}
	return rootNode
}

func rotateTree(rootNode *Node, cn *Node, pn *Node, gn *Node, path *Path) *Node {
	var (
		grandNode  *Node = gn
		parentNode *Node = pn
		chaildNode *Node = cn
		rotateNode *Node
	)
	// 祖父ノードの左か右かによって親ノード回転ロジックは反転させる。
	// また、親ノード回転には1回転のケースと2回転のケースがあり、子ノードがどちらにあるかで回転数が決まる
	if grandNode.left == parentNode {
		// 親ノードの右に子ノードがある場合であれば2回転のケースとなる
		if parentNode.right == chaildNode {
			grandNode.left = leftRotate(parentNode)
		}
		// 色の塗り替え、回転
		grandNode.withRed()
		grandNode.left.withBlack()
		rotateNode = rightRotate(grandNode)
	} else {
		// 親ノードの左に子ノードがある場合であれば2回転のケースとなる
		if parentNode.left == chaildNode {
			grandNode.right = rightRotate(parentNode)
		}
		// 色の塗り替え、回転
		grandNode.withRed()
		grandNode.right.withBlack()
		rotateNode = leftRotate(grandNode)
	}

	// 祖父ノードがルートだった場合は、回転済み祖父ノードをルートとして扱う
	if grandNode == rootNode {
		return rotateNode
	}

	// 祖父ノードがルートでない場合は、回転済み祖父ノードで上書く
	gGrandNode := path.pop()
	if gGrandNode.left == grandNode {
		gGrandNode.left = rotateNode
	} else {
		gGrandNode.right = rotateNode
	}
	return rootNode
}

func rightRotate(n *Node) *Node {
	ln := n.left
	n.left = ln.right
	ln.right = n
	return ln
}

func leftRotate(n *Node) *Node {
	rn := n.right
	n.right = rn.left
	rn.left = n
	return rn
}

func main() {
	// 引数を取得
	var values []int
	flag.Parse()
	args := flag.Args()
	for _, s := range args {
		if i, err := strconv.Atoi(s); err == nil {
			values = append(values, i)
		}
	}

	// ツリーの作成
	var n *Node
	for _, v := range values {
		n = addTree(n, v)
	}

	// ツリーの出力
	var valueStore [][]int
	printTree(n.value, 0, 0, n, "", valueStore)
}
