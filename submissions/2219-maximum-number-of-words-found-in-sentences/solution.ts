function mostWordsFound(sentences: string[]): number {
    let maxWords = 0;

    sentences.forEach(sentence => {
        const splitted = sentence.split(' ').filter(s => s != '')
        const numWords = splitted.length;
        if (numWords > maxWords) {
            maxWords = numWords;
        }
    });

    return maxWords;
};
