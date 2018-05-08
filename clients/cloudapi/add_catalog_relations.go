package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AddCatalogRelationsRequest struct {
	requests.RpcRequest
	CatalogId string                          `position:"Query" name:"CatalogId"`
	ApiIds    string                          `position:"Query" name:"ApiIds"`
	ApiList   *AddCatalogRelationsApiListList `type:"Repeated" name:"ApiList"`
}

func (req *AddCatalogRelationsRequest) Invoke(client *sdk.Client) (resp *AddCatalogRelationsResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "AddCatalogRelations", "apigateway", "")
	resp = &AddCatalogRelationsResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddCatalogRelationsResponse struct {
	responses.BaseResponse
	RequestId common.String
}

type AddCatalogRelationsApiListList []string

func (list *AddCatalogRelationsApiListList) UnmarshalJSON(data []byte) error {
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
