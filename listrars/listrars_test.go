package listrars

import (
	"testing"
)

type testpair struct {
	filename  string
	extracted bool
}

var tests = []testpair{
	{"imarar.rar", true},
	{"/home/user/bar/foo.rar", true},
	{"/home/lisa/give/meTHIS.RaR", true},
	{"/home/lisa/give/meTHIS01.RAR", true},
	{"/tmp/yes_part_01.RAR", true},
	{"/tmp/yes_part01.RAR", true},
	{"/tmp/yes.dummy.dumb2part001.rar", true},
	{"/tmp/should.soulder.part1.rar", true},
	{"/tmp/should.soulder.part001.rar", true},
	{"/tmp/should.soulder.part002.rar", false},
	{"nope.part2.rar", false},
	{"jusi.part0010.rar", false},
	{"notme.pArt03.RAR", false},
	{"iamadirectory", false},
	{"andiamatextfile.txt", false},
	{"withPart.part01.txt", false},
}

func (tp testpair) Name() string {
	return tp.filename
}

func TestGetRars(t *testing.T) {
	intf := make([]interface{}, len(tests))
	for i, v := range tests {
		intf[i] = v
	}
	extractedFiles, err := GetRars(intf)
	if err != nil {
		t.Error(err)
	}

	for _, testElem := range tests {
		expected := testElem.extracted
		found := false
		for _, e := range extractedFiles {
			if testElem.Name() == e {
				found = true
				break
			}
		}
		if found && !expected {
			t.Error("Did not expect to extract ", testElem.Name())
		} else if !found && expected {
			t.Error("Extracted ", testElem.Name(), ", but did not expect to do so")
		}
	}

}
