// Code generated by Soft build; DO NOT EDIT.
package soft

func NewSoft() *Soft {
	s :=
		&Soft{
			Name:      "xx-软件",
			Alias:     "别名",
			Author:    "作者",
			Copyright: "All rights reserved",
			Version: []Version{
				{
					Version:   "0.0.5",
					Status:    Base,
					Hash:      "dc02885ceec94e4e664f3679336eb631",
					GitHash:   "",
					CreatedAt: "2020.08.28 16:58:14",
					Log:       "update",
				},
				{
					Version:   "0.0.1",
					Status:    Base,
					Hash:      "46c9ab35431609ae4950d5ec6d1d7d2a",
					GitHash:   "6effa2b0971e2449c9fb223b2c966d48cb5ea09f",
					CreatedAt: "2019.12.03 16:54:52",
					Log:       "init",
				},
			},
		}
	return s
}
