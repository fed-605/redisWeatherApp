package provider

func validate(baseURL, apiKey, location string) error {
	if baseURL == "" {
		return errEmptyUrl
	}
	if apiKey == "" {
		return errInvalidApiKey
	}
	if location == "" {
		return errEmptyLocation
	}
	return nil
}
