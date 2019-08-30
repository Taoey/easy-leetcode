//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。 
//
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。 
//
// 示例: 
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
// 
// Related Topics 数组 哈希表

//微信公众号：【坂本先生】，leetcode刷题，算法，编程思考，Java/Go语言相关技术，一起进步吧！

package main
//go：两数之和


//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
    var result []int
    intMap := make(map[int]int)

    for index,num := range nums{
        another := target-num
        if _,ok := intMap[another];ok{
            result = append(result,intMap[another],index)
            return result
        }else {
            intMap[num]=index
        }
    }
    result = append(result, 0, 0)
    return result
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
    
}