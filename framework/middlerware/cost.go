package middlerware

import (
	"fmt"
	"github.com/hktime/web-server-demo/framework"
	"time"
)

func Cost()framework.ControllerHandler{
	return func(c *framework.Context) error {
		start := time.Now()
		c.Next()
		fmt.Printf("api uri: %v, time cost %vs\n", c.GetRequest().RequestURI, time.Now().Sub(start).Seconds())
		return nil
	}
}
