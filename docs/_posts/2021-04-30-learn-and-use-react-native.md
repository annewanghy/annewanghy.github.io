---
layout: post
title: "Something small tips I learn when I using react-native"
date: 2021-04-30 19:47:55 +0800
categories: Learning
---

### Hide Warning

there are lots of yellow warning during developing, need to click one by one, use the below code to hide all of them

```js
console.disableYellowBox = true;
```

### Flex

Unlike website, react-native `flex direction` is `column` by default, need to manually set `flexDirection:row` to make it in one line

### why words doesn't display

In react-native, if you use `<View>some texts</View>`, it won't display, because text must be wrapped by `<Text>`, so must write in `<View><Text></Text></View>`

### Use Redux to manage data together

when you can change it, all the children will update

```ts
import { connect } from "react-redux";
import { Dispacth } from "redux"; // use its type

const Component = connect(
  // first is data in state
  // second is dispatch
  // callback => pass need data as props to C
  ({ data: { order } }, { dispatch }) => ({ order, dispatch })
)(C);

const C = ({
  order,
  dispatch,
}: {
  order: OrderType;
  dispatch: Dispatch<any>;
}) => {
  return (
    <View>
      <Text>{order}</Text>
      <Button onPress={() => dispatch(xxx)}>click me to do something</Button>
    </View>
  );
};
```

### Use StyleSheet to create common style

if you want to write some style, you can write inline style, but cannot be used commonly
use `StyleSheet` can help write re-used styles

```tsx
import { StyleSheet } from "react-native"

const styles = StyleSheet.create({
   container: {
      display: "flex"
   }
})

<View style={[styles.container, otherStyle]}>
xxxx
</View>
```

### List

1. SectionList
2. FlatList
   both of them are been wrapped by `ScrollView`, so they are scrollable
   
   
### 如何加阴影
props:
```
shadowColor: Sets the drop shadow color.
shadowOffset: Sets the drop shadow offset.
shadowOpacity: Sets the drop shadow opacity (multiplied by the color’s alpha component).
shadowRadius: Sets the drop shadow blur radius.
```

栗子
```
shadowColor: "#000",
shadowOpacity: 0.4,
shadowRadius: 10,
shadowOffset: { width: 0, height: 10 },
```

### 如何在点击外部view层关闭dropdown
方法就是在View的onTouchEnd的时候关闭dropdown
```jsx
const [showSorter, setShowSorter] = useState(false);

<View
onTouchEnd={() => {
   if (showSorter) {
     setShowSorter(false);
   }
}}
>
{showSorter && <SorterContent />}
</View>
```

### 怎么查看ipad的log
首先需要准备一根连接线,把电脑和iPad连接起来,然后可以打开xcode -> window -> Devices and simulator, 选中iPad, 打开console就可以log啦 
![image](https://user-images.githubusercontent.com/18532655/129028726-ddf0144d-aeae-4af3-9cc0-9b57d5e33c6b.png)
![image](https://user-images.githubusercontent.com/18532655/129028347-19942ad7-725c-4249-86f9-ddc00055ef84.png)
![image](https://user-images.githubusercontent.com/18532655/129028452-93a15d54-9cfb-43b3-81ec-db510e85e24e.png)
