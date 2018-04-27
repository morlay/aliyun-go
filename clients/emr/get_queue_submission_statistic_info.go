package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetQueueSubmissionStatisticInfoRequest struct {
	requests.RpcRequest
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	ApplicationType string `position:"Query" name:"ApplicationType"`
	FinalStatus     string `position:"Query" name:"FinalStatus"`
}

func (req *GetQueueSubmissionStatisticInfoRequest) Invoke(client *sdk.Client) (resp *GetQueueSubmissionStatisticInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetQueueSubmissionStatisticInfo", "", "")
	resp = &GetQueueSubmissionStatisticInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetQueueSubmissionStatisticInfoResponse struct {
	responses.BaseResponse
	RequestId           string
	QueueSubmissionList GetQueueSubmissionStatisticInfoClusterStatQueueSubmissionList
}

type GetQueueSubmissionStatisticInfoClusterStatQueueSubmission struct {
	Queue      string
	Submission int64
}

type GetQueueSubmissionStatisticInfoClusterStatQueueSubmissionList []GetQueueSubmissionStatisticInfoClusterStatQueueSubmission

func (list *GetQueueSubmissionStatisticInfoClusterStatQueueSubmissionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetQueueSubmissionStatisticInfoClusterStatQueueSubmission)
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
