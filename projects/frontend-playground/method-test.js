function printName() {
    console.log(this.name)
}

const char1 = {
    name: "joe",
    getName: printName
}

char1.getName()