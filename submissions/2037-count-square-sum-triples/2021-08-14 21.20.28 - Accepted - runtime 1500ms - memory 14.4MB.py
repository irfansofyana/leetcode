class Solution:
    def countTriples(self, n: int) -> int:
        cnt = 0
        for a in range(1, n+1):
            for b in range(1, n+1):
                c = a ** 2 + b ** 2
                csqrt = sqrt(c)
                if int(csqrt) ** 2 == c and (csqrt >= 1 and csqrt <= n):
                    cnt += 1
        return cnt
        