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
