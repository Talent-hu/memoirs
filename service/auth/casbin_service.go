package auth

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"memoirs/global"
	"memoirs/model/vo"
	"sync"
)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

func (casbinService *CasbinService) UpdateCasbin(roleCode string, casbinInfos []vo.CasbinInfo) error {
	casbinService.ClearCasbin(0, roleCode)
	rules := [][]string{}
	for _, v := range casbinInfos {
		rules = append(rules, []string{roleCode, v.Path, v.Method})
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

func (casbinService *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.DB.Model(&gormAdapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

func (casbinService *CasbinService) GetPolicyPathByRoleId(roleCode string) (pathMaps []vo.CasbinInfo) {
	e := casbinService.Casbin()
	list := e.GetFilteredPolicy(0, roleCode)
	for _, v := range list {
		pathMaps = append(pathMaps, vo.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

var (
	syncEnforcer *casbin.SyncedEnforcer
	once         sync.Once
)

func (casbinService *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		adapter, _ := gormAdapter.NewAdapterByDB(global.DB)
		text := `
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
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			global.Log.Error("记载字符串模式失败", zap.Error(err))
			return
		}
		syncEnforcer, _ = casbin.NewSyncedEnforcer(m, adapter)
	})
	_ = syncEnforcer.LoadPolicy()
	return syncEnforcer
}
