### 描述
给你一个仅包含小写英文字母和 '?' 字符的字符串 s，请你将所有的 '?' 转换为若干小写字母，使最终的字符串不包含任何 连续重复 的字符。

注意：你 不能 修改非 '?' 字符。

题目测试用例保证 除 '?' 字符 之外，不存在连续重复的字符。

在完成所有转换（可能无需转换）后返回最终的字符串。

如果有多个解决方案，请返回其中任何一个。

可以证明，在给定的约束条件下，答案总是存在的。

```
输入：s = "?zs"
输出："azs"
解释：该示例共有 25 种解决方案，从 "azs" 到 "yzs" 都是符合题目要求的。
     只有 "z" 是无效的修改，因为字符串 "zzs" 中有连续重复的两个 'z' 。
```

### 条件

- 1 <= s.length <= 100
- s 仅包含小写英文字母和 '?' 字符

### 思路
从题目的要求“不能有连续重复的字符”出发，可以直接得出下面的限制

```
任何 ？不能同前后两个字符相同
要决定 ？的值，需要先确定它前后两个字符
```
最简单的做法就是直接将 ？ 从 'a' 开始尝试跟前后两个值比较

