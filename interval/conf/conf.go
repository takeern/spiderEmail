package conf

type SendInfo struct {
	Ac		string
	Ps 		string
	Host	string
}

const DB_USER = "root"
const DB_PSAAWD = "maskTakeern"
const DB_IP = "47.103.12.134"
// const DB_IP = "127.0.0.1"
const DB_PORT = "3306"
const DB_DATABASE = "spider"
const DB_CHARSET = "utf8"

const WAIT_SEND_EMAIL_TIME = 60 * 60 * 3 // 60 * 50
const WAIT_SPIDER_TIME = 60 * 2
const SPIDER_TIMEOUT = 15 
const HTTP_TRY_REQUEST_TIMES = 2
const RETRY_REGISTER_TIMES = 10

const HOST_PORT = "6010"
const MASTER_HOST = "47.103.12.134:" + HOST_PORT
const MASTER_TOKEN = "random"
const SLAVE_PORT = "6011"

const DB_URL = "http://wwwijetchorg/"
const SPIDER_URL = "http://dpi-proceedings.com/index.php/dtem/article/download/31137/29718"

const (
	RegisterCodeSuccess = 0
	RegisterMsgSuccess = " register Success "
	RegisterCodeError = -1
	RegisterMsgErrorRepeat = " this ip Repeat registered "
)

const (
	Retry_Spider_Times = 20
	Retry_Send_Email_Times = 10
)

var SendList = [...][4]SendInfo{
	{
		// {
		// 	Ac: "tq123456@foxmail.com",
		// 	Ps: "dazjvvhpevbqbbah",
		// 	Host: "smtp.qq.com:587",
		// },
		{
			Ac: "publish_house@yeah.net",
			Ps: "Tq6614118",
			Host: "smtp.yeah.net:25",
		},
		{
			Ac: "publish_house@sohu.com",
			Ps: "tq6614118",
			Host: "smtp.sohu.com:25",
		},
		{
			Ac:	"takeern@163.com",
			Ps: "tq123456",
			Host: "smtp.163.com:25",
		},
		{
			Ac: "publishhouse@sina.com",
			Ps: "9822fc8955387b92",
			Host: "smtp.sina.com:587",
		},
	},
	{
		{
			Ac: "yuyuanzhou2010@sina.com",
			Ps: "yuyuanzhou2010",
			Host: "smtp.sina.com:587",
		},
		{
			Ac: "suncangjun2010@163.com",
			Ps: "Abcd123",
			Host: "smtp.163.com:25",
		},
		{
			Ac: "lijinman2010@126.com",
			Ps: "Abcd123",
			Host: "smtp.126.com:25",
		},
	},
}

var RecieveList = [...]string{
	"tq123456@foxmail.com",
	"870307181@qq.com",
	"89070310@qq.com",
}

const (
	SEND_EMAIL = 1000
	SPIDER_EMAIL = 1001
)


