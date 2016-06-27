#数据结构
数组，结构，哈希表，存储速度更快，使用空间更少
##数组
- Common Lisp 的数组至少可以达到七个维度，每个维度至少可以容纳 1023 个元素。  
- :initial-element 实参是选择性的。如果有提供这个实参，整个数组会用这个值作为初始值。若试着取出未初始化的数组内的元素，其结果为未定义（undefined）。
>(setf arr (make-array '(2 3) :initial-element nil))//make-array构造2x3数组  
>(aref arr 0 0取出0,0数组元素  
NIL  
要替换数组的某个元素，我们使用setf与aref：  
>(setf (aref arr 0 0) 'b)  
B  
>(aref arr 0 0)  
B  
- 要表示字面常量的数组（literal array），使用 #na 语法，其中 n 是数组的维度。举例来说，我们可以这样表示 arr 这个数组：  
\#2a((b nil nil) (nil nil nil))  
- 如果全局变量 *print-array* 为真，则数组会用以下形式来显示：
>(setf *print-array* t)  
T  
>arr  
\#2A((B NIL NIL) (NIL NIL NIL))  
###一维数组
>(setf vec (make-array 4 :initial-element nil))  
\#(NIL NIL NIL NIL)  
- 一维数组又称为向量（vector）。你可以通过调用 vector 来一步骤构造及填满向量，向量的元素可以是任何类型：  
>(vector "a" 'b 3)  
\#("a" b 3)  
- 可以用 aref 来存取向量，但有一个更快的函数叫做 svref ，专门用来存取向量。  
>(svref vec 0)  
NIL  

##字符和字符串
每个字符都有一个相关的整数 ── 通常是 ASCII 码，但不一定是。在多数的 Lisp 实现里，函数 char-code 返回与字符相关的数字，而 code-char 返回与数字相关的字符。  
字符比较函数 char< （小于）， char<= （小于等于)， char= （等于)， char>= （大于等于) ， char> （大于)，以及 char/= （不同)  
\>(sort "elbow" #'char<)//从小到达排列字符  
"below"  
\>(char "abc" 1)//取出字符  
\#\\b  
\>(equal "fred" "fred")  
T  
\>(equal "fred" "Fred")  
NIL  
\>(string-equal "fred" "Fred")//忽略大小写，比较字符串  
T  
\>(format nil "~A or ~A" "truth" "dare")//创建字符串  
"truth or dare"  
\>(concatenate 'string "not " "to worry")//将字符串串接起来  
"not to worry"  
###字符串，当作序列来操作
\>(mirror? "abba")//对称  
T  
**我们已经看过四种用来取出序列元素的函数： 给列表使用的 nth ， 给向量使用的 aref 及 svref ，以及给字符串使用的 char 。 Common Lisp 也提供了通用的 elt ，对任何种类的序列都有效**  
\>(elt '(a b c) 1)//elt效率不如其他特定的取出函数  
B  
###许多序列函数接受一个或多个，由下表所列的标准关键字参数：
参数      用途                    缺省值  
:key      应用至每个元素的函数    identity  
:test     作来比较的函数          eql  
:from-end 若为真，反向工作。      nil  
:start    起始位置                0  
:end      若有给定，结束位置      nil  
- 一个接受所有关键字参数的函数是 position ，返回序列中一个元素的位置，未找到元素时则返回 nil 。我们使用 position 来演示关键字参数所扮演的角色。  
>(position #\a "fantasia")  
1  
>(position #\a "fantasia" :start 3 :end 5)  
4  
- start 关键字参数是第一个被考虑的元素位置，缺省是序列的第一个元素。 :end 关键字参数，如果有给的话，是第一个不被考虑的元素位置  
- from-end 关键字参数  
>(position #\a "fantasia" :from-end t)  
7  
寻找关键字，从尾部找起，我们得到最靠近结尾的a的位置。但位置是像平常那样计算；而不是从尾端算回来的距离。  
- key 关键字参数是序列中每个元素在被考虑之前，应用至元素上的函数。如果我们说  
>(position 'a '((c d) (a b)) :key #'car)  
1  
那么我们要找的是，元素的 car 部分是符号 a 的第一个元素。  
- test 关键字参数接受需要两个实参的函数，并定义了怎样是一个成功的匹配。缺省函数为 eql 。如果你想要匹配一个列表，你也许想使用 equal 来取代  
>(position '(a b) '((a b) (c d)))  
NIL  
>(position '(a b) '((a b) (c d)) :test #'equal)  
0  
- 要找到满足谓词的元素，其中谓词接受一个实参，我们使用 position-if 。它接受一个函数与序列，并返回第一个满足此函数的元素  
>(position-if #'oddp '(2 3 4 5))  
1  
**position-if 接受除了 :test 之外的所有关键字参数**

##结构
\>(defstruct point x y)  
这里定义了一个 point 结构，具有两个字段 x 与 y 。同时隐式地定义了 make-point 、 point-p 、 copy-point 、 point-x 及 point-y 函数。  
##哈希表
列表可以用来表示集合（sets）与映射（mappings）。但当列表的长度大幅上升时（或是 10 个元素），使用哈希表的速度比较快。你通过调用 make-hash-table 来构造一个哈希表，它不需要传入参数  
\>(setf ht (make-hash-table))  
\#\<Hash-Table BF0A96\>  
- 获取哈希值  
>(gethash 'color ht)  
NIL  
NIL  
在这里我们首次看到 Common Lisp 最突出的特色之一：一个表达式竟然可以返回多个数值。函数 gethash 返回两个数值。第一个值是与键值有关的数值，第二个值说明了哈希表是否含有任何用此键值来储存的数值。由于第二个值是 nil ，我们知道第一个 nil 是缺省的返回值，而不是因为 nil 是与 color 有关的数值。  
>(setf (gethash 'color ht) 'red)//赋值  
RED  
>(gethash 'color ht)  
RED  
T  
- 插入哈希  
>(setf (gethash 'apricot ht) t)//设置key=apricot的哈希，空集合  
T  
>(gethash 'apricot ht)  
T  
T  
- 删除哈希  
>(remhash 'apricot ht)//移除  
T  
- 迭代  
哈希表有一个迭代函数： maphash ，它接受两个实参，接受两个参数的函数以及哈希表。该函数会被每个键值对调用，没有特定的顺序  
>(setf (gethash 'shape ht) 'spherical  
        (gethash 'size ht) 'giant)  
GIANT  
>(maphash #'(lambda (k v)  
               (format t "~A = ~A~%" k v))  
           ht)  
SHAPE = SPHERICAL  
SIZE = GIANT  
COLOR = RED  
NIL 
maphash 总是返回 nil ，但你可以通过传入一个会累积数值的函数，把哈希表的词条存在列表里。
###哈希表可以容纳任何数量的元素，但当哈希表空间用完时，它们会被扩张。如果你想要确保一个哈希表，从特定数量的元素空间大小开始时，可以给 make-hash-table 一个选择性的 :size 关键字参数。做这件事情有两个理由：因为你知道哈希表会变得很大，你想要避免扩张它；或是因为你知道哈希表会是很小，你不想要浪费内存。 :size 参数不仅指定了哈希表的空间，也指定了元素的数量。平均来说，在被扩张前所能够容纳的数量。所以  
(make-hash-table :size 5)  
会返回一个预期存放五个元素的哈希表。  
和任何牵涉到查询的结构一样，哈希表一定有某种比较键值的概念。预设是使用 eql ，但你可以提供一个额外的关键字参数 :test 来告诉哈希表要使用 eq ， equal ，还是 equalp  
