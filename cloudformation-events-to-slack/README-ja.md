cloudformation-events-to-slack
===

CloudFormation Stack に関連するイベント (リソースのステータス変更、スタックのステータス変更、ドリフト検出) を、 Amazon SNS を通じてメールで通知する構成の実装例です。

# 主なリソース

## SNS::Topic

## SNS::Subscription

- Protocol: email

## Events::Rule

- EventPattern
  - Source: `aws.cloudformation`
  - Detail-Type:
    - `CloudFormation Resource Status Change`
    - `CloudFormation Stack Status Change`
    - `CloudFormation Drift Detection Status Change`

# 使い方

## 通知先のメールアドレスを環境変数に設定

```bash
export EMAIL_FOR_SUBSCRIBE='your-email@example.com'
```

## SNS::Topic と Events::Rule を構築する Stack をデプロイ

```bash
cdk synth CloudformationEventsToSlackStack
```

```bash
cdk deploy CloudformationEventsToSlackStack
```

Then, you will receive a confirmation email to the email address you set up.

## 通知のテストのための Stack をデプロイ (S3 バケットを作成)

1. 作成する S3 Bucket のバケット名を環境変数に設定

    ```bash
    export TMP_BUCKET_NAME=='your-bucket-name-for-notification-test'
    ```

2. S3 Bucket 作成のための Stack をデプロイ (通知が届きます)

    ```bash
    cdk synth NotificationTestStack
    ```

    ```bash
    cdk deploy NotificationTestStack
    ```

# Author

[michimani210](https://twitter.com/michimani210)