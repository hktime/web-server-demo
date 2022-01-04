package framework

import (
	"errors"
	"strings"
)

type Tree struct {
	root *node
}

func NewTree()*Tree{
	root := newNode()
	return &Tree{root: root}
}

type node struct {
	isLast bool
	segment string
	handler ControllerHandler
	children []*node
}

func newNode()*node{
	return &node{
		isLast: false,
		segment: "",
		children: []*node{},
	}
}

func (tree *Tree) AddRoute(uri string, handler ControllerHandler)error{
	n := tree.root
	if n.matchNode(uri) != nil{
		return errors.New("route exists: " + uri)
	}
	segments := strings.Split(uri, "/")
	for idx, segment := range segments{
		if !isWildSegment(segment){
			segment = strings.ToUpper(segment)
		}
		isLast := idx == len(segments) -1
		var objNode *node
		childNodes := n.filterChildNodes(segment)
		if len(childNodes) > 0{
			for _, cnode := range childNodes{
				if cnode.segment == segment{
					objNode = cnode
					break
				}
			}
		}
		if objNode == nil{
			cnode := newNode()
			cnode.segment = segment
			if isLast{
				cnode.isLast = true
				cnode.handler = handler
			}
			n.children = append(n.children, cnode)
			objNode = cnode
		}
		n = objNode
	}
	return nil
}

func (tree *Tree) FindHandler(uri string) ControllerHandler{
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil{
		return nil
	}
	return matchNode.handler
}

// 判断segment是否是通用的，即是否以:开头
func isWildSegment(segment string)bool{
	return strings.HasPrefix(segment, ":")
}

func (n *node) filterChildNodes(segment string) []*node{
	if len(n.children) == 0{
		return nil
	}
	if isWildSegment(segment){
		return n.children
	}
	nodes := make([]*node, 0, len(n.children))
	for _, cnode := range n.children{
		if isWildSegment(cnode.segment){
			nodes = append(nodes, cnode)
		}else if cnode.segment == segment{
			nodes = append(nodes, cnode)
		}
	}
	return nodes
}

func (n *node) matchNode(uri string) *node{
	segments := strings.SplitN(uri, "/", 2)
	// 第一个部分用来匹配下一层子节点
	segment := segments[0]
	if !isWildSegment(segment){
		segment = strings.ToUpper(segment)
	}
	cnodes := n.filterChildNodes(segment)
	if cnodes == nil || len(cnodes) == 0{
		return nil
	}
	if len(segments) == 1{
		for _, tn := range cnodes{
			if tn.isLast{
				return tn
			}
		}
		return nil
	}
	for _, tn := range cnodes{
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil{
			return tnMatch
		}
	}
	return nil
}