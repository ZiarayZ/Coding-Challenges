//write a contains() function, assuming for strings?

String.prototype.contains = function (value) {
    let char = 0;
    let location = -1;
    for (let i = 0; i < this.length; i++) {
        if (this[i] === value[char]) {
            if (char == 0) {
                location = i;
            }
            char++;
        } else {
            char = 0;
            location = -1;
        }
        if (char == value.length) {
            return true;//or return location as index
        }
    }
    return false;//or return -1 as index
}

testingString = "TODO list: create games";
console.log(testingString.contains("todo"));
console.log(testingString.contains("list: cr"));
console.log(testingString.contains("TODO"));