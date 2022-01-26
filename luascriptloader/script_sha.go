package luascriptloader

const (

	// 加载状态
	ScriptStatusUnLoaded = 0 // 未加载
	ScriptStatusReady    = 1 // 已加载可用
	ScriptStatusLoading  = 2 // 正在加载
)

type Shas struct {
	Sha    string
	Script string // 脚本SHA1校验值
	Status int64  // 脚本加载状态，0：未加载；1：已加载可用；2：正在加载
	Name   string // 脚本名称
}

//项目启动的时候，请把所有需要用到的lua脚本赋值给ScriptsDefine，在使用的时候就只有map读，不用加锁
var ScriptsDefine = map[string]*Shas{}
