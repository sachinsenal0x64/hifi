package middleware

import (
	"hifi/types"
	"sync"
)

var (
	songMap = make(map[string]types.SubsonicSong)
	songMu  sync.RWMutex

	albumYearMap = make(map[string]string)
	albumYearMu  sync.RWMutex

	artistsCache = make(map[string]map[int]types.SubsonicArtist)
	artistsOrder = map[string][]int{}
	artistsMu    sync.RWMutex

	artistInfoCache = make(map[string]types.SubsonicArtistInfo)
	artistInfoMu    sync.RWMutex

	artistsWithAlbumsCache = make(map[string]map[string]types.SubsonicArtistWithAlbums)
	artistMu               sync.RWMutex
)
