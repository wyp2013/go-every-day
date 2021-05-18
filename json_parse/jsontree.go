package json_parse

import (
	"encoding/json"
	"errors"
	"fmt"
	parser "github.com/buger/jsonparser"
)

const RootKey = "optionsxxx"

type Left struct {
	Type   string  `json:"type"`
	Field  string  `json:"field"`
}

type Item struct {
	Left   Left        `json:"left"`
	Op     string      `json:"op"`
	Right  interface{} `json:"right"`
}

type Node struct {
	Id          string
	Conjunction string
	Data        *Item
	Children    []*Node
}

func (n *Node) IsLeftNode() bool {
	return len(n.Children) == 0
}

func NewNode() *Node {
	return &Node{}
}

func parserGet(data []byte, key string) ([]byte, error) {
	if data, _, _, err := parser.Get(data, key); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

type Tree struct {
	root *Node
}

func (t *Tree) BuildTree(jsonData []byte) error {
	if err := t.checkParam(jsonData); err != nil {
		return err
	}

	if data, err := parserGet(jsonData, RootKey); err != nil {
		return err
	} else {
		t.root = NewNode()
		t.jsonParse(data, t.root)
	}

	return nil
}


func (t *Tree) checkParam(jsonData []byte) error {
	if len(jsonData) == 0 {
		return errors.New("empty json")
	}

	var v interface{}
	err := json.Unmarshal(jsonData, &v)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tree) jsonParse(data[] byte, parent *Node) {
	if id, err := parserGet(data, "id"); err == nil {
		parent.Id = string(id)
	}

	if conjunction, err := parserGet(data, "conjunction"); err == nil {
		parent.Conjunction = string(conjunction)
	} else {
		// todo 如果err不为空，根据业务要求具体处理
	}

	if children, err := parserGet(data, "children"); err == nil {
		// 不是叶子节点
		t.jsonParseArray(children, parent)
	} else {
		// 是叶子节点
		t.makeLeafNode(children, parent)
	}
}

func (t *Tree) jsonParseArray(data []byte, parent *Node) {

	parser.ArrayEach(data, func(value []byte, dataType parser.ValueType, offset int, err error) {
		if _, err := parserGet(value, "children"); err == nil {
			branchNode := t.makeBranchNode(parent)
			t.jsonParse(value, branchNode)
		} else {
			t.makeLeafNode(value, parent)
		}
	})
}

func (t *Tree) makeLeafNode(data []byte, parent *Node) (*Node, error) {
	leaf := NewNode()
	if err := json.Unmarshal(data, &leaf.Data); err != nil {
		// todo 打日志
	}

	parent.Children = append(parent.Children, leaf)
	return leaf, nil
}

func (t *Tree) makeBranchNode(parent *Node) *Node {
	node := NewNode()
	parent.Children = append(parent.Children, node)

	return node
}

// 广度遍历打印测试
func (t *Tree) bfsTree() {
	if t.root == nil {
		return
	}

	var queue []*Node
	queue = append(queue, t.root)

	for ;len(queue) > 0; {
		node := queue[0]
		queue = queue[1:]

		if node.IsLeftNode() {
			fmt.Println( "leafeNode: ", node.Data)
		} else {
			fmt.Println("branchNode",node.Conjunction, node.Id, node.Children)
			for _, v := range node.Children {
				queue = append(queue, v)
			}
		}
	}
}


