package dict

// Consumer is used to traversal dict, if it returns false the traversal will be break
type Consumer func(key string, val interface{}) bool

// Dict is interface of a key-value data structure
type Dict interface {
	Get(key string) (val interface{}, exists bool)
	Len() int
	Put(key string, val interface{}) (result int)
	PutIfAbsent(key string, val interface{}) (result int)
	PutIfExists(key string, val interface{}) (result int)
	Remove(key string) (result int)
	ForEach(consumer Consumer)
	Keys() []string
	RandomKeys(limit int) []string
	RandomDistinctKeys(limit int) []string
	Clear()
}

type EasyDict struct {
	m map[string]interface{}
}

func SoEasy() *EasyDict {
	return &EasyDict{m: map[string]interface{}{}}
}

// Get
//
//	@Description: get value
//	@receiver dict
//	@param key
//	@return val
//	@return exists
func (dict *EasyDict) Get(key string) (val interface{}, exists bool) {
	val, ok := dict.m[key]
	return val, ok
}

// Len
//
//	@Description: return length
//	@receiver dict
//	@return int
func (dict *EasyDict) Len() int {
	if dict.m == nil {
		panic("container is nil")
	}
	return len(dict.m)
}

// Put
//
//	@Description: set key-value
//	@receiver dict
//	@param key
//	@param val
//	@return result
func (dict *EasyDict) Put(key string, val interface{}) (result int) {
	_, existed := dict.m[key]
	dict.m[key] = val
	if existed {
		return 0
	}
	return 1
}

// PutIfAbsent
//
//	@Description: puts value if the key is not exists and returns the number of updated key-value
//	@receiver dict
//	@param key
//	@param val
//	@return result
func (dict *EasyDict) PutIfAbsent(key string, val interface{}) (result int) {
	_, existed := dict.m[key]
	if existed {
		return 0
	}
	dict.m[key] = val
	return 1
}

// PutIfExists
//
//	@Description: puts value if the key is exists and returns the number of inserted key-value
//	@receiver dict
//	@param key
//	@param val
//	@return result
func (dict *EasyDict) PutIfExists(key string, val interface{}) (result int) {
	_, existed := dict.m[key]
	if existed {
		dict.m[key] = val
		return 1
	}
	return 0
}

// Remove
//
//	@Description: removes the key and return the number of deleted key-value
//	@receiver dict
//	@param key
//	@return result
func (dict *EasyDict) Remove(key string) (result int) {
	_, existed := dict.m[key]
	delete(dict.m, key)
	if existed {
		return 1
	}
	return 0
}

// ForEach
//
//	@Description: traversal the dict
//	@receiver dict
//	@param consumer
func (dict *EasyDict) ForEach(consumer Consumer) {
	for k, v := range dict.m {
		if !consumer(k, v) {
			break
		}
	}
}

// Keys
//
//	@Description: returns all keys in dict
//	@receiver dict
//	@return []string
func (dict *EasyDict) Keys() []string {
	result := make([]string, len(dict.m))
	i := 0
	for k := range dict.m {
		result[i] = k
		i++
	}
	return result
}

// RandomKeys
//
//	@Description: randomly returns keys of the given number, may contain duplicated key
//	@receiver dict
//	@param limit
//	@return []string
func (dict *EasyDict) RandomKeys(limit int) []string {
	result := make([]string, limit)
	for i := 0; i < limit; i++ {
		for k := range dict.m {
			result[i] = k
			break
		}
	}
	return result
}

// RandomDistinctKeys
//
//	@Description:
//	@receiver dict
//	@param limit
//	@return []string
func (dict *EasyDict) RandomDistinctKeys(limit int) []string {
	size := limit
	if size > len(dict.m) {
		size = len(dict.m)
	}
	result := make([]string, size)
	i := 0
	for k := range dict.m {
		if i == size {
			break
		}
		result[i] = k
		i++
	}
	return result
}

// Clear
//
//	@Description: removes all keys in dict
//	@receiver dict
func (dict *EasyDict) Clear() {
	*dict = *SoEasy()
}
