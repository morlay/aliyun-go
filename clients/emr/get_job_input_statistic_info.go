package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type GetJobInputStatisticInfoRequest struct {
	requests.RpcRequest
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *GetJobInputStatisticInfoRequest) Invoke(client *sdk.Client) (resp *GetJobInputStatisticInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetJobInputStatisticInfo", "", "")
	resp = &GetJobInputStatisticInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetJobInputStatisticInfoResponse struct {
	responses.BaseResponse
	RequestId    string
	Total        int
	PageNumber   int
	PageSize     int
	JobInputList GetJobInputStatisticInfoClusterStatJobInputList
}

type GetJobInputStatisticInfoClusterStatJobInput struct {
	ApplicationId string
	JobId         string
	StartTime     int64
	FinishTime    int64
	Name          string
	Queue         string
	User          string
	State         string
	BytesInput    int64
}

type GetJobInputStatisticInfoClusterStatJobInputList []GetJobInputStatisticInfoClusterStatJobInput

func (list *GetJobInputStatisticInfoClusterStatJobInputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobInputStatisticInfoClusterStatJobInput)
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
