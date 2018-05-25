package util

import "fmt"

type TrieNode struct {
	Child *map[string]TrieNode
	Exist bool
}

func (n *TrieNode) IsMatch(words string) bool {
	if words == "" {
		return false
	} else {
		runes := []rune(words)
		key := string(runes[0])
		theMap := *n.Child

		if _, ok := theMap[key]; !ok {
			return false
		} else {
			theNode := theMap[key]
			runesLen := len(runes)

			if runesLen == 1 {
				return theNode.Exist
			} else {
				if theNode.Child != nil {
					return theNode.IsMatch(string(runes[1:]))
				} else {
					return false
				}
			}
		}
	}
}

func (n *TrieNode) IsExist(words string) (bool, string) {
	if words == "" {
		return false, ""
	} else {
		runes := []rune(words)
		key := string(runes[0])
		theMap := *n.Child
		existStr := key

		if _, ok := theMap[key]; !ok {
			return false, ""
		} else {
			theNode := theMap[key]
			runesLen := len(runes)

			if theNode.Exist || runesLen == 1 {
				if theNode.Exist {
					return true, existStr
				} else {
					return false, ""
				}
			} else {
				if theNode.Child != nil {
					bo, str := theNode.IsExist(string(runes[1:]))
					if bo {
						return bo, existStr + str
					} else {
						return false, ""
					}
				} else {
					return false, ""
				}
			}
		}
	}
}

func (n *TrieNode) Traversal(deep int) {
	for k, v := range *n.Child {
		fmt.Println(deep, k, v)
		if v.Child != nil {
			v.Traversal(deep + 1)
		}
	}
}

func (n *TrieNode) AddWord(words string) {

	runes := []rune(words)
	keyStr := string(runes[0])

	var exist bool
	var restStr string
	if len(runes) == 1 {
		exist = true
		restStr = ""
	} else {
		exist = false
		restStr = string(runes[1:])
	}

	if n.Child == nil {
		tm := make(map[string]TrieNode)
		n.Child = &tm
	}

	tmpMap := *n.Child

	if _, ok := tmpMap[keyStr]; !ok {
		tmpMap[keyStr] = TrieNode{nil, exist}
	} else {
		if exist {
			tm := tmpMap[keyStr]
			tm.Exist = exist
			tmpMap[keyStr] = tm
		}
	}

	n.Child = &tmpMap

	if len(restStr) > 0 {
		tm := tmpMap[keyStr]
		tm.AddWord(restStr)
		tmpMap[keyStr] = tm
	}
}
