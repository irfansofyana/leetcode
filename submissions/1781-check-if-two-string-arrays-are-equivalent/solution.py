class Solution:
    def arrayStringsAreEqual(self, word1: List[str], word2: List[str]) -> bool:
        words1 = ""
        for word in word1:
            words1 += word
        words2 = ""
        for word in word2:
            words2 += word
        return words1 == words2
