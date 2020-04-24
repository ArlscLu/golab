# 结构指针

## 图示
```

---[type]-----[value]------[key]-------
-------------[address]-----------------


```
## eg
```
标识        |类型       |        |说明
------------|-----------|-----------|-----------
address    |系统定义|
type       |值类型  |      说明value的类型
value      |值     |     
key        |变量名  |     key = key（普通）| key = *_(引用类型)

eg.
*int
0xc1||*int 0xc2 a
0xc2||int 10 ''

slice：
a = *a
0xc3||*ints    0xc4   a
0xc4||*ints  0xc5,0xc6
0xc5||int 10
0xc6||int 12

map
a = *a eg    mary=>*mary=>**mary
             marry.Name=>*(marry.Name)
0xc3||*(*int&*string&...)   0xc4   a
0xc4||*(int&string&...)  0xc5,0xc6,0xc7,0xc8,0xc9
0xc5||int 0
0xc6||int 0
0xc7||int 0
0xc8||int 0
0xc9||int 0

0xc10|| int 13 
```

## 代码
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




