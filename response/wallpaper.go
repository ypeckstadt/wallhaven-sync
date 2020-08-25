package response

type Wallpaper struct {
	ID         string   `json:"id"`
	URL        string   `json:"url"`
	ShortURL   string   `json:"short_url"`
	Views      int      `json:"views"`
	Favorites  int      `json:"favorites"`
	Source     string   `json:"source"`
	Purity     string   `json:"purity"`
	Category   string   `json:"category"`
	DimensionX int      `json:"dimension_x"`
	DimensionY int      `json:"dimension_y"`
	Resolution string   `json:"resolution"`
	Ratio      string   `json:"ratio"`
	FileSize   int      `json:"file_size"`
	FileType   string   `json:"file_type"`
	CreatedAt  string   `json:"created_at"`
	Colors     []string `json:"colors"`
	Path       string   `json:"path"`
	Thumbs     Thumbs   `json:"thumbs"`
}

type Thumbs struct {
	Large    string `json:"large"`
	Original string `json:"original"`
	Small    string `json:"small"`
}

type Meta struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
}

type WallpaperResponse struct {
	Wallpapers []Wallpaper `json:"data"`
	Meta       Meta        `json:"meta"`
}
