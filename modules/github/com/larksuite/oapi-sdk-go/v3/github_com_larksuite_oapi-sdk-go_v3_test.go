package lark

import (
	"github.com/ZenLiuCN/engine"
	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
	"github.com/ZenLiuCN/fn"

	"testing"
)

func TestSimple(t *testing.T) {
	v := engine.Get()
	defer v.Free()
	buf := engine.GetBytesBuffer()
	defer func() {

		println(buf.String())
	}()
	v.BufferConsole(buf)
	v.Debug = true
	fn.Panic1(v.RunTs(
		//language=typescript
		`
		import * as lark from 'github.com/larksuite/oapi-sdk-go/v3'
		import * as context from 'golang/context'
		import * as btv1 from 'github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1'
		import * as core from 'github.com/larksuite/oapi-sdk-go/v3/core'
		
		`))
}
