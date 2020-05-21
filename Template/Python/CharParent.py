from Parent import Parent


class CharParent(Parent):
    def __init__(self, char):

        self.char = char

    def openP(self):
        print("Open Char Parent")

    def close(self):
        print("\nClose Char Parent")

    def println(self):
        print(self.char, end='')
