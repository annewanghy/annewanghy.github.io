### Use curl to send a post request

```js
curl -X POST -H "Content-Type: application/json" \
 -d '{"product_code":"AF5087L","bizs":[1,2,3], "size": 12}' \
https://api-prod.lacoste.jp/api/theplant.ec.api.recommendation.RecommendedProductsService/GetSimilarProducts
```

[browser avoid cors](https://alfilatov.com/posts/run-chrome-without-cors/)

```js
open -n -a /Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --args --user-data-dir="/tmp/chrome_dev_test" --disable-web-security
```
