# API 文档

### 认证模式

JWT，放在 HTTP 请求的头信息Authorization字段

```
Authorization: Bearer <JWT token>
```

### API 定义

API 前缀：`/api/v1`

### 1. 注册

请求：

```
POST /user
{
    "username": "用户名",
    "password": "密码",
    "email": "邮箱地址"
}
```

响应：

```
{
    "id": "用户 ID"
}
```

### 2. 登录（获取JWT令牌）

请求：

```
GET /user/token?username="用户名"&password="密码"
```

响应：

```
{
    "token": "令牌",
    "expire_time": 123456789 // 令牌到期时间
}
```

### 3. 验证邮箱

请求：

```
GET /user/verify?id="用户ID"&code="验证码"
```

响应：无

URL会发送到待验证的邮箱中

### 4. 获取用户信息

权限：

```
default：获取当前用户信息
admin：获取任何用户信息（不指定ID时可获取所有用户）
```

请求：

```
GET /user/:id
```

响应：

```
{
    "result": [
        {
            "id": "用户 ID",
            "username": "用户名",
            "email": "邮箱地址",
            "role": "权限组"
        }
        // ...
    ]
}
```

### 5. 修改用户信息

权限：

```
default：修改当前用户的用户名、密码、邮箱地址（需重新验证）
admin：修改任何用户的用户名、密码、邮箱地址、权限组
不修改留空即可
```

请求：

```
PUT /user/:id
{
    "id": "用户 ID",
    "username": "用户名",
    "password": "密码",
    "email": "邮箱地址",
    "role": "权限组"
}
```

响应：

```
{
	"id": "用户 ID",
	"username": "用户名",
	"email": "邮箱地址",
	"role": "权限组"
}
```

### 6. 删除用户

权限：

```
需要admin权限
```

请求：

```
DELETE /user/:id
```

响应：无

### 7. 通过微信登录

说明：CODE由微信登录成功返回，作为换取access_token的票据

请求：

```
GET /user/tokenWX?code="CODE"
```

响应：

```
{
    "token": "令牌",
    "expire_time": 123456789 // 令牌到期时间
}
```

### 8. 绑定微信

权限：

```
需要当前用户令牌
```

请求：

```
GET /user/bindW?code="CODE"
```

响应：无

