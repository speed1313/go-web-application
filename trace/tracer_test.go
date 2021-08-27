
package trace

import (
	"bytes"
	"testing"
)

//名前がTestで始まり, *testing.T型の引数を一つ受け取る関数は全てユニットテストとみなされる.テストを実行するとこの条件を満たす関数が全て呼び出される.
func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	tracer.Trace("こんにちは, traceパッケージ")
	if buf.String() != "こんにちは, traceパッケージ\n" {
		t.Errorf("'%s'という誤った文字列が出力されました", buf.String())
	}

}

func TestOff(t *testing.T){
	var silentTracer Tracer
	silentTracer.Trace("データ")
}
