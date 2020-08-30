学习笔记

+ 树
    - 树的面试题解法一般都是用递归。原因：1树的节点是用递归的方式定义的；2.定义的重复性。
+ 递归
    - 递归终结条件
    - 处理当前层逻辑
    - 下探到下一层
    - 清理当前层（可选）
+ 写递归的注意事项
    - 不要进行人肉递归（看函数本身开始写）
    - 找到最近最简方法，将其拆解称可重复解决的问题（重复子问题）
    - 数学归纳法思维
+ 递归魔板
```
public void recur(int level, int param){
	// terminator
	if (level>MAX_LEVEL){
		// process result
		return ;
	}
	
	// process current logic
	process(level, param);
	
	// drill down
	recur(level:level+1,newParam);
	
	// restore current status
}
```

+ 回溯
    - 回溯算法好难！https://leetcode-cn.com/problems/permutations/solution/hui-su-suan-fa-xiang-jie-by-labuladong-2/