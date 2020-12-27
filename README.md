# マッピングの作成

- PUT or POST
- https://{エンドポイント}/{index名}
- mapping.json

# マッピングの確認

- GET
- https://{エンドポイント}/{index名}/{type名}/_mapping

# サンプルデータ

- PUT or POST
- https://{エンドポイント}/{index名}/{type名}
- sample_doc.json

# 自動マッピングの禁止

- PUT or POST
- https://{エンドポイント}/{index名}/_mapping
- dynamic_strict.json

# 検索

- GET or POST
- https://{エンドポイント}/{検索対象index名}/_search
- https://{エンドポイント}/{検索対象index名,検索対象index名}/_search
- https://{エンドポイント}/{検索対象index名*}/_search
- https://{エンドポイント}/_search
- query.json