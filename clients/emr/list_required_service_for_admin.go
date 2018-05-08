package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListRequiredServiceForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	EmrVersion      string `position:"Query" name:"EmrVersion"`
	ServiceNameList string `position:"Query" name:"ServiceNameList"`
}

func (req *ListRequiredServiceForAdminRequest) Invoke(client *sdk.Client) (resp *ListRequiredServiceForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListRequiredServiceForAdmin", "", "")
	resp = &ListRequiredServiceForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListRequiredServiceForAdminResponse struct {
	responses.BaseResponse
	RequestId        common.String
	NeedOtherService bool
	ServiceList      ListRequiredServiceForAdminServiceList
}

type ListRequiredServiceForAdminService struct {
	ServiceName        common.String
	ServiceDisplayName common.String
}

type ListRequiredServiceForAdminServiceList []ListRequiredServiceForAdminService

func (list *ListRequiredServiceForAdminServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRequiredServiceForAdminService)
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
