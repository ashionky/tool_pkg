package luascriptloader

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type LuaScriptLoader struct {
	redisClient *redis.Client
}

func NewLuaScriptLoader(redisClient *redis.Client) *LuaScriptLoader {
	return &LuaScriptLoader{
		redisClient: redisClient,
	}
}

func (r *LuaScriptLoader) ExecScript(ctx context.Context, name string, keys []string, args []interface{}) (res interface{}, err error) {
	loader, ok := ScriptsDefine[name]
	if !ok {
		return nil, errors.New("unknown script name when loading redis scripts")
	}
	if loader.Sha == "" {
		if loader.Status == ScriptStatusUnLoaded {
			loader.Status = ScriptStatusLoading
			loader.Sha, err = r.redisClient.ScriptLoad(ctx, loader.Script).Result()
			if err != nil {
				return nil, errors.New("failed to load script when loading redis scripts: " + err.Error())
			}
			loader.Status = ScriptStatusReady
			res, err = r.redisClient.EvalSha(ctx, loader.Sha, keys, args...).Result()
			if err != nil {
				err = errors.New("failed to exec script: " + err.Error())
			}
			return res, err
		} else if loader.Status == ScriptStatusLoading {
			for {
				if loader.Status == ScriptStatusLoading {
					err = r.waitingConsumeScript(loader, time.Now().Unix())
					if err != nil {
						return nil, errors.New("waiting consume script failed of " + loader.Name)
					} else {
						break
					}
				}
			}
		}
	}
	res, err = r.redisClient.EvalSha(ctx, loader.Sha, keys, args...).Result()
	if err != nil {
		err = errors.New("failed to exec script:" + err.Error())
	}
	return res, err
}

func (r *LuaScriptLoader) waitingConsumeScript(loader *Shas, execTime int64) (err error) {
	t := time.NewTimer(time.Microsecond * 10)
	defer t.Stop()
	for {
		<-t.C
		if loader.Status == ScriptStatusUnLoaded {
			return errors.New("script not load when waiting consume script of " + loader.Name)
		} else if loader.Status == ScriptStatusLoading {
			err = r.waitingConsumeScript(loader, execTime)
			return err
		} else {
			return nil
		}
	}
}
