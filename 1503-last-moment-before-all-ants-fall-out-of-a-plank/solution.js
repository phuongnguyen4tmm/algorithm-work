/**
 * @param {number[]} arr
 * @return {boolean}
 */
var canMakeArithmeticProgression = function (arr) {
    arr = arr.sort((a, b) => a - b);
    console.log(arr)
    const current = Math.abs(arr[0] - arr[1]);
    for (let i = 0; i < arr.length-1; i++) {
        if (Math.abs(arr[i] - arr[i + 1]) !== current) {
            return false;
        }
    }
    return true;
};