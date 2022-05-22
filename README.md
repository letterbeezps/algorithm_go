# Example Data Structure Go

[TOC]

总结在leetcode中使用的常见数据结构与算法，并给出一个基本的go语言实现案例和练习题。

## 排序 Sort

### **快速排序 quickSort** [example](Sort/quickSort/main.go)

[解析](https://studygolang.com/articles/35264)

**练习** leetcode

* [leetcode215](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)

### **快速排序非递归版本 quickSortByStack** [example](Sort/quickSortByStack/main.go)

[解析](https://studygolang.com/articles/35278)

### **归并排序 mergeSort** [example](Sort/mergeSort/main.go)

[解析](https://studygolang.com/articles/35298)

**练习** leetCode

* [leetcode21](https://leetcode-cn.com/problems/merge-two-sorted-lists/)

### **堆排序 heapSort** [example](Sort/heapSort/main.go)

[优先队列](Queue/PriorityQueue/main.go)
[解析](https://studygolang.com/articles/35556)

**练习** leetCode

* [leetCode23](https://leetcode-cn.com/problems/merge-k-sorted-lists/)

## 搜索 Search

### **二分查找 binarySearch** [example](Search/binarySearch/main.go)

[解析](https://studygolang.com/articles/35347)

* [leetCode33](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/) | [ans](Search/binarySearch/learn/leetCode33/main.go)

## 字符串

### **KMP** [example](String/Kmp/main.go)

**练习** leetcode

* [leetCode28](https://leetcode-cn.com/problems/implement-strstr/)

## 栈

### **Stack** [example](Stack/main.go)

**练习** leetcode

* [leetcode20](https://leetcode-cn.com/problems/valid-parentheses/)

### **至换** [example](Array/permeation/main.go)

## 树

### 二叉树

#### **前序遍历** [example](Tree/binaryTree/preorder/main.go)

**练习** leetcode

* [leetcode144](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)

#### **中序遍历** [example](Tree/binaryTree/inorder/main.go)

**练习** leetcode

* [leetcode94](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)

#### **后序遍历** [example](Tree/binaryTree/postorder/main.go)

**练习** leetcode

* [leetcode145](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)

## 图

### 最短路径

#### **Dijkstra** [example](Graph/dijkstra/main.go)

Dijkstra可以求一个带权（权重不能为负数）图内，某个点到所有点的最短路径大小

**练习** leetcode

* [leetcode743](https://leetcode.cn/problems/network-delay-time/)

## 数学运算

**大数相乘** [example](Math/Multiply/main.go)

* [leetcode43](https://leetcode-cn.com/problems/multiply-strings/)

**幂运算** [example](Math/Pow/main.go)

* [leetcode50](https://leetcode-cn.com/problems/powx-n/)

**二进制加法** [example](Math/BinaryAdd/main.go)

* [leetcode67](https://leetcode-cn.com/problems/add-binary/)

**求根号** [example](Math/Sqrt/main.go)

* [leetcode69](https://leetcode-cn.com/problems/sqrtx/)

**位运算实现整除** [example](Math/Divide/main.go)

* [leetcode](https://leetcode.cn/problems/divide-two-integers/)
