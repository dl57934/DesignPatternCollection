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
