function validSolution(board){
    // verify rows
    board.forEach(row => {
      if (!verify(row)) return false
    })
    
    // verify columns
    for (i = 0; i < 9; i++) {
      col = board.reduce((acc, rec) => [ ...acc, rec[i]], [])
      if (!verify(col)) return false
    }
    
    //verify blocks
    for (i = 0; i < 9; i++) {
      block = []
      for (j = 0; j < 3; j++) {
        block = [...block, 
                 ...board[j + 3 * ~~(i/3)]
                      .slice(0 + 3 * i%3, 3 + 3 * i%3)]
      }
      if (!verify(block)) return false
    }
    
    return true
  }
  
  function verify(block){
    const SUM_1_TO_9 = 45
    return block.reduce((acc, rec) => acc + rec, 0) === SUM_1_TO_9 
  }