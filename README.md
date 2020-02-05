## Deploy

```bash
gcloud functions deploy Pinger --runtime go113 --trigger-http --timeout 10
```

## Test
```bash
gcloud functions describe Pinger
```

## Running a local test UDP server
```bash
cd server
go run main.go
```

## Testing the UDP server
```bash
echo "spacemesh" > /dev/udp/127.0.0.1/7555
```