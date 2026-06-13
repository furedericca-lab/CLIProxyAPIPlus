package auth

func init() {
	registerRefreshLead("kiro", func() Authenticator { return NewKiroAuthenticator() })
	registerRefreshLead("github-copilot", func() Authenticator { return NewGitHubCopilotAuthenticator() })
	registerRefreshLead("gitlab", func() Authenticator { return NewGitLabAuthenticator() })
	registerRefreshLead("codebuddy", func() Authenticator { return NewCodeBuddyAuthenticator() })
	registerRefreshLead("codebuddy-intl", func() Authenticator { return NewCodeBuddyIntlAuthenticator() })
	registerRefreshLead("cursor", func() Authenticator { return NewCursorAuthenticator() })
}
