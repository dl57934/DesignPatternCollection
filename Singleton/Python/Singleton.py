class Singleton:
    _singleton = None

    def __new__(cls, *args, **kwargs):
        if not cls._singleton:
            cls._singleton = super(Singleton, cls).__new__(
                cls, *args, **kwargs)
        return cls._singleton
