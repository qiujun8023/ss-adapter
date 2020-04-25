shadowsocks-libev adapter for ss-panel
===========

## 部署

### 使用 Docker + Docker Compose 部署

- 配置 docker-compose.yml

```bash
cat > ./docker-compose.yml << \EOF
version: '3'
services:
  libev:
    image: shadowsocks/shadowsocks-libev
    restart: always
    network_mode: host
    user: root
    command: ss-manager --manager-address 127.0.0.1:8839 -s 0.0.0.0 -m aes-256-cfb -u -v
  adapter:
    image: qiujun8023/shadowsocks:libev
    restart: always
    network_mode: host
    environment:
      APP_MANAGER_SERVER: '127.0.0.1'
      APP_MANAGER_PORT: 8839
      APP_API_URL: https://example.com/
      APP_NODE_ID: 1
      APP_NODE_TOKEN: token
      APP_SYNC_INTERVAL: 30
EOF
```

- 运行

```bash
docker-compose up -d
```

### 环境变量说明

| 字段   | 描述   |
|:----|:----|
| APP_MANAGER_SERVER   | ss-manager 对应的地址   |
| APP_MANAGER_PORT   | ss-manager 对应的端口   |
| APP_API_URL   | ss-panel 对应的地址   |
| APP_NODE_ID   | ss-panel 对应的节点编号   |
| APP_NODE_TOKEN   | ss-panel 对应的节点Token   |
| APP_SYNC_INTERVAL   | shadowsocks 与 ss-panel 的交互时间间隔   |