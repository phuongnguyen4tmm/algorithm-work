/**
 * @param {string} s
 * @param {number[]} indices
 * @return {string}
 */
var restoreString = function (s, indices) {
    let arr = "";
    for (let i = 0; i < indices.length; i++) {
        for (let j = 0; j < indices.length; j++) {
            if(indices[j] === i) {
                arr += s[j]
            }
        }
    }
    return arr;
};