# task_app_server
- タスク管理
- 認証機能

## API

GET /api/tasks 全てのタスクの取得をする <br> 

POST /api/tasks タスクの追加をする <br> 
{ "name" : "掃除"　}　<br>

PUT /api/tasks タスクを完了させる <br> 
DELETE /api/tasks タスクを削除する <br> 

POST /api/register　ユーザー登録 <br> 
{
    "name": "abcd",
    "password": "agdsa43"
} <br>

POST /api/login JWTの取得

![task-management SVG](./task-management.svg "Sample SVG Image")

## 使用技術
- Go 1.22
- mySQL
- Echo
- Gorm
- Docker