class Solution:
    def finalValueAfterOperations(self, operations: List[str]) -> int:
        ans = 0
        increments = ['++X', 'X++']
        decrements = ['--X', 'X--']

        for op in operations:
            if op in increments:
                ans += 1
            elif op in decrements:
                ans -= 1

        return ans        
