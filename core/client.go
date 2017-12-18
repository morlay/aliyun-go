package core

import (
	"net/url"
	"io/ioutil"
	"github.com/morlay/aliyun-go/core/util"
	"net/http"
	"encoding/json"
	"github.com/sirupsen/logrus"
)

type Client struct {
	AccessKeyId     string // Access Key Id
	AccessKeySecret string // Access Key Secret
	SecurityToken   string
	Endpoint        string
	Version         string
	ServiceCode     string
	UserAgent       string
	RegionID        string
}

func (client *Client) Request(action string, args interface{}, response interface{}) error {
	request := Request{}
	request.init(client.Version, action, client.AccessKeyId, client.SecurityToken, client.RegionID)
	query := util.ConvertToQueryValues(request)
	util.SetQueryValues(args, &query)

	signature := util.CreateSignatureForRequest("GET", &query, client.AccessKeySecret+"&")

	requestURL := client.Endpoint + "?" + query.Encode() + "&Signature=" + url.QueryEscape(signature)

	logrus.Debugf("%s %s", "GET", requestURL)

	httpReq, err := http.NewRequest("GET", requestURL, nil)

	if err != nil {
		return GetClientError(err)
	}

	httpReq.Header.Set("User-Agent", httpReq.UserAgent()+" "+client.UserAgent)
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return GetClientError(err)
	}
	statusCode := httpResp.StatusCode

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)

	if err != nil {
		return GetClientError(err)
	}

	if statusCode >= 400 && statusCode <= 599 {
		errorResponse := ErrorResponse{}
		err = json.Unmarshal(body, &errorResponse)
		ecsError := &Error{
			ErrorResponse: errorResponse,
			StatusCode:    statusCode,
		}
		return ecsError
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		logrus.Errorf("%s", body)
		return GetClientError(err)
	}

	return nil
}

func GetClientErrorFromString(str string) error {
	return &Error{
		ErrorResponse: ErrorResponse{
			Code:    "AliyunGoClientFailure",
			Message: str,
		},
		StatusCode: -1,
	}
}

func GetClientError(err error) error {
	return GetClientErrorFromString(err.Error())
}
