# Auto pull for git web hook

## Build

``` shell
go get github.com/gafeng/autopullgo
go install github.com/gafeng/autopullgo
```

## Run

Write your config file in /etc/autopull/conf.toml:

``` toml
xxx = "github.com/xxx/xxx"
```

Then run:

``` shell
autopullgo
```

The program will listening on port 8920.
And the trigger url is:`your_ip:8920/trigger-git`

When catch a http request such as : `http://your_ip:8920/trigger-git?project=xxx`, the program will run `go get -u github.com/xxx/xxx` automatically.

