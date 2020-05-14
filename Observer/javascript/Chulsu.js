import Observer from "./Observer";

export default class Chulsu extends Observer {
  update(makingChange) {
    console.log(
      `Chulsu making Change son: ${makingChange.son} daughter ${makingChange.daughter}`
    );
  }
}
