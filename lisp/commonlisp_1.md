##求值
规则：  
由左至右对实参求值，将它们的数值传入函数，来返回整个表达式的值  

例子：  
(/ (- 7 1) (- 4 2))  
- Lisp 对 (- 7 1) 求值: 7 求值为 7 ， 1 求值为 1 ，它们被传给函数 - ，返回 6 
- Lisp 对 (- 4 2) 求值: 4 求值为 4 ， 2 求值为 2 ，它们被传给函数 - ，返回 2 
- 数值 6 与 2 被传入函数 / ，返回 3 

函数调用顺序，基本都如上述例子所描述的。  

##:abort
若lisp出错，它会打印一个错误讯息，接着带你到一种叫做中断循环（break loop）的顶层，:bort跳出中断循环顶层

##quote
quote接受一个实参，什么都不做，原封不动的返回  
不遵循求值规则，common lisp定义'做为quote缩写  
>'(+ 1 2)
(+ 1 2)

##数据类型
- integer
- string
- symbol
>'test  
test
- lists
>(list 'my (+ 1 2) "test")  
(my 3 "test")

##nil空，对自身求值

##list操作

- cons组合列表
>(cons a (b c d))  
A B C D  
-  car取出列表第一个值，cdr除了第一个值，返回其他所有元素
>(car '(a b c))  
A  
>(cdr '(a b c))  
(B C)  
- third取出第三个元素值
>(third '(a b c d))  
c  

##真假

当函数返回值为真假时，函数为predicate,在commonlisp中，此种函数通常用p结尾  
- t真,对自身求值  
- nil假

>(listp 25)  
nil  
>(listp '(1 2 3))  
T  
>(null nil)  
T  
>(not nil)  
T  

##if
任何非nil的值都为真  

>(if (listp '(1 2 3)) (+ 1 2) (+ 1 3))  
3  
>(if (listp '1) (+ 1 2) (+ 1 3))  
4  
>(if (listp '1) (+ 1 2) )  
NIL  
>(if 1 2 3)  
2  

##and or
所有实参真，返回最后一个参数  
and其中实参为假，or其中一个实参为真，则停止对后面的实参求值  
>(and t (+ 1 2))  
3 
 
