function numWaterBottles(numBottles: number, numExchange: number): number {
    let res = numBottles;

    while (numBottles >= numExchange) {
        res += Math.floor(numBottles/numExchange);
        numBottles = Math.floor(numBottles/numExchange) + numBottles%numExchange;
    }

    return res;
};