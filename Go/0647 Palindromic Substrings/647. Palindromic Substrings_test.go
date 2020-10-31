package main

import "testing"

type testForm struct {
	s      string
	output int
}

func test(data []testForm, t *testing.T) {
	for _, e := range data {
		got := countSubstrings(e.s)
		if got != e.output {
			t.Errorf("at input %s got %d expected %d\n", e.s, got, e.output)
		}
	}
}

func Test(t *testing.T) {
	data := []testForm{
		{
			s:      "abc",
			output: 3,
		},
		{
			s:      "abba",
			output: 6,
		},
		{
			s:      "aaa",
			output: 6,
		},
		{
			s:      "",
			output: 0,
		},
		{
			s:      "yemtwlmqfxzngwvxfjrhwmqrapqzivbueymakkzwcqlrjpoozjemqfrjttnpjrkamrsbtsfuzyqnreanyzqppxzqhcxhqnrupcbqjwjczatnexuzaljmukxiowvekzehixzaxqhgpkmtdecwepzyomewbhkxdzgkkrvnuapnsrnkxydzqqkppzptqbwveyfpdczzexnsqlylymbcgzvjzwoibkshsrschlyzsofjysppcpnrmucnrmkrbvtztjbhiutbxbqqjgetaeulfdhwwphzvsdogefijdzjjsvqsvkxfrlcgvxoxbgwtnjcpwxrlkoxykwshpkrrkqyyyhztognhpzawcyujfthvxziqjdutnlqukasztktsmbjervzowfnzauzetegoimcsnofjxdlzlelkwzbpjsusufkhacmcftudewzauuymjzzgfwafazfuuuabhtsxyvgnzkvjifazpobqfpbzgtxjmkwcmudecbboetcgfttivjwspymqbewpdgeizkbodclnwldpldhqrnekhmnsjnqtvfnkbpzxtwvfpezjoetotmpwaxeivjqzsmkfhjugsvtjzuidzeqramyxbfuslwjcnawxipnbhxbnmgvvchhvpbgpentwqhttukopfpcesvkylurbipkvdjupzeapxmdxyvswgkzxjeaxbmnintpaidsedwoterqwdnpwypkrmhwhvynawzjwlxqgcnwajrrukbdqxnjoagoifrgshlbgbefeufwlulcnsbgkxuyxyewznixkfniwipcmzlbvxmzraapuntclpzfltntbswelaxkltdazshnptihkfqtgdmsehqnregekbsbreecrqwpkednqvquldeyulukuzaewoyzurrcpeanjujhqelurledttzealbpikbgwxibflzbusfbhwfpdcfuxpbckpoliendrfohvobguhnjsswyregfwwvvdjcrbwohxgjtbpzstiaf",
			output: 1069,
		},
	}
	test(data, t)
}
