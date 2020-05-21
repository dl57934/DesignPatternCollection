# Python

Python의 경우 abstract 클래스가 특별히 없어서 메소드에 아래의 키워드를 사용해서 강제로 구현해주게 만들어 줄 수 있습니다.

```python
raise NotImplementedError
```

Parent.py
아래 코드를 보면 write 메소드에는 구현되지 않은 openP, close 등을 불러와서 크게 흐름은 정해져있습니다.

```python
class Parent:
    def openP(self):
        raise NotImplementedError

    def close(self):
        raise NotImplementedError

    def println(self):
        raise NotImplementedError

    def write(self):
        self.openP()
        for i in range(0, 5):
            self.println()
        self.close()
```

StringParent.py
직접적으로 openP, close, println 등을 직접 구현해주고 있습니다.

```python
from Parent import Parent

class StringParent(Parent):
    def __init__(self, string):
        self.string = string

    def openP(self):
        print("Open String")

    def close(self):
        print("Close String")

    def println(self):
        print(self.string)
```

main.py
main 문을 보면 최상위 클래스의 write 메소드를 호출해줍니다. 각 클래스에 openP, close, println을 어떻게 구현해주었는지에 따라 실행결과가 완전히 달라집니다.

```python
from StringParent import StringParent
from CharParent import CharParent

if __name__ == "__main__":
    stringParent = StringParent("string")
    charParent = CharParent("string")
    stringParent.write()
    charParent.write()
```