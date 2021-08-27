package main

import (
	"chat/trace"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
	"gopkg.in/ini.v1"
)

// ConfigList 設定ファイルから取得したデータを保持する構造体
type ConfigList struct {
    googleClientID     string
    googleClientSecret string
}
// Config 設定リスト保持変数
var Config ConfigList
// コンストラクタ
func init() {
    // ファイル読み込み
    cfg, err := ini.Load("conf/app.conf")
    if err != nil {
		log.Fatalln("設定ファイルを読み取れませんでした: ",err)
    }

    // 変数に設定
    Config = ConfigList{
        googleClientID:     cfg.Section("oauth").Key("googleClientID").String(),
        googleClientSecret: cfg.Section("oauth").Key("googleClientSecret").String(),
    }
}


type templateHandler struct {
	once     sync.Once
	filename string
	temp1    *template.Template
}

//serve template html file. template is compiled only once.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.temp1 =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	data:=map[string]interface{}{
		"Host": r.Host,
	}
	if authCookie,err:=r.Cookie("auth");err==nil{
		data["UserData"]=objx.MustFromBase64(authCookie.Value)
	}
	t.temp1.Execute(w, data)
}

func main() {
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() //フラグを解釈する
	gomniauth.SetSecurityKey("security_key")
	gomniauth.WithProviders(
		//facebook.New("クライアントID","秘密の値","http://localhost:8080/auth/callback/facebook"),
		//github.New("クライアントID","秘密の値","http://localhost:8080/auth/callback/github"),
		google.New(Config.googleClientID,Config.googleClientSecret,"http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	//if you want to set trace, uncomment the following
	r.tracer = trace.New(os.Stdout)
	http.Handle("/chat",MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login",&templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/",loginHandler)
	http.Handle("/room", r)
	//チャットルームを開始
	go r.run()
	log.Println("Webサーバを開始します。ポート: ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
