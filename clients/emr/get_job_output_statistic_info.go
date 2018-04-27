package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetJobOutputStatisticInfoRequest struct {
	requests.RpcRequest
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *GetJobOutputStatisticInfoRequest) Invoke(client *sdk.Client) (resp *GetJobOutputStatisticInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetJobOutputStatisticInfo", "", "")
	resp = &GetJobOutputStatisticInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetJobOutputStatisticInfoResponse struct {
	responses.BaseResponse
	RequestId     string
	Total         int
	PageNumber    int
	PageSize      int
	JobOutputList GetJobOutputStatisticInfoClusterStatJobOutputList
}

type GetJobOutputStatisticInfoClusterStatJobOutput struct {
	ApplicationId string
	JobId         string
	StartTime     int64
	FinishTime    int64
	Name          string
	Queue         string
	User          string
	State         string
	BytesOutput   int64
}

type GetJobOutputStatisticInfoClusterStatJobOutputList []GetJobOutputStatisticInfoClusterStatJobOutput

func (list *GetJobOutputStatisticInfoClusterStatJobOutputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobOutputStatisticInfoClusterStatJobOutput)
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
