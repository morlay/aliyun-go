package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetQueueInputStatisticInfoRequest struct {
	requests.RpcRequest
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
}

func (req *GetQueueInputStatisticInfoRequest) Invoke(client *sdk.Client) (resp *GetQueueInputStatisticInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetQueueInputStatisticInfo", "", "")
	resp = &GetQueueInputStatisticInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetQueueInputStatisticInfoResponse struct {
	responses.BaseResponse
	RequestId      string
	QueueInputList GetQueueInputStatisticInfoClusterStatQueueInputList
}

type GetQueueInputStatisticInfoClusterStatQueueInput struct {
	Queue      string
	BytesInput int64
}

type GetQueueInputStatisticInfoClusterStatQueueInputList []GetQueueInputStatisticInfoClusterStatQueueInput

func (list *GetQueueInputStatisticInfoClusterStatQueueInputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetQueueInputStatisticInfoClusterStatQueueInput)
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
