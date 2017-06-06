# Golang for...range
关键代码：
```
for k, r := range *rr {
    fmt.Printf("%dth r, id: %d, cpu: %f, mem: %f\n", k, r.ID, r.CPU, r.MEM)
    rs = append(rs, &r)
}
```


关键在`for...range`，for循环中的k,r是**临时变量**，`它们在确定之后地址不会改变`，但是值可以改变。rs = append(rs, &r)这里`&r`取r的地址一直是同一个值，循环结束后r的值是最后一次循环的值。

解决方法是是使用索引
```
for k, _ := range *rr {
    rs = append(rs, &(*r)[k])
}
```

[参考：聊聊Go中的Range关键字](https://xiaozhou.net/something-about-range-of-go-2016-04-10.html)