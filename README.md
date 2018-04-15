# HTTP-ICMP

### Build 
1. `$ cd $GOPATH/src/github.com/timdurward/http-icmp`
2. `$ make build`
3. `$ http-icmp`

### How to use:
POST:
```bash
curl -s -XPOST -d'{"hostname":"google.com", "count": 3}' http://localhost:8000/ping
``` 

Response:
```json
{
    "address": "google.com",
    "ip_address": {
        "IP": "216.58.217.46",
        "Zone": ""
    },
    "results": {
        "packets": {
            "sent": 3,
            "received": 3,
            "lost": 0
        },
        "statistics": {
            "minimum": 14989978,
            "maximum": 18361581,
            "average": 16327381
        }
    }
}
``` 

