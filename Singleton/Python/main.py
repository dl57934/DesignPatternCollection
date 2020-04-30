from Singleton import Singleton

if __name__ == "__main__":
    a = Singleton()
    b = Singleton()

    if a == b:
        print("같다")
