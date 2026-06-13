package main

import (
	"flag"

	"github.com/router-for-me/CLIProxyAPI/v7/internal/auth/kiro"
	"github.com/router-for-me/CLIProxyAPI/v7/internal/cmd"
	"github.com/router-for-me/CLIProxyAPI/v7/internal/config"
)

type plusLoginFlags struct {
	qwenLogin          bool
	kiloLogin          bool
	iflowLogin         bool
	iflowCookie        bool
	gitlabLogin        bool
	gitlabTokenLogin   bool
	cursorLogin        bool
	kiroLogin          bool
	kiroGoogleLogin    bool
	kiroAWSLogin       bool
	kiroAWSAuthCode    bool
	kiroImport         bool
	kiroIDCLogin       bool
	kiroCLILogin       bool
	kiroIDCStartURL    string
	kiroIDCRegion      string
	kiroIDCFlow        string
	githubCopilotLogin bool
	clineLogin         bool
	codeBuddyLogin     bool
	codeBuddyIntlLogin bool
}

func registerPlusLoginFlags(fs *flag.FlagSet, flags *plusLoginFlags) {
	fs.BoolVar(&flags.qwenLogin, "qwen-login", false, "Login to Qwen using OAuth")
	fs.BoolVar(&flags.kiloLogin, "kilo-login", false, "Login to Kilo AI using device flow")
	fs.BoolVar(&flags.iflowLogin, "iflow-login", false, "Login to iFlow using OAuth")
	fs.BoolVar(&flags.iflowCookie, "iflow-cookie", false, "Login to iFlow using Cookie")
	fs.BoolVar(&flags.gitlabLogin, "gitlab-login", false, "Login to GitLab Duo using OAuth")
	fs.BoolVar(&flags.gitlabTokenLogin, "gitlab-token-login", false, "Login to GitLab Duo using a personal access token")
	fs.BoolVar(&flags.cursorLogin, "cursor-login", false, "Login to Cursor using OAuth")
	fs.BoolVar(&flags.kiroLogin, "kiro-login", false, "Login to Kiro using Google OAuth")
	fs.BoolVar(&flags.kiroGoogleLogin, "kiro-google-login", false, "Login to Kiro using Google OAuth (same as --kiro-login)")
	fs.BoolVar(&flags.kiroAWSLogin, "kiro-aws-login", false, "Login to Kiro using AWS Builder ID (device code flow)")
	fs.BoolVar(&flags.kiroAWSAuthCode, "kiro-aws-authcode", false, "Login to Kiro using AWS Builder ID (authorization code flow, better UX)")
	fs.BoolVar(&flags.kiroImport, "kiro-import", false, "Import Kiro token from Kiro IDE (~/.aws/sso/cache/kiro-auth-token.json)")
	fs.BoolVar(&flags.kiroIDCLogin, "kiro-idc-login", false, "Login to Kiro using IAM Identity Center (IDC)")
	fs.BoolVar(&flags.kiroCLILogin, "kiro-cli-login", false, "Login to Kiro using native Kiro CLI OAuth flow")
	fs.StringVar(&flags.kiroIDCStartURL, "kiro-idc-start-url", "", "IDC start URL (required with --kiro-idc-login)")
	fs.StringVar(&flags.kiroIDCRegion, "kiro-idc-region", "", "IDC region (default: us-east-1)")
	fs.StringVar(&flags.kiroIDCFlow, "kiro-idc-flow", "", "IDC flow type: authcode (default) or device")
	fs.BoolVar(&flags.githubCopilotLogin, "github-copilot-login", false, "Login to GitHub Copilot using device flow")
	fs.BoolVar(&flags.clineLogin, "cline-login", false, "Login to Cline using OAuth")
	fs.BoolVar(&flags.codeBuddyLogin, "codebuddy-login", false, "Login to CodeBuddy using browser OAuth flow")
	fs.BoolVar(&flags.codeBuddyIntlLogin, "codebuddy-intl-login", false, "Login to CodeBuddy International (codebuddy.ai) using browser OAuth flow")
}

