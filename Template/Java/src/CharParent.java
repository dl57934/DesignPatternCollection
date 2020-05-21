public class CharParent extends Parent{
    private char letter;
    public CharParent(char letter){
        this.letter = letter;
    }
    @Override
    public void open() {
        System.out.print("<<<");
    }

    @Override
    public void close() {
        System.out.print(">>>");
    }

    @Override
    public void print() {
        System.out.print(this.letter);
    }
}
