### Nginx 配置文件深度解析（基于 Docker 微服务场景）

---

#### 配置文件核心结构概览

```bash
# 全局配置 → events 块 → http 块 → server 块 → location 块
# 层级关系：全局 → 网络模型 → HTTP 协议 → 虚拟主机 → 路由规则
```

---

### 一、全局配置块（Main Context）

```nginx
user root;                     
worker_processes auto;         
```

| 配置项                   | 作用说明                                         | 最佳实践                                                     |
| ------------------------ | ------------------------------------------------ | ------------------------------------------------------------ |
| `user root;`             | 定义 Nginx 工作进程的运行用户                    | **生产环境必须替换为普通用户**（如 `user nginx;`），避免 root 权限风险 |
| `worker_processes auto;` | 设置工作进程数量，`auto` 表示自动匹配 CPU 核心数 | 通常设为 `auto`，也可手动指定（如 4 核 CPU 写 `worker_processes 4;`） |

---

### 二、事件块（Events Context）

```nginx
events {
    worker_connections 1024; 
}
```

| 配置项               | 作用说明                           | 计算逻辑                                                   |
| -------------------- | ---------------------------------- | ---------------------------------------------------------- |
| `worker_connections` | 单个工作进程可同时处理的最大连接数 | **最大并发连接数 = worker_processes × worker_connections** |
|                      |                                    | 示例：4 进程 × 1024 = 4096 并发                            |

---

### 三、HTTP 块（HTTP Context）

#### 1. 压缩配置（Gzip）

```nginx
gzip on;
gzip_min_length 5k;
gzip_buffers 4 16k;
gzip_http_version 1.0;
gzip_comp_level 7;
gzip_types text/plain application/javascript text/css application/xml;
gzip_vary on;
```

| 配置项            | 作用说明                                         | 典型值                                  |
| ----------------- | ------------------------------------------------ | --------------------------------------- |
| `gzip on`         | 开启 Gzip 压缩                                   | `on`（必开）                            |
| `gzip_min_length` | 触发压缩的最小文件大小                           | 5KB（避免压缩小文件得不偿失）           |
| `gzip_comp_level` | 压缩级别（1-9），值越高压缩率越大但 CPU 消耗越高 | 折中选择 `6` 或 `7`                     |
| `gzip_types`      | 指定需要压缩的 MIME 类型                         | 按需添加（如 JSON：`application/json`） |

#### 2. 基础传输配置

```nginx
include mime.types;               
default_type application/octet-stream;
sendfile on;                      
keepalive_timeout 500;            
client_max_body_size 2000m;       
```

| 配置项                 | 作用说明                                                    | 注意事项                                                 |
| ---------------------- | ----------------------------------------------------------- | -------------------------------------------------------- |
| `include mime.types`   | 引入 MIME 类型定义文件（识别文件扩展名对应的 Content-Type） | 确保文件路径正确（容器内路径为 `/etc/nginx/mime.types`） |
| `client_max_body_size` | 允许客户端上传的最大文件体积                                | 根据业务需求调整（如视频上传设为 `4096m`）               |

---

### 四、虚拟主机块（Server Context）

#### 1. 监听与日志

```nginx
server {
    listen 80;
    access_log /var/log/nginx/access.log;
    error_log /var/log/nginx/error.log;
}
```

| 配置项       | 作用说明                          | 调试技巧                                            |
| ------------ | --------------------------------- | --------------------------------------------------- |
| `listen 80`  | 监听容器内部的 80 端口            | Docker 端口映射 `-p 9000:80` 表示外部通过 9000 访问 |
| `access_log` | 记录所有请求的访问日志            | 调试时开启，生产环境可对敏感路径关闭                |
| `error_log`  | 记录错误日志（含 502 等代理错误） | 使用 `tail -f error.log` 实时监控错误               |

#### 2. 通用路由配置

```nginx
location / {
    gzip_static on;  # 优先发送预压缩文件（如 .js.gz）
}
```

| 配置项           | 作用说明                                           | 使用场景                                   |
| ---------------- | -------------------------------------------------- | ------------------------------------------ |
| `gzip_static on` | 直接发送已预先压缩的静态文件（需配合 `gzip` 使用） | 适用于 CSS/JS 等静态资源，减少实时压缩开销 |

---

