

##关于 zhoushxGoDemo
  各种学习、实现 golang 的demo

##目录结构
	~/projGolang
		bin
		pkg
		src
			github.com				#	go get github.com 后的源码路径
			zhoushxGoDemo			#本demo所在目录
				calcProj			# 项目1
					app				#	子目录
						calc.go		#	使用相对路径将add.go import进来 
									#   	import "zhoushxGoDemo/calcProj/simplemath"
									#	此时需要将 ~/projGolang/src 添加到 $GOPATH 路径中
								
					simplemath		#	子目录
						add.go		#	内部公共模块代码


##环境变量 
	

	export GOARCH=amd64
	export GOROOT=/usr/local/go
	export GOBIN=$GOROOT/bin/
	export GOOS=linux
	export GOPATH=/home/zhoushx/projGolang
	export GOPATH=$GOPATH:/home/zhoushx/projGolang/src

	export PATH=$PATH:$HOME/.local/bin:$HOME/bin:$GOROOT/bin/:$HOME/projGolang/bin
