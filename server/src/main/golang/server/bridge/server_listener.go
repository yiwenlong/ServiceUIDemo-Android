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
package bridge

import "github.com/yiwenlong/server"

type MyServerListener interface {
	server.MyServerListener
}
