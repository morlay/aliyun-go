package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListBriefApConfigRequest struct {
	requests.RpcRequest
	SearchScan  int    `position:"Query" name:"SearchScan"`
	OrderCol    string `position:"Query" name:"OrderCol"`
	SearchName  string `position:"Query" name:"SearchName"`
	Length      int    `position:"Query" name:"Length"`
	SearchMac   string `position:"Query" name:"SearchMac"`
	PageIndex   int    `position:"Query" name:"PageIndex"`
	OrderDir    string `position:"Query" name:"OrderDir"`
	SearchModel string `position:"Query" name:"SearchModel"`
}

func (req *ListBriefApConfigRequest) Invoke(client *sdk.Client) (resp *ListBriefApConfigResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListBriefApConfig", "", "")
	resp = &ListBriefApConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListBriefApConfigResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
