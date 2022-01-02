impl Solution {
    pub fn num_water_bottles(num_bottles: i32, num_exchange: i32) -> i32 {
        let mut n = num_bottles;

        let mut res: i32 = n;

        while n >= num_exchange {
            let tmp = n/num_exchange;

            res += tmp;
            n = tmp + n % num_exchange;
        }

        return res;
    }
}