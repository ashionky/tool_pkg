/**
 * @Author pibing
 * @create 2022/1/25 2:43 PM
 */

package dirtyfilter

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	var arr = []string{"fuck", "file"} //脏词

	var Trie = NewTrie()
	Trie.Add(arr, nil) //建树

	result := Trie.PrefixSearch("f")
	fmt.Println(len(result))
	s, h := Trie.Check("fuck ddd ", "")
	fmt.Println(s, h)

}
