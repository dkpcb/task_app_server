# task_app_server
- タスク管理
- 認証機能

GET /api/tasks タスクの取得 <br> 

POST /api/tasks タスクの追加 <br> 
{ "name" : "掃除"　}　<br>

PUT /api/tasks タスクの完了 <br> 
DELETE /api/tasks タスクの削除 <br> 

POST /api/register　ユーザー登録 <br> 
{
    "name": "abcd",
    "password": "agdsa43"
} <br>

POST /api/login JWTの取得

- Go 1.22
- mySQL
- Echo
- Gorm
- Docker