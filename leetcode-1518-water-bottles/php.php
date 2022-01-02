class Solution {

    /**
     * @param Integer $numBottles
     * @param Integer $numExchange
     * @return Integer
     */
    function numWaterBottles($numBottles, $numExchange) {
        $res = $numBottles;

        while($numBottles >= $numExchange) {
            $tmp = intval($numBottles/$numExchange);

            $res += $tmp;
            $numBottles = $tmp + $numBottles % $numExchange;
        }

        return $res;
    }
}