/**
 * @Author pibing
 * @create 2021/12/29 5:53 PM
 */

package dirtyfilter

import (
	"strings"
)

type Node struct {
	//childs 用来当前节点的所有孩子节点
	childs map[rune]*Node
	term   bool //是否是终端节点
}

type Trie struct {
	root *Node
	size int
}

func NewNode() *Node {
	return &Node{
		childs: make(map[rune]*Node, 32),
	}
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

//假如我要把 敏感词： “tmd”
// Add([]{"tmd"}, nil)
func (p *Trie) Add(keywords []string, data interface{}) (err error) {
	for _, key := range keywords {
		key = strings.TrimSpace(key)
		node := p.root
		runes := []rune(key)
		for _, r := range runes {
			ret, ok := node.childs[r]
			if !ok {
				ret = NewNode()
				node.childs[r] = ret
			}

			node = ret
		}

		node.term = true
	}

	return
}

// findNode("tmd")
func (p *Trie) findNode(key string) (result *Node) {

	node := p.root
	chars := []rune(key)
	for _, v := range chars {
		ret, ok := node.childs[v]
		if !ok {
			return
		}

		node = ret
	}

	result = node
	return
}

func (p *Trie) collectNode(node *Node) (result []*Node) {

	if node == nil {
		return
	}

	if node.term {
		result = append(result, node)
		return
	}

	var queue []*Node
	queue = append(queue, node)

	for i := 0; i < len(queue); i++ {
		if queue[i].term {
			result = append(result, queue[i])
			continue
		}

		for _, v1 := range queue[i].childs {
			queue = append(queue, v1)
		}
	}

	return
}

//根据前缀查询脏词数量
func (p *Trie) PrefixSearch(key string) (result []*Node) {

	node := p.findNode(key)
	if node == nil {
		return
	}

	result = p.collectNode(node)
	return
}

// text = "我们都喜欢王八蛋"
// replace = "***"
func (p *Trie) Check(text, replace string) (result string, hit bool) {
	if replace == "" {
		replace = "***"
	}
	// 把text转换为rune数组，rune类型代表一个UTF-8字符，
	// 当需要处理中文、日文或者其他复合字符时，则需要用到rune类型
	chars := []rune(text)
	if p.root == nil {
		return
	}

	var left []rune // 检测后需返回的内容
	node := p.root  //从字典树的根开始查询
	start := 0
	for index, v := range chars {
		ret, ok := node.childs[v]
		if !ok { //不存在当前节点，则保留原字符
			left = append(left, chars[start:index+1]...)
			start = index + 1 //移动对比内容chars的角标
			node = p.root     //返回字典数的根节点继续开始寻找
			continue
		}

		node = ret
		if ret.term { //存在当前v，且节点没有子树，则替换当前v
			hit = true    // 是否存在脏词
			node = p.root //返回字典树的根，继续查询下个字符
			left = append(left, ([]rune(replace))...)
			start = index + 1 //移动对比内容chars的角标
			continue
		}
	}

	result = string(left)
	return
}
