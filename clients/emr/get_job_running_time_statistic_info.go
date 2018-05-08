package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetJobRunningTimeStatisticInfoRequest struct {
	requests.RpcRequest
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *GetJobRunningTimeStatisticInfoRequest) Invoke(client *sdk.Client) (resp *GetJobRunningTimeStatisticInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetJobRunningTimeStatisticInfo", "", "")
	resp = &GetJobRunningTimeStatisticInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetJobRunningTimeStatisticInfoResponse struct {
	responses.BaseResponse
	RequestId       common.String
	Total           common.Integer
	PageNumber      common.Integer
	PageSize        common.Integer
	RunningTimeList GetJobRunningTimeStatisticInfoClusterStatJobRunningTimeList
}

type GetJobRunningTimeStatisticInfoClusterStatJobRunningTime struct {
	ApplicationId common.String
	JobId         common.String
	StartTime     common.Long
	FinishTime    common.Long
	Name          common.String
	Queue         common.String
	User          common.String
	State         common.String
	RunningTime   common.Long
}

type GetJobRunningTimeStatisticInfoClusterStatJobRunningTimeList []GetJobRunningTimeStatisticInfoClusterStatJobRunningTime

func (list *GetJobRunningTimeStatisticInfoClusterStatJobRunningTimeList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetJobRunningTimeStatisticInfoClusterStatJobRunningTime)
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
