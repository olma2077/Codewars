function digPow(n, p){
    sum = n.toString()
            .split('')
            .map(x => parseInt(x))
            .reduce((acc, rec, idx) => acc += rec ** (p + idx),
                    0)
    return sum % n ? -1 : sum / n
}
