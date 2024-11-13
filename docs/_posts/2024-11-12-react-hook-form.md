1. `type=number|date` (validateAsNumber|validateAsDate)
2. touched表示被触碰过, dirty表示和之前的value不一样(touchFields|dirtyFields|isDirty)
3. mode: onSumit(default)|onTouched|onChange|all
4. watch(value|[value]) 实时监听real-time-listening
5. getField|setField
6. phoneNumber: [1,2,3], phoneNumber.0, phoneNumber.1
7. soical{twitter, facebook}, social.twitter, social.facebook
8. required, pattern, maxLength, minLength, max, min, validate: (value) => true|string
9. `<form onSumbit={handleSubmit(onSumbit(data), onError(error))} noValidate>`
10. control, react-hook-form dev tool
11. formState: isValid|isDirty submit button
12. register => name, value, onChange, onBlur
13. `<DevTool control={control>`
14. isSubmitting|isSubmitted|isSubmitSuccessful
15. submitCount计算成功的次数
16. useWatch,需要包含在FormProvider或者传递control参数, 对于性能有帮助, 集中到component的re-render, 减少外部的re-render
