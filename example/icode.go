/*
 * Copyright 2018 DE-labtory
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"os"

	"github.com/DE-labtory/sdk"
	"github.com/DE-labtory/sdk/example/handler"
	"github.com/DE-labtory/sdk/logger"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Port int `short:"p" long:"port" description:"set port"`
}

func main() {
	logger.EnableFileLogger(true, "./icode.log")
	parser := flags.NewParser(&opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		logger.Error(nil, "fail parse args: "+err.Error())
		os.Exit(1)
	}

	exHandler := &handler.HandlerExample{}
	ibox := sdk.NewIBox(opts.Port)
	ibox.SetHandler(exHandler)
	ibox.On(30)
}
