aws-cdk-go-examples
===

AWS CDK で Go 言語を使った場合の実装例を置いているリポジトリです。

# Examples

## cloudformation-events-to-slack

特定の CloudFormation Stack に関連するイベントを、 Amazon SNS を通じてメールで通知する構成の実装例です。

詳細は [cloudformation-events-to-slack/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/cloudformation-events-to-slack) を読んでください。

## eventbridge-scheduler-cfn

EventBridge Scheduler で設定したスケジュールによって Lambda 関数を定期実行する構成の実装例です。 Scheduler のリソースは `CfnResource` を使って定義します。

詳細は [eventbridge-scheduler-cfn/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/eventbridge-scheduler-cfn) を読んでください。

## scheduled-lambda-function

Lambda 関数を EventBridge で設定した CRON 式によって定期実行する構成の実装例です。

詳細は [scheduled-lambda-function/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/scheduled-lambda-function) を読んでください。

## sqs-to-lambda-with-dlq

デッドレターキューを設定した SQS Queue へのメッセージ送信をトリガーに起動する Lambda 関数と、 デッドレターキューを処理する別の Lambda 関数を構成する実装例です。

詳細は [sqs-to-lambda-with-dlq/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/sqs-to-lambda-with-dlq) を読んでください。

## step-functions-with-sdk-integration

StepFunctions でいくつかの AWS SDK インテグレーションを利用したステートマシンを構築する実装例です。

詳細は [step-functions-with-sdk-integration/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/step-functions-with-sdk-integration) を読んでください。

# Memo

[AWS CDK を Go で書いたときのメモ - zenn](https://zenn.dev/michimani/scraps/3fb7f8675ef22e)

# License

[MIT](https://github.com/michimani/aws-cdk-go-examples/blob/main/LICENSE)

# Author

[michimani210](https://twitter.com/michimani210)