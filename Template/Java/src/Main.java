public class Main {
    public static void main(String []args){
        Parent stringChild = new StringParent("Template Pattern");
        Parent charChild = new CharParent('t');

        stringChild.write();
        charChild.write();
    }
}
