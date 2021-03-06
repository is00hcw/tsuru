// Copyright 2015 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hipache_test

import (
	"fmt"
	"net/url"

	"github.com/tsuru/tsuru/router"
	_ "github.com/tsuru/tsuru/router/hipache"
)

func Example() {
	router, err := router.Get("hipache")
	if err != nil {
		panic(err)
	}
	err = router.AddBackend("myapp")
	if err != nil {
		panic(err)
	}
	url, err := url.Parse("http://10.10.10.10:8080")
	if err != nil {
		panic(err)
	}
	err = router.AddRoute("myapp", url)
	if err != nil {
		panic(err)
	}
	addr, _ := router.Addr("myapp")
	fmt.Println("Please access:", addr)
	err = router.RemoveRoute("myapp", url)
	if err != nil {
		panic(err)
	}
	err = router.RemoveBackend("myapp")
	if err != nil {
		panic(err)
	}
}
