# Super Basic Proxy

## Configuration

```json
{
  "port": 8080,
  "routes": [
    {
      "path_prefix": "/api",
      "target": "http://localhost:3000"
    },
    {
      "path_prefix": "/static",
      "target": "http://localhost:8081"
    }
  ],
  "headers": {
    "add": {
      "X-Proxy-By": "sb-proxy"
    },
    "remove": ["X-Powered-By"]
  },
  "timeout_ms": 30,
}
```
