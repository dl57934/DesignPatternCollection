public class Culsu implements Observer {
    public void update(MakingChange makingChange) {
        System.out.println("Culsu has son: " + makingChange.getSon() + " and daughter: " + makingChange.getDaughter());
    }
}