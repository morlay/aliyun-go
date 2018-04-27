package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeStackRelationRequest struct {
	requests.RpcRequest
	ResourceOwnerId   int64                                      `position:"Query" name:"ResourceOwnerId"`
	EmrVersion        string                                     `position:"Query" name:"EmrVersion"`
	ClusterId         string                                     `position:"Query" name:"ClusterId"`
	StackVersion      string                                     `position:"Query" name:"StackVersion"`
	StackVersionLists *DescribeStackRelationStackVersionListList `position:"Query" type:"Repeated" name:"StackVersionList"`
	EmrVersionLists   *DescribeStackRelationEmrVersionListList   `position:"Query" type:"Repeated" name:"EmrVersionList"`
}

func (req *DescribeStackRelationRequest) Invoke(client *sdk.Client) (resp *DescribeStackRelationResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeStackRelation", "", "")
	resp = &DescribeStackRelationResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeStackRelationResponse struct {
	responses.BaseResponse
	RequestId            string
	Success              bool
	EmrStackRelationList DescribeStackRelationEmrStackRelationList
}

type DescribeStackRelationEmrStackRelation struct {
	EmrVersion   string
	StackVersion string
}

type DescribeStackRelationStackVersionListList []string

func (list *DescribeStackRelationStackVersionListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeStackRelationEmrVersionListList []string

func (list *DescribeStackRelationEmrVersionListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type DescribeStackRelationEmrStackRelationList []DescribeStackRelationEmrStackRelation

func (list *DescribeStackRelationEmrStackRelationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeStackRelationEmrStackRelation)
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
