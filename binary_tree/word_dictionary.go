package binary_tree

type WordDictionary struct {
	IsEnd    bool
	Children []*WordDictionary
	Chat     byte
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{Children: make([]*WordDictionary, 0)}
}

/** Adds a word into the data structure. */
func (w *WordDictionary) AddWord(word string) {
	length := len(word)
	cur := w
	for i := 0; i < length; i++ {
		char := word[i]
		j, childLen := 0, len(cur.Children)
		for ; j < childLen; j++ {
			if cur.Children[j].Chat == char {
				cur = cur.Children[j]
				break
			}
		}
		if j == childLen {
			child := &WordDictionary{
				Chat: char, Children: make([]*WordDictionary, 0),
			}
			cur.Children = append(cur.Children, child)
			cur = child
		}
	}
	cur.IsEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (w *WordDictionary) Search(word string) bool {
	length := len(word)
	cur := w
	for i := 0; i < length; i++ {
		char := word[i]
		j, childLen := 0, len(cur.Children)
		if char == '.' {
			for ; j < childLen; j++ {
				if true == cur.Children[j].Search(word[i+1:]) {
					return true
				}
			}
			return false
		} else {
			for ; j < childLen; j++ {
				if cur.Children[j].Chat == char {
					cur = cur.Children[j]
					break
				}
			}
			if j == childLen {
				return false
			}
		}
	}
	return cur.IsEnd
}
