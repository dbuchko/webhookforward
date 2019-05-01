# webhookforward
simple app to show how you could forward a webrequest to another url
<p>
Cross compile for linux diego cells using:
  <p>
env GOOS=linux GOARCH=amd64 go build
    <p>
Push to PCF using binary build pack with:
      <p>
cf push whproxy -c'./webhookforward' -b binary_buildpack -k 64m -m 64m

Set the address to foward to:

cf set-env whproxy FORWARD_URL https://myslackwebhookurl

<p>
  Can test by curling the app like:
  <p>
  curl --data-binary '{\"Test Message\":\"Test Message Body\"}'  -H "Content-Type: application/json" --request POST https://whproxy.<Apps Route>
<p>
or for a more complex sample:
curl --data-binary "@sampleMessage"\
  -H "Content-Type: application/json" --request POST https://wh\
proxy.<Apps Route>
