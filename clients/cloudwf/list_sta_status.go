package cloudwf

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListStaStatusRequest struct {
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

func (r ListStaStatusRequest) Invoke(client *sdk.Client) (response *ListStaStatusResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ListStaStatusRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("cloudwf", "2017-03-28", "ListStaStatus", "", "")

	resp := struct {
		*responses.BaseResponse
		ListStaStatusResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ListStaStatusResponse

	err = client.DoAction(&req, &resp)
	return
}

type ListStaStatusResponse struct {
	RequestId string
	Success   bool
	Message   string
	Data      string
	ErrorCode int
	ErrorMsg  string
}
