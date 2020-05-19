## golang 环境安装
安装包下载地址为：https://golang.org/dl/。 

默认情况下 .msi 文件会安装在 c:\Go 目录下。你可以将 c:\Go\bin 目录添加到 Path 环境变量中。添加后你需要重启命令窗口才能生效。

设置环境变量<br>
path中添加 C:\Go\bin<br>
GOPATH设置 D:\godev 并创建src bin pkg三个目录<br>
重新打开窗口 使用go env 命令 查看gopath 和goroot是否正确

## goland的破解
goland打开 创建任意一个go文件,将百度网盘链接: https://pan.baidu.com/s/1e0s2lbpBn9v9CQiRyzZNZQ 提取码: vduk 的jetbrains-agent.jar
文件直接拖入 IDEA 界面中，当然如果是其他IDE的话，也同理。 然后根据提示重启

重新打开 点击 Help 下的 Register 按钮查看软件激活信息

## git自动补全
打开git bash<br>
```shell script
$ vim /etc/profile.d/git-prompt.sh

```
将其修改为如下内容
```shell script
if test -f /etc/profile.d/git-sdk.sh
then
	TITLEPREFIX=SDK-${MSYSTEM#MINGW}
else
	TITLEPREFIX=$MSYSTEM
fi

if test -f ~/.config/git/git-prompt.sh
then
	. ~/.config/git/git-prompt.sh
else
	PS1='\[\033]0;Bash\007\]'      # 窗口标题
	PS1="$PS1"'\n'                 # 换行
	PS1="$PS1"'\[\033[32;1m\]'     # 高亮绿色
	PS1="$PS1"'➜  '               # unicode 字符，右箭头
	PS1="$PS1"'\[\033[33;1m\]'     # 高亮黄色
	PS1="$PS1"'\W'                 # 当前目录
	if test -z "$WINELOADERNOEXEC"
	then
		GIT_EXEC_PATH="$(git --exec-path 2>/dev/null)"
		COMPLETION_PATH="${GIT_EXEC_PATH%/libexec/git-core}"
		COMPLETION_PATH="${COMPLETION_PATH%/lib/git-core}"
		COMPLETION_PATH="$COMPLETION_PATH/share/git/completion"
		if test -f "$COMPLETION_PATH/git-prompt.sh"
		then
			. "$COMPLETION_PATH/git-completion.bash"
			. "$COMPLETION_PATH/git-prompt.sh"
			PS1="$PS1"'\[\033[31m\]'   # 红色
			PS1="$PS1"'`__git_ps1`'    # git 插件
		fi
	fi
	PS1="$PS1"'\[\033[36m\] '      # 青色
fi

MSYS2_PS1="$PS1"

```
