package merkledag
/* Has(key []byte) (bool, error)：用于检查指定的键是否存在于存储中，返回布尔值表示是否存在以及可能的错误。
Put(key, value []byte) error：用于向存储中添加或更新指定的键值对，如果操作成功则返回nil，否则返回错误。
Get(key []byte) ([]byte, error)：用于获取指定键对应的值，同时返回值和可能的错误。
Delete(key []byte) error：用于删除存储中指定的键值对，成功则返回nil，否则返回错误
 */
type KVStore interface {
	Has(key []byte) (bool, error)
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
}
