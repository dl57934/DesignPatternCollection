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
