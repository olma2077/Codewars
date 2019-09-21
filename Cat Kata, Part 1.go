package kata

import "math"

// Measure distance, if cat is missing distance is infinite.
func Ruler(a, b [2]int) float64 {
    if a[0] == 0 || b[0] == 0 {
        return math.Inf(1)
    } else {
        return math.Sqrt(math.Pow(float64(a[0]-b[0]), 2) + math.Pow(float64(a[1]-b[1]), 2))
    }
}

func PeacefulYard(yard []string, minDistance int) bool {
    var L, M, R [2]int
    // Walk 2D array of chars to get cats' coordinates.
    // As missing cat will have {0, 0} coordinates, 
    // add 1 to found cats' coordinates.
    for i, s := range yard {
        for j, c := range s {
            switch c {
            case 'L':
                L = [2]int{i+1, j+1}
            case 'R':
                R = [2]int{i+1, j+1}
            case 'M':
                M = [2]int{i+1, j+1}
            }
        }
    }
    return Ruler(L, M) >= float64(minDistance) && 
           Ruler(L, R) >= float64(minDistance) &&
           Ruler(M, R) >= float64(minDistance)
}
