# back

CHOOSE 的后端服务。使用 [go migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) 进行模式迁移，所以在运行之前需要运行：

```sh
# 安装 migrate 工具。
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# 模式迁移。
bash db/migrate.sh
```

## API
API 遵守 [谷歌的 API 风格](https://cloud.google.com/apis/design/standard_methods?hl=zh-cn)，因此我们提供以下接口（资源类型表示都为 JSON）：

- CreateAnswer，`POST /api/v1/answers`，请求模式见 [answer](#answer)，响应模式同见 [answer](#answer)。
- ListAnswer，`GET /api/v1/answers?user_name={user_name}`，响应模式为模式 [answer](#answer) 的列表。
- GetQuestion，`GET /api/v1/questions/{id}`，响应模式见 [question](#question)。

### answer
Example：

```json
{
    "id": "99f9",
    "user_name": "foo",
    "question_id": "ff9f",
    "answer": "I think..."
}
```

Note：在 CreateAnswer 请求中，字段 `id` 会被忽略，故 **应该** 避免携带 `id` 字段。

### question
Example：

```json
{
    "id": "ff9f",
    "context": "To be or not to be...",
    "question": "How about..."
}
```
