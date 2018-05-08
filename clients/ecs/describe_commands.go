package ecs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCommandsRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Description          string `position:"Query" name:"Description"`
	Type                 string `position:"Query" name:"Type"`
	CommandId            string `position:"Query" name:"CommandId"`
	PageNumber           int64  `position:"Query" name:"PageNumber"`
	PageSize             int64  `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	Name                 string `position:"Query" name:"Name"`
}

func (req *DescribeCommandsRequest) Invoke(client *sdk.Client) (resp *DescribeCommandsResponse, err error) {
	req.InitWithApiInfo("Ecs", "2014-05-26", "DescribeCommands", "ecs", "")
	resp = &DescribeCommandsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCommandsResponse struct {
	responses.BaseResponse
	RequestId  common.String
	TotalCount common.Long
	PageNumber common.Long
	PageSize   common.Long
	Commands   DescribeCommandsCommandList
}

type DescribeCommandsCommand struct {
	CommandId      common.String
	Name           common.String
	Type           common.String
	Description    common.String
	CommandContent common.String
	WorkingDir     common.String
	Timeout        common.Long
}

type DescribeCommandsCommandList []DescribeCommandsCommand

func (list *DescribeCommandsCommandList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCommandsCommand)
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
