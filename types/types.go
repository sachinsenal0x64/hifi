package types

func MetaBanner() SubsonicWrapper {
	var w SubsonicWrapper
	w.Subsonic.Status = "ok"
	w.Subsonic.Version = "1.15.0"
	w.Subsonic.Type = "hifi"
	w.Subsonic.ServerVersion = "0.19.0"
	w.Subsonic.OpenSubsonic = true
	return w
}

// -------------------- TRANSFORM --------------------

type TidalSearchResponse struct {
	Items []struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Duration int    `json:"duration"`
		Explicit bool   `json:"explicit"`
		Artist   struct {
			ID      int    `json:"id,omitempty"`
			Name    string `json:"name,omitempty"`
			Picture string `json:"picture,omitempty"`
		} `json:"artist"`
		Album struct {
			ID    int    `json:"id,omitempty"`
			Title string `json:"title,omitempty"`
			Cover string `json:"cover,omitempty"`
		} `json:"album"`
	} `json:"items"`
}

type TidalArtistResponse struct {
	Items []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	} `json:"items"`
}

// Subsonic Response Format
type SubsonicSong struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Album       string `json:"album"`
	Artist      string `json:"artist"`
	Duration    int    `json:"duration"`
	CoverArt    string `json:"coverArt"`
	Type        string `json:"type"`
	IsVideo     bool   `json:"isVideo"`
	Parent      string `json:"parent"`
	ContentType string `json:"contentType"`
	Suffix      string `json:"suffix"`
	ArtistID    string `json:"artistId"`
	AlbumID     string `json:"albumId"`
}

type SubsonicArtist struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CoverArt    string `json:"coverArt"`
	AlbumCount  int    `json:"albumCount"`
	UserRating  int    `json:"userRating,omitempty"`
	ArtistImage string `json:"artistImageUrl,omitempty"`
}

type SubsonicAlbum struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Artist   string `json:"artist"`
	CoverArt string `json:"coverArt"`
}

type SubsonicSearchResult struct {
	Artist []SubsonicArtist `json:"artist,omitempty"`
	Album  []SubsonicAlbum  `json:"album,omitempty"`
	Song   []SubsonicSong   `json:"song,omitempty"`
}

type SubsonicArtists struct {
	IgnoredArticles string                    `json:"ignoredArticles,omitempty"`
	Index           []SubsonicArtistIndexItem `json:"index,omitempty"`
}

type SubsonicArtistIndexItem struct {
	Name   string           `json:"name,omitempty"`
	Artist []SubsonicArtist `json:"artist,omitempty"`
}

type SubsonicWrapper struct {
	Subsonic struct {
		Status        string                `json:"status"`
		Version       string                `json:"version"`
		Type          string                `json:"type"`
		ServerVersion string                `json:"serverVersion"`
		OpenSubsonic  bool                  `json:"openSubsonic"`
		SearchResult3 *SubsonicSearchResult `json:"searchResult3,omitempty"`
		Artists       *SubsonicArtists      `json:"artists,omitempty"`
		Song          *SubsonicSong         `json:"song,omitempty"`
	} `json:"subsonic-response"`
}

type PlaybackInfo struct {
	Manifest string `json:"manifest,omitempty"`
}

type ManifestData struct {
	Urls []string `json:"urls,omitempty"`
}
