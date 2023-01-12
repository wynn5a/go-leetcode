package solution1807

import "testing"

func TestCase1(t *testing.T) {
	actual := evaluate("(name)is(age)yearsold", [][]string{{"name", "bob"}, {"age", "two"}})
	expected := "bobistwoyearsold"
	if actual != expected {
		t.Errorf("Test failed, expected: %s, actual: %s \n", expected, actual)
	}
}

func TestCase2(t *testing.T) {
	actual := evaluate("hi(name)", [][]string{{"a", "b"}})
	expected := "hi?"
	if actual != expected {
		t.Errorf("Test failed, expected: %s, actual: %s \n", expected, actual)
	}
}

func TestCase3(t *testing.T) {
	actual := evaluate("(a)(a)(a)aaa", [][]string{{"a", "yes"}})
	expected := "yesyesyesaaa"
	if actual != expected {
		t.Errorf("Test failed, expected: %s, actual: %s \n", expected, actual)
	}
}

func TestCase4(t *testing.T) {
	actual := evaluate("gxzhwnx(ygk)ovmmhpfdyikllmq(jf)ewfuftlfpnldvvr(x)gu", [][]string{{"xwjusmky", "j"}, {"wmrgfhrg", "t"}, {"najfujbd", "n"}, {"nzcighye", "e"}, {"rjfqtqjs", "z"}, {"oufqgmod", "k"}, {"lybhjwsa", "g"}, {"rolgbcxd", "d"}, {"txpqekel", "e"}, {"rmeomfpo", "q"}, {"mwuqxfto", "o"}, {"fsppasht", "b"}, {"gfeoblsk", "z"}, {"afkhkiij", "p"}, {"tghzzgou", "e"}, {"ygafplfb", "z"}, {"cjsqqrtr", "z"}, {"sbsgsopm", "j"}, {"lcthvwgc", "o"}, {"x", "bitw"}, {"iefkefzx", "t"}, {"nosqpfql", "g"}, {"okxdmbxo", "i"}, {"psjtivfk", "e"}, {"uugtkfsw", "v"}, {"gjvevakd", "n"}, {"wagkprbb", "l"}, {"ygk", "qlz"}, {"dehlqief", "h"}, {"qjhrwnoo", "q"}, {"kimampwo", "j"}, {"mamrxrxo", "f"}, {"sslbxllo", "h"}, {"drkkpvru", "k"}, {"xsvjtkzs", "c"}, {"ntgilyux", "s"}, {"prlepcqk", "v"}, {"dwdnboel", "r"}, {"fdpfoeod", "r"}, {"lyyftify", "a"}, {"nokqdfmr", "n"}, {"pwkqzfsw", "p"}, {"hcxcbqqp", "s"}, {"jf", "kud"}, {"wuvrjcoy", "l"}, {"moepwhbx", "u"}, {"xskrszok", "p"}, {"rbzvcrho", "k"}, {"ldmlqxsw", "g"}, {"zgbhjksm", "h"}, {"pkncicvu", "k"}})
	expected := "gxzhwnxqlzovmmhpfdyikllmqkudewfuftlfpnldvvrbitwgu"
	if actual != expected {
		t.Errorf("Test failed, expected: %s, actual: %s \n", expected, actual)
	}
}
