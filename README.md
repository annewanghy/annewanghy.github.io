# annewanghy.github.io

个人博客 **huiying's blog**，基于 [Jekyll](https://jekyllrb.com/) 构建，通过 GitHub Pages 部署，线上地址：<https://annewanghy.github.io/>

## 目录结构

```
.
├── docs/            # 站点源码（GitHub Pages 的发布目录）
│   ├── _config.yml  # Jekyll 配置
│   ├── _posts/      # 博客文章（按 YYYY-MM-DD-标题.md 命名）
│   ├── _layouts/    # 页面布局模板
│   ├── _includes/   # 可复用的页面片段（head、header、footer 等）
│   ├── css/         # 样式（cayman / normalize / 自定义）
│   ├── index.md     # 首页
│   ├── about.md     # 关于页
│   └── 404.html     # 404 页面
└── note/            # 个人学习笔记（不参与站点构建）
```

## 部署说明

GitHub Pages 的发布源**必须**设置为 `master` 分支的 `/docs` 目录（站点源码所在位置），否则根路径会返回 404。

设置路径：仓库 → **Settings → Pages → Build and deployment**，Source 选 `Deploy from a branch`，Branch 选 `master`，文件夹选 `/docs`。

推送到 `master` 后 GitHub 会自动重新构建并发布，约 1–2 分钟后生效。

## 本地预览

需要 Ruby 环境（Jekyll 4.2，依赖见 `docs/Gemfile`）。

```bash
cd docs
bundle install            # 首次安装依赖
bundle exec jekyll serve  # 启动本地服务，默认 http://localhost:4000
```

## 新增文章

在 `docs/_posts/` 下新建文件，文件名遵循 `YYYY-MM-DD-标题.md` 格式，开头加上 Front Matter：

```markdown
---
layout: post
title: 文章标题
date: 2026-06-17
---

正文内容……
```
