# Watchdog

Kontroluje nakonfigurovane zavislosti a na zaklade testu, upravuje stav prob.

## Run
```bash
./watchdog -cfg example.yaml
```


## REST Endpointy
```yaml
GET: /_healthz/live
GET: /_healthz/ready
```



## Konfiguracni jazyk
```yaml
port: 8080
checkIntervalSec: 10
initialStatus:
  live: false
  ready: false
rules:
- name: ping-root-port
  errorMsg: root.cz does not respond on port 80
  address: 91.213.160.188
  port: 80
  method: ping
  required:
  - ready
  - live
- name: broken-tcp-ping-port
  errorMsg: root.cz does not respond on port 8880
  address: 91.213.160.188
  port: 8880
  method: ping
  required:
  - ready
  - live
- name: google-dns
  errorMsg: Network is unreachable
  address: 8.8.8.8
  method: resolve
  required:
  - ready
- name: seznam-http-ping
  errorMsg: Check HTTPS on www.seznam.cz
  address: https://www.seznam.cz/
  method: http-get
  required:
  - ready
- name: broken-http-ping
  errorMsg: Check HTTPS on www.seznam.cz
  address: https://www.seznamsss.cz/
  method: http-get
  required:
  - ready%
```


## TODO

 - validator yamlu (vubec nic ted nekontroluju)
 - doplneni DNS kontroly
 - sjednoceni logovani
 - doplnit testy
 - doplnit k8s example manifest
 - Zlepsit dokumentaci
