
## git status 文件乱码问题
 通过将Git配置变量 core.quotepath 设置为false，就可以解决中文文件名称在这些Git命令输出中的显示问题，
``` shell script
git config --global core.quotepath false
```

## git log 中文乱码问题
先看下 LANG 环境变量是否为统一字符编码：
```shell script
$ echo $LANG;
$ locale 
```
输出结果为空
执行export LANG="zh_CN.UTF-8"命令，问题能否解决？

如果不能，再试下修改 git config
```shell script
git config --global i18n.commitencoding utf-8
git config --global i18n.logoutputencoding utf-8
export LESSCHARSET=utf-8

```

