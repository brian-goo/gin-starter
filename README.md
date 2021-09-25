## golang gin starter

```bash
# start a local dev server
$ npm install -g nodemon
$ make dev
```

### middleware
- cors   
- api key authentication     
set TOKEN in makefile   
- auth0 authentication   
set AUTH0_ISS and AUTH0_AUD in makefile   
the public key from jwks.json will be cached in memory for 12 hours   

### log rotation

### graceful restart / shutdown
forked https://github.com/fvbock/endless

### http "HEAD" method will be handled with "GET" method
i.e. no need to implement HEAD method in a router level