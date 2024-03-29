package main

/**
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

来源：力扣（leetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
TODO 核心思路：将两个数组用分割线隔开，分割线需要满足以下条件：
		1.分割线左右两个数组的元素个数大致相同
		2.分割线左边的数组最大值≤分割线右边数组的元素最大值（最重要）
		3.如果总数为偶数，则取(分割线左边最大与分割线右边最小)/2
			如果总数为奇数，则取分割线左边最大的数
		2 4 6 | 15
		1 7 |8 10 17		中位数是7
	-
		3 8 |9 10
		2 4  6| 12 18 20	中位数是(8+9)/2 = 8.5
	-
		在核心思路的前提下，我们得知，当数组A和数组B的长度分别为m和n时
		偶数情况下，分割线左边的数据个数sizeLeft = (m+n)/2，又因为计算机是向下取整的，所以也等于(m+n)/2
		奇数情况下，分割线左边的数据个数sizeLeft = (m+n+1)/2
		根据这个规律，我们只要确定其中一个数组的分割线下标，就能确定另外一个数组的分割线下标(两者的和为sizeLeft)
	如何保证2.？
		数组A分割线左边的值必须≤数组B分割线右边的最小值
		数组B分割线左边的值必须≤数组A分割线右边的最小值
	当发现分割线不满足2. 我们需要适当调整分割线的位置
	-
	Q1:如何确定分割线？
	Q2:如何调整分割线的位置？
	-
	A2:
		2 4 |6 15
		1 7 8 |10 17
		很明显，8>6，不符合要求，这个时候需要将右侧最小值的分割线右移一位，左侧最大值的分割小左移一位，再递归判断
		2 4 6| 15
		1 7 |8 10 17
		-
		2 4 6 8| 10 17
		1 |7 15
		很明显，8>7，此时依旧要将右侧最小值的分割线右移一位，左侧最大值的分割小左移一位，再递归判断
		2 4 6 |810 17
		1 7 |15
*/
func main() {

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

}
