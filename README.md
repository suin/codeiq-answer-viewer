# CodeIQ Answer Viewer

[![Build Status](https://drone.io/github.com/suin/codeiq-answer-viewer/status.png)](https://drone.io/github.com/suin/codeiq-answer-viewer/latest)

CodeIQ出題者のための解答CSVを閲覧するWebアプリです。
Goで書かれています。

![](https://raw.githubusercontent.com/suin/codeiq-answer-viewer/master/image.png)

## 使い方

Macで下記コマンドを実行してウェブサーバを立てる

```
curl https://drone.io/github.com/suin/codeiq-answer-viewer/files/artifacts/darwin_amd64/codeiq-answer-viewer.tar.gz | tar xzvf -
./codeiq-answer-viewer
```

ポート 51019(こー・ど・あい・きゅう) で立ち上がるのでブラウザで開く

```
open http://localhost:51019
```

CodeIQ管理画面からダウンロードしてきた解答CSVをアップロードする



## コントリビューション

必要なもの

* coffee コマンド(CoffeeScript)
* make コマンド
* go 1.2

チェックアウトしてきたら依存ライブラリを入れる:

```
go get ./...
```


makeコマンドを実行するとデバッグ用の環境が立ち上がります:

```
make
```

※ assets の変更は make し直すと反映されます(Goのバイナリに含める都合上、動的ロードではない)

## ライセンス

MITライセンス

