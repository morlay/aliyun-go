# aliyun-go

Generated aliyun sdks for go lang 

## How to

```go
package main

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/morlay/aliyun-go/clients/ecs"
	"fmt"
)

func main() { 
    // 创建通用 Client 实例
    client, err := sdk.NewClientWithAccessKey(
        "<your-region-id>", 			// 您的可用区ID
        "<your-access-key-id>", 		// 您的 Access Key ID
        "<your-access-key-secret>")		// 您的 Access Key Secret
    if err != nil {
    	// 异常处理
    	panic(err)
    }
    // 创建 API 请求并设置参数
    req := ecs.DescribeInstancesRequest{
    	IoOptimized: "true",
    }
    // 请求接口
    response, err := req.Invoke(client)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
}
```

### Custom Build

``` sh
git clone ...

# initial
git submodule update --init
gradle 

# update
git submodule update --remote
gradle
```

## Issues

* batchcompute 和别的产品代码结构不同， 不做处理了