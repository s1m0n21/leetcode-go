/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/4/1 17:40
 */

package _count_common_words_with_one_occurrence

import (
	"testing"

	"github.com/s1m0n21/leetcode-go/utils"
)

func TestAnswer(t *testing.T) {
	type input struct {
		worlds1, worlds2 []string
	}
	testCase := utils.NewTestCase(t, func(i input) int {
		return countWords(i.worlds1, i.worlds2)
	})

	testCase.SetAndRun(input{[]string{"leetcode", "is", "amazing", "as", "is"}, []string{"amazing", "leetcode", "is"}}, 2)
	testCase.SetAndRun(input{[]string{"b", "bb", "bbb"}, []string{"a", "aa", "aaa"}}, 0)
	testCase.SetAndRun(input{[]string{"a", "ab"}, []string{"a", "a", "a", "ab"}}, 1)
	testCase.SetAndRun(input{[]string{"bjxzvssdoq", "oom", "lxrrvf", "aoeselhvrnw", "awnornqyztqlza", "bjxqkapuvaw", "wibxruerngdzgjd", "rezrwdzvllpbjpnikhzraz", "pswmnrsepudx", "nlicjldpeia", "glg", "nllxfcjjitmsuugmr", "cl", "pysmpgjakkjnusfopphb", "zxlwcdjpn", "xktsfnchwrdesnf", "qptnoxxgrjmvr", "exlfwjfsbsirbbkyqjtinrrwuhh", "rqbnghajxygilgdjejopyuwyjqrx", "vrjkqsicuqoalqyaxkaaogxbf", "ixnlltqbpygmpjuspom", "izajsxotcbhzdnkujwgdzo", "b", "lighabre", "i", "ljqqbfddipvcooh", "hboedpepeeunx", "bkhzhiefammwqkhvampokd", "ptlozguwmyyp", "loeshsjgazzwvs", "kyrltbdzlymjxtvwiiq", "fk", "mbjpgwsahkgkehlcoqbhunqchxj", "nfyuvlrmiturheb", "cyqwsiysmoirurj", "sciqruywy", "podsrhmsozan", "nlyadkrxhdbup", "gdugldwghzt", "wpjm", "gjobdekmjisjgadkwwemnmco", "dkjdtimdghvlhuetxyaklk", "iwqylhdwqbwaqdouowoodhs", "mn"}, []string{"eeormvovhzslwsqgzthlgntgzc", "zfwownznh", "suxrkdbjdjjtbkjucsbyk", "u", "y", "lbjooktoctgwbbptiffytquha", "dcsxrghgpultkatbecjadbespvww", "vwduylshcpaiu", "rtcxwctvquaiuwkgvdx", "a", "szearxmdqcismljmihbtkcirztdnrc", "htgmuxtxdunsvfizb", "hybe", "nsegkgwcvopncmfpaahhhjeuqjosv", "jtarnnpppxtzmopixeijqqahkd", "hazcgrrnpourkyoeanodejiptne", "kurhokvhixihe", "ljwycewmecfqdhtxiokjn", "qgjzzvpyvwetlsvcsw", "aunns", "nwcnfrzzvxafkfjfnczummtubikji", "nipiygnvlfntgpxfedj", "mgnt", "xvjehufvaqouhztnmts", "sjtbrfjwtqxakqktxjaljrbwfoxvz", "dfeujeikfrtrpiafrgxvjlkpxtog", "u", "ggbcxoasodaqaazulrxjleecexey", "inedrgssajhpygfvozigohis", "pevxwgfzxebybfe", "cgy", "fnhvlx", "vxfybaebkoq", "xvhx", "mxbqjtanctljewwjjlbpkgbtsm", "mlwagamcikbcpuexhikmizp", "qeiomipvsoqlsnhylulirrcwtqga", "bwemqcgyusuauwlpbjjru", "iimcbidtndh", "lpjejlkmxtlbyvnscy", "dlspriicnyykdsyvswlgktavwloq", "dib", "qoptbduulgqwquvhdvmwdz", "xrtxghrbfrhpzduxeljnctgg", "schmbsaupayqnpchn", "kah", "itotymryqufqpozrwmvsl", "gurb", "xsyocxcmwvqmnmxthfemmu", "pkfdutm", "hkbwxwjxyuld", "ukdqszfjckdunnhpevw", "kqfwytdvnvjrchiwprcqkfntqticsc", "zjmsfwjddgjiypsmagdrujb", "gn", "ebncbjvhpbjecbrizdpv", "nbfehcktwswih", "sttmqcdmobdgtgkyxydyovahknjbsn", "sryyufrtocf", "eiicpwknxrzqylqpybhfd", "pey", "njimttradeoa", "wcogjdfr", "prva", "tyxdmxgw", "wluzocppg", "kzm", "wbyyperlkflaoxyxftzwxvmemof", "snzpclbulddnmmjmpdurcybo", "mowxgpmzojtmympmt", "uvtnojjahrovzmlukf", "sykhtgejlmbzshhneoyqr", "ib", "haqkkizidifykwijm", "csjtexnr", "yvgr", "vzcxbtlthrynnawxnkxdptp", "yvxrmscsckv"}}, 0)
}
