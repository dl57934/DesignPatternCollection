from Observer import Observer

class Chulsu(Observer):
    def update(self, makingChange):
        print("Chulsu has son: {} and daughter: {}".format(makingChange.son, makingChange.daughter))