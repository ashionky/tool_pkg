/**
 * @Author pibing
 * @create 2022/1/24 6:37 PM
 */

package luascriptloader

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

func luatest() string {
	script := `
       if #KEYS ~= 1 or #ARGV ~= 1 then
			return -1
		end
        redis.pcall('set',KEYS[1],ARGV[1] )
        return 1
        
`
	return script
}
func TestNewLuaScriptLoader(t *testing.T) {
	ScriptsDefine["test"] = &Shas{
		Sha:    "",
		Script: luatest(),
		Status: ScriptStatusUnLoaded,
		Name:   "test",
	}

	var ctx = context.Background()

	var client *redis.Client
	load := NewLuaScriptLoader(client)
	res, err := load.ExecScript(ctx, "test", []string{"key"}, []interface{}{"badf"})

	fmt.Println(res, err)
}
