function bouncingBall(h,  bounce,  window) {
    if (h <= 0
        || bounce <= 0
        || bounce >= 1
        || window >= h) return -1
    
    let res = 0
    for (i = h; i > window; i *= bounce) {
      res += 2
    }
    return res - 1
}
