package main

import (
    "encoding/base32"
    "encoding/base64"
    "fmt"
    "log"
    "net/url"
    "os"
    "strings"

    "google.golang.org/protobuf/proto"

    "github.com/Singert/google_authenticator_parser/OtpMigration" // 修改为你生成 .pb.go 的路径
)

func main() {
    // 1. 原始 QR URL，或你自己提取后的 data 参数
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "用法: %s <otpauth-migration URL>\n", os.Args[0])
        os.Exit(1)
    }

    rawURL := os.Args[1]
    // 2. 提取 & 解码 data 参数
    u, err := url.Parse(rawURL)
    if err != nil {
        log.Fatalf("无法解析 URL: %v", err)
    }

    base64Str := u.Query().Get("data")
    decodedURL, err := url.QueryUnescape(base64Str)
    if err != nil {
        log.Fatalf("URL 解码失败: %v", err)
    }

    binaryData, err := base64.StdEncoding.DecodeString(decodedURL)
    if err != nil {
        log.Fatalf("Base64 解码失败: %v", err)
    }

    // 3. 使用 Protobuf 解析数据
    var payload OtpMigration.MigrationPayload
    if err := proto.Unmarshal(binaryData, &payload); err != nil {
        log.Fatalf("Protocol Buffers 解析失败: %v", err)
    }

    // 4. 遍历账号并提取 Secret
    for _, param := range payload.OtpParameters {
        secretBase32 := base32.StdEncoding.EncodeToString(param.Secret)
        fmt.Println("==========")
        fmt.Printf("Account: %s\n", param.GetName())
        fmt.Printf("Issuer : %s\n", param.GetIssuer())
        fmt.Printf("Secret : %s\n", strings.TrimRight(secretBase32, "=")) // 去除 padding
    }
}
