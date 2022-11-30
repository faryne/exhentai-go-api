package category

type Type int

const (
	Doujinshi Type = 1 << (1 + iota) // 2 / 4 / 8 / 16 ...
	Manga
	ArtistCG
	GameCG
	Western
	NonH
	ImageSet
	Cosplay
	AsianPorn
	Misc
)
