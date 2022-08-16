//Given an array of distinct elements and a number x, find if there is a pair with product equal to x.

function productNLOGN(value, array) {
    if (value != 0) {
        for (let i = 0; i < array.length; i++) {
            if ((array[i] > 0 && array[i] <= value) || (array[i] < 0 && array[i] >= value)) {
                for (let j = i; j < array.length; j++) {
                    if (array[i] * array[j] == value) {
                        return [i, j];
                    }
                }
            }
        }
    } else {
        if (array[0] == 0) {
            return [0, 1];
        }
        for (let i = 1; i < array.length; i++) {
            if (array[i] == 0) {
                return [0, i];
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