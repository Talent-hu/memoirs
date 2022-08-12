
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





