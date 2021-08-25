function DNAStrand(dna){
    return dna.split('')
              .map(x => x === 'A' ? 'T' :
                        x === 'T' ? 'A' :
                        x === 'G' ? 'C' : 'G')
              .join('')
    
}
