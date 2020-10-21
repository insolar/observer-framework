// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/observer-framework/blob/master/LICENSE.md.

package configuration

type Observer struct {
	Profefe     Profefe
}

type Profefe struct {
	StartAgent bool   `insconfig:"true| if true, start the profefe agent"`
	Address    string `insconfig:"http://127.0.0.1:10100| Profefe collector public address to send profiling data"`
	Labels     string `insconfig:"host,localhost| Application labels. For example, region,europe-west3,dc,fra"`
}
