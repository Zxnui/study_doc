#函数
##全局函数
\> (fboundp '+)//+号是否和函数绑定  
T  
\> (symbol-function '+)//返回+号绑定的函数  
\#\<Compiled-function + 17BA4E\>  
可通过 symbol-function 给函数配置某个名字：  
(setf (symbol-function 'add2)  
  #'(lambda (x) (+ x 2)))  
新的全局函数可以这样定义，用起来和 defun 所定义的函数一样：  
\> (add2 1)  
3  
