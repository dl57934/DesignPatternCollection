class Parent:
    def openP(self):
        raise NotImplementedError

    def close(self):
        raise NotImplementedError

    def println(self):
        raise NotImplementedError

    def write(self):
        self.openP()
        for i in range(0, 5):
            self.println()
        self.close()
