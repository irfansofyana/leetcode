class Solution:
    def reverseWords(self, s: str) -> str:
        reversedWords = map(lambda p: p[::-1], s.split(" "))
        return " ".join(reversedWords)