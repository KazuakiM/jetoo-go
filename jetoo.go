package jetoo

import (
	"net/url"
	"strconv"
)

func Parse(rawurl string) DataSet {
	return fillResult(execParse(rawurl))
}

func execParse(rawurl string) (result DataSet) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return
	}

	pass, _ := u.User.Password()

	port, err := strconv.Atoi(u.Port())
	if err != nil {
		return
	}

	result = DataSet{
		Scheme:   u.Scheme,
		Host:     u.Hostname(),
		Port:     port,
		User:     u.User.Username(),
		Pass:     pass,
		Path:     u.Path,
		Query:    u.RawQuery,
		Fragment: u.Fragment,
	}
	return
}

func fillResult(result DataSet) DataSet {
	return result
}
