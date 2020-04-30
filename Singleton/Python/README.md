# Python Singleton

python의 경우 객체를 생성될때 사용되는 __new__ 함수를 사용하여 singleton을 만들어 줄 수 있다.
__new__ class Method이기 때문에 부모가 아닌 자기자신의 클래스를 참조할 수 있다.

그래서 객체를 생성해줄때 멤버 변수가 None이면 인스턴스를 할당해주고 아니면 처음에 할당해준 멤버 변수를 반환해줍니다.

1. 맴버변수가 null이라면 __new__를 이용하여 인스턴트 생성 후 반환.
2. 멤버변수가 1번에 의해 인스턴트가 만들어졌으면 바로 반환.

main.py

```python
from Singleton import Singleton

if __name__ == "__main__":
    a = Singleton()
    b = Singleton()

    if a == b:
        print("같다")
```

Singleton.py

```python
class Singleton:
    _singleton = None

    def __new__(cls, *args, **kwargs):
        if not cls._singleton:
            cls._singleton = super(Singleton, cls).__new__(
                cls, *args, **kwargs)
        return cls._singleton
```