package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId     common.String
	Total         common.Integer
	PageNumber    common.Integer
	PageSize      common.Integer
	JobOutputList GetJobOutputStatisticInfoClusterStatJobOutputList
}

type GetJobOutputStatisticInfoClusterStatJobOutput struct {
	ApplicationId common.String
	JobId         common.String
	StartTime     common.Long
	FinishTime    common.Long
	Name          common.String
	Queue         common.String
	User          common.String
	State         common.String
	BytesOutput   common.Long
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
