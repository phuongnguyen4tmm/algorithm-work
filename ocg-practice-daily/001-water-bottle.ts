// submission: https://leetcode.com/problems/water-bottles/submissions/858882533/`
function numWaterBottles(numBottles: number, numExchange: number): number {
    let res = numBottles;

    while (numBottles >= numExchange) {
        const exchangedBottle = Math.floor(numBottles/numExchange);
        const emptyBottleRest = numBottles % numExchange;
        res += exchangedBottle;

        numBottles = exchangedBottle + emptyBottleRest;
    }

    return res;
};
