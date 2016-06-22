#列表
- eql,eql 只有在它的参数是相同对象时才返回真
- equal,若它的参数打印出的值相同时，返回真

####将一个列表或变量赋值给其他变量或列表时，复制的是指针，可以eql等同.Lisp 有时会选择一个折衷的表示法，而不是指针。举例来说，因为一个小整数所需的内存空间，少于一个指针所需的空间，一个 Lisp 实现可能会直接处理这个小整数，而不是用指针来处理。

##copy-list
接受一个列表，然后返回此列表的复本。新的列表会有同样的元素，但是装在新的 Cons 对象里  
永远equal无法eql

##append
返回任何数目的列表串接  

##load载入程序
(load "test.lisp")  
某些场合，lisp文件，结尾可能是.lsp

##存储
- nth,获取元素  
>(nth 0 '(a,b,c))  
A
- nthcdr,找到第n个cdr  
>(nthcdr 2 '(a,b,c))  
(C)
- last,返回最后一个cons  
>(last '(a,b,c))  
(C)
##树
cons对象就是一棵树，car左子树，cdr右子树
##member
>(member 'b '(a,b,c))  
(B,C)
- member可以接受关键字key test  
>(member '(a) '((a) (z)) :test #'equal)  
((A) (Z))  
>(member 'a '((a b) (c d)) :key #'car)  
((A B) (C D))  
>(member 2 '((1) (2)) :key #'car :test #'equal)  
((2))  
>(member 2 '((1) (2)) :test #'equal :key #'car)  
((2))
##adjoin
>(adjoin 'b '(a b c))  
(A B C)  
>(adjoin 'z '(a b c))  
(Z A B C)
##union并集 intersection交集 complement补集
>(union '(a b c) '(c b s))  
(A C B S)  
>(intersection '(a b c) '(b b c))  
(B C)  
>(set-difference '(a b c d e) '(b e))  
(A C D)  
##序列
>(length '(a b c))  
3  
>(subseq '(a b c d) 1 2)//复制列表  
(B)  
>(subseq '(a b c d) 1)  
(B C D)  
>(reverse '(a b c))//颠倒  
(C B A)  
>(sort '(0 2 1 3 8) #'>)//排序，小心sort是破坏性的，它被允许修改被传入的序列  
(8 3 2 1 0)  
>(every #'> '(1 3 5) '(0 2 4))  
T  
##栈
>(setf x '(b))  
(B)  
>(push 'a x)//入栈  
(A B)  
>(pop x)//出栈  
(B)  
##关联
>(setf tran '((+ . "add") (- . "subtract")))  
((+ . "add") (- . "subtract"))  
>(assoc '+ trans)//根据给定的key，取出值  
(+ . "add")  
>(assoc '* trans)  
NIL  
##
