
### Keybase Webhook Notification Handler

This tool accepts Alertmanager webhook notification HTTP POST requests and forwards them to Keybase.


 #### Alertmanager Configuration
```yaml
global:
  resolve_timeout: 5m
receivers:
- name: noop-receiver
- name: keybase
  webhook_configs:
    - url: http://keybase-alertmanager.default:3000/webhook
      send_resolved: true
route:
  receiver: keybase
  group_by:
  - job
  routes:
  - receiver: watchdog
    match:
      alertname: Watchdog
    group_wait: 15s
    group_interval: 30s
    repeat_interval: 1m
  repeat_interval: 1h
receivers:
- name: noop-receiver
- name: keybase
  webhook_configs:
  - send_resolved: true
    url: http://keybase-alertmanager.default:3000/webhook
- name: watchdog
  webhook_configs:
  - url: http://keybase-alertmanager.default:3000/watchdog
  ```