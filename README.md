# login-app-api

　このアプリケーションは、github.com/Shakkuuu/login-app-api-use を使用するためのAPIです。  
　dockerでGoのサーバーとmysqlを起動させ、localhost:8081 にアクセスしてGETやPOST、PUT、DELETEを行います。  

## 使い方

1. dockerで起動  
``` docker compose up ```  

* localhost:8081で起動される
* mysqlは3307ポートで起動される

## entity

* User
ID,Name,Password,Ticket,Coin  
(Coin:Coin entityに移行したため、実質未使用)  

* Coin
ID,Name,Qty,Speed,Speedneed  
(Name:ユーザーごとの所持コイン識別用,Speed:minigameでのワンクリックごとの取得コイン枚数,Speedneed:minigameでのワンクリックごとの取得コイン枚数アップグレードに必要なコイン枚数)  

* Coin
ID,Name,Title,Text  
(Name:ユーザーごとのメモ識別用)  

## users

GET  

* "/" 全てのユーザーのユーザー情報を取得する
* "/showid/:id" idを指定してユーザー情報の取得
* "/showname/:username" usernameを指定してユーザー情報の取得
* "/tiadd/:username" usernameを指定して、チケットの追加
* "/tisub/:username" usernameを指定して、チケットの削除

POST  

* "/" ユーザーの追加

PUT  

* "/:id" ユーザー情報の更新

DELETE  

* "/:id" ユーザーの削除

## memos

GET  

* "/" 全てのメモを取得する
* "/showname/:username" usernameを指定してそのユーザーのメモの取得

POST  

* "/" メモの追加

DELETE  

* "/:username" usernameで指定して、指定したユーザーのメモを全て削除する

## gamecoin

GET  

* "/" 全てのユーザーのコイン情報を取得する
* "/showname/:username" usernameを指定してコイン情報の取得

POST  

* "/" ユーザーにコイン登録

PUT  

* "/:username" ユーザーのコイン,スピード,スピードアップに必要なコイン枚数 の増減

DELETE  

* "/:username" ユーザーからコイン情報の削除

## 開発環境

* macbook air M1
* Visual Studio Code
* go 1.19.5 darwin/arm64
* docker v3
* mysql v8.0
* 使用パッケージ  
標準  
``` fmt, time ```  
その他  
```github.com/go-sql-driver/mysql, github.com/jinzhu/gorm, github.com/gin-gonic/gin```

## 今後実装予定のもの

* メモごとの編集と削除機能
* 未使用のUser.Coinの整理
* Ticketの有効活用
* minigameの機能追加
* :idで指定しているところをusernameでリクエストできるように修正
