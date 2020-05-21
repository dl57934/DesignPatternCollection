# Java Template Pattern

Parent.java

abstract class를 사용하여 abstract method들과 큰 흐름을 정해주는 메소드를 설정해줍니다. 여기서는 write()입니다.

```java
public abstract class Parent {
    public abstract void open();
    public abstract void close();
    public abstract void print();


    public void write(){
        open();
        for(int i = 0; i < 5; i++)
            print();
        close();
    }
}
```

StringParent.java

Parent를 상속 받은 후 직접 구현해줍니다. 

```java
public class StringParent extends Parent{
    private String words;

    public StringParent(String words){
        this.words = words;
    }

    @Override
    public void open() {
        System.out.println("--- Open String Parent ---");
    }

    @Override
    public void close() {
        System.out.println("--- Close String Parent ---");
    }

    @Override
    public void print() {
        System.out.print("|");
        System.out.print(this.words);
        System.out.println("|");
    }
}
```

Main.java

Main.java에서는 StringParent, CharParent의 객체를 만들어 준 후 write를 실행해줍니다.
객체의 타입을 Parent로 해준 이유는 다형성을 위해서 해주었습니다.

```java
public class Main {
    public static void main(String []args){
        Parent stringChild = new StringParent("Template Pattern");
        Parent charChild = new CharParent('t');

        stringChild.write();
        charChild.write();
    }
}
```

output 

StringParent와 CharParent에서 최상위 클래스에서 받은 클래스를 어떻게 구현해주는지에 따라서 다른 결과값이 나옵니다.

```java
--- Open String Parent ---
|Template Pattern|
|Template Pattern|
|Template Pattern|
|Template Pattern|
|Template Pattern|
--- Close String Parent ---
<<<ttttt>>>
```