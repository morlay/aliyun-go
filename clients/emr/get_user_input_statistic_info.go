package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type GetUserInputStatisticInfoRequest struct {
	requests.RpcRequest
	FromDatetime    string `position:"Query" name:"FromDatetime"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ToDatetime      string `position:"Query" name:"ToDatetime"`
}

func (req *GetUserInputStatisticInfoRequest) Invoke(client *sdk.Client) (resp *GetUserInputStatisticInfoResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "GetUserInputStatisticInfo", "", "")
	resp = &GetUserInputStatisticInfoResponse{}
	err = client.DoAction(req, resp)
	return
}

type GetUserInputStatisticInfoResponse struct {
	responses.BaseResponse
	RequestId     common.String
	UserInputList GetUserInputStatisticInfoClusterStatUserInputList
}

type GetUserInputStatisticInfoClusterStatUserInput struct {
	User       common.String
	BytesInput common.Long
}

type GetUserInputStatisticInfoClusterStatUserInputList []GetUserInputStatisticInfoClusterStatUserInput

func (list *GetUserInputStatisticInfoClusterStatUserInputList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]GetUserInputStatisticInfoClusterStatUserInput)
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
