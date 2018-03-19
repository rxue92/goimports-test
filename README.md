# 现象
在windows上使用goland时，切换分支时有的文件会自动变蓝（表示有改动），用git status和gif diff --check显示只有分行符不同

# 问题重现
1. 随便准备一个git仓库
2. git core.autocrlf设为true
3. 建立新分支（比如develop）
4. 在新分支上做出修改并提交
5. 切换到master分支。如果用的是goland的terminal，切换后hide一下terminal，让goland加载一下文件；如果用的是git bash，需要打开一下goland。
6. 在develop分支上修改过的文件在master分支上会被标蓝。此时用git status会发现修改过得文件有changes not staged for commit。

由于切换分支时要求文件要么放入缓存区要么是未修改的，所以此时是无法再切换分支的。

# 解决方法
切换分支后，运行git add .，就可以切换分支了。

运行git add后命令行会显示
```
warning: LF will be replaced by CRLF in functions/test_funcs.go.
The file will have its original line endings in your working directory.
```
由于只是分行符的不同，此时并不能commit（nothing to commit, working tree clean）。

# 分析
如果仅用git bash进行分支切换，或者用goland的terminal进行切换后不隐藏terminal，切换分支后不会出现类似问题。
因此猜测是goland根据分支加载文件时对分行符进行了改动。
