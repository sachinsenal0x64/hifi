package types

import (
	"sync"

	"github.com/valkey-io/valkey-go"
)

func MetaBanner() SubsonicWrapper {
	var w SubsonicWrapper
	w.Subsonic.Status = "ok"
	w.Subsonic.Version = "1.15.0"
	w.Subsonic.Type = "hifi"
	w.Subsonic.ServerVersion = "0.19.0"
	w.Subsonic.OpenSubsonic = true
	return w
}

type Router struct {
	Valkey valkey.Client
	Mem    map[string]string
	MemMu  sync.Mutex
}

// -------------------- TRANSFORM --------------------

// Subsonic User
type Person struct {
	UserName string
	PassWord string
}

type ExploreItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Cover string `json:"cover"`
}

type PlaybackInfo struct {
	Manifest string `json:"manifest,omitempty"`
}

type ManifestData struct {
	Urls []string `json:"urls,omitempty"`
}

type TidalNew struct {
	Rows []struct {
		Modules []struct {
			PagedList struct {
				Items []struct {
					ID    int    `json:"id"`
					Title string `json:"title"`
					Cover string `json:"cover"`
				} `json:"items"`
			} `json:"pagedList"`
		} `json:"modules"`
	} `json:"rows"`
}

type TidalSearchResponse struct {
	Tracks struct {
		Items []struct {
			ID       int    `json:"id,omitempty"`
			Title    string `json:"title,omitempty"`
			Duration int    `json:"duration,omitempty"`
			Album    struct {
				ID          int    `json:"id"`
				Title       string `json:"title"`
				Cover       string `json:"cover"`
				ReleaseDate string `json:"releaseDate"`
			} `json:"album"`
			Artist []struct {
				ID      int    `json:"id,omitempty"`
				Name    string `json:"name,omitempty"`
				Picture string `json:"picture,omitempty"`
			} `json:"artists"`
		} `json:"items"`
	} `json:"tracks"`
	Albums struct {
		Items []struct {
			ID          int    `json:"id,omitempty"`
			Title       string `json:"title,omitempty"`
			Cover       string `json:"cover,omitempty"`
			ReleaseDate string `json:"releaseDate,omitempty"`
			Duration    int    `json:"duration,omitempty"`
			Artist      []struct {
				ID      int    `json:"id,omitempty"`
				Name    string `json:"name,omitempty"`
				Picture string `json:"picture,omitempty"`
			} `json:"artists"`
		} `json:"items"`
	} `json:"albums"`

	Artists struct {
		Items []struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Picture string `json:"picture"`
		} `json:"items"`
	} `json:"artists"`
}

type TidalAlbumBannerResponse struct {
	ID       int    `json:"id,omitempty"`
	Title    string `json:"title,omitempty"`
	Duration int    `json:"duration,omitempty"`
	Artist   struct {
		ID int `json:"id"`
	} `json:"artist"`
}

type TidalAlbumResponse struct {
	TotalNumberOfItems int `json:"totalNumberOfItems"`
	Items              []struct {
		Item struct {
			ID              int    `json:"id"`
			Title           string `json:"title"`
			Duration        int    `json:"duration"`
			Explicit        bool   `json:"explicit"`
			StreamStartDate string `json:"streamStartDate,omitempty"`
			ReleaseDate     string `json:"releaseDate,omitempty"`
			TrackNumber     int    `json:"trackNumber,omitempty"`
			Cover           string `json:"cover,omitempty"`
			Artist          struct {
				ID      int    `json:"id,omitempty"`
				Name    string `json:"name,omitempty"`
				Picture string `json:"picture,omitempty"`
			} `json:"artist"`
			Album struct {
				ID           int    `json:"id"`
				Title        string `json:"title"`
				Cover        string `json:"cover"`
				VibrantColor string `json:"vibrantColor"`
				VideoCover   string `json:"videoCover"`
			} `json:"album"`
		} `json:"item"`
	} `json:"items"`
}

type TidalArtistResponse struct {
	Items []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
	} `json:"items"`
}

type TidalTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type TidalArtistAlbumsResponse struct {
	Header struct {
		Biography struct {
			Text string `json:"text,omitempty"`
		} `json:"biography"`
	} `json:"header"`
	Item struct {
		Data struct {
			ID                         int    `json:"id"`
			Name                       string `json:"name"`
			Picture                    string `json:"picture"`
			Popularity                 int    `json:"popularity"`
			SelectedAlbumCoverFallback string `json:"selectedAlbumCoverFallback"`
		} `json:"data"`
	} `json:"item"`
	Items []struct {
		ModuleId string `json:"moduleId,omitempty"`
		Items    []struct {
			Data struct {
				ID           int    `json:"id"`
				Title        string `json:"title"`
				Duration     int    `json:"duration"`
				Cover        string `json:"cover"`
				VibrantColor string `json:"vibrantColor"`
				VideoCover   string `json:"videoCover"`
				ReleaseDate  string `json:"releaseDate"`
				Type         string `json:"type"`
				Artists      []struct {
					ID      int    `json:"id"`
					Name    string `json:"name"`
					Picture string `json:"picture"`
					Main    bool   `json:"main"`
				} `json:"artists"`
			} `json:"data"`
		} `json:"items"`
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
	Parent      int    `json:"parent"`
	Year        string `json:"year,omitempty"`
	Track       int    `json:"track,omitempty"`
	ContentType string `json:"contentType"`
	Suffix      string `json:"suffix"`
	ArtistID    string `json:"artistId"`
	AlbumID     string `json:"albumId"`
}

type SubsonicArtist struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	CoverArt string `json:"coverArt"`
}

type SubsonicAlbum struct {
	ID          string         `json:"id"`
	Album       string         `json:"album,omitempty"`
	Parent      int            `json:"parent"`
	Title       string         `json:"title"`
	Name        string         `json:"name"`
	IsDir       bool           `json:"isDir"`
	CoverArt    string         `json:"coverArt"`
	SongCount   int            `json:"songCount"`
	Duration    int            `json:"duration"`
	PlayCount   int            `json:"playCount,omitempty"`
	ArtistID    int            `json:"artistId"`
	Artist      string         `json:"artist"`
	Year        string         `json:"year,omitempty"`
	TrackNumber int            `json:"trackNumber,omitempty"`
	Genre       string         `json:"genre,omitempty"`
	Song        []SubsonicSong `json:"song,omitempty"`
}

type SubsonicSearchResult struct {
	Artist []SubsonicArtist `json:"artist,omitempty"`
	Album  []SubsonicAlbum  `json:"album,omitempty"`
	Song   []SubsonicSong   `json:"song,omitempty"`
}

type SubsonicArtists struct {
	Index []SubsonicArtistIndexItem `json:"index,omitempty"`
}

type SubsonicArtistIndexItem struct {
	Name   string           `json:"name,omitempty"`
	Artist []SubsonicArtist `json:"artist,omitempty"`
}

type SubsonicAlbumList struct {
	Album []SubsonicAlbum `json:"album"`
}

type SubsonicArtistInfo struct {
	ID        string `json:"id"`
	Biography string `json:"biography,omitempty"`
}

type SubsonicArtistWithAlbums struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	CoverArt   string          `json:"coverArt,omitempty"`
	AlbumCount int             `json:"albumCount,omitempty"`
	Album      []SubsonicAlbum `json:"album,omitempty"`
}

type SubsonicWrapper struct {
	Subsonic struct {
		Status        string                    `json:"status"`
		Version       string                    `json:"version"`
		Type          string                    `json:"type"`
		ServerVersion string                    `json:"serverVersion"`
		OpenSubsonic  bool                      `json:"openSubsonic"`
		SearchResult3 *SubsonicSearchResult     `json:"searchResult3,omitempty"`
		Artists       *SubsonicArtists          `json:"artists,omitempty"`
		Song          *SubsonicSong             `json:"song,omitempty"`
		Album         *SubsonicAlbum            `json:"album,omitempty"`
		AlbumList2    *SubsonicAlbumList        `json:"albumList2,omitempty"`
		ArtistInfo    *SubsonicArtistInfo       `json:"artistInfo,omitempty"`
		Artist        *SubsonicArtistWithAlbums `json:"artist,omitempty"`
	} `json:"subsonic-response"`
}
