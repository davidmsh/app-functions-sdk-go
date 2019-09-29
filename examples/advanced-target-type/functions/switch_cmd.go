//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package functions

import (
	"context"

	"github.com/edgexfoundry/app-functions-sdk-go/appcontext"
)

const (
	deviceId  = "7cccb1a0-8e26-4feb-95c7-423e49559017"
	commandId = "6663264d-c62e-49c6-bec3-96964570a4f9"

	jsonSwitchOn  = "{\"SwitchButton\": \"true\"}"
	jsonSwitchOff = "{\"SwitchButton\": \"false\"}"
)

type Switch struct {
	Status string `json:"status"`
}

func SendSwitchCommand(edgexcontext *appcontext.Context, params ...interface{}) (bool, interface{}) {
	edgexcontext.LoggingClient.Debug("Format Phone Number 2")

	if len(params) < 1 {
		// We didn't receive a result
		return false, nil
	}

	if edgexcontext.CommandClient == nil {
		edgexcontext.LoggingClient.Error("Command client is available")
		return false, nil
	}

	sw, ok := params[0].(Switch)

	if !ok {
		edgexcontext.LoggingClient.Error("Invalid switch")
		return false, nil
	}

	var cmd string

	ctx := context.WithValue(context.Background(), "", "")

	switch status := sw.Status; status {
	case "on":
		cmd = jsonSwitchOn
	case "off":
		cmd = jsonSwitchOff
	default:
		edgexcontext.LoggingClient.Error("Invalid switch status: " + status)
		return false, nil
	}

	r, err := edgexcontext.CommandClient.Put(deviceId, commandId, cmd, ctx)

	if err == nil {
		edgexcontext.LoggingClient.Debug("Response : " + r)
	} else {
		edgexcontext.LoggingClient.Error("Error sending request " + err.Error())
	}

	return true, cmd
}
