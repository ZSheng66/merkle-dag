package merkledag

import "hash"

func Add(store KVStore, node Node, h hash.Hash) []byte {
    // 将对象序列化为字节数组
    data, err := Encode(node)
    if err != nil {
        log.Errorf("Failed to encode node: %v", err)
        return nil
    }

    // 对字节数组进行哈希计算并记录哈希值
    h.Reset()
    _, err = h.Write(data)
    if err != nil {
        log.Errorf("Failed to hash data: %v", err)
        return nil
    }
    hashVal := h.Sum(nil)
    
    // 将哈希值和字节数组写入到 KVStore 中
    err = store.Put(hashVal, data)
    if err != nil {
        log.Errorf("Failed to put data into KVStore: %v", err)
        return nil
    }

    // 返回哈希链根节点
    return hashVal
}
