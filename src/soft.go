package soft

import (
	"fmt"
)

type Soft struct {
	Name      string  `json:"name"`
	Alias     string  `json:"alias"`
	Author    string  `json:"author"`
	Version   Version `json:"version"`
	Copyright string  `json:"copyright"`
	Inherit   bool    `json:"inherit"`
}

func (s *Soft) SimpleVersion() string {
	return fmt.Sprintf("%s %s", s.Name, s.Version.Version)
}

func (s *Soft) FullVersion() string {
	return fmt.Sprintf("%s %s_%s_%s", s.Name, s.Version.Version, s.Version.updatedAt[:10], s.Version.Status)
}

func (s *Soft) fullVersion() string {
	return fmt.Sprintf("%s_%s_%s", s.Version.Version, s.Version.updatedAt[:10], s.Version.Status)
}

func (s *Soft) Info() string {
	return fmt.Sprintf(
		"name: %s\nversion: %s\nalias: %s\nauthor: %s\nlatestlog: %s\nsignature: %s\nlasttime: %s\ncopyright: %s",
		s.Name, s.fullVersion(), s.Alias, s.Author, s.Version.Log, s.Version.hash,
		s.Version.updatedAt, s.Copyright,
	)
}

func (s *Soft) JSON() string {
	return fmt.Sprintf(
		`&Soft{
			Name:      "%s",
			Alias:     "%s",
			Author:    "%s",
			Copyright: "%s",
			Version: Version{
				Version:   "%s",
				Status:    %s,
				updatedAt: "%s",
				Log:       "%s",
			},
			Inherit: true,
		}`,
		s.Name, s.Alias, s.Author, s.Copyright, s.Version.Version, s.Version.Status, s.Version.updatedAt, s.Version.Log,
	)
}
