#Fail2ban使用，自动激活iptable，防止ssh破解

##启用防火墙firewalld
- 状态systemctl status firewalld
- 重启systemctl restart firewalld
- 开放一些服务端口firewall-cmd --add-port=80/tcp --permanent
- 重新加载，其修改生效firewall-cmd --reload
- 查看防火墙基本情况firewall-cmd --list-all
- 注意，若使用docker管理一些服务，先启动firewalld再重启docker服务，注意docker服务重启后，里面运行的容器也都需要重新启动。

##安装Fail2ban，对sshd进行限制，防破解
- CentOS内置源并未包含fail2ban，需要先安装epel源yum -y install epel-release
- 安装fial2ban，yum install fail2ban
- vi /etc/fail2ban 配置10分钟内，若登陆5次失败，则将ip列为黑名单，屏蔽10分钟
- 启动systemctl start fail2ban
- 测试一下
- 查看被屏蔽状态fail2ban-client status sshd
- 删除被屏蔽的ip,使其能够正常使用fail2ban-client set sshd unbanip 192.168.1.202