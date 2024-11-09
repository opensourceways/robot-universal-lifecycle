// Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"flag"
	"github.com/opensourceways/robot-framework-lib/config"
	"github.com/opensourceways/robot-framework-lib/framework"
	"github.com/opensourceways/server-common-lib/logrusutil"
	"os"
)

const component = "robot-universal-lifecycle"

func main() {
	logrusutil.ComponentInit(component)

	opt := new(robotOptions)
	cnf, token := opt.gatherOptions(flag.NewFlagSet(os.Args[0], flag.ExitOnError), os.Args[1:]...)
	if opt.shutdown {
		return
	}

	bot := newRobot(cnf, token)
	framework.StartupServer(
		framework.NewServer(bot, opt.service, config.ServerAdditionOptions{HandlePath: "/" + opt.handlePath}),
		opt.service,
		config.ServerAdditionOptions{},
	)
}
