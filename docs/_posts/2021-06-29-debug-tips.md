---
layout: post
title: "Tips for debugging 更好的debug"
date: 2021-05-31 19:47:55 +0800
categories: Debug Tips
---

### block js and domain

有时候找不到问题，可以尝试用排除法，这时候 chrome 的 block js and block domian 就起大作用了，可以通过一个个排除来确定最后的原因，右键就可以开启 block 啦
![image](https://user-images.githubusercontent.com/18532655/122753288-9dee5f80-d2c4-11eb-8d8c-feac5823dec9.png)
![image](https://user-images.githubusercontent.com/18532655/122753306-a5ae0400-d2c4-11eb-829f-ae3a198d6943.png)

### Safari 浏览器

chrome 的 source map 不是很好，突然发现 Safari 好用很多，一些 bundle 之后的 js 都变成了数字+字母的 js，而 Safari 可以看到源文件
