func canConstruct(ransomNote string, magazine string) bool {
    magazineArr := make([]int, 26)
    ransomNoteArr := make([]int, 26)
    
    for i := 0; i < len(magazine); i++ {
        ch := int(magazine[i]) - 97
        magazineArr[ch]++
    }
    
    for i := 0; i < len(ransomNote); i++ {
        ch := int(ransomNote[i]) - 97
        ransomNoteArr[ch]++
    }
    
    for i := 0; i < 26; i++ {
        if ransomNoteArr[i] > magazineArr[i] {
            return false
        }
    }
    
    return true
}