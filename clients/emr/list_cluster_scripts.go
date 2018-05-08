package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListClusterScriptsRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *ListClusterScriptsRequest) Invoke(client *sdk.Client) (resp *ListClusterScriptsResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterScripts", "", "")
	resp = &ListClusterScriptsResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterScriptsResponse struct {
	responses.BaseResponse
	RequestId      common.String
	ClusterScripts ListClusterScriptsClusterScriptList
}

type ListClusterScriptsClusterScript struct {
	Id        common.String
	Name      common.String
	StartTime common.Long
	EndTime   common.Long
	Path      common.String
	Args      common.String
	Status    common.String
}

type ListClusterScriptsClusterScriptList []ListClusterScriptsClusterScript

func (list *ListClusterScriptsClusterScriptList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterScriptsClusterScript)
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
