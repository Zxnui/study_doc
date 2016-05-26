#特殊权限

##档案隐藏属性

- chattr,完整的用在ext2/ext3/ext4上，其他格式，可能无法完整支持
opt:  增加+，减少-，设置=  
param:  
A:atime不会改变  
S:同步写入磁盘  
a:只能增加资料，无法修改和删除  
c:自动压缩后，在存储文件  
i:不能被删除/改名/设定快捷方式，也无法写入或新增  
s:档案一旦被删除，永远无法找回  
d:dump时，档案不会被备份  
u:被删除，但是文件还在磁盘中，可以轻易找回  
  
chattr +a /var/log/java_test.log
- lsattr,查看属性

##特殊权限
- set UID(s替换第一组x)
ls -la /usr/bin/passwd  
-rwsr-xr-x 其中s，就是被set UID(此时s是用在用户权限x的位置处),用户test运行passwd时，能够将自己修改的密码保存到/etc/shadow中;但是/etc/shadow权限是---------,只有root能够操作，对passwd进行set UID的作用就是让test在运行passwd的时候，能够底层使用root去操作/etc/shadow  
同时注意，set UID只能对binary program使用，不能对shell script使用。

- set GID(s替换第二组x)
当s被用在群组权限x处时，就是被set GID了  
ls -la /usr/bin/locate  
和SUID不同，SGID可以对档案和目录使用。对binary program使用，用户在执行指令时，能够活动群组的支援;用在目录上，用户在此目录新建档案，新建的档案拥有群组为此目录的群组

- sticky Bit(t替换第三组x)
简称SBIT,只正对目录有效  
举例：  
drwxrwxrwx zxnui student fileName  
目录fileName，用户zxnui和群组student都对其下面的文件拥有删除/更名/搬移的权限，若权限变为drwxrwxrwt,此时用户zxnui对目录的权限变为t，zxnui只能对目录fileName自己建立的档案或目录进行删除/更名/搬移的权限，无法删除student中其他人建立的档案。  

##特殊权限修改
SUID 4  
SGID 2  
SBIT 1  

-rwxrwxrwx fileName  
chmod 4777 fileName ---> -rwsrwxrwx  
chmod 2777 fileName ---> -rwxrwsrwx  
chmod 1777 fileName ---> -rwxrwxrwt  
若fileName权限-rw-rw-rw-,此时chmod 7666 fileName，档案权限就会变为-rwSrwSrwT，大写的S/T出现的原因，s/t权限都是建立的x权限上的，若档案没有x权限，S/T就表示空的意思
-
