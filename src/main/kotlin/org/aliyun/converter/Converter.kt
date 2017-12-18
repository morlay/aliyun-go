package org.aliyun.converter

var sdkModelPath = Regex("aliyun-openapi-java-sdk/([^\\/]+)/src/main/java/com/aliyuncs/([^\\/]+)/model/v([^\\/]+)$")

class Converter {
    companion object {
        @JvmStatic
        fun main(args: Array<String>) {
            val sdks = SDK.collectAll()

            sdks.forEach { s ->
                generateSDK(s.value)
            }

//            val sdk = sdks.getValue("vpc")
//            generateSDK(sdk)
        }

        fun generateSDK(sdk: SDK) {
            println("generating sdk ${sdk.name} ...")

            val parsedCodes = loadAndParseAllJavaFile(sdk.baseDir)

            var canGenerate = false

            parsedCodes.forEach { parsedCode ->
                parsedCode.value.imports.forEach { i ->
                    if (i.name.asString() == "com.aliyuncs.RpcAcsRequest") {
                        canGenerate = true
                    }
                }
            }

            if (!canGenerate) {
                println("${sdk.name} is not support")
                return
            }

            parsedCodes.forEach { parsedCode ->
                parsedCode.value.types.forEach { type ->
                    if (type.isTopLevelType) {
                        type.members.forEach { member ->
                            if (member.isConstructorDeclaration) {
                                val constructorDeclaration = member.asConstructorDeclaration()
                                constructorDeclaration.body.statements.forEach { stmt ->
                                    if (stmt.isExplicitConstructorInvocationStmt) {
                                        val argsForSuper = stmt.asExplicitConstructorInvocationStmt().arguments
                                        if (argsForSuper[2].isStringLiteralExpr) {
                                            sdk.addOperation(Operation(argsForSuper[2].asStringLiteralExpr().asString()))
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }

            parsedCodes.forEach({ parsedCode ->
                val imports = mutableMapOf<String, List<String>>()

                parsedCode.value.imports.forEach { i ->
                    val importPaths = i.name.asString().split(".")
                    imports.set(importPaths[importPaths.size - 1], importPaths)
                }

                sdk.operations.forEach { op ->
                    parsedCode.value.types.forEach { type ->
                        val schema = Schema.fromTypeDecl(type, imports)
                        when (schema.name) {
                            op.key + "Request" -> {
                                op.value.setParam(schema)
                            }
                            op.key + "Response" -> {
                                op.value.setResp(schema)
                            }
                            else -> {
                                sdk.def(schema)
                            }
                        }
                    }
                }
            })

            val baseDir = "clients/${sdk.name}"

            writeFile("${baseDir}/client.go", sdk.goClient())

            var operations = ""

            sdk.operations.forEach { op ->
                operations += op.value.goClientMethod(sdk.name)
            }

            writeFile("${baseDir}/operations.go", sdk.withPkg(operations))

            Runtime.getRuntime().exec("goimports -w ./${baseDir}")
        }
    }
}
