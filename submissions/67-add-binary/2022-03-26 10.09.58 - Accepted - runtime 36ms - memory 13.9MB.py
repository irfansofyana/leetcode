class Solution:
    def addZeros(self, a: str, num_of_zeros: int) -> str:
        for i in range(num_of_zeros):
            a = '0' + a
        return a
    
    def addBinary(self, a: str, b: str) -> str:
        if len(a) < len(b):
            a = self.addZeros(a, len(b)-len(a))
        else:
            b = self.addZeros(b, len(a)-len(b))
        
        idx = len(a)-1
        take_away = 0
        
        ans = ''
        while idx >= 0:
            s = (int(a[idx]) + int(b[idx]) + take_away)
            ans = str(s % 2) + ans
            take_away = s // 2
            idx -= 1
        
        if take_away > 0:
            ans = '1' + ans
        
        return ans