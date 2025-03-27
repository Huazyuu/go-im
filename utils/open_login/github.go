package open_login

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type GitHubConf struct {
	ClientID     string
	ClientSecret string
	Redirect     string
}

// GitHubInfo GitHub返回的字段
type GitHubInfo struct {
	Login     string `json:"login"`      // 用户名
	Name      string `json:"name"`       // 昵称
	AvatarURL string `json:"avatar_url"` // 头像
	ID        string `json:"id"`         // 用户唯一标识
}

// GitHubLogin GitHub需要的凭证参数
type GitHubLogin struct {
	clientID     string
	clientSecret string
	redirect     string
	code         string
	accessToken  string
}

func NewGitHubLogin(code string, conf GitHubConf) (githubInfo GitHubInfo, err error) {
	if code == "" {
		return GitHubInfo{}, errors.New("code is empty")
	}
	githubLogin := &GitHubLogin{
		clientID:     conf.ClientID,
		clientSecret: conf.ClientSecret,
		redirect:     conf.Redirect,
		code:         code,
	}
	err = githubLogin.GetAccessToken()
	if err != nil {
		return githubInfo, err
	}
	githubInfo, err = githubLogin.GetUserInfo()
	if err != nil {
		return githubInfo, err
	}
	logx.Info(githubInfo)
	return githubInfo, nil
}

// GetAccessToken 获取访问令牌
func (g *GitHubLogin) GetAccessToken() error {
	client := &http.Client{Timeout: 30 * time.Second}

	data := url.Values{}
	data.Set("client_id", g.clientID)
	data.Set("client_secret", g.clientSecret)
	data.Set("code", g.code)
	data.Set("redirect_uri", g.redirect)

	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return fmt.Errorf("create request failed: %v", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		logx.Error(err)
		return fmt.Errorf("request failed: %v", errors.New("系统错误"))
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %v", err)
	}

	var result struct {
		AccessToken string `json:"access_token"`
		Error       string `json:"error"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("parse json failed: %v", err)
	}

	if result.AccessToken == "" {
		return fmt.Errorf("empty access token: %s", result.Error)
	}

	g.accessToken = result.AccessToken
	return nil
}

// GetUserInfo 获取用户信息
func (g *GitHubLogin) GetUserInfo() (GitHubInfo, error) {
	var info GitHubInfo
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return info, fmt.Errorf("create request failed: %v", err)
	}

	req.Header.Add("Authorization", "token "+g.accessToken)
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return info, fmt.Errorf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return info, fmt.Errorf("unexpected status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return info, fmt.Errorf("read body failed: %v", err)
	}

	if err := json.Unmarshal(body, &info); err != nil {
		return info, fmt.Errorf("parse json failed: %v", err)
	}

	return info, nil
}
