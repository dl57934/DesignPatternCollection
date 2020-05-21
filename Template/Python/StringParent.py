from Parent import Parent


class StringParent(Parent):
    def __init__(self, string):
        self.string = string

    def openP(self):
        print("Open")

    def close(self):
        print("Close")

    def println(self):
        print(self.string)
