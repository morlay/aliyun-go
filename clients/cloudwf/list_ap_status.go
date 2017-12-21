package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApStatusRequest struct {
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

func (r ListApStatusRequest) Invoke(client *sdk.Client) (response *ListApStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListApStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		ListApStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListApStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListApStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
