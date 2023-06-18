# 徽章

[![CI](https://github.com/xh-polaris/meowchat-api/actions/workflows/static-analysis.yml/badge.svg)](https://github.com/xh-polaris/meowchat-api/actions/workflows/static-analysis.yml)
[![Build](https://github.com/xh-polaris/meowchat-api/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/xh-polaris/meowchat-api/actions/workflows/docker-publish.yml)

# 指南

**启动服务**

```bash
make start # 或者仅输入make
```

在启动服务之前，请先通过设置环境变量CONFIG_PATH来指定配置文件路径，如果没有设置的话，默认使用``etc/config.yaml``

例如这样指定配置文件路径：

```bash
make start CONFIG_PATH='etc/config.yaml'
```

**更新服务**

```bash
make update
```


