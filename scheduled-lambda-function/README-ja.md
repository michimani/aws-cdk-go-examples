scheduled-lambda-function
===

Lambda 関数を EventBridge で設定した CRON 式によって定期実行する構成の実装例です。

# 主なリソース

## Lambda::function

- Runtime: `go1.x`
- Timeout: 30 sec
- Memory: 128 MB
- Operation:
  - リクエスト ID と実行時の Unix 時刻とともに `Hello AWS CDK with Golang.` と標準出力に出力します。
  - `RequestResponse` で起動した場合は、下記のようなレスポンスを返します。

    ```json
    {
      "message": "Hello AWS CDK with Golang.",
      "requestId": "71bd4fd1-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
      "timestamp": 1658236882
    }
    ```

## Events::Rule

- 下記の CRON 式で指定されたスケジュールルールです。

  ```bash
  0 */4 * * ? *
  ```

- 前述した Lambda 関数を Rule のターゲットとします。

# 使い方

## Lambda 関数のビルド

```bash
make build
```

## テスト

```bash
go test .
```

## デプロイ

```bash
cdk synth
```

```bash
cdk deploy
```

## 手動で Lambda 関数を実行

```bash
aws lambda invoke \
--function-name hello-aws-cdk-golang-function \
--invocation-type RequestResponse \
--region ap-northeast-1 \
out && cat out
```

# Author

[michimani210](https://twitter.com/michimani210)