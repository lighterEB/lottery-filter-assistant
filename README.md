lottery-filter-assistant/
├── cmd/                      # 应用程序入口
│   └── main.go
│
├── config/                   # 配置文件
│   ├── config.go            # 配置加载器
│   ├── window.toml          # 窗口配置
│   └── menu.toml            # 菜单配置
│
├── internal/                 # 内部包
│   ├── app/                 # 应用程序核心
│   │   └── application.go   # 应用程序实例
│   │
│   ├── domain/              # 领域模型
│   │   ├── entity/         # 实体定义
│   │   │   ├── lottery.go
│   │   │   └── record.go
│   │   └── repository/     # 仓储接口
│   │       └── lottery_repository.go
│   │
│   ├── infrastructure/      # 基础设施
│   │   ├── persistence/    # 数据持久化
│   │   │   └── sqlite/
│   │   └── api/           # 外部API调用
│   │       └── lottery_api.go
│   │
│   ├── initialize/         # 初始化相关
│   │   ├── app.go
│   │   ├── menu.go
│   │   └── components.go
│   │
│   ├── service/           # 业务服务
│   │   └── lottery_service.go
│   │
│   └── ui/               # UI相关
│       ├── components/   # UI组件
│       │   ├── table.go
│       │   ├── chart.go
│       │   └── form.go
│       ├── pages/        # 页面
│       │   ├── analysis.go
│       │   └── history.go
│       └── theme/        # 主题
│           └── custom_theme.go
│
├── pkg/                  # 可导出的包
│   ├── utils/           # 通用工具
│   │   ├── file.go
│   │   └── time.go
│   └── calculator/      # 计算相关
│       └── statistics.go
│
├── assets/              # 静态资源
│   ├── images/
│   └── icons/
│
├── docs/                # 文档
│   └── api.md
│
├── test/               # 测试
│   └── mock/
│
├── go.mod
└── go.sum