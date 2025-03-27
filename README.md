# `github.com/boostgo/collection`

# Get started

Print event numbers from int slice

```go
package main

import (
	"fmt"
	
	"github.com/boostgo/collection/slicex"
)

func main() {
	numbers := slicex.Of([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	numbers.
		Filter(func(value int) bool {
			return value%2 == 0
		}).
		Each(func(idx int, value int) {
			fmt.Print(value, " ")
		})
	// 2 4 6 8 10

	fmt.Println()
	fmt.Println(numbers) // [1,2,3,4,5,6,7,8,9,10]
}
```

or without creating OfSlice instance

```go
package main

import (
	"fmt"
	
	"github.com/boostgo/collection/slicex"
)

func main() {
	fmt.Println(slicex.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(value int) bool {
		return value%2 == 0
	})) // [2 4 6 8 10]
}

```

# Slice

### Chunk

Divide slice to chunks by provided size

```go
package main

import (
	"fmt"
	
	"github.com/boostgo/collection/slicex"
)

func main() {
	chunks := slicex.Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	for _, chunk := range chunks {
		fmt.Println(chunk)
	}
	// [1 2 3]
	// [4 5 6]
	// [7 8 9]
	// [10]
}

```

### Iterator

Iterator to iterate by provided slice

**Methods:**
- **Next() bool** - state of current iterator index. Returns false if iterator comes to last index
- **Value() (T, bool)** - returns value by current index and state if value returned (returns false if index reached last index)
- **MustValue() T** - returns value by current index without state
- **Skip(count int) Iterator[T]** - skips iterations by provided skip count
- **Reverse() Iterator[T]** - reverse current iterator
- **Shuffle(source ...rand.Source) Iterator[T]** - shuffle current iterator. Could be provided custom rand.Source interface implementation
- **Each(fn func(int, T)) Iterator[T]** - iterate each element of iterator
- **Reset() Iterator[T]** - resets current index to 0
- **Source() [T]** - returns origin slice

```go
package main

import (
	"fmt"

	"github.com/boostgo/collection/slicex"
)

func main() {
	iterator := slicex.Iterate([]string{"value #1", "value #2", "value #3", "value #4", "value #5"})
	iterator.Skip(2)
	iterator.Each(func(idx int, value string) {
		fmt.Println(value)
	})

	// value #3
	// value #4
	// value #5
}

```

# Map

- **Keys(map[K comparable, V any]) []K** - returns all keys from map 
- **Values(map[K comparable, V any]) []V** - returns all values from map
- **Copy(map[K comparable, V any]) map[K]V** - returns copy of map

```go
package main

import (
	"fmt"

	"github.com/boostgo/collection/mapx"
)

func main() {
	m := map[string]string{
		"user1": "role1",
		"user2": "role2",
		"user3": "role2",
	}

	fmt.Println("keys:", mapx.Keys(m))
	fmt.Println("values:", mapx.Values(m))
	fmt.Println("copy:", mapx.Copy(m))

	// keys: [user1 user2 user3]
	// values: [role1 role2 role2]
	// copy: map[user1:role1 user2:role2 user3:role2]
}

```

### Async Map

Thread safe map with sync.Locker or sync.RWLocker support

```go
package main

import (
	"fmt"
	
	"github.com/boostgo/collection/mapx"
)

type Person struct {
	LastName  string
	FirstName string
}

func main() {
	asyncMap := mapx.NewAsyncMap[int, Person]()
	asyncMap.Store(1, Person{
		LastName:  "Johnson",
		FirstName: "John",
	})
	asyncMap.Store(2, Person{
		LastName:  "Johnson",
		FirstName: "Adam",
	})
	asyncMap.Each(func(id int, person Person) bool {
		fmt.Println(id, person.LastName, person.FirstName)
		return true
	})
	// 1 Johnson John
	// 2 Johnson Adam
}

```
