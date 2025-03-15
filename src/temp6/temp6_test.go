package temp6

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Value struct {
	Val       string
	expireSec int64 // 过期时间
}

// 带ttl的KV存储，需要线程安全
type TTLKV struct {
	store sync.Map
}

func CreateTTLKV() *TTLKV {

	store := sync.Map{}

	return &TTLKV{
		store: store,
	}

}

func (kv *TTLKV) Get(key string) (string, bool) {

	value, ok := kv.store.Load(key)

	if !ok {
		return "", false
	}

	valueEntry := value.(*Value)
	if valueEntry.expireSec <= 0 {
		return valueEntry.Val, true
	}

	now := time.Now().Unix()
	if valueEntry.expireSec <= now {
		// 已过期-》删除
		kv.store.Delete(key)
		return "", false
	}

	return valueEntry.Val, true
}

func (kv *TTLKV) Put(key string, value string, expireSec int64) {
	now := time.Now().Unix()
	val := &Value{
		Val:       value,
		expireSec: now + expireSec,
	}

	kv.store.Store(key, val)
}

func TestTTLKey(t *testing.T) {

	key := "testKey"
	value := "testValue"

	ttlKV := CreateTTLKV()
	ttlKV.Put(key, value, 2) // 2s
	getedValue, ok := ttlKV.Get(key)
	if ok {
		fmt.Println("1 -> get : ", getedValue)
	}

	time.Sleep(3 * time.Second)
	getedValue, ok = ttlKV.Get(key)
	if ok {
		fmt.Println("2 -> get : ", getedValue)
	} else {
		fmt.Println("3 -> nothing")
	}

}
