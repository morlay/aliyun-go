package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type UpdateDomainToDomainGroupRequest struct {
	requests.RpcRequest
	DataSource    int                                      `position:"Query" name:"DataSource"`
	UserClientIp  string                                   `position:"Query" name:"UserClientIp"`
	FileToUpload  string                                   `position:"Body" name:"FileToUpload"`
	DomainNames   *UpdateDomainToDomainGroupDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	Replace       string                                   `position:"Query" name:"Replace"`
	Lang          string                                   `position:"Query" name:"Lang"`
	DomainGroupId int64                                    `position:"Query" name:"DomainGroupId"`
}

func (req *UpdateDomainToDomainGroupRequest) Invoke(client *sdk.Client) (resp *UpdateDomainToDomainGroupResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "UpdateDomainToDomainGroup", "", "")
	resp = &UpdateDomainToDomainGroupResponse{}
	err = client.DoAction(req, resp)
	return
}

type UpdateDomainToDomainGroupResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type UpdateDomainToDomainGroupDomainNameList []string

func (list *UpdateDomainToDomainGroupDomainNameList) UnmarshalJSON(data []byte) error {
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
