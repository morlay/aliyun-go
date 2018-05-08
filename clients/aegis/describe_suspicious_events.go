package aegis

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeSuspiciousEventsRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
}

func (req *DescribeSuspiciousEventsRequest) Invoke(client *sdk.Client) (resp *DescribeSuspiciousEventsResponse, err error) {
	req.InitWithApiInfo("aegis", "2016-11-11", "DescribeSuspiciousEvents", "vipaegis", "")
	resp = &DescribeSuspiciousEventsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeSuspiciousEventsResponse struct {
	responses.BaseResponse
	RequestId        common.String
	PageSize         common.Integer
	TotalCount       common.Integer
	CurrentPage      common.Integer
	HttpStatusCode   common.Integer
	SuspiciousEvents DescribeSuspiciousEventsSuspiciousEventList
}

type DescribeSuspiciousEventsSuspiciousEventList []common.String

func (list *DescribeSuspiciousEventsSuspiciousEventList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
