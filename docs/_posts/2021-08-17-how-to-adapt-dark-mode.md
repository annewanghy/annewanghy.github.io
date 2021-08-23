---
layout: post
title: "How to Adapt Dark Mode"
date: 2021-08-17 19:47:55 +0800
categories: Learning

---

### Dark mode is a feature in IOS 13

### IOS13有了暗黑模式Dark Mode

1. 禁止Dark Mode
   因为我们的app是很早开发的, 所以如果一下子匹配dark mode需要修改很多, 因此暂时先禁止掉
   具体操作
   向 `info.plist` 文件里面写 https://stackoverflow.com/questions/58395926/how-to-force-disable-ios-dark-mode-in-react-native

   ```swift
   <key>UIUserInterfaceStyle</key>
   <string>Light</string>
   ```
2. 怎么在模拟器上开启dark mode
![image](https://user-images.githubusercontent.com/18532655/130399204-e8345182-18f5-4c42-8fbc-b3910bf019e5.png)

