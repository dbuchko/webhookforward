# webhookforward
simple app to show how you could forward a webrequest to another url
<p>
Cross compile for linux diego cells using:
  <p>
env GOOS=linux GOARCH=amd64 go build
    <p>
Push to PCF using binary build pack with:
      <p>
cf push whproxy -c'./proxy' -b binary_buildpack -k 64m -m 64m

Set the address to foward to:

cf set-env whproxy FORWARD_URL https://myslackwebhookurl
