---
layout: post
title: "How to catch and report react native app error"
date: 2022-01-21 17:47:55 +0800
categories: Learning
---

### 错误的分类

1. react-native 开发的 app 分为两个部分

   1. javascript
   2. native

2. 所以错误也分为两个场景

   1. javascript 的错误
   2. native 的错误

3. 一般的 javascript 错误是可以被 catch 到并且显示成错误提示, 但是有些不能 catch 的错误, 尤其是打包之后的 app 内部产生了 js 错误就不能 debug 了.所以要把他 report 出来, 方便知道哪个位置错了

4. native 的错误更加难以 debug, 而且 iOS 的 app, 一旦出现了 native 的代码错误, 就会导致 app 再也打不开了,所以 native 的错误更加要重视

### 怎么捕捉这些错误

因为有开源社区的支持, 我发现了一个很好用的包[react-native-exception-handler](https://github.com/a7ul/react-native-exception-handler)

他提供了两个 [expection handler](https://github.com/a7ul/react-native-exception-handler#usage), 也就是错误处理方式

```js
import {getJSExceptionHandler, setNativeExceptionHandler} from 'react-native-exception-handler';
```

其中 js 可以传递错误进来, 看到详细的错误, native 拿到的是 error message, 是一个字符串.

这样我们可以对两种错误都进行处理了

### 发送到第三方的错误平台

然后这个错误处理方法的参数是一个函数, 因为可以把它抽象出来, 对其进行处理, 如果是比较严重的错误, 可以把它 report 到第三方, 并且弹窗提示用户.

如果是不严重的, 就可以把它打下日志, 开发的过程中遇到了就解决下.

```js
const reporter = (error) => {
  // Logic for reporting to devs
  // Example : Log issues to github issues using github apis.
  console.log(error); // sample
};

const errorHandler = (e, isFatal) => {
  if (isFatal) {
    // report出来
    reporter();

    // 弹窗提示
    Alert.alert(
        'Unexpected error occurred',
        `Error: ${(isFatal) ? 'Fatal:' : ''} ${e.name} ${e.message}`
    );
  } else {
    console.log(e); 
  }
};

setJSExceptionHandler(errorHandler);
```

错误的 report 有很多种选择, 可以根据自己的选择配置项目.
