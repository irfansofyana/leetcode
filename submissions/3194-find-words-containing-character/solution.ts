function findWordsContaining(words: string[], x: string): number[] {
    let idx = [];

    for (let i = 0; i < words.length; i++) {
        if (words[i].indexOf(x) != -1) {
            idx.push(i);
        } 
    }

    return idx;
};
