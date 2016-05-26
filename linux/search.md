#搜索

##指令
- which，寻找PATH内的路径中的指令
which -a passwd  
which -a mysql  

##档案
###whereis，通过特定的目录中寻找档案名
-l    :可以列出 whereis 会去查寻的主要目录  
-b    :只找 binary 格式的档案  
-m    :只找在说明文档 manual 下的档案  
-s    :只找 source 来源档案  
-u    :搜寻不在上述三個項目中的其他特殊档案  

###locate,搜索资料库在/var/lib/mlocate备注的档案，资料库一天搜索一次数据，手动执行资料库的搜索：updated,不过资料库搜索整理的时间还是挺长的。
-i  ：忽略大小写的差异；  
-c  ：只输出档案数量  
-l  ：输出几行，例如输出五行 -l 5  
-S  ：输出 locate查到的详细信息  
-r  ：正则查找  

###find,遍历硬盘搜索，全，但是慢
find会搜索次目录

- 时间
-mtime n  :n为数字，n天前的（一天内）改动过的档案  
-mtime +n :n天前改动过的档案  
-mtime -n :n天内改动过的档案  
-newer file :file为已存在的档案，列出比file新的档案

- 使用人
-uid n:用户id相关  
-gid n:群组id相关  
-user name  
-group name  
-nouser  
-nogroup  

- 权限及名称
-name fileName  
-size [+-] SIZE  
-type TYPE :TYPE类型有f,b,c,d,l,s,d等  
-perm mode :若文件权限-rwxrwxrwx,mode=777  
-perm -mode :包括mode权限的文件，例如mode=111,若文件权限777,则文件被查出来，因为777囊括了111  
-perm /mode :mode=777,若文档700,则文档被查出，因为文档有权限7  
