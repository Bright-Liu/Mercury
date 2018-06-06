# mercury

# 说明

用来模拟批量创建中车风电数据，主要有如下几个功能点：
* 可以设置风机总台数
* 可以设置每条记录中指标数量
* 每个文件里面有10条记录，每条记录中前8位是用来存放时间戳，其余指标均为4位float数据
* 必须制定数据存放根路径，脚本会自动在根路径下面创建风场和风机文件夹，每个风场下面默认255个风机

# 编译、运行、帮助等

```bash
 bright@mbp  ~/Downloads  cd ~/03.GitLab/SkyData/mercury
 bright@mbp  ~/03.GitLab/SkyData/mercury   master ●  ll
total 16
-rw-r--r--  1 bright  staff   429B May 22 15:35 README.md
-rw-r--r--  1 bright  staff   3.4K May 22 15:17 main.go
 bright@mbp  ~/03.GitLab/SkyData/mercury   master ●  go build main.go
 bright@mbp  ~/03.GitLab/SkyData/mercury   master ●  ll
total 4512
-rw-r--r--  1 bright  staff   429B May 22 15:35 README.md
-rwxr-xr-x  1 bright  staff   2.2M May 22 15:35 main
-rw-r--r--  1 bright  staff   3.4K May 22 15:17 main.go
 bright@mbp  ~/03.GitLab/SkyData/mercury   master ●  ./main -help
Usage of ./main:
  -machine int
    	风机总台数 (default 10000)
  -metric int
    	每条记录中指标总数 (default 40)
  -root string
    	根路径,用来保存模拟风场和风机数据
 ✘ bright@mbp  ~/03.GitLab/SkyData/mercury   master ●  ./main
root folder path can't be blank!
 bright@mbp  ~/03.GitLab/SkyData/mercury   master ●  ./main -root "/Users/bright/Downloads/test"
[2018-05-22 15:36:27] create file /Users/bright/Downloads/test/dat/20180522153627.dat , cost 9 ms!
Copy dat file to 10000 folders, cost 2299 ms!
[2018-05-22 15:36:37] create file /Users/bright/Downloads/test/dat/20180522153637.dat , cost 0 ms!
Copy dat file to 10000 folders, cost 2472 ms!
```
