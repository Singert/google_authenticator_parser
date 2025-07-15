太好了，既然测试成功，以下是你需要的：

---

## ✅ `README.md`

````markdown
# Google Authenticator QR Code Decoder (Go)

本项目提供一个 Go 程序，用于解析 Google Authenticator 导出的迁移二维码（`otpauth-migration://` URL），并提取出账号信息和 Base32 编码的 Secret。

## 功能简介

- 从 `otpauth-migration://offline?data=...` URL 中提取 `data` 字段；
- 解码 URL 编码 + Base64 编码；
- 使用 Protocol Buffers 解码为 `MigrationPayload`；
- 输出每个账号的账号名、发行方、Base32 格式 Secret。

## 使用方式

```bash
go run main.go '<otpauth-migration://offline?data=...>'
````

例如：

```bash
go run main.go 'otpauth-migration://offline?data=Ci8KCke7Wn1dzBf7...（省略）'
```

## 项目结构

```
google_authenticator_parser/
├── go.mod
├── main.go
├── README.md
└── OtpMigration/
    ├── OtpMigration.proto
    └── OtpMigration.pb.go
```

## 编译依赖

确保安装以下工具：

* [protoc](https://github.com/protocolbuffers/protobuf/releases)
* `protoc-gen-go` 插件：

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## 参考资料

本项目参考了 Jimmy0w0 的博客文章：
👉 [https://jimmy0w0.me/zh-Hans/posts/google-authenticator-qr-code-data-zh-Hans/](https://jimmy0w0.me/zh-Hans/posts/google-authenticator-qr-code-data-zh-Hans/)

> 感谢其深入分析 Google Authenticator QR 导出机制，帮助我们快速实现兼容。

---

MIT License


---

## ✅ Git commit message

=
````

 
