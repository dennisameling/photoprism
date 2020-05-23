package form

// AlbumSearch represents search form fields for "/api/v1/albums".
type AlbumSearch struct {
	Query    string `form:"q"`
	ID       string `form:"id"`
	Slug     string `form:"slug"`
	Name     string `form:"name"`
	Favorite bool   `form:"favorite"`
	Count    int    `form:"count" binding:"required" serialize:"-"`
	Offset   int    `form:"offset" serialize:"-"`
	Order    string `form:"order" serialize:"-"`
}

func (f *AlbumSearch) GetQuery() string {
	return f.Query
}

func (f *AlbumSearch) SetQuery(q string) {
	f.Query = q
}

func (f *AlbumSearch) ParseQueryString() error {
	return ParseQueryString(f)
}

func NewAlbumSearch(query string) AlbumSearch {
	return AlbumSearch{Query: query}
}
