# js Template 패턴

js의 경우에도 특별히 abstract class가 존재하지 않기 때문에 메소드를 선언만 해놓고 구현을 하지는 않았습니다.

Parent.js

```javascript
export default class Parent {
  open() {}
  close() {}
  print() {}

  write() {
    this.open();
    for (let i = 0; i < 5; i++) this.print();
    this.close();
  }
}
```

StringParent.js

Parent를 상속받아 구현해주었다.

```js
import Parent from "./Parent";

export default class StringParent extends Parent {
  constructor(input) {
    super();
    this.input = input;
  }

  open() {
    super.open();
    console.log("OPEN String Parent\n");
  }

  close() {
    super.close();
    console.log("CLOSE String Parent");
  }

  print() {
    super.print();
    console.log(this.input);
  }
}
```

index.js

클래스를 통해서 객체를 만든 후 최상위 클래스의 write를 호출해주어서 write의 순서대로 StringParent에 구현해준대로 실행됩니다.

```js
import StringParent from "./StringParent";

const strParent = new StringParent("String Parent에요");

strParent.write();
```

output

```js
OPEN String Parent

String Parent에요
String Parent에요
String Parent에요
String Parent에요
String Parent에요
CLOSE String Parent
```