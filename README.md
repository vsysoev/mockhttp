# mockhttp
Simple http mock server for testing purpose. Need to mock backend for simulating
particular behavior.
Behavior is defined in yml config file. Sample is
```yaml
/hello:                           // entrypoint
  status: 200                     // status code
  delay: 200ms                    // delay in time.Duration
  body: '{"status":"OK"}'         // body as string
/world:
  status: 404
  delay: 1s
```
