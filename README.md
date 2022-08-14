
## 生成swagger命令
---
```markdown
swag init
```
---

gitee地址： https://gitee.com/hutiancai/memo.git

### 试题系统

---
```markdown
Username:38hjj6hmo9uj
Password:pscale_pw_mo2G89p_l__PczcNgf9CBA2nYI4Xj8MuPkac444wd0w
```
---

## rbac模型 
---
```markdown
    [request_definition]
    r = sub, obj, act
    
    [policy_definition]
    p = sub, obj, act
    
    [role_definition]
    g = _, _
    
    [policy_effect]
    e = some(where (p.eft == allow))
    
    [matchers]
    m = r.sub == p.sub && ParamsMatch(r.obj,p.obj) && r.act == p.act
```
---


### minio对象存储
```shell
1 nohup sudo MINIO_ACCESS_KEY=minioadmin MINIO_SECRET_KEY=123456789 
/opt/mino/mino server --address=0.0.0.0:9000 --console-address=0.0.0.0:9001 
--config-dir /usr/local/mino/etc/ /usr/local/mino/data/ > /usr/local/mino/mino.log 2>&1&

```

* minio启动脚本
```shell
#!/usr/bin/env bash
# 指定minio的账号
export MINIO_ACCESS_KEY=minio
# 指定minio的密码
export MINIO_SECRET_KEY=minio123321
nohup /opt/mino/bin/mino server  /opt/mino/data > /opt/mino/log/mino.log 2>&1 &
```

minio web客户端访问地址： http://101.132.251.60:9000