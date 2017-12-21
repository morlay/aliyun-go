package crm

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindAllContacterRequest struct {
	KpId int64 `position:"Query" name:"KpId"`
}

func (r FindAllContacterRequest) Invoke(client *sdk.Client) (response *FindAllContacterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		FindAllContacterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Crm", "2015-03-24", "FindAllContacter", "", "")

	resp := struct {
		*responses.BaseResponse
		FindAllContacterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.FindAllContacterResponse

	err = client.DoAction(&req, &resp)
	return
}

type FindAllContacterResponse struct {
	Success       bool
	ResultCode    string
	ResultMessage string
	Data          FindAllContacterContacterInfoList
}

type FindAllContacterContacterInfo struct {
	ContacterId       int64
	KpId              int64
	CustomerId        int64
	ContacterName     string
	ContacterEmail    string
	ContacterMobile   string
	ContacterPosition string
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
