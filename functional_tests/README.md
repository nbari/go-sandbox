http://stackoverflow.com/q/39690509/1135424

coverage binary:

    go test -c -coverpkg=. -o myProgram

./myProgram -test.coverprofile=/tmp/profile

$ go tool cover -html /tmp/profile -o /tmp/profile.html
$ open /tmp/profile.html
