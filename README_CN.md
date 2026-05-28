# CLIProxyAPIPlus 维护分支

这是我们维护的 CLIProxyAPIPlus 分支，不是上游文档镜像。

当前维护策略：

- 以 `HsnSaboor/CLIProxyAPIPlus` 作为 Plus 基线。
- `router-for-me/CLIProxyAPI` 只作为后续选择性补丁来源。
- 不直接合并 router 中删除 Plus provider 的改动。
- 不同步上游的 `README*`、`AGENTS.md`、`CLAUDE.md` 或 `docs/**` 文档内容。

## 文档边界

- `docs/<scope>/`：当前活跃 scope 的合同、计划、检查清单和证据。
- `docs/archive/<scope>/`：已完成 scope 的存档。
- `.codex/wiki/**`：长期维护知识、决策、经验和代码库说明。

`docs/` 可以继续按照 repo-task-driven workflow 放我们的 scope 和归档，但不要放上游文档镜像。

## 已完成 scope

- `docs/archive/hsnsaboor-clean-root/`

该 scope 已完成：以 HsnSaboor 当前 `upstream/main` 为干净根，再把我们的维护策略作为新的 commit 叠加上去。

## 开发命令

```bash
gofmt -w .
go build -o cli-proxy-api ./cmd/server
go build -o test-output ./cmd/server && rm test-output
go test ./...
```

详细维护规则见：

- `README.md`
- `AGENTS.md`
- `.codex/wiki/reference/upstream-plus-maintenance.md`
