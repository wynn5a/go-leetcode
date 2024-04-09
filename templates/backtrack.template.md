
# Backtrack algorithm template

```go
func backtrack(candidates, path, answers) {
    if (满足结束条件) {
        answers.add(path)
        return
    }
    for( 选择：选择列表（可以动态生成) ) {
        做选择
        backtrack(candidates, path, answers)
        撤销选择
    }
}
```