//
// Copyright 2020 Yiwenlong(wlong.yi#gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may obtain a copy of the License at
// you may not use this file except in compliance with the License.
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package main

import (
	"fmt"
	"github.com/yiwenlong/server/bridge"
	"time"
)

type ServerHandler struct {
}

func (lh *ServerHandler) Log(message string) {
	fmt.Printf("MyLog ==>  %s", message)
}

func (lh *ServerHandler) OnServerStart() {
	fmt.Printf("MyLog ==> OnServerStart\n")
}

func (lh *ServerHandler) OnServerStop() {
	fmt.Printf("MyLog ==> OnServerStop\n")
}

func (lh *ServerHandler) OnServerError(msg string) {
	fmt.Printf("MyLog ==> OnServerStop\n")
}

func main() {
	bridge.RegisterLogHandler(&ServerHandler{})
	bridge.InitServer("localhost", 8080)
	bridge.RegisterServerListener(&ServerHandler{})
	bridge.BootServer()
	time.Sleep(1 * time.Minute)
	bridge.StopServer()
}
