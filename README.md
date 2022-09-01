To run this example, ensure you have the `AWS sam` cli client installed and Docker. This was confirmed working with SAM `1.55.0` and Docker `20.10.14`.

```bash
sam build
sam local invoke -e events/event.json
```

If you have AWS credentials loaded into your environment, you can also push to your AWS account with:

```bash
sam deploy --guided
``

This is an example project for a blog post on dave.dev and video on YouTube. This was generated also for the AWS Community Builders programme by David Gee (dave.dev / built.fm) in spare time.
