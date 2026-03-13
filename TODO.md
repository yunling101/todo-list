# TODO.md - Todo List 技术方案

## 技术栈

- **前端：** React + Vite
- **后端：** Go + Gin 框架
- **数据库：** SQLite (轻量级，无需额外安装)

## API 设计

### Base URL
```
http://localhost:8080/api/v1
```

### 接口列表

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | /todos | 获取所有 Todo |
| POST | /todos | 创建新 Todo |
| GET | /todos/:id | 获取单个 Todo |
| PUT | /todos/:id | 更新 Todo |
| DELETE | /todos/:id | 删除 Todo |
| PATCH | /todos/:id/toggle | 切换 Todo 完成状态 |

### 数据结构

#### Todo
```json
{
  "id": "string (UUID)",
  "title": "string (必填, 1-100字符)",
  "completed": "boolean",
  "created_at": "string (RFC3339)",
  "updated_at": "string (RFC3339)"
}
```

#### 请求/响应示例

**POST /todos**
```json
// Request
{
  "title": "学习 Go"
}

// Response
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "学习 Go",
  "completed": false,
  "created_at": "2026-03-13T01:30:00Z",
  "updated_at": "2026-03-13T01:30:00Z"
}
```

**PUT /todos/:id**
```json
// Request
{
  "title": "学习 Go - 更新",
  "completed": true
}

// Response
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "title": "学习 Go - 更新",
  "completed": true,
  "created_at": "2026-03-13T01:30:00Z",
  "updated_at": "2026-03-13T01:35:00Z"
}
```

## 目录结构

```
/home/ubuntu/.openclaw/project/
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   ├── config/
│   │   └── config.go
│   ├── models/
│   │   └── todo.go
│   ├── handlers/
│   │   └── todo.go
│   ├── database/
│   │   └── database.go
│   └── tests/
│       └── todo_test.go
│
└── frontend/
    ├── src/
    │   ├── App.jsx
    │   ├── App.css
    │   ├── components/
    │   │   ├── TodoList.jsx
    │   │   ├── TodoItem.jsx
    │   │   └── TodoForm.jsx
    │   └── api/
    │       └── todo.js
    ├── index.html
    ├── package.json
    └── vite.config.js
```

## 开发约定

1. 后端端口：8080
2. 前端端口：5173 (Vite 默认)
3. 前端调用后端需配置代理
4. 统一使用 UTF-8 编码
