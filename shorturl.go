package shorturl

import "net/url"

func ShortUrl(params ...interface{}) (string, error)  {

	//Check if the Long URL is valid
	_, err := url.ParseRequestURI(params[0].(string))
	if err != nil {
		return "", err
	}

	//Check if the Base URL is valid
	_, err = url.ParseRequestURI(params[1].(string))
	if err != nil {
		return "", err
	}

	return "", nil
}
