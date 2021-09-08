function divisors(integer) {
    res = [ ...Array(Math.floor(integer / 2) + 1).keys() ]
    .filter( x => x <= 1 ? false :
                 integer % x === 0 )
    return res.length === 0 ? `${integer} is prime` : res
}
