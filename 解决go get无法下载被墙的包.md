#### 方式一：

```bash
# 使用gopm(Go Package Manager)代替go下载,是go上的包管理工具，十分好用
# 1. 下载安装gopm
go get -u github.com/gpmgo/gopm
# 2. 使用gopm安装被墙的包
gopm get github.com/Shopify/sarama
```

---

#### 方式二：

```bash
#  golang 在 github 上建立了一个镜像库，如 https://github.com/golang/net 即是 https://golang.org/x/net 的镜像库.获取 golang.org/x/net 包（其他包类似），其实只需要以下步骤：
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/net.git
```
#### 方式三：

使用国内网站打包好，然后下载打包好的压缩包解压安装到本地
```bash
https://www.golangtc.com/download/package

# 1.在上诉网站里按照提示下载这个包，并解压到本地文件夹
# 2.在goland IDE的terminal里进行go install cloud.google.com/go/storage
# 3.第二步会报错，提示有些包在指定路径没找到，于是按照提示给的链接再在第一步的网站里找到并下载，按照路径放在对应的文件夹内
# 4.所有提示缺少的包都装好以后就可以顺利的install了。
```
