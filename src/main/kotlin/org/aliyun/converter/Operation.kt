package org.aliyun.converter


class Operation(
        var requestType: String,
        val product: String,
        val version: String,
        val action: String,
        val serviceCode: String,
        var method: String,
        val uriPattern: String
) {
    private var parameters = Schema("")
    private var response = Schema("")

    fun setParam(s: Schema) {
        this.parameters = s
    }

    fun setResp(s: Schema) {
        this.response = s
    }

    fun goClientMethod(): String {
        val subTypePrefix = this.action
        val sideCodes = mutableMapOf<String, String>()

        var codes = ""

        codes += """
type ${this.action}Request ${this.parameters.goType(subTypePrefix, true, sideCodes)}
"""
        when (requestType) {
            "roa" -> {
                codes += """
func (r ${this.action}Request) Invoke(client *sdk.Client) (response *${this.action}Response, err error) {
	req := struct {
		*requests.RoaRequest
		${this.action}Request
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("${this.product}", "${this.version}", "${this.action}", "${this.uriPattern}","${this.serviceCode}", "")
    req.Method = "${this.method}"

	resp := struct {
		*responses.BaseResponse
		${this.action}Response
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.${this.action}Response

	err = client.DoAction(&req, &resp)
	return
}
"""
            }
            "rpc" -> {
                codes += """
func (r ${this.action}Request) Invoke(client *sdk.Client) (response *${this.action}Response, err error) {
	req := struct {
		*requests.RpcRequest
		${this.action}Request
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("${this.product}", "${this.version}", "${this.action}", "${this.serviceCode}", "")

	resp := struct {
		*responses.BaseResponse
		${this.action}Response
	}{
		BaseResponse: &responses.BaseResponse{},
	}
    response = &resp.${this.action}Response

	err = client.DoAction(&req, &resp)
	return
}
"""
            }
        }

        this.parameters.getAllDefinitions().forEach { s ->
            codes += """
type ${subTypePrefix}${s.key} ${s.value.goType(subTypePrefix, true, sideCodes)}
"""
        }

        codes += """
type ${this.action}Response ${this.response.goType(subTypePrefix, false, sideCodes)}
"""
        this.response.getAllDefinitions().forEach { s ->
            codes += """
type ${subTypePrefix}${s.key} ${s.value.goType(subTypePrefix, false, sideCodes)}
"""
        }

        sideCodes.forEach { s ->
            codes += s.value
        }

        return codes
    }
}