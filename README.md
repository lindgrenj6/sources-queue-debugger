# sources-queue-debugger

Much like https://github.com/lindgrenj6/availability-dummy - this is used for debugging sources-api running in k8s. I usually run it from minikube but it can be ran from anywhere that runs clowder or bonfire.

### To setup in your local bonfire config
Add this snippet to your local configuration:
```yaml
- name: sources-queue-debugger
  host: github
  repo: lindgrenj6/sources-queue-debugger
  path: deploy/clowdapp.yaml
```

From there there should be:
- `sources-queue-debugger-event-stream`
- `sources-queue-debugger-status`
- `sources-queue-debugger-superkey-requests`
- `sources-queue-debugger-notifications`
- `sources-queue-debugger-satellite-operations`

deployments running pods!

### Tailing the pod logs
If you tail any of those pod logs you'll get info pertaining to the topic they're set up for such as `platform.sources.event-stream` topic, but in JSON format, which I usually pipe to jq.
The output looks like this:
```json
{
  "topic": "platform.sources.event-stream",
  "headers": [
    {
      "key": "x-rh-sources-account-number",
      "value": "1234"
    },
    {
      "key": "x-rh-sources-org-id",
      "value": "1234"
    },
    {
      "key": "x-rh-identity",
      "value": "aaaaa"
    },
    {
      "key": "event_type",
      "value": "ApplicationAuthentication.create"
    },
    {
      "key": "encoding",
      "value": "json"
    }
  ],
  "body": {
    "application_id": 62,
    "authentication_id": 93,
    "created_at": "2022-08-15 20:26:29 UTC",
    "id": 62,
    "paused_at": null,
    "tenant": "",
    "updated_at": "2022-08-15 20:26:29 UTC"
  }
}
```

It's really useful to be able to see every event that is being produced!

### HTTP Interface
Each pod also has a http server running on port `8000` which serves info about the current session via `/info`, so if you run a `GET :8000/info` request you'll get a nice count of events on the queue, output is like this:
```json
{
    "Application.create": 4,
    "ApplicationAuthentication.create": 4,
    "Authentication.create": 4,
    "Source.create": 4
}
```
which signifies there were 4 events of each produced since start.

To reset the counters, simply call `/clear` and they will be reset to 0 (or delete the pod).
