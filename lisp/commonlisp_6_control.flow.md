#控制流

##区块
###progn,顺序求值，返回最后一个表达式的值  
\>(progn  
    (format t "a")  
    (format t "b")  
    (+ 11 12))  
ab  
23  
###block,带有名字及紧急出口的 progn  
\>(block head//block区块名字  
    (format t "Here we go.")  
    (return-from head 'idea)//返回，跳出block，之后的表达式不会被求值  
    (format t "We'll never see this."))  
Here we go.  
IDEA  
- 也有一个 return 宏，它把传入的参数当做封闭区块 nil 的返回值：  
>(block nil  
    (return 27))  
27  
- 使用 defun 定义的函数主体，都隐含在一个与函数同名的区块，所以你可以：  
(defun foo ()  
  (return-from foo 27))  
###tagbody,基本不会被使用
\>(tagbody  
    (setf x 0)  
    top  
      (setf x (+ x 1))  
      (format t "~A " x)  
      (if (< x 10) (go top)))//若x<10.跳到top，重新执行  
1 2 3 4 5 6 7 8 9 10  
NIL  
##lambda
\>((lambda (x) (+ x 1)) 3)  
4  
##条件
- if.条件为真，执行第一个实参，假执行第二个实参
- when,条件为真，执行主体
- unless,条件为假，执行主体
- cond,允许多个条件  
(defun  test (a b)  
  (cond ((> a b) a)//若\a>b,则输出a，否则  
        (t b)//t是真，直接执行b  
  )  
)  
若符合条件后，没有表达式，则返回条件式的值  
>(cond (99))  
99  
- case,eql判断
(defun month-length (mon)  
  (case mon  
    ((jan mar may jul aug oct dec) 31)  
    ((apr jun sept nov) 30)  
    (feb (if (leap-year) 29 28))  
    (otherwise "unknown month")))//otherwise或t，如果没有子句符合时，或是子句只包含键值时，返回nil  
- typecase和case类似，不过eql比较改为typep
##迭代
- do，do隐含了block和tagbody，因此可以使用return,return-from,go  
>(dolist (x '(a b c d) 'done)  
    (format t "~A " x))  
A B C D  
DONE  
>(dotimes (x 5 x)  
  (format t "~A " x))  
0 1 2 3 4  
5  
- 函数 mapc 和 mapcar 很像，但不会 cons 一个新列表作为返回值，所以使用的唯一理由是为了副作用。它们比 dolist 来得灵活，因为可以同时遍历多个列表：
>(mapc #'(lambda (x y)  
          (format t "~A ~A  " x y))  
      '(hip flip slip)  
      '(hop flop slop))  
HIP HOP  FLIP FLOP  SLIP SLOP  
(HIP FLIP SLIP)  
总是返回 mapc 的第二个参数。  
##多值
- values 函数返回多个数值。它一个不少地返回你作为数值所传入的实参：
>(values 'a nil (+ 2 4))  
A  
NIL  
6  
- 如果一个 values 表达式，是函数主体最后求值的表达式，它所返回的数值变成函数的返回值
>((lambda () ((lambda () (values 1 2)))))  
1  
2  
- 然而若只预期一个返回值时，第一个之外的值会被舍弃：  
>(let ((x (values 1 2)))  
    x)  
1  
###要接收多个数值，我们使用 multiple-value-bind :
\>(multiple-value-bind (x y z) (values 1 2 3)  
    (list x y z))  
(1 2 3)  
\>(multiple-value-bind (x y z) (values 1 2)  
    (list x y z))  
(1 2 NIL)  
- 你可以借由 multiple-value-call 将多值作为实参传给第二个函数：  
>(multiple-value-call #'+ (values 1 2 3))  
6  
还有一个函数是 multiple-value-list :  
>(multiple-value-list (values 'a 'b 'c))  
(A B C)  
看起来像是使用 #'list 作为第一个参数的来调用 multiple-value-call 。
##中止
- catch throw  
(defun super ()  
  (catch 'abort  
    (sub)  
    (format t "We'll never see this.")))  
(defun sub ()  
  (throw 'abort 99))  
表达式依序求值，就像它们是在 progn 里一样。在这段代码里的任何地方，一个带有特定标签的 throw 会导致 catch 表达式直接返回：  
>(super)  
99  
- error
>(progn  
    (error "Oops!")  
    (format t "After the error."))  
Error: Oops!  
       Options :abort, :backtrace  
>>  
译注：2 个 >> 显示进入中断循环了。  
- unwind-protect,防止代码被 throw 与 error 打断,一个 unwind-protect 接受任何数量的实参，并返回第一个实参的值。然而即便是第一个实参的求值被打断时，剩下的表达式仍会被求值：  
>(setf x 1)  
1  
>(catch 'abort  
    (unwind-protect  
      (throw 'abort 99)  
      (setf x 2)))  
99  
>x  
2  

