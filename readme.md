
> 练习步骤

1. 5~10分钟思考

2. 有思路开始写代码，不然马上看题解

3. 默写背诵熟练

4. 闭卷自己写

5. 反复回顾至少五遍

6. 看官方和国际站的优秀代码

> 刷题误区

1. 只一遍

> 核心思维-升维

空间换时间

> 算法加速的四维？

增加一个维度，比如单链表搜索O(n)，升级维度变成跳表，空间增加了O(n)，但是复杂度 O(logn)

> log(n)

只要折半操作基本都是 log(n), 比如二分搜索，二叉搜索树

> 左右向中间收敛

两个指针分别指向两段，向中间移动

> 没有思路的提怎么办？

1. 看能不能暴力

2. 看基本的情况

3. 归纳基本情况, 找最近重复问题

```
比如 爬楼梯

无法暴力

看 一个台阶，两个台阶，三个台阶的情况

归纳，发现第三个台阶是之前一和两个台阶的和
```

> 基本思想

所有问题解决就是找重复



> 递归复杂度

化成递归树，看每次递归多少个节点，比如斐波拉契每次是上一层递归的2^level， 所以复杂度为 O(2^n)


> 递归本质

本质是函数帮你创建栈，所以可以用栈来解决,手动维护栈, 所以二叉树遍历可以用栈来解决
