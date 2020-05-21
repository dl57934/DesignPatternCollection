# Go Template Pattern

Golang에서는 명확히 abstract Class가 존재하지 않는다. 아니 클래스 자체가 존재하지 않는다. 그래서 비슷하게 구현하기 위해서 interface와 struct에 method를 만들어서 구현해주었습니다.

Parent.go

interface를 만들어 주었다. 다른 언어들과 달리 abstract class와 같은 것은 존재하지 않기 때문에 main 문에 추가로 처리 해주었습니다.

```go
package Test

type Parent interface{
	Open()
	Close()
	Print()
}
```

StringParent.go

Open, Close, Print 등을 직접구현하여 Parent를 타입으로 사용할 수 있습니다. go에서는 interface의 메소드들을 구현해줄 때 직접적으로 상속하기 보다는 interface의 함수들을 직접적으로 다 구현해준 객체이면 interface의 타입으로 사용할 수 있습니다. main을 통해 확인할 수 있다.

```go
package Test

import "fmt"

type StringParent struct {
	Words string
}

func (word StringParent) Open() {
	fmt.Println("Open String")
}

func (word StringParent) Close() {
	fmt.Println("Close String")
}

func (word StringParent) Print() {
	fmt.Println(word.Words)
}
```

main.go

StringParent의 구현된 것을 writeLine에 넣어주는 것을 보아 StringParent가 Parent로 받아질 수 있음을 알 수 있다. 그래서 차례대로 Open, Close, Print를 출력해준다.

```go
package main

import "main.go/Test"

func main() {
	stringParent := Test.StringParent{"String"}
	charParent := Test.CharParent{"Char"}
	writeLine(stringParent, charParent)
}

func writeLine(parents ...Test.Parent) {
	for _, parent := range parents {
		parent.Open()
		parent.Print()
		parent.Close()
	}
}
```


Output

```go
Open String
String
Close String
open
Char
close
```