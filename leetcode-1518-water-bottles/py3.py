class Solution:
    def numWaterBottles(self, numBottles: int, numExchange: int) -> int:
        res = numBottles
        while numBottles >= numExchange:
            res += int(numBottles/numExchange)
            numBottles = int(numBottles/ numExchange) + numBottles % numExchange

        return int(res)