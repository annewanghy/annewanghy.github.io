Something small tips I learn when I using react-native

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
  ({data: { order }}, {dispatch}) => ({order, dispatch})
)(C)
  
const C = ( {order, dispatch}:{order: OrderType, dispatch: Dispatch<any>} ) => {
    return <View>
       <Text>{order}</Text>
       <Button onPress={() => dispatch(xxx)}>click me to do something</Button>
     </View>
}
```

### Use StyleSheet to create common style
if you want to write some style, you can write `style={{display: "flex"}}`, it's inline, and cannot be used commonly
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


