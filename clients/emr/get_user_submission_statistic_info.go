package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetUserSubmissionStatisticInfoRequest struct {
	requests.RpcRequest
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
	ApplicationType string `position:"Query" name:"ApplicationType"`
	FinalStatus     string `position:"Query" name:"FinalStatus"`
}

func (req *GetUserSubmissionStatisticInfoRequest) Invoke(client *sdk.Client) (resp *GetUserSubmissionStatisticInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetUserSubmissionStatisticInfo", "", "")
	resp = &GetUserSubmissionStatisticInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetUserSubmissionStatisticInfoResponse struct {
	responses.BaseResponse
	RequestId          common.String
	UserSubmissionList GetUserSubmissionStatisticInfoClusterStatUserSubmissionList
}

type GetUserSubmissionStatisticInfoClusterStatUserSubmission struct {
	User       common.String
	Submission common.Long
}

type GetUserSubmissionStatisticInfoClusterStatUserSubmissionList []GetUserSubmissionStatisticInfoClusterStatUserSubmission

func (list *GetUserSubmissionStatisticInfoClusterStatUserSubmissionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserSubmissionStatisticInfoClusterStatUserSubmission)
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
