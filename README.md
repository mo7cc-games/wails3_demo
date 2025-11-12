# wails3_demo

## 项目介绍
基于 Wails3 框架的桌面应用示例项目

包含无边框，前后端窗口通信等等

## 环境要求

- Go 1.25+
- Node.js 24+

```bash
# 安装最新的 Wails3
go install github.com/wailsapp/wails/v3/cmd/wails3@latest

```

# 常用 CLI 命令

```bash
npm install -g del-cli

# 清理并更新全部依赖
npm run re-update

# 查看 Wails3 版本
wails3 version

# 更新依赖- 小心执行，可能会导致 wails3 无法使用
go get -u ./...
go mod tidy

# 安装项目脚本依赖
npm install

# 安装 frontend 依赖
cd frontend && npm install && cd ..

# 启动开发环境
npm run dev-app

# 构建发布版本
npm run build-app

# 清理目录
npm run clear-dir

# 同步当前分支到远端
npm run git-sync <message>

```
