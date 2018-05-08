package crm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type FindAllContacterRequest struct {
	requests.RpcRequest
	KpId int64 `position:"Query" name:"KpId"`
}

func (req *FindAllContacterRequest) Invoke(client *sdk.Client) (resp *FindAllContacterResponse, err error) {
	req.InitWithApiInfo("Crm", "2015-03-24", "FindAllContacter", "", "")
	resp = &FindAllContacterResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindAllContacterResponse struct {
	responses.BaseResponse
	Success       bool
	ResultCode    common.String
	ResultMessage common.String
	Data          FindAllContacterContacterInfoList
}

type FindAllContacterContacterInfo struct {
	ContacterId       common.Long
	KpId              common.Long
	CustomerId        common.Long
	ContacterName     common.String
	ContacterEmail    common.String
	ContacterMobile   common.String
	ContacterPosition common.String
}

type FindAllContacterContacterInfoList []FindAllContacterContacterInfo

func (list *FindAllContacterContacterInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindAllContacterContacterInfo)
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
