Check the HTTP headers

    $ curl -sI 0:8000
    HTTP/1.1 200 OK
    Connection: keep-alive
    Content-Length: 12
    Content-Type: text/plain
    Server: gophr
    Date: Sun, 04 Jun 2017 08:53:38 GMT

    $ curl -sI 0:8000 | wc  -c
    141
