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
	"github.com/yiwenlong/server"
	"log"
	"time"
)

type Listener struct {
}

func (l *Listener) OnServerStart() {
	fmt.Printf(">>>> start <<<<\n")
}

func (l *Listener) OnServerStop() {
	fmt.Printf(">>>> stop <<<<\n")
}

func (l *Listener) OnServerError(msg string) {
	fmt.Printf(">>>> error <<<<\n")
}

func main() {
	myServer := server.NewMyServer("localhost", 8080)
	listener := Listener{}
	myServer.RegisterListener(&listener)
	myServer.Boot()
	time.Sleep(time.Second)
	log.Printf("Server is running: %v\n", myServer.IsRunning())
	time.Sleep(5 * time.Second)
	log.Printf("Server is running: %v\n", myServer.IsRunning())
	myServer.Stop()
	time.Sleep(time.Second)
}
