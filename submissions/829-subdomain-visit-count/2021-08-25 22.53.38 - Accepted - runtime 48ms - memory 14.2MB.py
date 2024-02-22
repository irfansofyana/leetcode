class Solution:
    def subdomainVisits(self, cpdomains: List[str]) -> List[str]:
        cnt = {}
        for cpdomain in cpdomains:
            space = cpdomain.find(' ')
            rep = int(cpdomain[:space])
            
            domains = cpdomain[space+1:].split('.')
            idx = len(domains) - 1
            curr = ""
            while (idx >= 0):
                if idx == len(domains)-1:
                    curr = domains[idx]
                else:
                    curr = domains[idx] + "." + curr
                if curr not in cnt:
                    cnt[curr] = rep
                else:
                    cnt[curr] += rep
                idx -= 1
        
        ans = []
        for k in cnt:
            ans.append(f"{cnt[k]} {k}")
        
        return ans