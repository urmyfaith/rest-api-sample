package main

import (
    "github.com/simpleton/tinker-api/routers"
    "github.com/facebookgo/grace/gracehttp"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/labstack/echo/middleware"
***REMOVED***

func main(***REMOVED*** {
    server := echo.New(***REMOVED***
    server.Use(middleware.Logger(***REMOVED******REMOVED***
    server.Use(middleware.Recover(***REMOVED******REMOVED***

    //Run the API
    api := routers.NewAPIRouter(server***REMOVED***
    api.Init(***REMOVED***

    std := standard.New(":8300"***REMOVED***
    std.SetHandler(server***REMOVED***
    gracehttp.Serve(std.Server***REMOVED***
}
