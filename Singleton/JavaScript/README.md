# JavaScript Singleton

javascript에서 singleton에서 오로지 함수를 사용하는 것과 클래스를 사용하는것 둘로 나누어 코드를 짤수 있다. 그러나 편의성에서는 클래스쪽이 더 좋은 것 같다. 그래도 둘다 알아보자!!

- es7 static을 제공해주기 때문에 static으로 class안에 변수를 만들어서 이 변수에 인스턴스가 할당됐으면 바로 반환해주고 아니면 할당 후 반환해주면 된다.
- es6에서는 class 밖에 변수를 선언한 것을 사용하여 인스턴스를 할당한 후 반환해주면된다.
  
```js
let instance; //es6

class Singleton {
  static instance; //es7
  constructor() {
    if (!instance) instance = this;
    return instance;
  }
}

const A = new Singleton();
const B = new Singleton();

if (A === B) console.log("같다");
else console.log("다르다");
```

함수를 사용해서 Singleton 만들기

우선 Singleton 인스턴스가 사용할 함수들을 먼저 정의합니다. 예제에서는 create()
그 후 getInstance 함수를 만들어서 인스턴스를 반환해줍니다.

printer 함수를 할당받는 것이 아닌 함수를 실행한 결과를 받고 있습니다. 그래서 함수가 한번만 실행되어 printer가 매번 동일하니 getInstance의 결과가 동일합니다.

```js
const printer = (() => {
  let instance;
  const create = () => {
    const print = () => {};
    const turnOn = () => {};

    return {
      print,
      turnOn,
    };
  };

  const getInstance = () => {
    if (!instance) {
      instance = create();
    }
    return instance;
  };

  return {
    getInstance,
  };
})();

const A = printer.getInstance();
const B = printer.getInstance();

if (A === B) console.log("같다");
else console.log("다르다");
```