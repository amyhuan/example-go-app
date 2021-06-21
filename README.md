# example-go-app

This app was to see why a regex was failing to capture lines of juniper cli output. It turned out the bug was somewhere else in the moby canary code, because this code snippet works just fine. It captures the IP in the string given correctly

Make the binary: `go build .`
