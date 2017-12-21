package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ProfileTagRequest struct {
	requests.RpcRequest
	Idtype     int64  `position:"Query" name:"Idtype"`
	BeginDate  string `position:"Query" name:"BeginDate"`
	EndDate    string `position:"Query" name:"EndDate"`
	AppType    int    `position:"Query" name:"AppType"`
	Tag        string `position:"Query" name:"Tag"`
	Agsid      int64  `position:"Query" name:"Agsid"`
	AreaNumber int    `position:"Query" name:"AreaNumber"`
}

func (req *ProfileTagRequest) Invoke(client *sdk.Client) (resp *ProfileTagResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ProfileTag", "", "")
	resp = &ProfileTagResponse{}
	err = client.DoAction(req, resp)
	return
}

type ProfileTagResponse struct {
	responses.BaseResponse
	Success   bool
	Data      string
	ErrorCode int
	ErrorMsg  string
}
