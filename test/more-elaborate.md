# This is a Markdown document.
## It is more elaborate than some

## Is the backend written in go?

Yes, *of course* it is written in go.

## Does it run on containers?

Yes, *of course* it is containerized.

## Ok.

The backend that produced this document is a multi-tiered web
architecture compromised of:

  - [DynamoDB][dynamo] as a datastore
  - [EC2 Container Service][ecs] to run containers

The containers currently running are dockerized Go applications:

  - [dynago-docs][dynago-docs] as a content-agnostic documenter server
  - [renderdog][renderdog] as as front-end html rendering agent

## What's Next?
There are many steps left:

 0. The above projects are hackjob prototypes; make them more robust.
 1. Integrate [golang-sso-google][sso] for authentication
 2. Write the [Pumpkin][pumpkin] front-end

## Blah blah

    content, err := module.PublicFunction(input)
    if err != nil {
        panic(err)
    }

[dynamo]: https://aws.amazon.com/dynamodb
[ecs]: https://aws.amazon.com/ecs
[dynago-docs]: https://github.com/jamesonwilliams/dynago-docs
[renderdog]: https://github.com/jamesonwilliams/renderdog
[sso]: https://github.com/jamesonwilliams/golang-sso-google
[pumpkin]: https://github.com/jamesonwilliams/pumpkin

