# openapi
梵响开放接口

# 示例

```go
package main

import (
	"fmt"

	"github.com/funxdata/openapi"
	"github.com/funxdata/openapi/geoip"
)

func main() {
	cli := openapi.NewClient("account-key", "account-secret")

	geoinfo, err := cli.GeoIP().GetGeo(cli.Context, &geoip.IP{Ip: "103.46.244.190"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(geoinfo)
	fmt.Println(geoinfo.City, geoinfo.Isp)
}

```
