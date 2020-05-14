public class Main {
    public static void main(String[] args) {
        Culsu ch = new Culsu();
        Younghee yh = new Younghee();

        MakingChange mc = new MakingChange();
        mc.addObserver(ch);
        mc.addObserver(yh);

        mc.makingSon();
        mc.makingDaughter();
    }
}