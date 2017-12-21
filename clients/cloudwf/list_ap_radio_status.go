package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApRadioStatusRequest struct {
	requests.RpcRequest
	SearchDisabled      int    `position:"Query" name:"SearchDisabled"`
	OrderCol            string `position:"Query" name:"OrderCol"`
	SearchName          string `position:"Query" name:"SearchName"`
	SearchChannelEquals int    `position:"Query" name:"SearchChannelEquals"`
	Length              int    `position:"Query" name:"Length"`
	SearchMac           string `position:"Query" name:"SearchMac"`
	SearchApgroupName   string `position:"Query" name:"SearchApgroupName"`
	PageIndex           int    `position:"Query" name:"PageIndex"`
	OrderDir            string `position:"Query" name:"OrderDir"`
	SearchApStatus      int    `position:"Query" name:"SearchApStatus"`
}

func (req *ListApRadioStatusRequest) Invoke(client *sdk.Client) (resp *ListApRadioStatusResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApRadioStatus", "", "")
	resp = &ListApRadioStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListApRadioStatusResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
