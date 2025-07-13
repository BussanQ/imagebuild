// Package seaweed
// @author Administrator
// @date 2025/07/13
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	TargetPath string `json:"targetPath"`
	Mounts     []struct {
		MountPoint string `json:"mountPoint"`
		Options    struct {
			URL          string `json:"url"`
			AccessKey    string `json:"accessKey"`
			AccessSecret string `json:"accessSecret"`
		} `json:"options"`
	} `json:"mounts"`
}

//export AWS_ACCESS_KEY_ID=$(cat $akId)
//export AWS_SECRET_ACCESS_KEY=$(cat $akSecret)

const script = `#!/bin/sh
set -ex
mkdir -p $targetPath
weed mount -filer=$url -dir=$targetPath -filer.path=$remotePath
`

func main() {
	// 读取配置文件
	fileContent, err := os.ReadFile("/etc/fluid/config/config.json")
	if err != nil {
		fileContent, err = os.ReadFile("/etc/fluid/config.json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading config file: %v\n", err)
			os.Exit(1)
		}
	}

	// 解析 JSON 数据
	var config Config
	if err := json.Unmarshal(fileContent, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// 创建并写入 shell 脚本文件
	scriptFile, err := os.Create("mount.sh")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating script file: %v\n", err)
		os.Exit(1)
	}
	defer scriptFile.Close()

	// 写入变量部分到脚本文件
	fmt.Fprintf(scriptFile, "targetPath=\"%s\"\n", config.TargetPath)
	fmt.Fprintf(scriptFile, "url=\"%s\"\n", config.Mounts[0].Options.URL)

	mountPoint := config.Mounts[0].MountPoint
	if strings.HasPrefix(mountPoint, "s3://") {
		mountPoint = mountPoint[len("s3://"):]
	}
	fmt.Fprintf(scriptFile, "remotePath=\"/%s\"\n", mountPoint)
	//fmt.Fprintf(scriptFile, "akId=\"%s\"\n", config.Mounts[0].Options.AccessKey)
	//fmt.Fprintf(scriptFile, "akSecret=\"%s\"\n", config.Mounts[0].Options.AccessSecret)

	// 写入脚本主体
	fmt.Fprint(scriptFile, script)

	// 设置脚本文件权限为可执行
	if err := os.Chmod("mount.sh", 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error setting script file permissions: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Script generated successfully.")
}
