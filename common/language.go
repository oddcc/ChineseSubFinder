package common

const (
	SubNameKeywordChineseSimple = "chs"
	SubNameKeywordTraditional	 = "cht"
)

// Language 语言类型，注意，这里默认还是查找的是中文字幕，只不过下载的时候可能附带了其他的
type Language int
const (
	Unknow                     Language = iota // 未知语言
	ChineseSimple                              // 简体中文
	ChineseTraditional                         // 繁体中文
	ChineseSimpleEnglish                       // 简英双语字幕
	ChineseTraditionalEnglish                  // 繁英双语字幕
	English                                    // 英文
	Japanese                                   // 日语
	ChineseSimpleJapanese                      // 简日双语字幕
	ChineseTraditionalJapanese                 // 繁日双语字幕
	Korean                                     // 韩语
	ChineseSimpleKorean                        // 简韩双语字幕
	ChineseTraditionalKorean                   // 繁韩双语字幕
)


// 需要符合 emby 的格式要求，在后缀名前面
const (
	Emby_unknow = ".unknow"					// 未知语言
	Emby_chi 	= ".chi"					// 简体
	Emby_chs 	= ".chs"					// 简体
	Emby_cht 	= ".cht"					// 繁体
	Emby_chs_en = ".chs_en"                 // 简英双语字幕
	Emby_cht_en = ".cht_en"                	// 繁英双语字幕
	Emby_en 	= ".en"                     // 英文
	Emby_jp 	= ".jp"						// 日语
	Emby_chs_jp = ".chs_jp"                 // 简日双语字幕
	Emby_cht_jp = ".cht_jp"                	// 繁日双语字幕
	Emby_kr 	= ".kr"                     // 韩语
	Emby_chs_kr = ".chs_kr"                 // 简韩双语字幕
	Emby_cht_kr = ".cht_kr"                	// 繁韩双语字幕
)

const (
	MathLangChnUnknow = "未知语言"
	MatchLangDouble   = "双语"
	MatchLangChs      = "简"
	MatchLangCht      = "繁"
	MatchLangChsEn    = "简英"
	MatchLangChtEn    = "繁英"
	MatchLangEn       = "英"
	MatchLangJp       = "日"
	MatchLangChsJp    = "简日"
	MatchLangChtJp    = "繁日"
	MatchLangKr       = "韩"
	MatchLangChsKr    = "简韩"
	MatchLangChtKr    = "繁韩"
)

func (l Language) String() string {
	switch l {
	case ChineseSimple:
		return MatchLangChs
	case ChineseTraditional:
		return MatchLangCht
	case ChineseSimpleEnglish:
		return MatchLangChsEn
	case ChineseTraditionalEnglish:
		return MatchLangChtEn
	case English:
		return MatchLangEn
	case Japanese:
		return MatchLangJp
	case ChineseSimpleJapanese:
		return MatchLangChsJp
	case ChineseTraditionalJapanese:
		return MatchLangChtJp
	case Korean:
		return MatchLangKr
	case ChineseSimpleKorean:
		return MatchLangChsKr
	case ChineseTraditionalKorean:
		return MatchLangChtKr
	default:
		return MathLangChnUnknow
	}
}