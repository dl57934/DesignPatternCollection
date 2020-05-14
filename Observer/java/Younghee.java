public class Younghee implements Observer {
    public void update(MakingChange makingChange) {
        System.out
                .println("Younghee has son: " + makingChange.getSon() + " and daughter: " + makingChange.getDaughter());
    }

}