interface Animal {}

class Cat implements Animal {
  walk () {
    console.log('walking')
  }
}

class Fish implements Animal {
  swim () {
    console.log('swimming')
  }
}

function addFish(animals: Array<Animal>) {
  animals.push(new Fish())
}

const cats: Array<Cat> = [new Cat(), new Cat()]
for (const cat of cats) {
  cat.walk()
}
console.log("adding fish")
addFish(cats)
for (const cat of cats) {
  cat.walk()
}