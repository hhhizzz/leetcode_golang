package _208

import "testing"

func TestTrie(t *testing.T) {

	word := "hello"
	prefix := "he2"
	obj := Constructor()
	obj.Insert(word)

	param_2 := obj.Search(word)
	param_3 := obj.StartsWith(prefix)
	println(param_2)
	println(param_3)
	println(obj.Search(prefix))

}
