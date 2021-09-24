function spinWords(string) {
    return string.split(' ')
        .map(item =>
            item.length < 5 ?
            item :
            item.split('')
                .reverse()
                .join('')
        )
        .join(' ')
}