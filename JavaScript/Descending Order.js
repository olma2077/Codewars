function descendingOrder(n){
    return eval(n.toString()
                 .split('')
                 .sort((a, b) => b - a)
                 .join('')
               )
}
