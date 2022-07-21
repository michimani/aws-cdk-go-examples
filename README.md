aws-cdk-go-examples
===

This is a repository of example implementations of using AWS CDK with the Go language.

---

AWS CDK で Go 言語を使った場合の実装例を置いているリポジトリです。

## cloudformation-events-to-slack

This is an example implementation of a configuration in which CloudFormation stack creation, modification, and deletion events are notified via email by Amazon SNS.

Please read [cloudformation-events-to-slack/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/cloudformation-events-to-slack).

--- 

CloudFormation の Stack の作成・変更・削除のイベントを、 Amazon SNS を通じてメールで通知する構成の実装例です。

詳細は [cloudformation-events-to-slack/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/cloudformation-events-to-slack) を読んでください。

## scheduled-lambda-function

This is an example implementation of a Lambda function that is executed periodically according to a schedule specified by a CRON expression.

Please read [scheduled-lambda-function/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/scheduled-lambda-function).

---

Lambda 関数を EventBridge で設定した CRON 式によって定期実行する構成の実装例です。

詳細は [scheduled-lambda-function/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/scheduled-lambda-function) を読んでください。

## step-functions-with-sdk-integration

This is an example implementation of building a StepFunctions state machine that translates text using some AWS SDK integration.

Please read [step-functions-with-sdk-integration/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/step-functions-with-sdk-integration).

---

StepFunctions でいくつかの AWS SDK インテグレーションを利用したステートマシンを構築する実装例です。

詳細は [step-functions-with-sdk-integration/README.md](https://github.com/michimani/aws-cdk-go-examples/blob/main/step-functions-with-sdk-integration) を読んでください。

# License

[MIT](https://github.com/michimani/aws-cdk-go-examples/blob/main/LICENSE)

# Author

[michimani210](https://twitter.com/michimani210)