### Docker 怎么看文件
当线上环境出了问题, 但是本地看是好的, 找不到原因, 于是就把docker下载下来, 对比下文件来找到问题.

```
docker pull image // 先拉下来代码
docker image ls // 列出所有image, 找到对应image的imageID
docker run -it --rm <imageID> sh // 进入 docker 内, 用shell查看文件内容

最后如何退出呢
command+D退出.
```
