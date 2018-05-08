package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListStaStatusRequest struct {
	requests.RpcRequest
	OrderCol          string `position:"Query" name:"OrderCol"`
	SearchGroupName   string `position:"Query" name:"SearchGroupName"`
	SearchStatus      int    `position:"Query" name:"SearchStatus"`
	Length            int    `position:"Query" name:"Length"`
	SearchUsername    string `position:"Query" name:"SearchUsername"`
	OrderDir          string `position:"Query" name:"OrderDir"`
	SearchProtocal    string `position:"Query" name:"SearchProtocal"`
	SearchSsid        string `position:"Query" name:"SearchSsid"`
	SearchApName      string `position:"Query" name:"SearchApName"`
	SearchIp          string `position:"Query" name:"SearchIp"`
	PageIndex         int    `position:"Query" name:"PageIndex"`
	SearchMac         string `position:"Query" name:"SearchMac"`
	SearchDescription string `position:"Query" name:"SearchDescription"`
}

func (req *ListStaStatusRequest) Invoke(client *sdk.Client) (resp *ListStaStatusResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListStaStatus", "", "")
	resp = &ListStaStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListStaStatusResponse struct {
	responses.BaseResponse
	RequestId common.String
	Success   bool
	Message   common.String
	Data      common.String
	ErrorCode common.Integer
	ErrorMsg  common.String
}
