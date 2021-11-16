# セットアップ

```sh
./init.sh
```

# 掃除

```sh
./clean.sh
```
# クエリ

※6系以降 **Content-Type** 必須

## ヘルスチェック

```sh
curl 'http://localhost:9200/_cluster/health?pretty=true'
```

## マッピングの作成

- PUT or POST
- https://{エンドポイント}/{index名}
- mapping.json

## マッピングの確認

- GET
- https://{エンドポイント}/{index名}/{type名}/_mapping

## サンプルデータ

- PUT or POST
- https://{エンドポイント}/{index名}/{type名}
- sample_doc.json

```sh
curl -H "Content-Type: application/json" -XPOST localhost:9200/index/type --data-binary @sample_doc_2.json
```

## 自動マッピングの禁止

- PUT or POST
- https://{エンドポイント}/{index名}/_mapping
- dynamic_strict.json

## 検索

- GET or POST
- https://{エンドポイント}/{検索対象index名}/_search
- https://{エンドポイント}/{検索対象index名,検索対象index名}/_search
- https://{エンドポイント}/{検索対象index名*}/_search
- https://{エンドポイント}/_search
- query.json

```sh
curl -H "Content-Type: application/json" -XGET localhost:9200/index/_search -d @match_all.json
```

```sh
curl -H "Content-Type: application/json" -XGET localhost:9200/index/_search -d @random_score.json
```