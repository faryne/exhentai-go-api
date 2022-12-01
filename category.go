package exhentai_go_api

type CategoryType int

const (
	CategoryDoujinshi CategoryType = 1 << (1 + iota) // 2 / 4 / 8 / 16 ...
	CategoryManga
	CategoryArtistCG
	CategoryGameCG
	CategoryWestern
	CategoryNonH
	CategoryImageSet
	CategoryCosplay
	CategoryAsianPorn
	CategoryMisc
)

var Categories = [10]CategoryType{
	CategoryDoujinshi,
	CategoryManga,
	CategoryArtistCG,
	CategoryGameCG,
	CategoryWestern,
	CategoryNonH,
	CategoryImageSet,
	CategoryCosplay,
	CategoryAsianPorn,
	CategoryMisc,
}
