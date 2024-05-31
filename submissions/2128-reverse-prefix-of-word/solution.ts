function reversePrefix(word: string, ch: string): string {
    const pos = word.indexOf(ch)
    if (pos === -1) {
        return word
    } 
    
    const s1 = word.substr(0, pos+1)
    const s2 = word.substr(pos+1)
    return reverse(s1) + s2
};

function reverse(s: string): string {
    return s.split("").reverse().join("")
}
