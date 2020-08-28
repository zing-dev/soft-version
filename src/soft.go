package soft

import (
	"fmt"
)

type Soft struct {
	Name      string    `json:"name"`
	Alias     string    `json:"alias"`
	Author    string    `json:"author"`
	Version   []Version `json:"versions"`
	Copyright string    `json:"copyright"`
}

func (s *Soft) SimpleVersion() string {
	version := s.Version[0]
	return fmt.Sprintf("%s %s", s.Name, version.Version)
}

func (s *Soft) FullVersion() string {
	version := s.Version[0]
	return fmt.Sprintf("%s %s_%s_%s", s.Name, version.Version, version.CreatedAt[:10], version.Status)
}

func (s *Soft) fullVersion() string {
	version := s.Version[0]
	return fmt.Sprintf("%s_%s_%s", version.Version, version.CreatedAt[:10], version.Status)
}

func (s *Soft) Info() string {
	version := s.Version[0]
	return fmt.Sprintf(
		"name: %s\nversion: %s\nalias: %s\nauthor: %s\nlatestlog: %s\nhash: %s\ngit signature: %s\nlasttime: %s\ncopyright: %s",
		s.Name, s.fullVersion(), s.Alias, s.Author, version.Log, version.Hash, version.GitHash,
		version.CreatedAt, s.Copyright,
	)
}

func (s *Soft) List() string {
	str := fmt.Sprintf("name: %s\nalias: %s\nauthor: %s\n", s.Name, s.Alias, s.Author)
	str += fmt.Sprintf("version:\n")
	for k, v := range s.Version {
		str += fmt.Sprintf("\tlog: %s\n\thash: %s\n\tgithash: %s\n\tlasttime: %s\n", v.Log, v.Hash, v.GitHash, v.CreatedAt)
		if k+1 != len(s.Version) {
			str += "\n"
		}
	}
	str += fmt.Sprintf("copyright: %s", s.Copyright)
	return str
}

func (s *Soft) JSON() string {
	str := fmt.Sprintf(`
	&Soft{
		Name:      "%s",
		Alias:     "%s",
		Author:    "%s",
		Copyright: "All rights reserved",
		Version: []Version{`, s.Name, s.Alias, s.Author)
	for _, v := range s.Version {
		str += fmt.Sprintf(`
			{
				Version:   "%s",
				Status:    %s,
				Hash:   "%s",
				GitHash:   "%s",
				CreatedAt: "%s",
				Log:       "%s",
			},`, v.Version, v.Status, v.Hash, v.GitHash, v.CreatedAt, v.Log)
	}

	str += `	
		},
}`
	return str
}
