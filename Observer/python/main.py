from Chulsu import Chulsu 
from Younghee import Younghee  
from MakingChange import MakingChange 

if __name__ == '__main__':
    cs = Chulsu()
    yh = Younghee()
    mc = MakingChange()
    mc.addObserver(yh)
    mc.addObserver(cs)
    mc.addSon()
    mc.addDaughter()
