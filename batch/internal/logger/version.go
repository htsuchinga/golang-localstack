package logger

import "fmt"

// ビルドハッシュ(go build時に値を置き換える)
var BuildHash = "xxxxxxxxxxxxxxx"

// ビルド日時(go build時に値を置き換える)
var BuildDate = "YYYY-MM-DDThh:mm:ssJST"

// Goコンパイラのバージョン(go build時に値を置き換える)
var GoVersion = "go version goX.XX.XX os/arch/"

// バージョン情報文字列を生成
func Version() string {
	hash := BuildHash
	if len(BuildHash) > 7 {
		hash = BuildHash[:7]
	}

	compiler := GoVersion
	if len(GoVersion) > 13 {
		compiler = GoVersion[:13]
	}

	return fmt.Sprintf("varsion:%s build:%s golang:%s", hash, BuildDate, compiler)
}
