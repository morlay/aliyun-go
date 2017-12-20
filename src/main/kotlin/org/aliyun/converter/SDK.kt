package org.aliyun.converter

import java.io.File

class SDK(val name: String, val version: String, val baseDir: String) {
    companion object {
        fun collectAll(): Map<String, SDK> {
            val dirs = glob(sdkModelPath, File("aliyun-openapi-java-sdk"))

            val sdks = mutableMapOf<String, SDK>()

            for (dir in dirs) {
                val names = dir.split("/")
                val group = names[1].replace(Regex("^aliyun-java-sdk-"), "")
                val versionDates = Regex("^v([0-9]{4})([0-9]{2})([0-9]{2})").matchEntire(names[9])!!.groupValues
                val version = "${versionDates[1]}-${versionDates[2]}-${versionDates[3]}"

                sdks[group] = SDK(group, version, dir)
            }

            return sdks
        }
    }

    val operations = mutableMapOf<String, Operation>()
    private val definitions = mutableMapOf<String, Schema>()

    fun addOperation(op: Operation) {
        this.operations[op.action] = op
    }

    fun def(s: Schema) {
        this.definitions[s.name] = s
    }

    fun withPkg(code: String): String {
        return """
package ${this.name.replace("-", "_")}

import (
    "encoding/json"
    "github.com/morlay/aliyun-go/core"
)

${code}
            """
    }

    fun goClient(): String {
        val name = toUpperCamelCase(this.name)
        val endpoint = "https://${this.name}.aliyuncs.com"

        return this.withPkg("""
func New${name}Client(key string, secret string, regionId string) *${name}Client {
    return &${name}Client{
        Client: core.Client{
            Endpoint: "${endpoint}",
            Version: "${this.version}",
            RegionID: regionId,
            AccessKeyId: key,
            AccessKeySecret: secret,
        },
    }
}

type ${name}Client struct {
    core.Client
}
""")
    }
}
