package merkledag

import (
	"encoding/json"
	"strings"
)

// Hash2File 从 KVStore 中读取与给定 hash 对应的数据，并根据 path 返回对应的文件内容
func Hash2File(store KVStore, hash []byte, path string) []byte {
	// 递归函数，用于在树结构中查找对应的文件内容
	var findFileInTree func(interface{}, string) []byte
	findFileInTree = func(node interface{}, path string) []byte {
		switch n := node.(type) {
		case map[string]interface{}:
			segments := strings.Split(path, "/")
			if len(segments) == 0 {
				return []byte("Invalid path")
			}
			if child, ok := n[segments[0]]; ok {
				if len(segments) == 1 {
					// 找到对应的文件内容
					switch content := child.(type) {
					case string:
						return []byte(content)
					case []byte:
						return content
					default:
						return []byte("Invalid content type")
					}
				} else {
					// 继续向下查找
					return findFileInTree(child, strings.Join(segments[1:], "/"))
				}
			} else {
				return []byte("File not found")
			}
		default:
			return []byte("Invalid tree structure")
		}
	}

	// 从 KVStore 中读取 hash 对应的数据
	treeData, err := store.Get(hash)
	if err != nil {
		// 处理获取数据失败的情况，这里简单返回错误信息
		return []byte("Failed to retrieve data for the given hash")
	}

	// 解析树结构数据并查找对应的文件内容
	var tree interface{}
	err = json.Unmarshal(treeData, &tree)
	if err != nil {
		// JSON 解析出错，返回错误信息
		return []byte("Error parsing tree data")
	}

	// 根据 path 在树结构中查找文件内容并返回
	return findFileInTree(tree, path)
}