#基础指令
大小写敏感

##开机关机
-  sync	
将内存中数据保存到硬盘，这是由于计算机出于性能考虑，有些常用数据读取到内存中，在你进行操作后，关机。此时，有时候可能数据还在内存里，没有保存到硬盘上，因此，手动的sync一下，有利于数据保护(其实shutdown,reboot,halt等指令已经自动sync了)。sync指令只会同步对应用户的数据，root用户能够保存所有数据。  
-  shutdown  
-k	不关机，发出关机警告  
-r	系统服务都停掉后，重启  
-h	系统服务都停掉后，关机  
-c	取消进行的shutdown相关指令  

shutdown -h now		现在关机  
shutdown -h 20:25	今天20：25关机,若21：00下达指令，明天的20：25关机  
shutdown -h +10		10分钟关机  
shutdowm -k now "This system will shutdown"  

-  reboot,halt,poweroff

##编码
locale  
echo LANG  
LANG=en_US.UTF-8  
export LC_ALL=en_US.UTF-8  

##帮助，查询指令文档
date -- help  
date -h  
man date  
info date  
file /usr/bin/passwd查看档案使用到的库，可以判断档案大体类型

##时间
date  
date +%Y%m%d  
date +%Y/%m/%d  
date +%H:%M  
cal  
cal 2016  
cal 10 2016  

##空间
- df针对整个档案
df -h

- du针对目录
-a 所有
-h G/M显示容量
-s 总量
-S 不包括子目录，总量

##检测
who  
w  
netstat  
ps -aux|grep tomcat  

##增删改查
mkdir dirName 
rmdir dirName   
cd dirName  
touch fileName.c  
rm -rf fileName.c  
mv ./oldFileName.c ./newFileName.c  
cp ./oldFileName.c ./newFileName.c  
ll  
ls  
ls -al   
pwd  
ln -s /from /to	软链接

##读取
tail -f fileName
vi fileName  
more fileName  
less fileName 
head -n 100 fileName  
tail -n 100 fileName   
od  fileName  
echo 可以在屏幕上显示任何东西
