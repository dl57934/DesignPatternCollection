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
