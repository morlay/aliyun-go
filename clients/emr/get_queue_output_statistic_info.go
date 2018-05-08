package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetQueueOutputStatisticInfoRequest struct {
	requests.RpcRequest
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
}

func (req *GetQueueOutputStatisticInfoRequest) Invoke(client *sdk.Client) (resp *GetQueueOutputStatisticInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetQueueOutputStatisticInfo", "", "")
	resp = &GetQueueOutputStatisticInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetQueueOutputStatisticInfoResponse struct {
	responses.BaseResponse
	RequestId       common.String
	QueueOutputList GetQueueOutputStatisticInfoClusterStatQueueOutputList
}

type GetQueueOutputStatisticInfoClusterStatQueueOutput struct {
	Queue       common.String
	BytesOutput common.Long
}

type GetQueueOutputStatisticInfoClusterStatQueueOutputList []GetQueueOutputStatisticInfoClusterStatQueueOutput

func (list *GetQueueOutputStatisticInfoClusterStatQueueOutputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetQueueOutputStatisticInfoClusterStatQueueOutput)
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
