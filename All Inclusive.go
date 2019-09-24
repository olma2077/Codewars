package kata

func ContainAllRots(strng string, arr []string) bool { 
    for i := 0; i < len(strng); i++ {
        pop := false
        rot := strng[i:] + strng[:i]
        for _, s := range(arr) {
            if s == rot {
                // We have to pop out of upper loop to continue
                pop = true
                break
            }
        }
        if !pop {
            return false
        }
    }
    return true
}
