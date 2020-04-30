public class Main {
    public static void main(String[] args) {
        Singleton singletonA = Singleton.getSingleton();
        Singleton singletonB = Singleton.getSingleton();
        if (singletonA.equals(singletonB)) {
            System.out.println("같다");
        } else {
            System.out.println("다르다");
        }
    }
}
