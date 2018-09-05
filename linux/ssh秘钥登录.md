#ssh修改端口和使用密钥(密钥应加密)
##密钥生成
- 链接250服务器
- 切换root用户
- ssh-keygen
- 填写密钥位置和名字/data/ssh/qyj
- 填写密钥加密码qyjmed2015,.
- 完成密钥生成，/data/ssh里面有俩文件，qyj和qyj.pub，及私钥和公钥

##公钥传输
- 每个服务进行公钥传输
- ssh-copy-id -i /data/ssh/qyj.pub root@192.168.1.29
- 填写服务器root登陆密码xxxxxxxx
- 完成密钥传输，若出现权限问题，请打开目标服务器sshd密码登陆的权限，否则无法实现密钥注入

##登陆方式
- 完成密钥注入后
- 远程登陆目标服务器192.168.1.29
- vi /etc/ssh/sshd.config
- 调转到最后一行，将密码登陆方式关闭 PasswordAuthentication no
- 重启sshd服务，使修改生效systemctl restart sshd
- 最后，服务无法通过密码登陆，只能通过密钥登陆了

##多服务器
之前的三步，是对单个服务器进行的操作，在多服务器的情况下，每个服务器单独传输，显然不太合适，也不方便，使用脚本，统一完成公钥传输过程  
- vi ip.conf
- 在文件中一次写入ip和对应的密码，保存
- yum install expect
- 写脚本vi expect_sshkey.exp，导入密钥
- 写脚本vi run.sh，从文件中读取ip和密码，再加上密钥，循环批量执行expect_sshkey.exp
- 运行脚本sh run.sh,完成批量公钥的导入

##脚本
###expect_sshkey.exp
```
#!/usr/bin/expect
if { $argc !=3 } {
  send_user"expect_sshkey.exp file host password\n"
  exit
}

#define var
set file [lindex $argv 0]
set host [lindex $argv 1]
set password [lindex $argv 2]
spawn ssh-copy-id -i $file "root@$host"
expect {
   "yes/no"   {send "yes\r";exp_continue}
   "*password"   {send "$password\r"}
}
expect eof

exit -onexit {
  send_user "ok!\n"
}
```
###run.sh
```
#!/bin/bash
function ssh_copy(){
    echo "ip = $ip,password = $pwd"
}
while IFS='' read -r line || [[ -n "$line" ]]; do
    IFS=‘:’ read -r ip pwd <<< "$line"
    ssh_copy ip pwd
    /data/ssh/expect_sshkey.exp /data/ssh/qyj.pub $ip $pwd
done < "/data/ssh/ip.conf"
```