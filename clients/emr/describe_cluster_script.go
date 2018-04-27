package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeClusterScriptRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
}

func (req *DescribeClusterScriptRequest) Invoke(client *sdk.Client) (resp *DescribeClusterScriptResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DescribeClusterScript", "", "")
	resp = &DescribeClusterScriptResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeClusterScriptResponse struct {
	responses.BaseResponse
	RequestId           string
	ScriptNodeInstances DescribeClusterScriptScriptNodeInstanceList
}

type DescribeClusterScriptScriptNodeInstance struct {
	NodeId    string
	NodeIp    string
	StartTime int64
	EndTime   int64
	Status    string
}

type DescribeClusterScriptScriptNodeInstanceList []DescribeClusterScriptScriptNodeInstance

func (list *DescribeClusterScriptScriptNodeInstanceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeClusterScriptScriptNodeInstance)
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
