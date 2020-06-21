// Copyright 2020 astaxie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/metric"
)

func main() {
	// we start admin service
	// Prometheus will fetch metrics data from admin service's port
	beego.BConfig.Listen.EnableAdmin = true

	beego.BConfig.AppName = "my app"

	ctrl := &MainController{}
	beego.Router("/hello", ctrl, "get:Hello")
	beego.RunWithMiddleWares(":8080", metric.PrometheusMiddleWare)
	// after you start the server
	// and GET http://localhost:8080/hello
	// access http://localhost:8088/metrics
	// you can see something looks like:
	// http_request_beego_sum{appname="my app",duration="1002",env="prod",method="GET",pattern="/hello",server="beegoServer:1.12.1",status="200"} 1002
	// http_request_beego_count{appname="my app",duration="1002",env="prod",method="GET",pattern="/hello",server="beegoServer:1.12.1",status="200"} 1
	// http_request_beego_sum{appname="my app",duration="1004",env="prod",method="GET",pattern="/hello",server="beegoServer:1.12.1",status="200"} 1004
	// http_request_beego_count{appname="my app",duration="1004",env="prod",method="GET",pattern="/hello",server="beegoServer:1.12.1",status="200"} 1
}

type MainController struct {
	beego.Controller
}

func (ctrl *MainController) Hello() {
	time.Sleep(time.Second)
	ctrl.Ctx.ResponseWriter.Write([]byte("Hello, world"))
}
