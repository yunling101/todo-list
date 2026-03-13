# Todo List 应用

一个简单的 Todo List 管理应用，采用前后端分离架构。

## 技术栈

| 端 | 技术 |
|----|------|
| 前端 | React + Vite |
| 后端 | Go + Gin |
| 数据库 | SQLite |

## 功能特性

- ✅ 创建 Todo
- ✅ 查看 Todo 列表
- ✅ 编辑 Todo
- ✅ 删除 Todo
- ✅ 标记完成/未完成
- ✅ 筛选（全部/待完成/已完成）

## 项目结构

```
todo-list/
├── backend/              # Go 后端
│   ├── main.go          # 主程序
│   ├── go.mod           # Go 依赖
│   └── todos.db         # SQLite 数据库
│
├── frontend/             # React 前端
│   ├── src/
│   │   ├── App.jsx      # 主应用
│   │   ├── App.css      # 样式
│   │   ├── api/         # API 调用
│   │   └── components/ # 组件
│   ├── vite.config.js   # Vite 配置
│   └── package.json
│
├── PROJECT.md           # 项目文档
├── TASKS.md             # 任务看板
├── TODO.md              # 技术方案
└── README.md            # 项目说明
```

## API 接口

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | /api/v1/todos | 获取所有 Todo |
| POST | /api/v1/todos | 创建新 Todo |
| GET | /api/v1/todos/:id | 获取单个 Todo |
| PUT | /api/v1/todos/:id | 更新 Todo |
| DELETE | /api/v1/todos/:id | 删除 Todo |
| PATCH | /api/v1/todos/:id/toggle | 切换完成状态 |

## 快速启动

### 后端

```bash
cd backend
go run main.go
```

后端服务将在 http://localhost:8080 运行

### 前端

```bash
cd frontend
npm install
npm run dev
```

前端应用将在 http://localhost:5173 运行

## 开发团队

| 角色 | 描述 |
|------|------|
| XiaMiClaw | 技术主管（任务协调、代码审核） |
| backend | 后端开发 |
| frontend | 前端开发 |
| testing | 测试工程师 |

## 许可证

MIT
