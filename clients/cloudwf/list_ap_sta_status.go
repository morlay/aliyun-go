package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListApStaStatusRequest struct {
	requests.RpcRequest
	OrderCol       string `position:"Query" name:"OrderCol"`
	SearchProtocal string `position:"Query" name:"SearchProtocal"`
	SearchSsid     string `position:"Query" name:"SearchSsid"`
	SearchIp       string `position:"Query" name:"SearchIp"`
	Length         int    `position:"Query" name:"Length"`
	SearchUsername string `position:"Query" name:"SearchUsername"`
	SearchMac      string `position:"Query" name:"SearchMac"`
	PageIndex      int    `position:"Query" name:"PageIndex"`
	Id             int64  `position:"Query" name:"Id"`
	OrderDir       string `position:"Query" name:"OrderDir"`
}

func (req *ListApStaStatusRequest) Invoke(client *sdk.Client) (resp *ListApStaStatusResponse, err error) {
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListApStaStatus", "", "")
	resp = &ListApStaStatusResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListApStaStatusResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
