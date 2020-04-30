# golang Singleton

golang에서는 class가 없기 때문에 구조체를 사용하여 구현한다.
인스턴스를 생성하는 함수에서 sync.Once를 사용한다. 이는 등록한 함수를 딱 한번만 실행하게 한다. 
그래서 인스턴스를 할당해주는 역할의 함수를 등록한다.

```go
package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age  int
}

var singleton *User
var once sync.Once

func GetUser() *User {
	once.Do(func() {
		singleton = &User{"hoony", 24}
	})
	return singleton
}

func (user *User) getName() string {
	return user.Name
}

func (user *User) setName(name string) {
	user.Name = name
}

func main() {
	user1 := GetUser()
	user2 := GetUser()

	if user1 == user2 {
		fmt.Println("같다")
	}
}
```