# ether-subscription-app

A simple app written in `Go` for consuming messages from the `newTxs` stream from `bloXroute`

## Config/Setup

The following configs are in the `resources` directory and need to be updated:

### bloxroute_config.json

| Field                           | Description                                                                                                                                         |
|---------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------|
| `websocketsCloudApiBaseUri`     | bloXroute websocket URL for the cloud-api. See https://docs.bloxroute.com/streams/working-with-streams/creating-a-subscription#websocket-connection |
| `authorizationHeader`           | This can be found in the `Account` section (after logging in) at https://portal.bloxroute.com/                                                      |

### main_config.json

The following configs create the websocket subscription filter. These can be combined or set independently. For example:
* only the `toAddressFilter`, and you will only receive messages for transactions sent _to_ that address
* only the `fromAddressFilter`, and you will only receive messages for transactions sent _from_ that address
* both the `toAddressFilter` and `fromAddressFilter`, and you will only receive messages for transactions sent from the specified address that are sent to the specified address

| Field                             | Description                                   |
|-----------------------------------|-----------------------------------------------|
| `subscriptionFilters.toAddress`   | Filter for address transactions are sent to   |
| `subscriptionFilters.fromAddress` | Filter for address transactions are sent from |

## Running the App

### System with Go Installed
If you have `go` setup on your machine, you can simply run the following:

```shell
go build -o main && ./main
```

### Via Docker
1. Build the image
    ```shell
   docker build -t ether-subscription-app  .
   ```
2. Run the docker image in a container
    ```shell
   docker run ether-subscription-app
    ```