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
