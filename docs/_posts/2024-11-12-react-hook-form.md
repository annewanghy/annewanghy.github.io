---
layout: post
title: "learn react-hook-form"
date: 2024-11-12 15:47:55 +0800
categories: Learning
---


1. `type=number|date` `(validateAsNumber|validateAsDate)`
3. touched表示被触碰过, `dirty`表示和之前的`value`不一样`(touchFields|dirtyFields|isDirty)`
4. `mode: onSumit(default)|onTouched|onChange|all`
5. `watch(value|[value])` 实时监听`real-time-listening`
6. `getField|setField`
7. `phoneNumber: [1,2,3]`, `phoneNumber.0`, `phoneNumber.1`
8. `soical{twitter, facebook}`, `social.twitter`, `social.facebook`
9. `required, pattern, maxLength, minLength, max, min, validate: (value) => true|string`
10. `<form onSumbit={handleSubmit(onSumbit(data), onError(error))} noValidate>`
11. control, `react-hook-form dev tool`
12. formState: `isValid|isDirty submit button`
13. register => `name, value, onChange, onBlur`
14. `<DevTool control={control>`
15. `isSubmitting|isSubmitted|isSubmitSuccessful`
16. `submitCount`计算成功的次数
17. `useWatch`,需要包含在FormProvider或者传递control参数, 对于性能有帮助, 集中到component的re-render, 减少外部的re-render
