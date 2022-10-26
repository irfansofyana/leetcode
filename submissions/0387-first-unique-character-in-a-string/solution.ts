function firstUniqChar(s: string): number {
    const freqMap = {};
    
    for (let i = 0; i < s.length; i++) {
        if (s[i] in freqMap) {
            freqMap[s[i]]++;
        } else {
            freqMap[s[i]] = 1;
        }
    }
    
    for (let i = 0; i < s.length; i++) {
        if (freqMap[s[i]] == 1) {
            return i;
        }
    }
    
    return -1;
};
