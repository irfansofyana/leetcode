func longestCommonPrefix(strs []string) string {
    var lenString int = 1
    var numString int = len(strs)
    var flag bool = true
    
    for flag {
        if lenString > len(strs[0]) {
            flag = false;
            break
        }
        var referenceStr string = strs[0][0:lenString]
        
        for i := 1; i < numString; i++ {
            if lenString > len(strs[i]) {
                flag = false;
                break
            }
            var currStr string = strs[i][0:lenString]
            if currStr != referenceStr {
                flag = false
                break
            }
        }
        if flag {
            lenString += 1
        }
    }
    
    if lenString == 1 {
        return ""
    } else {
        return strs[0][0:lenString-1]
    }
}