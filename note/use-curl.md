### Use curl to send a post request

```js
curl -X POST -H "Content-Type: application/json" \
 -d '{"product_code":"AF5087L","bizs":[1,2,3], "size": 12}' \
https://api-prod.lacoste.jp/api/theplant.ec.api.recommendation.RecommendedProductsService/GetSimilarProducts
```
