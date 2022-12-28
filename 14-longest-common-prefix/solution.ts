function longestCommonPrefix(strs: string[]): string {
    let prefixString = "";
    let minLength = strs[0].length;
    for(let i = 0 ; i< strs.length; i++) {
        minLength = Math.min(minLength, strs[i].length);
    }
    for(let i = 0 ; i< minLength; i++) {
        const current = strs[0][i];
        console.log(strs[0][i])
        for(let j = 1; j < strs.length; j++) {
            if (strs[j][i] !== current) {
                return prefixString;
            }
        }
        prefixString += current;
        console.log(prefixString)
    }
    return prefixString;
};