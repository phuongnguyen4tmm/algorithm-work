object Solution {
    def numWaterBottles(numBottles: Int, numExchange: Int): Int = {
        var n = numBottles
        var res = n

        while (n >= numExchange) {
            var tmp = (n/numExchange).toInt
            res += tmp
            n = tmp + n % numExchange
        }

        return res
    }
}