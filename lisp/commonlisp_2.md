##函数
>(defun our-third(x) (car (cdr (cdr (x)))))   
>(out-third '(a b c d))  
c  
##递归
某些事务，通过某种通用规则，计算结果，若结果不满足要求，则继续通过这种规则来计算;此规则称为递归
##format 输出
t表示输出被送到缺省的地方,~A被实参填充的位置，～%换行  
>(format t "~A + ~A = ~A.~%" 1 2 (+ 1 2))  
1 + 2 = 3.  
NIL  
##read 输入
read 是一个完整的 Lisp 解析器（parser）。不仅是可以读入字符，然后当作字符串返回它们。它解析它所读入的东西，并返回产生出来的 Lisp 对象  
(defun askem (string)  
	(format t "~A" string)  
	(read)) 
##let 引入局部变量
>(let ((x 1)(y 2))  
	(+ x y))  
3
##全局 defparameter
全局变量  
>(defparmeter *test* 12)  
全局常量  
>(defparmeter limit (+ *test* 1))  
函数boundp，用来检查某些符号是否为全局变量或常量  
>(boundp '*glob*)  
T
##setf 赋值
>(setf *test* 98)  
98  
- 如果 setf 的第一个实参是符号（symbol），且符号不是某个局部变量的名字，则 setf 把这个符号设为全局变量：  
>(setf x (list 'a 'b 'c))  
(A B C)  
>(setf (car x) 'n)  //setf 第一个参数(car x)为表达式或变量名时，第二个实参的值被插入第一个实参引用的位置  
N  
>x  
(N B C)  
- setf可以同时给某些变量赋值
>(setf a 'b  
       c 'd  
       e 'f)  
等同  
>(setf a 'b)  
>(setf c 'd)  
>(setf e 'f)  
##函数式编程
函数式编程意味着撰写利用返回值而工作的程序，而不是修改东西。它是 Lisp 的主流范式。大部分 Lisp 的内置函数被调用是为了取得返回值，而不是副作用。

##迭代
- do
(do (i start (+ i 1)) ((> i end) 'done) (format t "test is ~A ~%" i))  
- progn
progn 接受任意数量的表达式，依序求值，并返回最后一个表达式的值  

##函数做为对象function
function简写#'  

##apply
apply 接受一个函数和实参列表，并返回把传入函数应用在实参列表的结果  
apply 可以接受任意数量的实参，只要最后一个实参是列表即可  
>(apply #'+ 1 2 '(3 4 5))  
15  
##funcall
funcall和apply一样，区别是：不需要把实参包装成列表。  
>(funcall #'+ 1 2 3)  
6
