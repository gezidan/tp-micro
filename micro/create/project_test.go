package create

import (
	"strings"
	"testing"

	"github.com/xiaoenai/tp-micro/micro/info"
)

func TestGenerator(t *testing.T) {
	info.Init("test")
	proj := NewProject([]byte(src))
	proj.gen()
	t.Logf("main.go:\n%s", proj.codeFiles["main.go"])
	t.Logf("config.go:\n%s", proj.codeFiles["config.go"])
	t.Logf("args/const.gen.go:\n%s", proj.codeFiles["args/const.gen.go"])
	t.Logf("args/type.gen.go:\n%s", proj.codeFiles["args/type.gen.go"])
	t.Logf("logic/tmp_code.gen.go:\n%s", proj.codeFiles["logic/tmp_code.gen.go"])
	t.Logf("logic/model/init.go:\n%s", proj.codeFiles["logic/model/init.go"])
	t.Logf("api/pull_handler.gen.go:\n%s", proj.codeFiles["api/pull_handler.gen.go"])
	t.Logf("api/push_handler.gen.go:\n%s", proj.codeFiles["api/push_handler.gen.go"])
	t.Logf("api/router.gen.go:\n%s", proj.codeFiles["api/router.gen.go"])
	t.Logf("sdk/rpc.gen.go:\n%s", proj.codeFiles["sdk/rpc.gen.go"])
	t.Logf("sdk/rpc_test.gen.go:\n%s", proj.codeFiles["sdk/rpc.gen_test.go"])
	for k, v := range proj.codeFiles {
		if strings.HasPrefix(k, "logic/model") {
			t.Logf("%s:\n%s", k, v)
		}
	}
}

const src = `// package __TPL__ is the project template
package __TPL__

// __API__PULL__ register PULL router:
//  /home
//  /math/divide
type __API__PULL__ interface {
	Home(*struct{}) *HomeResult
	Math
}

// __API__PUSH__ register PUSH router:
//  /stat
type __API__PUSH__ interface {
	Stat(*StatArg)
}

// __MYSQL__MODEL__ create mysql model
type __MYSQL__MODEL__ struct {
	User
	Log
	Device
}

// __MONGO__MODEL__ create mongodb model
type __MONGO__MODEL__ struct {
	Meta
}

// Math controller
type Math interface {
	// Divide handler
	Divide(*DivideArg) *DivideResult
}

// HomeResult home result
type HomeResult struct {
	Content string // text
}

type (
	// DivideArg divide api arg
	DivideArg struct {
		// dividend
		A float64
		// divisor
		B float64 ` + "`param:\"<range: 0.01:100000>\"`" + `
	}
	// DivideResult divide api result
	DivideResult struct {
		// quotient
		C float64
	}
)

// StatArg stat handler arg
type StatArg struct {
	Ts int64 // timestamps
}

// User user info
type User struct {
	Id   int64  ` + "`key:\"pri\"`" + `
	Name string ` + "`key:\"uni\"`" + `
	Age  int32
}

type Log struct {
	Text string
}

type Device struct {
	UUID string ` + "`key:\"pri\"`" + `
}

type Meta struct {
	Hobby []string
	Tags  []string
}
`
