# 剑指offer

## [面试题05. 替换空格](https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/)

```
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
示例 1：
输入：s = "We are happy."
输出："We%20are%20happy."
限制：
0 <= s 的长度 <= 10000
```



解题思路：


日常解题法：本次提交使用比较简单的遍历替换法

```go
// 时间复杂度：O(n)遍历字符串s一遍
// 空间复杂度：O(n)，额外创建字符数组，长度为s的长度的3倍
func replaceSpace(s string) string {
    result := []rune{}
    for _,item :=range s{
        if item == ' '{
           result = append(result, []rune{'%','2','0'}...)
        }else{
            result = append(result,item)
        }
    }
    return string(result)
}
```

这种解法其实有一点不好，golang进行切片扩容时会有性能问题。



面试装逼法：推荐使用计数替换法，golang strings.replace 方法中使用的就是这种方法进行的字符串替换



```
// TODO
```



## [面试题07. 重建二叉树](https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/)

```
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
```



主要思路：

- 为中序遍历建立map映射
- 利用中序遍历和前序遍历规律对中序遍历数组进行分割，划分为不同的“子数组”进行递归操作

注意：公共变量的初始化写法

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 公共变量
var dict map[int]int   // 中序遍历映射字典
var preorderGlobal []int  // 全局前序遍历

// 主函数
func buildTree(preorder []int, inorder []int) *TreeNode {
    preorderGlobal = preorder
    dict =map[int]int{}
    for index,item := range inorder{
        dict[item] = index
    }
    return helper(0,0,len(inorder)-1)
}

// 递归函数 三个参数：前序遍历中root节点的下标，中序遍历左边界，中序遍历右边界
func helper(preRootIndex,inLeftIndex,InRightIndex int) *TreeNode{
  	// 递归出口
    if inLeftIndex>InRightIndex{
        return nil
    }
  	// 获取root节点在中序遍历中的下标位置
    inRootIndex := dict[preorderGlobal[preRootIndex]]
    root:= TreeNode{
        Val: preorderGlobal[preRootIndex],
        Left:helper(preRootIndex+1,inLeftIndex,inRootIndex-1), // 左子树构建
        Right:helper(preRootIndex+(inRootIndex-inLeftIndex+1),inRootIndex+1,InRightIndex), // 右子树构建：rootindex = 前序遍历root节点下标+左子树长度
    }
    return &root
}
```



## [面试题55 - I. 二叉树的深度](https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof/)



```
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
```



解题思路：

利用广度优先遍历思想，进行二叉树的层序遍历，每遍历一层深度+1

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root ==nil{
        return 0
    }
	// dfs 队列
	queue := []*TreeNode{}
	deep :=0
	queue = append(queue,root)

	for len(queue)>0{
    // 每次获取一层的结点，deep数量+1
		deep += 1
    len :=len(queue) 
		for i:=0; i<len;i++{
			top := queue[0]
			queue = queue[1:]

			if top.Left !=nil{
				queue = append(queue, top.Left)
			}
			if top.Right !=nil{
				queue = append(queue, top.Right)
			}
		}

	}
	return deep
}
```





