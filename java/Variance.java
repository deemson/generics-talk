import java.util.ArrayList;
import java.util.List;

interface Animal {
    String makeSound();
}

class Cat implements Animal {
    @Override
    public String makeSound() {
        return "meow";
    }
}

class Dog implements Animal {
    @Override
    public String makeSound() {
        return "woof";
    }
}



public class Variance {
    public static void makeAllAnimalsDoSounds(List<? extends Animal> animals) {
        for (Animal animal: animals) {
            System.out.println(animal.makeSound());
        }
    }

    public static void main(String[] args) {
        List<Cat> cats = new ArrayList<>();
        cats.add(new Cat());
        cats.add(new Cat());
        makeAllAnimalsDoSounds(cats);
    }
}
