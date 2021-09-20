function iqTest(numbers){
    num = numbers.split(' ').map(x => parseInt(x))
    even = num.filter(x => !(x % 2))
    odd = num.filter(x => x % 2)
    return even.length === 1 
            ? num.indexOf(even[0]) + 1
            : num.indexOf(odd[0]) + 1
}
