package arrays

func ContainsDuplicate(input []int) bool {

    varMap := make(map[int]int)

    for _, v := range input{
        if varMap[v] == 1 {
            return true
        }
        varMap[v] = 1
    }

    return false
}
