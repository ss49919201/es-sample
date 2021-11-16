## 特定の条件に一致した場合に重みをつける

```
GET sample/_search
{
  "query": {
    "bool": {
      "should": [
        {
          "terms": {
            "blood.keyword": [
              "B"
            ], 
            "boost": 2.0
          }
        },
        {
          "terms": {
            "blood.keyword": [
              "O",
              "A"
            ]
          }
        }
      ]
    }
  },
  "sort" : [
    "_score",
    { "number": { "order" : "desc" }}
  ],
  "size": 20
}
```