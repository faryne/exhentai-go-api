package exhentai_go_api

import (
	"fmt"
	"strings"
)

type LanguageType string

const (
	LangAfrikaans  LanguageType = "afrikaans"
	LangAlbanian   LanguageType = "albanian"
	LangArabic     LanguageType = "arabic"
	LangAramaic    LanguageType = "aramaic"
	LangArmenian   LanguageType = "armenian"
	LangBengali    LanguageType = "bengali"
	LangBosnian    LanguageType = "bosnian"
	LangBulgarian  LanguageType = "bulgarian"
	LangBurmese    LanguageType = "burmese"
	LangCatalan    LanguageType = "catalan"
	LangCebuano    LanguageType = "cebuano"
	LangChinese    LanguageType = "chinese"
	LangCree       LanguageType = "cree"
	LangCreole     LanguageType = "creole"
	LangCroatian   LanguageType = "croatian"
	LangCzech      LanguageType = "czech"
	LangDanish     LanguageType = "danish"
	LangDutch      LanguageType = "dutch"
	LangEnglish    LanguageType = "english"
	LangEsperanto  LanguageType = "esperanto"
	LangEstonian   LanguageType = "estonian"
	LangFinnish    LanguageType = "finnish"
	LangFrench     LanguageType = "french"
	LangGeorgian   LanguageType = "georgian"
	LangGerman     LanguageType = "german"
	LangGreek      LanguageType = "greek"
	LangGujarati   LanguageType = "gujarati"
	LangHebrew     LanguageType = "hebrew"
	LangHindi      LanguageType = "hindi"
	LangHmong      LanguageType = "hmong"
	LangHungarian  LanguageType = "hungarian"
	LangIcelandic  LanguageType = "icelandic"
	LangIndonesian LanguageType = "indonesian"
	LangIrish      LanguageType = "irish"
	LangItalian    LanguageType = "italian"
	LangJapanese   LanguageType = "japanese"
	LangJavanese   LanguageType = "javanese"
	LangKannada    LanguageType = "kannada"
	LangKazakh     LanguageType = "kazakh"
	LangKhmer      LanguageType = "khmer"
	LangKorean     LanguageType = "korean"
	LangKurdish    LanguageType = "kurdish"
	LangLadino     LanguageType = "ladino"
	LangLao        LanguageType = "lao"
	LangLatin      LanguageType = "latin"
	LangLatvian    LanguageType = "latvian"
	LangMarathi    LanguageType = "marathi"
	LangMongolian  LanguageType = "mongolian"
	LangNdebele    LanguageType = "ndebele"
	LangNepali     LanguageType = "nepali"
	LangNorwegian  LanguageType = "norwegian"
	LangOromo      LanguageType = "oromo"
	LangPapiamento LanguageType = "papiamento"
	LangPashto     LanguageType = "pashto"
	LangPersian    LanguageType = "persian"
	LangPolish     LanguageType = "polish"
	LangPortuguese LanguageType = "portuguese"
	LangPunjabi    LanguageType = "punjabi"
	LangRomanian   LanguageType = "romanian"
	LangRussian    LanguageType = "russian"
	LangSango      LanguageType = "sango"
	LangSanskrit   LanguageType = "sanskrit"
	LangSerbian    LanguageType = "serbian"
	LangShona      LanguageType = "shona"
	LangSlovak     LanguageType = "slovak"
	LangSlovenian  LanguageType = "slovenian"
	LangSomali     LanguageType = "somali"
	LangSpanish    LanguageType = "spanish"
	LangSwahili    LanguageType = "swahili"
	LangSwedish    LanguageType = "swedish"
	LangTagalog    LanguageType = "tagalog"
	LangTamil      LanguageType = "tamil"
	LangTelugu     LanguageType = "telugu"
	LangThai       LanguageType = "thai"
	LangTibetan    LanguageType = "tibetan"
	LangTigrinya   LanguageType = "tigrinya"
	LangTurkish    LanguageType = "turkish"
	LangUkrainian  LanguageType = "ukrainian"
	LangUrdu       LanguageType = "urdu"
	LangVietnamese LanguageType = "vietnamese"
	LangWelsh      LanguageType = "welsh"
	LangYiddish    LanguageType = "yiddish"
	LangZulu       LanguageType = "zulu"
)

type language struct {
	Languages []LanguageType
}

func NewLanguage() PropertyManagement {
	var l = language{
		Languages: make([]LanguageType, 0),
	}
	return &l
}

func (l *language) Add(input []interface{}) {
	for _, v := range input {
		if inArray(v, l.Languages) == false {
			l.Languages = append(l.Languages, v.(LanguageType))
			continue
		}
	}
}

func (l *language) Remove(input []interface{}) {
	for k, v := range input {
		if inArray(v, l.Languages) == true {
			l.Languages = append(l.Languages[:k], l.Languages[k+1:]...)
			break
		}
	}
}

func (l *language) String(builder *strings.Builder) {
	if len(l.Languages) == 0 {
		builder.WriteString(fmt.Sprintf("l:%s ", string(l.Languages[0])))
	} else {
		for _, v := range l.Languages {
			builder.WriteString(fmt.Sprintf("~l:%s ", string(v)))
		}
	}
}
