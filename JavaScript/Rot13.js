const rotate = (a, z, c) => c < a || c > z
                            ? c
                            : c + 13 <= z
                                ? c + 13
                                : c + 12 - (z - a)

function rot13(message) {
    return message
        .split('')
        .map(char => String.fromCharCode(
            rotate('a'.charCodeAt(),
                'z'.charCodeAt(),
                rotate('A'.charCodeAt(),
                    'Z'.charCodeAt(),
                    char.charCodeAt()
                )
            )
        ))
        .join('')
}
