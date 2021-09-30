var moveZeros = function (arr) {
    return arr.filter(a => a !== 0).concat(arr.filter(a => a === 0))
}
