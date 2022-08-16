//Given an array of distinct elements and a number x, find if there is a pair with product equal to x.

function productNLOGN(value, array) {
    for (let i = 0; i < array.length; i++) {
        for (let j = i; j < array.length; j++) {
            if (array[i] * array[j] == value) {
                return [i, j];
            }
        }
    }
    return [0, 0];
}

function product(value, array) {
    // calculate the factors of value, then compare to array
    return [0, 0];
}

console.log(productNLOGN(45, [1,5,15,18,46,9]));
console.log(product(45, [1,5,15,18,46,9]));