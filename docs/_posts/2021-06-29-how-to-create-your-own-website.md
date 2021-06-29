---
layout: post
title: "How to create your own website by using github"
categories: Learn Tips
---

1. github 创建一个 `username.github.io` 必须是这样的格式
2. 选择主题，创建一个 `index.md`，博客就搭建完成了 ✅，访问https://username.github.io可以看到对应的内容
3. 可以新建一个`about.md`, 这样可以新建一个 about 页面
4. 博客要写博文，但是一个个这样手动链接未免太麻烦，于是我们需要工具帮忙自动链接，并生成导航栏

   1. clone 项目到本地
   2. 安装 ruby 和 gem
   3. 给 gem 安装 bundler
      `sudo gem install jekyll bundler`
   4. `jekyll new . --force`，不用担心，只是添加，不会把之前创建的文件给丢掉
   5. 配置一下`_config.yml` 自己的 name
   6. `jekyll server --trace`，启动一个本地服务器，访问 localhost:4000，就可以看到 jekyll 为我们创建的首页了
   7. 在`\_posts` 目录下按照`2021-06-29-how-to-start-your-own-blog.md`的格式命名文件
   8. 文件的首部要写上使用的模版
      ```md
      ---
      layout: post
      title: "How to create your own website by using github"
      categories: Learn Tips
      ---
      ```

5. 最后把代码 merge 回 master 就可以看到网站成功部署上去了
![image](https://user-images.githubusercontent.com/18532655/123804342-b7745480-d91f-11eb-96cb-fc34cdd0c9ff.png)
