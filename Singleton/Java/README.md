# Java Singleton

Java의 경우 일반적으로 New를 통해 인스턴트를 만들어준다. 하지만 Singleton은 매번 객체를 만들 수 없게 생성자를 만들 수 없게 한다.

그래서 class에 미리 인스턴트를 만든 후 생성자를 private로 막아서 새로운 객체를 만들 수 없게 합니다.
또 미리 만든 인스턴트를 호출할 수 있는 static method를 만들어서 인스턴트를 사용합니다.

Main.java

```java
public class Main {
    public static void main(String[] args) {
        Singleton singletonA = Singleton.getSingeton();
        Singleton singletonB = Singleton.getSingeton();
        if (singletonA.equals(singletonB)) {
            System.out.println("같다");
        } else {
            System.out.println("다르다");
        }
    }
}
```

Singleton.java

```java
public class Singleton {
    private static Singleton singleton = new Singleton();

    private Singleton(){ }

    public static Singleton getSingleton() {
        return singleton;
    }
}
```