func (flags plusLoginFlags) Active() bool {
	return flags.qwenLogin ||
		flags.kiloLogin ||
		flags.iflowLogin ||
		flags.iflowCookie ||
		flags.gitlabLogin ||
		flags.gitlabTokenLogin ||
		flags.cursorLogin ||
		flags.kiroLogin ||
		flags.kiroGoogleLogin ||
		flags.kiroAWSLogin ||
		flags.kiroAWSAuthCode ||
		flags.kiroImport ||
		flags.kiroIDCLogin ||
		flags.kiroCLILogin ||
		flags.githubCopilotLogin ||
		flags.clineLogin ||
		flags.codeBuddyLogin ||
		flags.codeBuddyIntlLogin
}

func handlePlusLoginCommand(cfg *config.Config, options *cmd.LoginOptions, flags plusLoginFlags, useIncognito, noIncognito bool) bool {
	switch {
	case flags.githubCopilotLogin:
		cmd.DoGitHubCopilotLogin(cfg, options)
	case flags.codeBuddyLogin:
		cmd.DoCodeBuddyLogin(cfg, options)
	case flags.codeBuddyIntlLogin:
		cmd.DoCodeBuddyIntlLogin(cfg, options)
	case flags.clineLogin:
		cmd.DoClineLogin(cfg, options)
	case flags.qwenLogin:
		cmd.DoQwenLogin(cfg, options)
	case flags.kiloLogin:
		cmd.DoKiloLogin(cfg, options)
	case flags.iflowLogin:
		cmd.DoIFlowLogin(cfg, options)
	case flags.iflowCookie:
		cmd.DoIFlowCookieAuth(cfg, options)
	case flags.gitlabLogin:
		cmd.DoGitLabLogin(cfg, options)
	case flags.gitlabTokenLogin:
		cmd.DoGitLabTokenLogin(cfg, options)
	case flags.cursorLogin:
		cmd.DoCursorLogin(cfg, options)
	case flags.kiroLogin:
		runKiroLoginCommand(cfg, useIncognito, noIncognito, func() { cmd.DoKiroLogin(cfg, options) })
	case flags.kiroGoogleLogin:
		runKiroLoginCommand(cfg, useIncognito, noIncognito, func() { cmd.DoKiroGoogleLogin(cfg, options) })
	case flags.kiroAWSLogin:
		runKiroLoginCommand(cfg, useIncognito, noIncognito, func() { cmd.DoKiroAWSLogin(cfg, options) })
	case flags.kiroAWSAuthCode:
		runKiroLoginCommand(cfg, useIncognito, noIncognito, func() { cmd.DoKiroAWSAuthCodeLogin(cfg, options) })
	case flags.kiroImport:
		kiro.InitFingerprintConfig(cfg)
		cmd.DoKiroImport(cfg, options)
	case flags.kiroIDCLogin:
		runKiroLoginCommand(cfg, useIncognito, noIncognito, func() {
			cmd.DoKiroIDCLogin(cfg, options, flags.kiroIDCStartURL, flags.kiroIDCRegion, flags.kiroIDCFlow)
		})
	case flags.kiroCLILogin:
		runKiroLoginCommand(cfg, useIncognito, noIncognito, func() { cmd.DoKiroCLILogin(cfg, options) })
	default:
		return false
	}
	return true
}

func runKiroLoginCommand(cfg *config.Config, useIncognito, noIncognito bool, run func()) {
	// Kiro auth defaults to incognito mode for multi-account support. Auth
	// commands exit after completion, so this config mutation does not leak into
	// server startup.
	setKiroIncognitoMode(cfg, useIncognito, noIncognito)
	kiro.InitFingerprintConfig(cfg)
	run()
}
