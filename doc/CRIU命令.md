# 热迁移系统CRIU命令使用手册

1）另一个终端查看测试程序pid：在docker容器内查看pid ps vax

```
  PID TTY      STAT   TIME  MAJFL   TRS   DRS   RSS %MEM COMMAND
     1 ?        Ss     0:00      4    93  2506   648  0.0 /bin/sh -c ./criu-server -Port 8080
     8 ?        Sl     0:00      6  5748 719503 33284  0.8 ./criu-server -Port 8080
    13 pts/0    Ss     0:00      4   885  3350  2232  0.0 /bin/bash
    33 pts/0    R+     0:00      0    86  5721  1292  0.0 ps vax
```

2）dump：设置检查点并中断程序

```
mkdir img
criu dump -t 8 -D img/ -o dump.log -v4 --tcp-established -j
```

4）恢复程序：

```
criu restore -D img/ -o restore.log -d --tcp-close -j -v4
```

5）检查进程是否使用原来的进程号继续运行：

```
ps vax
```

命令说明：

dump命令：当设置进程1597的checkpoint后，再查找该进程，发现进程不存在，即进程已经退出。查看快照文件目录，生成很多img文件，这些文件主要用于恢复应用。
-D（等价于–images-dir）：选项指定应用的快照文件保存目录
-j（等价于–shell-job）：表示该应用是一个通过shell启动的作业
-t：指定需要checkpoint的应用pid
当对应用设置checkpoint后，应用会自动退出，如果希望应用继续执行，需指定-R或–leave-running选项，restore命令：恢复后的程序从设置checkpoint的时间点继续运行，发现进程使用原来的进程号继续运行。

```
-D（等价于–images-dir）：选项指定应用的快照文件保存目录
-j（等价于–shell-job）：表示该应用是一个通过shell启动的作业
-t：指定需要checkpoint的应用pid
-o: 在恢复冻结的时候指定日志文件的名称
--tcp-established  检查点/恢复已建立的TCP连接
--skip-in-flight   跳过（忽略）正在运行的TCP连接
--tcp-close        不要转储或阻止已建立的tcp的状态
```



