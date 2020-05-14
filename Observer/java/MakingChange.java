import java.util.ArrayList;
import java.util.Iterator;

public class MakingChange {
    private int son = 0;
    private int daughter = 0;
    ArrayList<Observer> observers = new ArrayList();

    public void makingSon() {
        son += 1;
        allNotify();
    }

    public void makingDaughter() {
        daughter += 1;
        allNotify();
    }

    public void addObserver(Observer observer) {
        observers.add(observer);
    }

    public void allNotify() {
        Iterator<Observer> it = observers.iterator();

        while (it.hasNext()) {
            Observer observer = it.next();
            observer.update(this);
        }
    }

    int getSon() {
        return son;
    }

    int getDaughter() {
        return daughter;
    }
}