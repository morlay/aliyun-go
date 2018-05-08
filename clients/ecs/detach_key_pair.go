package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DetachKeyPairRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	InstanceIds          string `position:"Query" name:"InstanceIds"`
	KeyPairName          string `position:"Query" name:"KeyPairName"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *DetachKeyPairRequest) Invoke(client *sdk.Client) (resp *DetachKeyPairResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DetachKeyPair", "ecs", "")
	resp = &DetachKeyPairResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetachKeyPairResponse struct {
	responses.BaseResponse
	RequestId   common.String
	TotalCount  common.String
	FailCount   common.String
	KeyPairName common.String
	Results     DetachKeyPairResultList
}

type DetachKeyPairResult struct {
	InstanceId common.String
	Success    common.String
	Code       common.String
	Message    common.String
}

type DetachKeyPairResultList []DetachKeyPairResult

func (list *DetachKeyPairResultList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DetachKeyPairResult)
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
