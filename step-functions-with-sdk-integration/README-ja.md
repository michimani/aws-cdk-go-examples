step-functions-with-sdk-integration
===

StepFunctions でいくつかの AWS SDK インテグレーションを利用したステートマシンを構築する実装例です。

# 主なリソース

## Logs::LogGroup

- ステートマシンの実行ログを出力するための LogGroup です

## StepFunctions::StateMachine

- テキストを翻訳するステートマシンです
- 下記の AWS SDK インテグレーションを利用します:
  - Translate: translateText
  - S3: putObject

## S3::Bucket

- 翻訳結果を出力するための S3 Bucket です

# 使い方

## 作成する S3 Bucket のバケット名を環境変数に設定

```bash
export OUTPUT_BUCKET_NAME='output-bucket-name'
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

## ステートマシンを起動

```bash
aws stepfunctions start-execution \
--state-machine-arn $(
  aws stepfunctions list-state-machines \
  --query "stateMachines[?name=='sdk-integration-example-state-machine'].stateMachineArn" \
  --output text) \
--input file://testdata/statemachine_input.json
```

結果を確認

```bash
aws s3 cp "s3://${OUTPUT_BUCKET_NAME}/translate-result" -

# "こんにちは。これは、Go 言語で AWS CDK を使用するための実装例です。"
```

# Author

[michimani210](https://twitter.com/michimani210)