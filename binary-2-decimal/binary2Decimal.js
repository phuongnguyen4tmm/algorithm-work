function binary2Decimal(input) {
    let result = 0, len = input.length;

    for (let i = 0; i < len; i++) {
        let power = len - i - 1;
        let number = input.charCodeAt(i) - 48;
        console.log(number, input.charCodeAt(i))
        // result += Math.pow(power, number);
    }
    return result;
}