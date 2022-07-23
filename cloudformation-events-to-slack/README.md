cloudformation-events-to-slack
===

This is an example implementation of a configuration in which specified CloudFormation Stack events (changes of resource status, stack status and drift detection status) are notified via email by Amazon SNS.

# Resources

## SNS::Topic

## SNS::TopicPolicy

- Action: `sns:Publish`
- Effect: `Allow`
- Principal: `events.amazonaws.com`

## SNS::Subscription

- Protocol: email

## Events::Rule

- EventPattern
  - Source: `aws.cloudformation`
  - Detail-Type:
    - `CloudFormation Resource Status Change`
    - `CloudFormation Stack Status Change`
    - `CloudFormation Drift Detection Status Change`
  - Resources:
    - has prefix `NotificationTest`

# Usage

## Set email for subscribe for environment variable

```bash
export EMAIL_FOR_SUBSCRIBE='your-email@example.com'
```

## Deploy SNS and Events Stack

```bash
cdk synth CloudformationEventsToSlackStack
```

```bash
cdk deploy CloudformationEventsToSlackStack
```

Then, you will receive a confirmation email to the email address you set up.

## Deploy test stack for receiving notification (create a S3 Bucket)

1. Set bucket name for environment variable.

    ```bash
    export TMP_BUCKET_NAME=='your-bucket-name-for-notification-test'
    ```

2. Deploy test stack. (will notify)

    ```bash
    cdk synth NotificationTestStack
    ```

    ```bash
    cdk deploy NotificationTestStack
    ```

3. Deploy test stack. (will NOT notify)

    ```bash
    cdk synth NotNotificationTestStack
    ```

    ```bash
    cdk deploy NotNotificationTestStack
    ```

# Author

[michimani210](https://twitter.com/michimani210)