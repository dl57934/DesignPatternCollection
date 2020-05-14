import Observer from "./Observer";
export default class YoungHee extends Observer {
  update(makingChange) {
    console.log(
      `YoungHee making Change son: ${makingChange.son} daughter ${makingChange.daughter}`
    );
  }
}
