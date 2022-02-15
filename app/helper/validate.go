package helper

func TypeByPlatform(platform string) string {
	switch platform {
	case "ios":
		return "mobile"
	case "android":
		return "mobile"
	default:
		return "web"
	}
}
