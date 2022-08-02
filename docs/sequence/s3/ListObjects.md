```mermaid
sequenceDiagram
  autonumber
  actor user
  participant s3
  user ->> user: Region選択
  user ->> user: 操作対象AWS Resource選択
  user ->> s3: ListBucketsInput
  s3 -->> user: return buckets
  user ->> s3: ListObjectsV2Input
  s3 -->> user: return objects
```
