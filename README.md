# corsbypass
golang server at backend to bypass CORS

## Test:

/> corsnypass


/> curl http://localhost:94/anything
```
{
  "args": {},
  "data": "",
  "files": {},
  "form": {},
  "headers": {
    "Accept": "*/*",
    "Host": "httpbin.org",
    "User-Agent": "curl/7.55.1",
    "X-Amzn-Trace-Id": "Root=1-6203ef94-15e9d588254000896db5e80f"
  },
  "json": null,
  "method": "GET",
  "origin": "4.1.2.3",
  "url": "http://httpbin.org/anything"
}
```