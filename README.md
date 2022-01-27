# web-app-template
Frontend, Backend(Golang), RDBのWebappの開発用テンプレート

## コマンド

`Makefile`で定義
- `db/init`
    - データベース作製
    - migration前に実行する必要あり

## DB定義
- https://dbdiagram.io/home
    - E-R図で作成可能
    - SQLとしてエクスポートも可能

## 各サービスの役割
- db
    - データ永続化
    - RDBMSでもNoSQLでもなんでも
    - 
- adminer
    - データベース内のデータ確認など
    - コンソールからアクセスするより楽
- migration
    - [flyway](https://github.com/flyway/flyway-docker)を使用
    - `database/migration/sql/`以下にSQLを設置すると`db`にmigrationしてくれる
    - ファイル名の形式に注意
        - https://flywaydb.org/documentation/concepts/migrations#sql-based-migrations
- backend
    - 現在は`go 1.17.6`を使用
        - 最終的には`backend-dev`みたいなサービスを作ってその中にVSC serverを入れてその中で開発する形にしたい
        - localにgo入れるのあまりうれしくないので...
    - apiの配信
- go-swagger
    - `go-swagger`を利用してレスポンス用のgoの構造体を生成する
    - `./swagger`ディレクトリ内の`swagger.yml`をfront, backend双方で参照することで統一した形式で通信ができるようにする