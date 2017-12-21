package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApStatusRequest struct {
	requests.RpcRequest
	OrderCol          string `position:"Query" name:"OrderCol"`
	SearchName        string `position:"Query" name:"SearchName"`
	SearchGroupName   string `position:"Query" name:"SearchGroupName"`
	SearchStatus      int    `position:"Query" name:"SearchStatus"`
	SearchWanIp       string `position:"Query" name:"SearchWanIp"`
	SearchApModelName string `position:"Query" name:"SearchApModelName"`
	Length            int    `position:"Query" name:"Length"`
	OrderDir          string `position:"Query" name:"OrderDir"`
	SearchBssEquals   int    `position:"Query" name:"SearchBssEquals"`
	SearchSwVersion   int64  `position:"Query" name:"SearchSwVersion"`
	SearchCompanyName string `position:"Query" name:"SearchCompanyName"`
	SearchMac         string `position:"Query" name:"SearchMac"`
	PageIndex         int    `position:"Query" name:"PageIndex"`
}

func (req *ListApStatusRequest) Invoke(client *sdk.Client) (resp *ListApStatusResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApStatus", "", "")
	resp = &ListApStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListApStatusResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
