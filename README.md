# web-app-template
Frontend, Backend(Golang), RDBのWebappの開発用テンプレート

## 環境起動手順
1. 本リポジトリをfork
2. `.devcontainer/devcontainer.json`内の`service`propertyを変更する
    - フロントエンド開発時は`frontend`
    - バックエンド開発時は`backend`
    - Visual Studio Codeサーバーをインストールする対象のコンテナを指定する
3. `F1`からコマンドパレットを開き`Reopen in container`を実行、開いたサーバー内で作業
4. Migrationなどの実行時はコンテナの外のターミナルから`Make`コマンドを実行する

## コマンド
`Makefile`で定義
- `db/init`
    - データベース作製
    - migration前に実行する必要あり

## 各種実装方針
- https://12factor.net/ja/ を参考にする

## DB定義
- https://dbdiagram.io/home
    - E-R図で作成可能
    - SQLとしてエクスポートも可能

## 各サービスの役割
一部のサービスは各ディレクトリ内の`docker-compose.yml`内で定義されています

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
    - apiの配信
- go-swagger
    - `go-swagger`を利用してレスポンス用のgoの構造体を生成する
    - `./swagger`ディレクトリ内の`swagger.yml`をfront, backend双方で参照することで統一した形式で通信ができるようにする