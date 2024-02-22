class Solution:
    def reachNumber(self, target: int) -> int:
        lo = 0
        hi = 10**9
        limit = hi
        target = abs(target)
        
        while lo <= hi:
            mid = (lo + hi) // 2
            sumInt = mid * (mid + 1) // 2
            if sumInt < target:
                lo = mid + 1
            else:
                limit = min(limit, mid)
                hi = mid - 1
        
        while (target-(limit*(limit+1)//2)) % 2 != 0:
            limit += 1
        
        return limit
                