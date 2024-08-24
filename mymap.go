package bootcamp

// import "fmt"

type item struct {
	Key   string
	Value interface{}
}

type MyMap struct { // Define fields and any necessary structures here
	items []item
}

func (m *MyMap) Set(k string, v interface{}) {
	for i := range m.items {
		if m.items[i].Key == k {
			m.items[i].Value = v
			return
		}
	}

	m.items = append(m.items, item{Key: k, Value: v})
}

func (m *MyMap) Get(k string) interface{} {
	for i := range m.items {
		if m.items[i].Key == k {
			return m.items[i].Value
		}
	}
	return nil
}

func (m *MyMap) Has(k string) bool {
	for i := range m.items {
		if m.items[i].Key == k {
			return true
		}
	}
	return false
}

func (m *MyMap) Delete(k string) {
	for i := range m.items {
		if m.items[i].Key == k {
			m.items = append(m.items[:i], m.items[i+1:]...)
			return
		}
	}
}

func (m *MyMap) Items() []struct {
	Key   string
	Value interface{}
} {
	result := make([]struct {
		Key   string
		Value interface{}
	}, len(m.items))

	for i, item := range m.items {
		result[i] = struct {
			Key   string
			Value interface{}
		}{
			Key:   item.Key,
			Value: item.Value,
		}
	}

	return result
}

// func main() {
// 	myMap := MyMap{}

// 	myMap.Set("key1", 42)
// 	myMap.Set("key2", "value2")

// 	fmt.Println(myMap.Get("key1")) // 42
// 	fmt.Println(myMap.Get("key2")) // value2
// 	fmt.Println(myMap.Has("key1")) // true
// 	fmt.Println(myMap.Has("key3")) // false

// 	myMap.Delete("key2")
// 	fmt.Println(myMap.Has("key2")) // false

// 	items := myMap.Items()
// 	for _, item := range items {
// 		fmt.Printf("Key: %s, Value: %v\n", item.Key, item.Value) // Key: key1, Value: 42
// 	}
// }
