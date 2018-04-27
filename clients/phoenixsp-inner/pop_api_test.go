package phoenixsp_inner

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type PopApiTestRequest struct {
	requests.RpcRequest
	UserName string `position:"Query" name:"UserName"`
	Pwd      string `position:"Query" name:"Pwd"`
}

func (req *PopApiTestRequest) Invoke(client *sdk.Client) (resp *PopApiTestResponse, err error) {
	req.InitWithApiInfo("Phoenixsp-inner", "2016-08-05", "PopApiTest", "", "")
	resp = &PopApiTestResponse{}
	err = client.DoAction(req, resp)
	return
}

type PopApiTestResponse struct {
	responses.BaseResponse
	Code      string
	Data      string
	RequestId string
	Success   bool
	Message   string
}
