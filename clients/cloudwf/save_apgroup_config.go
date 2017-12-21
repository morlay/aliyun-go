package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveApgroupConfigRequest struct {
	requests.RpcRequest
	Country     string `position:"Query" name:"Country"`
	LogLevel    int    `position:"Query" name:"LogLevel"`
	Name        string `position:"Query" name:"Name"`
	EchoInt     int    `position:"Query" name:"EchoInt"`
	Scan        int    `position:"Query" name:"Scan"`
	Description string `position:"Query" name:"Description"`
	Id          int64  `position:"Query" name:"Id"`
	Dai         string `position:"Query" name:"Dai"`
	LogIp       string `position:"Query" name:"LogIp"`
}

func (req *SaveApgroupConfigRequest) Invoke(client *sdk.Client) (resp *SaveApgroupConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "SaveApgroupConfig", "", "")
	resp = &SaveApgroupConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveApgroupConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
