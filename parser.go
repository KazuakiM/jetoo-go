package jetoo

import (
	"net/url"
)

func Parse(rawurl string) DataSet {
	return fillResult(execParse(u))
}

func execParse(rawurl string) DataSet {
	u, err := url.Parse(rawurl)
	if err != nil {
		return DataSet{}
	}

	pass, err := u.User.Password()
	if err != nil {
		return DataSet{}
	}

	return DataSet{
		Scheme:   u.Scheme,
		Host:     u.Hostname(),
		Port:     u.Port(),
		User:     u.User.Username(),
		Pass:     pass,
		Path:     u.Path,
		Query:    u.RawQuery,
		Fragment: u.Fragment,
	}
}

func fillResult(result DataSet) DataSet {
	return result
}
