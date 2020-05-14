import Chulsu from "./Chulsu";
import YoungHee from "./Younghee";
import MakingChange from "./MakingChange";

const cs = new Chulsu();
const yh = new YoungHee();

const mc = new MakingChange();
mc.addObserver(cs);
mc.addObserver(yh);

mc.makingSon();
mc.makingDaughter();
