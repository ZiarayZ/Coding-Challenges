//find least number of coins required to make a desired amount, using only these values: 1, 5, 10, 25

//value = desired amount
function calculate_pennies(value) {
    if (!Number.isInteger(value) || value <= 0 || value === NaN) {
        return 0;
    }
    return Math.floor(value / 25) + Math.floor((value % 25) / 10) + Math.floor(((value % 25) % 10) / 5) + (((value % 25) % 10) % 5);
}
function calculate_dollars(value) {
    if (!(typeof value === "number") || value === NaN || value <= 0) {
        return 0;
    }
    return calculate_pennies(Math.floor(value*100));
}
//return number of coins

console.log(calculate_pennies(100));
console.log(calculate_pennies(121));
console.log(calculate_dollars("hello"));
console.log(calculate_dollars(1.23));
console.log(calculate_dollars(2.344536));