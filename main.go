package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 定义要拦截的域名列表
	adDomains := []string{
		"ourl.co",
	}

	// 生成广告过滤规则文件
	err := generateAdBlockRules("ad_block_rule_host.txt", adDomains)
	if err != nil {
		fmt.Println("Error generating ad block rules:", err)
	}

	// 生成 AdBlock 规则文件
	err = generateAdBlockAdblockRules("ad_block_rule_adblock.txt", adDomains)
	if err != nil {
		fmt.Println("Error generating ad block adblock rules:", err)
	}
}

// generateNoice 生成文件头部信息
func generateNoice(file *os.File, adDomains []string) error {
	// 写入文件头部信息
	_, err := file.WriteString("!Title: Sorana Ads Rule\n")
	if err != nil {
		return err
	}

	// 写入当前时间戳
	currentTime := time.Now().UTC().Format("2006-01-02T15:04:05.999999Z07:00")
	_, err = file.WriteString(fmt.Sprintf("!Last modified: %s\n", currentTime))
	if err != nil {
		return err
	}

	_, err = file.WriteString("!--------------------------------------\n")
	if err != nil {
		return err
	}

	// 写入总行数
	totalLines := len(adDomains) + 3 // +3 for the header lines
	_, err = file.WriteString(fmt.Sprintf("!Total lines: %d\n", totalLines))
	if err != nil {
		return err
	}

	// 写入版本信息
	_, err = file.WriteString("!Version: 0.0.1-alpha\n\n")
	if err != nil {
		return err
	}

	// 写入项目链接和许可证
	_, err = file.WriteString("!Homepage: https://github.com/JDruki/Sorana-Ads-Rule\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString("!License: https://github.com/JDruki/Sorana-Ads-Rule/blob/main/LICENSE\n\n")
	if err != nil {
		return err
	}
	return nil
}

// generateAdBlockRules 生成广告过滤规则文件
func generateAdBlockRules(filename string, adDomains []string) error {
	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入头部信息
	err = generateNoice(file, adDomains)
	if err != nil {
		return err
	}

	// 写入本地回环地址
	_, err = file.WriteString("127.0.0.1 localhost\n")
	if err != nil {
		return err
	}
	_, err = file.WriteString("::1 localhost\n\n")
	if err != nil {
		return err
	}

	// 写入要拦截的域名
	for _, domain := range adDomains {
		_, err = file.WriteString(fmt.Sprintf("127.0.0.1 %s\n", domain))
		if err != nil {
			return err
		}
	}

	fmt.Println("广告拦截规则文件已创建：", filename)
	return nil
}

// generateAdBlockAdblockRules 生成 AdBlock 规则文件
func generateAdBlockAdblockRules(filename string, adDomains []string) error {
	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入头部信息
	err = generateNoice(file, adDomains)
	if err != nil {
		return err
	}

	// 写入规则
	for _, domain := range adDomains {
		// 将域名转换为 AdBlock 规则格式
		rule := fmt.Sprintf("-*-*-*-*.%s^$script,third-party\n", domain)
		_, err = file.WriteString(rule)
		if err != nil {
			return err
		}
	}

	fmt.Println("AdBlock 规则文件已创建：", filename)
	return nil
}

