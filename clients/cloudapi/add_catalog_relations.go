package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddCatalogRelationsRequest struct {
	CatalogId string                          `position:"Query" name:"CatalogId"`
	ApiIds    string                          `position:"Query" name:"ApiIds"`
	ApiList   *AddCatalogRelationsApiListList `type:"Repeated" name:"ApiList"`
}

func (r AddCatalogRelationsRequest) Invoke(client *sdk.Client) (response *AddCatalogRelationsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddCatalogRelationsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "AddCatalogRelations", "apigateway", "")

	resp := struct {
		*responses.BaseResponse
		AddCatalogRelationsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddCatalogRelationsResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddCatalogRelationsResponse struct {
	RequestId string
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
