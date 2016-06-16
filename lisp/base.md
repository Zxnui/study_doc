## Common Lisp

推荐书籍:黑客与画家  
参考资料:http://acl.readthedocs.io/en/latest/index.html  
	On Lisp -田春译  

sudo apt-get install clisp  
clisp  
(+ 1 2)  
cd ~  
mkdir lisp  
cd lisp  
vi hello.lisp  
(defun helloworld () (format t "hello world"))  
clisp  
(load "~/clisp/hello.lisp")  
helloworld  
