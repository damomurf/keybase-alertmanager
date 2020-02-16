
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
  group_by:
  - job
  group_interval: 5m
  group_wait: 30s
  receiver: keybase
  repeat_interval: 12h
  routes:
  - match:
      alertname: Watchdog
    receiver: noop-receiver
```