##settings.xml

配置文件  

- <localRepository>${user.home}/.m2/repository</localRepository>本地仓库位置
- <interactiveMode>true</interactiveMode>true表示mave可以使用用户输入，默认true
- <usePluginRegistry>false</usePluginRegistry>true表示maven使用${user.home}/.m2/plugin-registry.xml管理插件版本，默认为false
- <offline>false</offline>true表示构建系统在离线模式下执行，默认为false

- 每个pluginGroup元素都包含一个groupId，当你在命令行中没有提供插件的groupid时，将会使用该列表。
<pluginGroups>
    <pluginGroup>org.mortbay.jetty</pluginGroup>
</pluginGroups>

- POM中的repositories和distributionManagement元素为下载和部署定义的仓库。一些设置如服务器的用户名和密码不应该和pom.xml一起分发。这种类型的信息应该存在于构建服务器上的settings.xml文件中。
<servers>
    <server>
        服务器的id，和repository/mirror中配置的id项匹配；
        <id>server001</id>

        服务器的认证信息；
        <username>my_login</username>
        <password>my_password</password>

        指定一个路径到一个私有key（默认为${user.home}/.ssh/id_dsa）和一个passphrase；
        <privateKey>${user.home}/.ssh/id_dsa</privateKey>
        <passphrase>some_passphrase</passphrase>

        设置文件和文件夹访问权限，对应unix文件权限值，如：664，后者775.
        <filePermissions>664</filePermissions>
        <directoryPermissions>775</directoryPermissions>

        <configuration></configuration>
    </server>
</servers>

- mirrors
<mirrors>
    <mirror>
        镜像的唯一标识和用户友好的名称；
        <id>planetmirror.com</id>
        <name>PlanetMirror Australia</name>

        镜像的url，用于代替原始仓库的url；
        <url>http://downloads.planetmirror.com/pub/maven2</url>
        <mirrorOf>central</mirrorOf>
    </mirror>
</mirrors>
