package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterServiceConfigTagRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ConfigTag       string `position:"Query" name:"ConfigTag"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *DescribeClusterServiceConfigTagRequest) Invoke(client *sdk.Client) (resp *DescribeClusterServiceConfigTagResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterServiceConfigTag", "", "")
	resp = &DescribeClusterServiceConfigTagResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterServiceConfigTagResponse struct {
	responses.BaseResponse
	RequestId     string
	ConfigTagList DescribeClusterServiceConfigTagConfigTagList
}

type DescribeClusterServiceConfigTagConfigTag struct {
	Tag       string
	TagDesc   string
	ValueList DescribeClusterServiceConfigTagValueList
}

type DescribeClusterServiceConfigTagValue struct {
	Value     string
	ValueDesc string
}

type DescribeClusterServiceConfigTagConfigTagList []DescribeClusterServiceConfigTagConfigTag

func (list *DescribeClusterServiceConfigTagConfigTagList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigTagConfigTag)
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

type DescribeClusterServiceConfigTagValueList []DescribeClusterServiceConfigTagValue

func (list *DescribeClusterServiceConfigTagValueList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterServiceConfigTagValue)
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
