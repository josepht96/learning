const obj = {
    name: "Alice",
    greet: function () {
        console.log(this.name)
        setInterval(function () {
            console.log(this.name);  // undefined! 'this' is different here
        }, 1000);
    }
};

const obj2 = {
    name: "Alice2",
    greet: function () {
        console.log(this.name)
        setInterval(() => {
            console.log(this.name);  // "Alice" - uses obj's 'this'
        }, 1000);
    }
};

const obj3 = {
    name: "Alice3",
    greet: function () {
        console.log(this.name)
        setInterval(() => {
            console.log(this.name)
        }, 1000)
    }
};

const obj4 = {
    name: "this is a full test",
    greet: function() {
        console.log("This is the function call")
        setInterval(() => {
            console.log(this.name)
        }, 1000)
    }
};

function test123() {
    console.log(this.test)
}

// obj.greet()
// obj2.greet()
// obj3.greet()
// obj4.greet()

// test123.call({ test: "name"})
obj4.greet.call({name: "test123"})