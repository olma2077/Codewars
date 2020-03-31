package kata

import "strings"

func Arrange(s string) string {
    sar := strings.Fields(s)
    for i := 1; i < len(sar); i += 2 {
        if len(sar[i]) < len(sar[i-1]) {
            sar[i], sar[i-1] = sar[i-1], sar[i]
        }
        sar[i-1] = strings.ToLower(sar[i-1])
        sar[i] = strings.ToUpper(sar[i])
        if i+1 < len(sar) {
            if len(sar[i]) < len(sar[i+1]) {
                sar[i], sar[i+1] = sar[i+1], sar[i]
            }
            sar[i] = strings.ToUpper(sar[i])
            sar[i+1] = strings.ToLower(sar[i+1])
        }
    }
    return strings.Join(sar, " ")
}
