# wails3_demo

## 项目介绍

这是一个 基于 Wails3 框架 的桌面应用程序示例，展示了如何使用 Wails3 创建跨平台的桌面应用。

包含悬浮球、系统托盘图标和菜单、上下文菜单、透明无边框窗口，前后端事件通信，全局事件监听 等功能。

目前仅在 Window 下完成了Demo。

## 环境要求

- Go 1.25 及以上版本
- Node.js 24 及以上版本
- npm 11 及以上版本
- Wails3 版本见 go.mod

# 常用 CLI 命令

```bash
# 安装 npm 全局工具
npm install -g ts-node del-cli

# 安装项目依赖并更新到最新版本
npm run re-update

# 仅安装项目依赖
npm install

# 仅安装前端依赖
cd frontend && npm install && cd ..

# 仅安装 go 依赖
go install github.com/wailsapp/wails/v3/cmd/wails3@latest

# 安装指定版本
go install github.com/wailsapp/wails/v3/cmd/wails3@v3.0.0-alpha.40
go mod tidy

# 启动开发模式 - 初次运行需要执行两边
npm run app-dev

# 打包
npm run app-build

# 同步当前分支到远端
npm run git-sync

# 清理工作区
npm run clear-dir

```
