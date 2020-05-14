
class MakingChange:
    def __init__(self):
        self.son = 0
        self.daughter = 0
        self.observers = []

    def addSon(self):
        self.son+=1
        self.allNotify()
    
    def addDaughter(self):
        self.daughter+=1
        self.allNotify()
    
    def allNotify(self):
        for observer in self.observers:
            observer.update(self)
    
    def addObserver(self, observer):
        self.observers.append(observer)