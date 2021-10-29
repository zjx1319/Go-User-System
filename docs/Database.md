# Database 文档

### 数据库结构：users

```
id		 int UNIQUE,
username varchar(20) UNIQUE,    -- 用户名
password char(32),				-- MD5加密后的密码
email	 varchar(50) UNIQUE,		-- 邮箱地址
verified boolean DEFAULT FALSE,			-- 是否已验证邮箱
role	 varchar(20) DEFAULT 'default',	-- 权限组
PRIMARY KEY(id),
```



### 数据库结构：email

```
id		 int REFERENCES users(id) ON DELETE CASCADE UNIQUE,
email	 varchar(50) REFERENCES users(email),
code	 char(32),				-- 邮箱验证码
```



### 数据库结构：wx

```
id		 int REFERENCES users(id) ON DELETE CASCADE UNIQUE,
wxname	 varchar(20),	-- 微信用户名
openid	 char(28),	-- 微信id
```



id本来想用Postgresql中serial自增的，发现插入未成功就会断层，干脆自己整一个

```sql
create or replace function setSerialVal() returns integer as $$
declare 
    nextid integer;
begin
    select max(id) into nextid from users;
    if nextid is null then nextid = 1;
    else nextid = nextid + 1;
    end if;
    return nextid;
end;
$$ language plpgsql;
```

