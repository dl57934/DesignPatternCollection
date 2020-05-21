from Parent import Parent


class StringParent(Parent):
    def __init__(self, string):
        self.string = string

    def openP(self):
        print("Open String")

    def close(self):
        print("Close String")

    def println(self):
        print(self.string)
