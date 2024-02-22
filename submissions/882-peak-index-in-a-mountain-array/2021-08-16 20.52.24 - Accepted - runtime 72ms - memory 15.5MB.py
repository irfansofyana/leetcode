class Solution:
    def peakIndexInMountainArray(self, arr: List[int]) -> int:
        left = 0
        right = len(arr) - 1
        
        while abs(left-right) >= 3:
            m1 = (2 * left + right) // 3
            m2 = (left + 2 * right) // 3
            if arr[m1] <= arr[m2]:
                left = m1
            else:
                right = m2
        
        return (left + right) // 2