var EmailModalList = [...]string{
	`
	<div id="qm_con_body"><div id="mailContentContainer" class="qmbox qm_con_body_content qqmail_webmail_only" style=""><div><p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="font-size:14.0pt;mso-bidi-font-size:
	10.5pt;font-family:宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:
	0pt">Journal of Innovation and Social Science Research<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="font-size:14.0pt;mso-bidi-font-size:
	10.5pt;font-family:宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:
	0pt">(</span><span lang="EN-US"><a href="http://www.jissr.net/" rel="noopener" target="_blank"><span style="mso-bidi-font-size:10.5pt;font-family:宋体;mso-bidi-font-family:宋体;
	mso-font-kerning:0pt">http://www.jissr.net</span></a></span><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:宋体;mso-bidi-font-family:宋体;
	color:black;mso-font-kerning:0pt">)is a high quality open access peer-reviewed
	research journal that publishes genuine and innovative research articles
	submitted by renowned scientists and researchers in the field of engineering. JISSR
	is dedicated to research practice, advocacy, education, and policy; hence JISSR
	provides a platform for researchers, academicians, professional, practitioners
	and students to impart and share knowledge in the form of high quality
	empirical and theoretical research papers, case studies, and literature
	reviews.<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt"><o:p>&nbsp;</o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">JISSR is
	interested in empirical, theoretical, methodological, and practice-oriented
	articles covering topics relevant to the field of engineering. Particular
	consideration is given to empirical articles using quantitative, qualitative,
	and mixed methodology. JISSR is widely preferred and popular among all research
	communities in the world.<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt"><o:p>&nbsp;</o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">The Editorial
	Board of JISSR encourages authors to submit manuscripts that contribute to
	knowledge through research and examination of methodology; and develop and advancement
	of theories that contribute to knowledge; in other to promote education and
	training of scientists.<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt"><o:p>&nbsp;</o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">The Editorial
	Board welcomes and encourages research communities worldwide to share their new
	ideas using this journal as a platform. This will help in the enhancement of
	knowledge and potentials for better prospects. Types of manuscript welcomed
	include high quality theoretical and empirical original research papers, case
	studies, review papers, conceptual framework, analytical and simulation models,
	technical note from researchers, academicians, professional, practitioners and
	students from all over the world. The scientific report should be of quality
	context that the Editorial Board would consider to be of interest to an
	international readership.<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt"><o:p>&nbsp;</o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">Additional
	information about the journal, including instructions for authors, is available
	on <a href="http://www.jissr.net." rel="noopener" target="_blank">www.jissr.ne<wbr>t.</a> <o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt"><o:p>&nbsp;</o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">Manuscripts should
	be submitted electronically through our submission portal or through the
	following e-mail addresses: <a href="mailto:editor@jissr.net" rel="noopener" target="_blank">editor@jissr<wbr>.net</a> <o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">We hope you will
	consider submitting your research to Journal of Innovation and Social Science
	Research.<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt"><o:p>&nbsp;</o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">Best regards,<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt"><o:p>&nbsp;</o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">Prof. Li,<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">For: Editorial
	Board Committee,<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">Journal of
	Innovation and Social Science Research<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">http:// </span><span lang="EN-US"><a href="http://www.jissr.net/" rel="noopener" target="_blank"><span style="mso-bidi-font-size:10.5pt;
	font-family:宋体;mso-bidi-font-family:宋体;mso-font-kerning:0pt">www.jissr.net</span></a></span><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:宋体;mso-bidi-font-family:
	宋体;color:black;mso-font-kerning:0pt"> <o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt"><o:p>&nbsp;</o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">unsubscribe:<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">Please reply with
	UNSUBSCRIBE as subject.<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">---------------------------------------------------------------------------------------------------------------------<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;line-height:17.85pt;
	mso-pagination:widow-orphan;tab-stops:45.8pt 91.6pt 137.4pt 183.2pt 229.0pt 274.8pt 320.6pt 366.4pt 412.2pt 458.0pt 503.8pt 549.6pt 595.4pt 641.2pt 687.0pt 732.8pt;
	background:white"><span lang="EN-US" style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:宋体;color:black;mso-font-kerning:0pt">Disclaimer: The
	CAN-SPAM Act of 2003 (Controlling the Assault of Non-Solicited Pornography and
	Marketing Act) establishes requirements for those who send commercial email;
	spells out penalties for spammers and companies whose products are advertised
	in spam, if they violate the law; and gives consumers the right to ask mailers
	to stop spamming them. The above mail is in accordance with the CAN-SPAM Act of
	2003. There are no deceptive subject lines and it is a manual process through
	our efforts on World Wide Web. If you send us an UNSUBSCRIBE email, we shall
	ensure that you do not receive any such mails from us again.<o:p></o:p></span></p>

	<p class="MsoNormal"><span lang="EN-US"><o:p>&nbsp;</o:p></span></p><br></div><style type="text/css">.qmbox style, .qmbox script, .qmbox head, .qmbox link, .qmbox meta {display: none !important;}</style></div></div>
	`,
	`
	<div id="qm_con_body"><div id="mailContentContainer" class="qmbox qm_con_body_content qqmail_webmail_only" style=""><div><p class="MsoNormal" align="center" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:center;mso-pagination:widow-orphan;background:white"><b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	color:black;background:#EFF5FB ">Call for Paper | Volume 06 | Issue-11 / </span></b><span lang="EN-US" style=" font-size:12.0pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt ">Journal of
	Innovation and Social Science Research - JISSR<b><o:p></o:p></b></span></p>

	<p class="MsoNormal" align="center" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:center;mso-pagination:widow-orphan;background:white"><b><span lang="EN-US" style=" font-size:12.0pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt ">Call for Paper |
	Volume 06 | Issue-11</span></b><span lang="EN-US" style=" font-size:12.0pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:auto;
	mso-pagination:widow-orphan;background:white"><span lang="EN-US" style=" font-size:12.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">Journal of Innovation and Social Science
	Research - JISSR having online ISSN 2591 6890. JISSR is a leading Open
	Access, Peer-Reviewed International Journal which provides rapid publication of
	your research articles and aims to promote the theory and practice along with
	knowledge sharing between researchers, developers, engineers, students, and
	practitioners working in and around the world in many areas like Sciences,
	Technology, Innovation, Engineering, Agriculture, Management and many more and
	it is recommended by all Universities, review articles and short communications
	in all subjects.<o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan;background:white"><b><span lang="EN-US" style=" font-size:12.0pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt ">JISSR Journal
	Details:</span></b><span lang="EN-US" style=" font-size:12.0pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<table class="MsoNormalTable" border="1" cellspacing="0" cellpadding="0" width="102%" style="width:102.08%;mso-cellspacing:0cm;background:white;mso-yfti-tbllook:
	1184;mso-padding-alt:0cm 0cm 0cm 0cm">
	<tbody><tr style="mso-yfti-irow:0;mso-yfti-firstrow:yes;height:13.55pt">
	<td width="35%" style="width:35.0%;padding:3.75pt 3.75pt 3.75pt 3.75pt;
	height:13.55pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:9.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:black;mso-font-kerning:0pt ">ISSN:<o:p></o:p></span></p>
	</td>
	<td style="padding:3.75pt 3.75pt 3.75pt 3.75pt;height:13.55pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:9.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:black;mso-font-kerning:0pt "><span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="2591-6890">2591-6890</span><o:p></o:p></span></p>
	</td>
	</tr>
	<tr style="mso-yfti-irow:1;height:14.1pt">
	<td width="35%" style="width:35.0%;padding:3.75pt 3.75pt 3.75pt 3.75pt;
	height:14.1pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:9.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:black;mso-font-kerning:0pt ">Journal Type:<o:p></o:p></span></p>
	</td>
	<td style="padding:3.75pt 3.75pt 3.75pt 3.75pt;height:14.1pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:9.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:black;mso-font-kerning:0pt ">International Open Access<o:p></o:p></span></p>
	</td>
	</tr>
	<tr style="mso-yfti-irow:2;height:27.1pt">
	<td width="35%" style="width:35.0%;background:#D0CECE;mso-background-themecolor:
	background2;mso-background-themeshade:230;padding:3.75pt 3.75pt 3.75pt 3.75pt;
	height:27.1pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:14.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:red;mso-font-kerning:0pt ">Frequency:<o:p></o:p></span></p>
	</td>
	<td style="background:#D0CECE;mso-background-themecolor:background2;
	mso-background-themeshade:230;padding:3.75pt 3.75pt 3.75pt 3.75pt;height:
	27.1pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:14.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:red;mso-font-kerning:0pt ">Monthly publication<o:p></o:p></span></p>
	</td>
	</tr>
	<tr style="mso-yfti-irow:3;height:27.1pt">
	<td width="35%" style="width:35.0%;background:#D0CECE;mso-background-themecolor:
	background2;mso-background-themeshade:230;padding:3.75pt 3.75pt 3.75pt 3.75pt;
	height:27.1pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:14.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:red;mso-font-kerning:0pt ">Review <o:p></o:p></span></p>
	</td>
	<td style="background:#D0CECE;mso-background-themecolor:background2;
	mso-background-themeshade:230;padding:3.75pt 3.75pt 3.75pt 3.75pt;height:
	27.1pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:14.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:red;mso-font-kerning:0pt ">Within <span style="border-bottom:1px dashed #ccc;" t="5" times="">7-8</span> Days<o:p></o:p></span></p>
	</td>
	</tr>
	<tr style="mso-yfti-irow:4;height:13.55pt">
	<td width="35%" style="width:35.0%;padding:3.75pt 3.75pt 3.75pt 3.75pt;
	height:13.55pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:9.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:black;mso-font-kerning:0pt ">Paper Publish:<o:p></o:p></span></p>
	</td>
	<td style="padding:3.75pt 3.75pt 3.75pt 3.75pt;height:13.55pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:9.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:black;mso-font-kerning:0pt ">Within <span style="border-bottom:1px dashed #ccc;" t="5" times="">2-3</span> Days after
	submitting the all documents<o:p></o:p></span></p>
	</td>
	</tr>
	<tr style="mso-yfti-irow:5;height:34.75pt">
	<td width="35%" style="width:35.0%;padding:3.75pt 3.75pt 3.75pt 3.75pt;
	height:34.75pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:9.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:black;mso-font-kerning:0pt ">Research Area:<o:p></o:p></span></p>
	</td>
	<td style="padding:3.75pt 3.75pt 3.75pt 3.75pt;height:34.75pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:9.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:black;mso-font-kerning:0pt ">Engineering, Technology,
	Pharmacy, Management, Biological Science, Applied Mathematics, Physics and
	Chemistry, Commerce, Arts, Medical Science and many more<o:p></o:p></span></p>
	</td>
	</tr>
	<tr style="mso-yfti-irow:6;height:27.1pt">
	<td width="35%" style="width:35.0%;background:#D0CECE;mso-background-themecolor:
	background2;mso-background-themeshade:230;padding:3.75pt 3.75pt 3.75pt 3.75pt;
	height:27.1pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:15.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:red;mso-font-kerning:0pt ">Publish Fee<o:p></o:p></span></p>
	</td>
	<td style="background:#D0CECE;mso-background-themecolor:background2;
	mso-background-themeshade:230;padding:3.75pt 3.75pt 3.75pt 3.75pt;height:
	27.1pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:15.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:red;mso-font-kerning:0pt ">150$/1050RMB<o:p></o:p></span></p>
	</td>
	</tr>
	<tr style="mso-yfti-irow:7;mso-yfti-lastrow:yes;height:27.1pt">
	<td width="35%" style="width:35.0%;background:#D0CECE;mso-background-themecolor:
	background2;mso-background-themeshade:230;padding:3.75pt 3.75pt 3.75pt 3.75pt;
	height:27.1pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:15.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:red;mso-font-kerning:0pt ">Index in <o:p></o:p></span></p>
	</td>
	<td style="background:#D0CECE;mso-background-themecolor:background2;
	mso-background-themeshade:230;padding:3.75pt 3.75pt 3.75pt 3.75pt;height:
	27.1pt">
	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;text-align:left;
	mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:15.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-bidi-font-family:
	Helvetica;color:red;mso-font-kerning:0pt ">CNKI<o:p></o:p></span></p>
	</td>
	</tr>
	</tbody></table>

	<p class="MsoNormal" style="mso-margin-top-alt:auto;mso-pagination:widow-orphan;
	background:white"><b><i><span lang="EN-US" style=" font-size:12.0pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;mso-font-kerning:
	0pt ">​</span></i></b><b><span lang="EN-US" style=" font-size:12.0pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;mso-font-kerning:
	0pt ">Editorial Office<o:p></o:p></span></b></p>

	<p class="MsoNormal" style="mso-margin-top-alt:auto;mso-pagination:widow-orphan;
	background:white"><span lang="EN-US" style=" font-size:12.0pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt ">Journal of
	Innovation and Social Science Research <o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;line-height:20.4pt;mso-pagination:widow-orphan;background:
	white"><span lang="EN-US"><a href="http://www.jissr.net/" rel="noopener" target="_blank"><span style=" font-size:
	12.0pt ; ; ; ;;mso-fareast-font-family:宋体;
	mso-font-kerning:0pt ">www.jissr.net</span></a></span><span lang="EN-US" style=" font-size:12.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;line-height:20.4pt;mso-pagination:widow-orphan;background:
	white"><span lang="EN-US" style=" font-size:12.0pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt ">Email: </span><span lang="EN-US"><a href="mailto:editor@jissr.net" rel="noopener" target="_blank"><span style=" font-size:12.0pt ; ; ; ;;mso-fareast-font-family:宋体;mso-font-kerning:
	0pt ">editor@jissr.net</span></a><o:p></o:p></span></p><br></div><style type="text/css">.qmbox style, .qmbox script, .qmbox head, .qmbox link, .qmbox meta {display: none !important;}</style></div></div>
	`,
	`
	<div id="qm_con_body"><div id="mailContentContainer" class="qmbox qm_con_body_content qqmail_webmail_only" style=""><div><p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><b><span style="mso-bidi-font-size:
	10.5pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:
	0pt">尊敬的作者：您好！</span></b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">&nbsp; &nbsp; &nbsp; &nbsp;</span></b><b><span style="mso-bidi-font-size:10.5pt;font-family:宋体;mso-bidi-font-family:Arial;
	color:black;mso-font-kerning:0pt">我社现有</span></b><b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:red;mso-font-kerning:0pt ">14</span></b><b><span style="mso-bidi-font-size:
	10.5pt;font-family:宋体;mso-bidi-font-family:Arial;color:red;mso-font-kerning:
	0pt">本国外英文期刊，</span></b><b><span style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">在对外征稿中，</span></b><b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt ">1</span></b><b><span style="mso-bidi-font-size:10.5pt;font-family:宋体;mso-bidi-font-family:Arial;
	color:black;mso-font-kerning:0pt">个月内快速见刊，有正规发票，</span></b><b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:red;mso-font-kerning:0pt ">1000</span></b><b><span style="mso-bidi-font-size:
	10.5pt;font-family:宋体;mso-bidi-font-family:Arial;color:red;mso-font-kerning:
	0pt">一篇，知网检索</span></b><b><span style="mso-bidi-font-size:10.5pt;font-family:
	宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">。</span></b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">&nbsp; &nbsp; &nbsp; &nbsp;</span></b><b><span style="mso-bidi-font-size:10.5pt;font-family:宋体;mso-bidi-font-family:Arial;
	color:black;mso-font-kerning:0pt">您可以进入期刊官网查看，也可以添加</span></b><b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:red;mso-font-kerning:0pt ">QQ <span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="980418101">980418101</span></span></b><b><span style="mso-bidi-font-size:10.5pt;font-family:宋体;mso-bidi-font-family:Arial;
	color:black;mso-font-kerning:0pt">了解详细期刊信息。</span></b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">&nbsp; &nbsp; &nbsp; &nbsp;</span></b><b><span style="mso-bidi-font-size:10.5pt;font-family:宋体;mso-bidi-font-family:Arial;
	color:black;mso-font-kerning:0pt">现在投稿可以申请优惠，投稿时请在邮件附言</span></b><b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt ">“</span></b><b><span style="mso-bidi-font-size:10.5pt;font-family:宋体;mso-bidi-font-family:Arial;
	color:black;mso-font-kerning:0pt">李编辑推荐</span></b><b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">”</span></b><b><span style="mso-bidi-font-size:
	10.5pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:
	0pt">。</span></b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;mso-font-kerning:
	0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><b><span style="font-size:
	22.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:
	0pt">加拿大期刊</span></b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:12.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">世纪科学出版公司是一家多学科的学术期刊出版商，《创新与社会科学研究杂志（<span lang="EN-US">JISSR</span>）》、《国际电力与能源工程杂志（<span lang="EN-US">IJPEE</span>）》是其出版物。两个期刊独立运作，并有自己的编辑委员会。</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><b><span lang="EN-US" style="font-size:12.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;
	mso-font-kerning:0pt">Journal of Innovation and Social Science Research (JISSR)</span></b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><b><span lang="EN-US" style="font-size:12.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;
	mso-font-kerning:0pt">ISSN: <span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="2591-6890">2591-6890</span></span></b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><b><span style="font-size:
	12.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:
	0pt">创新与社会科学研究杂志</span></b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">官网：</span><span lang="EN-US"><a href="http://www.jissr.net/" rel="noopener" target="_blank"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:#064977;mso-font-kerning:0pt">www.jissr.net</span></a></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">投稿邮箱：</span><span lang="EN-US"><a href="mailto:editor@jissr.net" rel="noopener" target="_blank"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:#064977;mso-font-kerning:0pt">editor@jissr.net</span></a></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">征稿范围：经济学，政治学，法学，伦理学，历史学，社会学，心理学，教育学，管理学，人类学，民俗学，新闻学，传播学，行为科学，女性主义学。</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><b><span lang="EN-US" style="font-size:12.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;
	mso-font-kerning:0pt">International Journal of Power and Energy Engineering
	(IJPEE)</span></b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;mso-font-kerning:
	0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><b><span lang="EN-US" style="font-size:12.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;
	mso-font-kerning:0pt">ISSN: <span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="2287-6464">2287-6464</span></span></b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><b><span style="font-size:
	12.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:
	0pt">国际电力与能源工程杂志</span></b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">官网：</span><span lang="EN-US"><a href="http://www.ijpee.net/" rel="noopener" target="_blank"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:#064977;mso-font-kerning:0pt">www.ijpee.net</span></a></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">投稿邮箱：</span><span lang="EN-US"><a href="mailto:editor@ijpee.net" rel="noopener" target="_blank"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:#064977;mso-font-kerning:0pt">editor@ijpee.net</span></a></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">征稿范围：能源科学技术，能源化学，能源地理学，能源计算与测量，储能技术，节能技术，一次能源，二次能源，能源系统工程，能源科学技术其他学科，工程热物理，动力机械工程，动力与电气工程其他学科，电力电子技术，能源技术经济学，能源经济学</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:14.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">知网检索页面：</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style="font-size:14.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;
	mso-font-kerning:0pt"><a href="http://scholar.cnki.net/WebPress/brief.aspx?dbcode=SJLF" rel="noopener" target="_blank">http://schol<wbr>ar.cnki.net/<wbr>WebPress/bri<wbr>ef.aspx?dbco<wbr>de=SJLF</a></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><b><span style="font-size:
	22.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:
	0pt">英国期刊</span></b><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style="font-size:14.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;
	mso-font-kerning:0pt">Bryan House</span><span style="font-size:14.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">出版社致力于推进全球科研与学术的研究和推广，并为实现科学各学科的知识拓展与推广而努力。我们努力为各科研机构、高校、教师、学生推荐优质学术资源。目前共有</span><span lang="EN-US" style="font-size:14.0pt;font-family:宋体;mso-bidi-font-family:Arial;
	color:red;background:yellow;mso-font-kerning:0pt">12</span><span style="font-size:14.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:red;
	background:yellow;mso-font-kerning:0pt">本</span><span style="font-size:14.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">专业期刊。</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:14.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">官网：</span><span lang="EN-US"><a href="http://www.bryanhousepub.org/" rel="noopener" target="_blank"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:#064977;mso-font-kerning:0pt">www.bryanhousepub.org</span></a></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">投稿邮箱：</span><span lang="EN-US"><a href="mailto:submission@bryanhousepub.org" rel="noopener" target="_blank"><span style="font-size:11.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:#064977;
	mso-font-kerning:0pt">submission@bryanhousepub.org</span></a></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:14.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">知网检索页面：</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US"><a href="http://scholar.cnki.net/WebPress/brief.aspx?dbcode=SJZN" rel="noopener" target="_blank"><span style="font-size:14.0pt;font-family:宋体;mso-bidi-font-family:Arial;mso-font-kerning:
	0pt">http://scholar.cnki.net/WebPress/brief.aspx?dbcode=SJZN</span></a></span><span lang="EN-US" style="font-size:14.0pt;font-family:宋体;mso-bidi-font-family:Arial;
	color:black;mso-font-kerning:0pt"><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">Journal of Research in Vocational
	Education (JRVE) ISSN: <span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="2408-5170">2408-5170</span></span><span lang="EN-US" style=" mso-bidi-font-size:
	10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">职业教育研究杂志</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">Journal of Progress in Civil Engineering
	(JPCE) ISSN: <span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="2322-0856">2322-0856</span></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">土木工程进展杂志</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">Journal of Petroleum and Mining
	Engineering (JPME) ISSN: 1110-6506</span><span lang="EN-US" style=" mso-bidi-font-size:
	10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">石油与采矿工程杂志</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">Journal of Contemporary Medical Practice
	(JCMP) ISSN: <span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="2006-2745">2006-2745</span></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">当代医学实践杂志</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:red;mso-font-kerning:0pt ">Journal of Research in Science and
	Engineering (JRSE) ISSN: <span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="1656-1996">1656-1996</span></span><span lang="EN-US" style=" mso-bidi-font-size:
	10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:red;mso-font-kerning:0pt">科学与工程研究杂志</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:red;mso-font-kerning:0pt">工程科学、自然科学类理工科都可投稿</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">Journal of Educational Research and
	Policies (JERP) ISSN: <span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="2006-1137">2006-1137</span></span><span lang="EN-US" style=" mso-bidi-font-size:
	10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">教育研究与政策杂志</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">Journal of Metallurgy and Materials
	Engineering (JMME) ISSN: <span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="2006-1919">2006-1919</span></span><span lang="EN-US" style=" mso-bidi-font-size:
	10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">冶金与材料工程杂志</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">International Journal of Environment
	Research (IJER) ISSN: <span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="1595-4080">1595-4080</span></span><span lang="EN-US" style=" mso-bidi-font-size:
	10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">国际环境研究杂志</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">Journal of Global Economy, Business and
	Finance (JGEBF) ISSN: <span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="2141-5595">2141-5595</span></span><span lang="EN-US" style=" mso-bidi-font-size:
	10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">全球经济、商业和金融杂志（</span><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt ">JGEBF</span><span style="font-size:11.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;
	mso-font-kerning:0pt">）</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">Journal of Energy Science (JES) ISSN:
	<span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="1689-8338">1689-8338</span></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;mso-font-kerning:
	0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">能源科学杂志</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">Journal of Social Science and Humanities
	(JSSH) ISSN: <span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="1811-1564">1811-1564</span></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">社会科学与人文杂志（</span><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt ">JSSH</span><span style="font-size:11.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;
	mso-font-kerning:0pt">）</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;mso-fareast-font-family:
	宋体;color:black;mso-font-kerning:0pt ">Journal of Agriculture and Horticulture
	(JAH) ISSN: <span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="1711-8239">1711-8239</span></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="mso-margin-top-alt:auto;mso-margin-bottom-alt:
	auto;text-align:left;mso-pagination:widow-orphan"><span style="font-size:11.0pt;
	font-family:宋体;mso-bidi-font-family:Arial;color:black;mso-font-kerning:0pt">农业与园艺杂志（</span><span lang="EN-US" style=" font-size:11.0pt ; ; ; ;;
	mso-fareast-font-family:宋体;color:black;mso-font-kerning:0pt ">JAH</span><span style="font-size:11.0pt;font-family:宋体;mso-bidi-font-family:Arial;color:black;
	mso-font-kerning:0pt">）</span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;mso-fareast-font-family:宋体;color:black;
	mso-font-kerning:0pt "><o:p></o:p></span></p>

	<p class="MsoNormal"><span lang="EN-US"><o:p>&nbsp;</o:p></span></p><br></div><style type="text/css">.qmbox style, .qmbox script, .qmbox head, .qmbox link, .qmbox meta {display: none !important;}</style></div></div>
	`,
	`
		<div id="qm_con_body"><div id="mailContentContainer" class="qmbox qm_con_body_content qqmail_webmail_only" style=""><div><p class="MsoNormal"><span lang="EN-US" style="font-size:16.0pt;color:red;
	background:black;mso-highlight:black">1000</span><span style="font-size:16.0pt;
	font-family:宋体;mso-ascii-font-family:Calibri;mso-ascii-theme-font:minor-latin;
	mso-fareast-font-family:宋体;mso-fareast-theme-font:minor-fareast;mso-hansi-font-family:
	Calibri;mso-hansi-theme-font:minor-latin;color:red;background:black;mso-highlight:
	black">一篇。最快</span><span lang="EN-US" style="font-size:16.0pt;color:red;
	background:black;mso-highlight:black">7</span><span style="font-size:16.0pt;
	font-family:宋体;mso-ascii-font-family:Calibri;mso-ascii-theme-font:minor-latin;
	mso-fareast-font-family:宋体;mso-fareast-theme-font:minor-fareast;mso-hansi-font-family:
	Calibri;mso-hansi-theme-font:minor-latin;color:red;background:black;mso-highlight:
	black">天上知网。</span><span style="font-size:16.0pt;font-family:宋体;mso-ascii-font-family:
	Calibri;mso-ascii-theme-font:minor-latin;mso-fareast-font-family:宋体;mso-fareast-theme-font:
	minor-fareast;mso-hansi-font-family:Calibri;mso-hansi-theme-font:minor-latin;
	color:#E7E6E6;mso-themecolor:background2;background:black;mso-highlight:black">知网检索后付款</span><span lang="EN-US" style="font-size:16.0pt;color:red;background:black;mso-highlight:
	black">. QQ</span><span style="font-size:16.0pt;font-family:宋体;mso-ascii-font-family:
	Calibri;mso-ascii-theme-font:minor-latin;mso-fareast-font-family:宋体;mso-fareast-theme-font:
	minor-fareast;mso-hansi-font-family:Calibri;mso-hansi-theme-font:minor-latin;
	color:red;background:black;mso-highlight:black">：</span><span lang="EN-US" style="font-size:16.0pt;color:red;background:black;mso-highlight:black"><span style="border-bottom:1px dashed #ccc;z-index:1" t="7" onclick="return false;" data="980418101">980418101</span></span><span lang="EN-US" style="font-size:16.0pt;color:red"><span style="mso-spacerun:yes">&nbsp; 
	</span><o:p></o:p></span></p>

	<p class="MsoNormal"><span lang="EN-US" style="font-size:14.0pt;color:red">14</span><span style="font-size:14.0pt;font-family:宋体;mso-ascii-font-family:Calibri;
	mso-ascii-theme-font:minor-latin;mso-fareast-font-family:宋体;mso-fareast-theme-font:
	minor-fareast;mso-hansi-font-family:Calibri;mso-hansi-theme-font:minor-latin;
	color:red">本专业期刊</span><span style="font-family:宋体;mso-ascii-font-family:Calibri;
	mso-ascii-theme-font:minor-latin;mso-fareast-font-family:宋体;mso-fareast-theme-font:
	minor-fareast;mso-hansi-font-family:Calibri;mso-hansi-theme-font:minor-latin">，能源电力、人文社科、医学、教育、石油天然气、材料、机械、土木等专业都有对应的期刊。</span><span lang="EN-US"><o:p></o:p></span></p>

	<p class="MsoNormal" align="left" style="text-align:left;mso-pagination:widow-orphan"><span lang="EN-US"><a href="https://mail.163.com/js6/www.jissr.net" rel="noopener" target="_blank"><span style="font-size:16.0pt;mso-bidi-font-family:Arial">www.jissr.net</span></a></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	color:black "><o:p></o:p></span></p>

	<p class="MsoNormal"><span lang="EN-US"><a href="http://www.bryanhousepub.org/" rel="noopener" target="_blank"><span style=" mso-bidi-font-size:10.5pt ; ; ; ; ">http://www.bryanhousepub.org</span></a></span><span lang="EN-US" style=" mso-bidi-font-size:10.5pt ; ; ; ;;
	color:black "><o:p></o:p></span></p>

	<p class="MsoNormal"><span lang="EN-US"><o:p>&nbsp;</o:p></span></p>

	<p class="MsoNormal"><span style="font-size:16.0pt;font-family:宋体;mso-ascii-font-family:
	Calibri;mso-ascii-theme-font:minor-latin;mso-fareast-font-family:宋体;mso-fareast-theme-font:
	minor-fareast;mso-hansi-font-family:Calibri;mso-hansi-theme-font:minor-latin;
	color:red">百度</span><span lang="EN-US" style="font-size:16.0pt;color:red">JISSR</span><span style="font-size:16.0pt;font-family:宋体;mso-ascii-font-family:Calibri;
	mso-ascii-theme-font:minor-latin;mso-fareast-font-family:宋体;mso-fareast-theme-font:
	minor-fareast;mso-hansi-font-family:Calibri;mso-hansi-theme-font:minor-latin;
	color:red">你就知道</span><span lang="EN-US" style="font-size:16.0pt;color:red"><o:p></o:p></span></p>

	<p class="MsoNormal"><span lang="EN-US" style="font-size:16.0pt;color:red"><o:p>&nbsp;</o:p></span></p>

	<p class="MsoNormal"><span style="font-family:宋体;mso-ascii-font-family:Calibri;
	mso-ascii-theme-font:minor-latin;mso-fareast-font-family:宋体;mso-fareast-theme-font:
	minor-fareast;mso-hansi-font-family:Calibri;mso-hansi-theme-font:minor-latin">职称评定，刷奖学金、项目结题、毕业保底、保研统统有效。</span><span lang="EN-US"><o:p></o:p></span></p>

	<p class="MsoNormal"><span style="font-family:宋体;mso-ascii-font-family:Calibri;
	mso-ascii-theme-font:minor-latin;mso-fareast-font-family:宋体;mso-fareast-theme-font:
	minor-fareast;mso-hansi-font-family:Calibri;mso-hansi-theme-font:minor-latin">投稿须知：</span><span lang="EN-US"><o:p></o:p></span></p>

	<p class="MsoNormal"><span lang="EN-US">1</span><span style="font-family:宋体;
	mso-ascii-font-family:Calibri;mso-ascii-theme-font:minor-latin;mso-fareast-font-family:
	宋体;mso-fareast-theme-font:minor-fareast;mso-hansi-font-family:Calibri;
	mso-hansi-theme-font:minor-latin">，从注册材料提交时间算起，见刊周期最长一个月。</span><span lang="EN-US"><o:p></o:p></span></p>

	<p class="MsoNormal"><span lang="EN-US">2. </span><span style="font-family:宋体;
	mso-ascii-font-family:Calibri;mso-ascii-theme-font:minor-latin;mso-fareast-font-family:
	宋体;mso-fareast-theme-font:minor-fareast;mso-hansi-font-family:Calibri;
	mso-hansi-theme-font:minor-latin">论文必须要有摘要、题目、作者、单位、必要的图标、结果、主要参考文献等。</span><span lang="EN-US"><o:p></o:p></span></p>

	<p class="MsoNormal"><span lang="EN-US">3. </span><span style="font-family:宋体;
	mso-ascii-font-family:Calibri;mso-ascii-theme-font:minor-latin;mso-fareast-font-family:
	宋体;mso-fareast-theme-font:minor-fareast;mso-hansi-font-family:Calibri;
	mso-hansi-theme-font:minor-latin">文章为全英文，页数无限制。</span><span lang="EN-US"><o:p></o:p></span></p>

	<p class="MsoNormal"><span lang="EN-US">4. </span><span style="font-family:宋体;
	mso-ascii-font-family:Calibri;mso-ascii-theme-font:minor-latin;mso-fareast-font-family:
	宋体;mso-fareast-theme-font:minor-fareast;mso-hansi-font-family:Calibri;
	mso-hansi-theme-font:minor-latin">请勿一稿多投，所有稿件将接受同行评审，审稿周期根据审稿老师的时间略有不同，审稿时间一般为：</span><span lang="EN-US"><span style="border-bottom:1px dashed #ccc;" t="5" times="">1-5</span></span><span style="font-family:宋体;mso-ascii-font-family:Calibri;
	mso-ascii-theme-font:minor-latin;mso-fareast-font-family:宋体;mso-fareast-theme-font:
	minor-fareast;mso-hansi-font-family:Calibri;mso-hansi-theme-font:minor-latin">天内。</span><span lang="EN-US"><o:p></o:p></span></p>

	<p class="MsoNormal"><span lang="EN-US">5. </span><span style="font-family:宋体;
	mso-ascii-font-family:Calibri;mso-ascii-theme-font:minor-latin;mso-fareast-font-family:
	宋体;mso-fareast-theme-font:minor-fareast;mso-hansi-font-family:Calibri;
	mso-hansi-theme-font:minor-latin">投稿流程：投稿→审稿→录用→注册→开具发票→电子版→纸质版。</span><span lang="EN-US"><o:p></o:p></span></p>

	<p class="MsoNormal"><span style="font-family:宋体;mso-ascii-font-family:Calibri;
	mso-ascii-theme-font:minor-latin;mso-fareast-font-family:宋体;mso-fareast-theme-font:
	minor-fareast;mso-hansi-font-family:Calibri;mso-hansi-theme-font:minor-latin">注：普刊长期征稿，知网可查，不送检</span><span lang="EN-US">EI</span><span style="font-family:宋体;mso-ascii-font-family:Calibri;
	mso-ascii-theme-font:minor-latin;mso-fareast-font-family:宋体;mso-fareast-theme-font:
	minor-fareast;mso-hansi-font-family:Calibri;mso-hansi-theme-font:minor-latin">，一般用于基金项目结题，学术论文发表，学术测评、毕业保底等，具体请作者查看自己高校的政策再做选择。</span><span lang="EN-US"><o:p></o:p></span></p><br></div><style type="text/css">.qmbox style, .qmbox script, .qmbox head, .qmbox link, .qmbox meta {display: none !important;}</style></div></div>
	`,
}