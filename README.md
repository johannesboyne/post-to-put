#POST to PUT

Transform a POST request to a PUT request.

##Usage

```bash
$ post-to-put -in 8080 -out http://127.0.0.1:8500
...
$ curl -X POST -d 'orange' http://127.0.0.1:8080/v1/kv/web/dronestatus 
```

The `POST` request is going to be transformed to a `PUT` request. Like in the example to set/update a Key-Value pair in consul.


##License

MIT
