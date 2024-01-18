# Cron Request

Cron request is written in Go that will send http requests on a schedule.

## Usage

1. Create a `config.yaml` file in the same directory as the binary with the following format:

```yaml
metadata:
  name: my-app
requests:
  - url: https://testapi.metopia.co/nft/mint-status
    method: GET
    name: mint-status
execution:
  interval: 10
```