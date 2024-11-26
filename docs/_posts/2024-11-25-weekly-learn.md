---
layout: post
title: "learn"
date: 2024-11-25 19:47:55 +0800
categories: Learning
---

1. 设置 Prettier 为默认格式化工具
	1.	打开 VS Code 设置：
  	•	点击右下角齿轮图标 → 选择 Settings。
  	•	或按快捷键 Cmd+。
	2.	搜索 default formatter，然后设置 Prettier 为默认格式化工具：
	  •	找到 Editor: Default Formatter，选择 esbenp.prettier-vscode（Prettier 插件）。
	3.	启用自动格式化：
	  •	搜索 format on save，勾选 Editor: Format On Save。
2. Safari: 时间只能是`2018-09-09`,不能是 `2018-9-9`, 因为不遵守 ISO 8601 规范
3. 对于测试的思考
   1. 分为黑盒测试(完全不懂代码, 只知道输入和输出的期望)和白盒测试(懂内部代码逻辑, 可以针对最复杂的逻辑进行测试)
   2. 关键是写好测试用例
   3. 知道功能, 输入, 和输出, 一一对应
   4. 实用工具, vitest(简单)和jest(复杂, 而且比较慢)
   5. 使用: describe描述一个测试, it/test一个具体的测试用例, expect期望, toBe简单对象匹配, 比如toBe(true|false|0|1), toEqual对象匹配, toMatch快照匹配(snapshots, 一个react components的render渲染出来的结果), toBeCall函数被调用, toBeCallWith函数被什么参数调用,not.toBe没有被, toBeInDocument在dom里面
   6. 每一个测试用例需要独立, 互不依赖也互不影响, 如果的确有, 那么在beforeEach或者afterEach, beforeAll/afterAll去清空依赖关系
   7. hook怎么测试, 借助工具`@testing-library/react`, hook一般不用, 除非里面需要用到react一些特性, 比如useEffect, useState这种, 简单的函数应该用最简单的function来做, hook的结果是
   ```js
    const {
      result: {
        current: { validateBirthday },
      },
    } = renderHook(() => useXxx({}))
   ```
   9. 测试的目的: 是为了让后面重构的时候, 可以更加安全避免改坏
   10. mock函数, 使用`jest.fn()`或者`vi.fn()`
