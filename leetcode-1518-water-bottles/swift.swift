class Solution {
    func numWaterBottles(_ numBottles: Int, _ numExchange: Int) -> Int {
        var n = numBottles
        var res: Int = n

        while n >= numExchange {
            var tmp: Int = Int(n/numExchange)

            res += tmp
            n = tmp + n % numExchange
        }

        return res
    }
}