# 结构指针

## 图示
```
----|---|-----|-----------|
  ta|   |     |           |
----|   |     |           |
  a | tv|   v | key1|key2 |
    |   |     |    引用    |
    |   |     |           |
----|---|-----|-----------|
```
## 定义
```
key                             变量名,引用即是变量别名
type@v = type of value          值类型|      与v共同
v = value                       值|         与tv共同
type@a = type of address        反应a对应的tv 
a = address                     系统定义

type@v,v  可以存  type@a, a
···
值类型：普通的 保存的是 实际值
···
···
引用类型：
指针变量 保存的是address
map/slice/channel/struct  保存的同指针 ，只不过赋了个别名  b = *a,并返回b
····

根据 tv+v的组合 形成了各类结构
  1.普通
  2.指针

```
## eg
```
0xc101 uint 8  a
0xc102 string "hello"  b
0xc103 []int []{1,2,3} c 
0xc104 *uint 0xc101 d

此时   d = *uint8 0xc101
      &d = **uint 01c104


```
```go
type  value  key
address

switch(type)
case *int :
  *int  0xc102412   a
  0xc222222
  other:
  int 10 aa
  0xc102412

case *map[int]string :
  *map[int]string    0x111111  c = *(0x111111)
  0x2222222
  other:
  map[int]string  map[int]string{1,2,3,4}  cc
  0x111111

case []int :
  *[]int  0x3333333  e = *(0x3333333)
  0x4444444
  other:
  [...]int  [3]int{6,7,8}    ee
  0x3333333


````




