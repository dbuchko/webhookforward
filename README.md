# webhookforward
simple app to show how you could forward a webrequest to another url
Cross compile for linux diego cells using:
env GOOS=linux GOARCH=amd64 go build
Push to PCF using binary build pack like:
cf push whproxy -c'./proxy' -b binary_buildpack -k 64m -m 64m