from Observer import Observer

class Younghee(Observer):
    def update(self, makingChange):
        print("Younghee has son: {} and daughter: {}".format(makingChange.son, makingChange.daughter))