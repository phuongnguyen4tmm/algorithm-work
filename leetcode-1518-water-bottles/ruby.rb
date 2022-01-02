# @param {Integer} num_bottles
# @param {Integer} num_exchange
# @return {Integer}
def num_water_bottles(num_bottles, num_exchange)
    res = num_bottles

    while num_bottles >= num_exchange
        div = num_bottles/num_exchange
        res += div.floor()
        num_bottles = div.floor() + num_bottles % num_exchange

    end

    return res

end