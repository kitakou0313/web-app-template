# web-app-template
Frontend, Backend(Golang), RDBのWebappの開発用テンプレート

## コマンド

`Makefile`で定義
- `db/init`
    - データベース作製
    - migration前に実行する必要あり

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
- api
    - apiの配信