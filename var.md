```
----|---|-----|-----------|
    |   |     |           |
  a | tv|   v | key1|key2 |
    |   |     |    引用    |
    |   |     |           |
----|---|-----|-----------|

a  , tv + v ,  [key1|key2]
         / 
 _ __  /
a  , tv + v ,  [key1|key2]
```

```
a = address           系统定义
tv = type of value    伴随v的变动而变动
v = value             可定义指向别的地址
key                   变量名,引用即是变量别名

```
0xc101 uint 8  a
0xc102 string "hello"  b
0xc103 []int []{1,2,3} c 
0xc104 *uint 0xc101 d

此时   d = *uint8 0xc101
      &d = **uint 01c104



对address  只能读  &&  （不能修改，但能创建）
对value  可读写

根据 tv+v的组合 形成了各类结构
  1.普通
  2.指针

引用是key的别名



