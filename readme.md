# ローカル開発環境の手引き

## SAMCLIのインストール

https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/sam-cli-requirements.html  
を参考にする。  

## 実行手順

### template.ymlの作成

`sam init --runtime go`で初期後、`template.yml` 以外は削除。  
`template.yml`を残すのは、ローカル環境下での動作確認を行うため。

### Netlify用の設定を実施

Netlify公式が用意してくれている以下サンプルから、  
`netlify.toml`や`Makefile`を元にビルド等の設定を実施。  

- https://github.com/netlify/aws-lambda-go-example

### ビルド＆lambda起動

`sam init〜`の際に作成された`Readme.md`にも記載ある通り、  
`GOOS=linux GOARCH=amd64 go build -o 〜`で実行ファイル生成。  

配置した実行ファイルのパスやファイル名を元に、  
`template.yml`の内容を編集後、`sam local start-api`で起動。
