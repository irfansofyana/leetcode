class Solution:
    def average(self, salary: List[int]) -> float:
        salary.sort()
        jum = sum(salary)
        jum -= (salary[0] + salary[-1])
        return jum / (len(salary)-2)
