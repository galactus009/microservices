* Renders a SVG image from the JSON Payload.
  This is done after attempting same in python which is taking 1 min. golang version takes 41 ms

* Testing
```curl -XPOST -H "Content-Type: application/json" localhost:8080/  --data"@testpayload.json"```
