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
