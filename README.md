# dotenv-to-secret

Convert .env to Kubernetes Secret Resource

---

# Usage

```shell
$ go get -u github.com/mrtc0/cmd/dotenv-to-secret
$ cat .env
USERNAME=mrtc0
PASSWORD=password
$ dotenv-to-secret
---
apiVersion: v1
kind: Secret
metadata:
  name: env
type: Opaque
data:
  PASSWORD: password
  USERNAME: mrtc0
```