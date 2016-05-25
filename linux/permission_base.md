#权限
-rwxrwxrwx 1 root root 1000 time fileName

##权限人
实例中root root  
第一个root,拥有档案的人  
第二个root,拥有档案的群组

##链接数
实例中，数字1 

##容量
实例中，1000表示档案容量，单位bytes

##档案名称
fileName,若名称前多个. 表示此文件为隐藏档案,例如.fileName

##读写
-rwxrwxrwx,其中-rwxrwxrwx共10个字符
- 第一个字符：- 档案;d 目录;l 链接档;b 供存储设备;c 一次性读取设备(鼠标/键盘);
- rwx,r 可读，w 可写,x 可执行;若无权限，则-代替，例如r-x，可读不可写可执行
- 剩余的rwx，实例中出现了3次，第一组：档案拥有者具有的权限;第二组:群组拥有权限;第三组:非本人，非群组，其他人权限

##root用户，拥有所有文件，最高权限

###权限修改
- chown zxnui:zxnui fileName	权限人修改
- chmod 777 fileName		权限修改
r:4  
w:2  
x:1  
举例：  
-rw-rw-r-- fileName  
chmod 111 fileName ---> ---x--x--x fileName  
chmod 222 fileName ---> --w--w--w- fileName  
chmod 555 fileName ---> -r--r--r-- fileName  
chmod 333 fileName ---> --wx-wx-wx fileName  
chmod 761 fileName ---> -rwxrw---x fileName  
chmod 770 fileName ---> -rwxrwx--- fileName  

- 符号改变权限
u:user g:group o:other a:all  
+:增加 -:减去 =:设定  
举例：  
-rw-rw-r-- fileName  
chmod u+x fileName ---> -rwxrw-r-- fileName  
chmod a-r fileName ---> --w--w---- fileName
chmod o+x fileName ---> -rw-rw-r-x fileName

###预设权限
- 设定 umask 002
umask  
umask -S  
0111  rw-rw-rw-  
0222  r-xr-xr-x  
0555  -wx-wx-wx  
0000  rwxrwxrwx  