### 五、微服务路由配置（Location Context）

#### 1. 鉴权服务路由（Auth Service）

```nginx
location /api/auth/ {
    # 代理到宿主机服务（关键配置）
    proxy_pass http://host.docker.internal:20261; 

    # 透传原始请求信息（核心头配置）
    proxy_set_header Host $http_host;             
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_set_header X-Original-URI $request_uri;
}
```

| 配置项                            | 作用说明                  | 必要性                                             |
| --------------------------------- | ------------------------- | -------------------------------------------------- |
| `proxy_pass`                      | 定义后端服务地址          | **必须使用 `host.docker.internal` 访问宿主机服务** |
| `proxy_set_header Host`           | 传递客户端原始域名        | 防止后端服务因域名错误拒绝响应                     |
| `proxy_set_header X-Real-IP`      | 传递客户端真实 IP         | 后端记录日志或 IP 限制时必备                       |
| `proxy_set_header X-Original-URI` | 透传未 Rewrite 的原始 URI | 用于后端校验请求路径                               |

#### 2. 需鉴权的用户服务路由（User Service）

```nginx
location /api/user/ {
    # 鉴权检查：先调用 Auth 服务（返回 200 才放行）
    auth_request /api/auth/; 

    # 代理到宿主机服务
    proxy_pass http://host.docker.internal:20262;

    # 头信息透传（与 Auth 服务一致）
    proxy_set_header Host $http_host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
}
```

| 配置项         | 作用说明             | 常见错误                                                     |
| -------------- | -------------------- | ------------------------------------------------------------ |
| `auth_request` | 定义鉴权接口路径     | **路径必须完全匹配**（如 `/api/auth/` vs `/api/auth` 可能导致 404） |
| `proxy_pass`   | 需鉴权的业务服务地址 | 确保所有服务地址均指向 `host.docker.internal`                |

---

### 六、Docker 特有配置注意事项

#### 1. 宿主机服务访问问题

```nginx
# 错误配置（容器内无法访问宿主机）
proxy_pass http://127.0.0.1:20261; 

# 正确配置（Docker 提供的特殊 DNS）
proxy_pass http://host.docker.internal:20261;
```

- **原理**：Docker 容器内的 `127.0.0.1` 指向容器自身，非宿主机。

- **替代方案**：若环境不支持 `host.docker.internal`，可通过启动参数传递宿主机 IP：

  ```bash
  docker run --add-host=host.docker.internal:host-gateway ...
  ```

#### 2. 服务绑定地址问题

- **错误现象**：Nginx 报 `502 Bad Gateway`

- **根本原因**：宿主机上的服务仅绑定到 `127.0.0.1`，拒绝容器连接。

- **修复方案**：确保服务监听 `0.0.0.0`（Go 语言示例）：

  ```go
  http.ListenAndServe("0.0.0.0:20261", nil)
  ```

---

### 七、完整配置优化建议

#### 1. 路径一致性修正

```diff
# 原配置（存在路径结尾不一致问题）
location /api/chat/ {
    auth_request /api/auth;   # ❌ 缺少结尾 /
    proxy_pass http://127.0.0.1:20263; # ❌ 未使用 host.docker.internal
}

# 修正后
location /api/chat/ {
    auth_request /api/auth/;  # ✅ 统一以 / 结尾
    proxy_pass http://host.docker.internal:20263; # ✅ 正确代理地址
}
```

#### 2. 生产环境安全加固

```nginx
user nginx;  # 改用非 root 用户
server_tokens off;  # 隐藏 Nginx 版本号
client_max_body_size 100m;  # 按需限制上传大小

# 限制敏感接口的 HTTP 方法
location /api/auth/ {
    limit_except POST { deny all; }  # 只允许 POST 请求
}
```

---

### 八、调试命令速查表

| 场景                   | 命令                                                         |
| ---------------------- | ------------------------------------------------------------ |
| 重新加载 Nginx 配置    | `docker exec my-nginx nginx -s reload`                       |
| 查看实时访问日志       | `docker exec my-nginx tail -f /var/log/nginx/access.log`     |
| 测试容器到宿主机连通性 | `docker exec my-nginx curl http://host.docker.internal:20261/health` |
| 检查配置语法错误       | `docker exec my-nginx nginx -t`                              |

---

