package arrays

import "sort"

func ContainsDuplicate(input []int, complexity string) bool {

    switch (complexity){
        case "time":
            varMap := make(map[int]int)

            for _, v := range input{
                if varMap[v] == 1 {
                    return true
                }
                varMap[v] = 1
            }

            return false

        case "space":
            sort.Ints(input)
            for i:=0; i<len(input)-1; i++{
                if input[i] == input[i+1]{
                    return true
                }
            }

            return false

    }
    
    return false
}
