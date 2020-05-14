class MakingChange {
  constructor() {
    this.son = 0;
    this.daughter = 0;
    this.observers = new Array();
  }

  makingSon() {
    this.son += 1;
    this.allNotify();
  }

  makingDaughter() {
    this.daughter += 1;
    this.allNotify();
  }

  allNotify() {
    this.observers.map((observer) => observer.update(this));
  }

  addObserver(observer) {
    this.observers.push(observer);
  }
}

export default MakingChange;
