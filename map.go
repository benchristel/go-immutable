package immutable

type Map struct {
	items map[string]interface{}
}

func NewMap() Map {
	return Map{}
}

func (m Map) Size() int {
	return len(m.items)
}

func (m Map) HasKey(key string) bool {
	_, exists := m.items[key]
	return exists
}

func (m Map) Get(key string) interface{} {
	return m.items[key]
}

func (m Map) Put(key string, value interface{}) Map {
	itemsCopy := make(map[string]interface{})
	for k, v := range m.items {
		itemsCopy[k] = v
	}
	itemsCopy[key] = value
	return Map{items: itemsCopy}
}

func (m Map) Remove(keyToRemove string) Map {
	var itemsCopy map[string]interface{}
	for k, v := range m.items {
		if k != keyToRemove {
			itemsCopy[k] = v
		}
	}
	return Map{items: itemsCopy}
}
