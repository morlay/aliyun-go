package push

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryAliasesRequest struct {
	requests.RpcRequest
	AppKey   int64  `position:"Query" name:"AppKey"`
	DeviceId string `position:"Query" name:"DeviceId"`
}

func (req *QueryAliasesRequest) Invoke(client *sdk.Client) (resp *QueryAliasesResponse, err error) {
	req.InitWithApiInfo("Push", "2016-08-01", "QueryAliases", "", "")
	resp = &QueryAliasesResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryAliasesResponse struct {
	responses.BaseResponse
	RequestId  string
	AliasInfos QueryAliasesAliasInfoList
}

type QueryAliasesAliasInfo struct {
	AliasName string
}

type QueryAliasesAliasInfoList []QueryAliasesAliasInfo

func (list *QueryAliasesAliasInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryAliasesAliasInfo)
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
