#!/usr/bin/env python3

import asyncio
import base64
import json
import os
from typing import Union

import httpx
import redis.asyncio as redis
import rich
import uvicorn
from dotenv import load_dotenv
from fastapi import FastAPI, HTTPException, Query
from fastapi.responses import FileResponse, HTMLResponse
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import Response

app = FastAPI(title="HiFi-RestAPI", version="v1.0", description="Tidal Music Proxy")
app.add_middleware(
CORSMiddleware,
allow_origins=["*"],
allow_credentials=False,
allow_methods=["*"],
expose_headers=["*"],
)

load_dotenv()

client_id = os.getenv("CLIENT_ID")
client_secret = os.getenv("CLIENT_SECRET")
access_token = os.getenv("TIDAL_TOKEN")
refresh_token = os.getenv("TIDAL_REFRESH")
redis_url = os.getenv("REDIS_URL")
redis_port = os.getenv("REDIS_PORT")
redis_password = os.getenv("REDIS_PASSWORD")
user_id = os.getenv("USER_ID")

client_id = "zU4XHVVkc2tDPo4t"
client_secret = "VJKhDFqJPqvsPVNBV6ukXTJmwlvbttP7wlMlrc72se4="


if os.path.exists("token.json"):
    with open("token.json", "r") as tok:
        token = json.loads(tok.read())

    refresh_token = token["refresh_token"]
    access_token = token["access_token"]


async def get_redis_connection():
    r = redis.Redis(
        host=redis_url or "localhost",
        port=int(redis_port or 6379),
        password=redis_password,
        db=0,
        ssl=True,
        ssl_cert_reqs="none",
        protocol=3,
        decode_responses=True,
    )
    return r


cached_tok = None


async def token_checker():
    r = await get_redis_connection()
    refresh_url = f"https://api.tidal.com/v2/feed/activities/?userId={user_id}"

    cached_tok = await r.get("access_token")
    await r.close()
    headers = {"authorization": f"Bearer {cached_tok}"}

    async with httpx.AsyncClient() as client:
        res2 = await client.get(url=refresh_url, headers=headers)
        rich.print(res2.json())

        if res2.status_code == 200:
            return res2.status_code


async def refresh():
    status = await token_checker()
    r = await get_redis_connection()
    if status == 200:
        cached_tok = await r.get("access_token")
        await r.close()
        tidal_token = cached_tok
        return tidal_token

    elif status != 200:
        cached_tok = None
        await r.delete("access_token")
        await r.close()

    if not await r.get("access_token"):
        await r.close()
        refresh_url = "https://auth.tidal.com/v1/oauth2/token"
        payload = {
            "client_id": client_id,
            "refresh_token": refresh_token,
            "grant_type": "refresh_token",
            "scope": "r_usr+w_usr+w_sub",
        }
        async with httpx.AsyncClient(http2=True) as client:
            try:
                res2 = await client.post(
                    url=refresh_url,
                    data=payload,
                    auth=(client_id, client_secret),
                )

                print(res2.http_version)

                # Assuming a successful response code is 200
                if res2.status_code == 200:
                    print_token = res2.json()
                    tida_token = print_token.get("access_token")
                    await r.set("access_token", tida_token)
                    await r.close()
                    tidal_token = tida_token
                    return tidal_token
                else:
                    return {"error": f"Failed to refresh token: {res2.status_code}"}

            except httpx.HTTPError as e:
                return {"error": f"HTTP error occurred: {str(e)}"}
            except Exception as e:
                return {"error": f"An error occurred: {str(e)}"}


async def auth():
    cids = client_id
    csec = client_secret

    url = "https://auth.tidal.com/v1/oauth2/token"

    payload = {
        "grant_type": "client_credentials",
        "client_id": cids,
        "client_secret": csec,
    }

    async with httpx.AsyncClient(http2=True) as client:
        res = await client.post(url=url, data=payload)

        access_token = res.json()["access_token"]
        expires_in = res.json()["expires_in"]
        token_type = res.json()["token_type"]
        out_res = {
            "access_token": access_token,
            "expires_in": expires_in,
            "token_type": token_type,
        }

    return out_res


@app.api_route("/", methods=["GET"], include_in_schema=False)
async def index():
    return {"HIFI-API": "v1.0", "Repo": "https://github.com/sachinsenal0x64/hifi"}

"""
Tidal Hi-Res streams are only avaliable over MPEG-DASH.

This endpoint returns the DASH manifest for the requested track ID and quality.
Configure Shaka Player to convert all HEAD requests to GET requests as this server doesn't support them:

```js
player.getNetworkingEngine().registerRequestFilter(function(type, request) {
    // Convert any HEAD requests to GET
    if (request.method === 'HEAD') {
        request.method = 'GET';
    }
});
```
"""
@app.api_route("/dash/", methods=["GET"])
async def get_hi_res(
    id: int,
    quality: str = "HI_RES_LOSSLESS"
):
    try:
        tokz = await refresh()
        tidal_token = tokz
        track_url = f"https://tidal.com/v1/tracks/{id}/playbackinfo?audioquality={quality}&playbackmode=STREAM&assetpresentation=FULL"
        
        payload = {
            "authorization": f"Bearer {tidal_token}",
        }
        async with httpx.AsyncClient(http2=True) as client:
            track_data = await client.get(url=track_url, headers=payload)
            final_data = track_data.json()
            decode_manifest = base64.b64decode(final_data["manifest"]) # returns dash
            return Response(content=decode_manifest, media_type=final_data["manifestMimeType"])
    except KeyError:
        raise HTTPException(
            status_code=404,
            detail="Quality not found. check API docs = https://github.com/sachinsenal0x64/Hifi-Tui?tab=readme-ov-file#-api-documentation"  
        )

@app.api_route("/track/", methods=["GET"])
async def get_track(
    id: int,
    quality: str = "LOSSLESS"
):
    try:
        if quality == "HI_RES_LOSSLESS":
            raise HTTPException(
                status_code=404,
                detail="HI_RES_LOSSLESS quality is not supported in this endpoint. Use /dash/ endpoint instead which returns MPEG-DASH.",
            )
        tokz = await refresh()
        tidal_token = tokz

        track_url = f"https://api.tidal.com/v1/tracks/{id}/playbackinfopostpaywall/v4?audioquality={quality}&playbackmode=STREAM&assetpresentation=FULL"

        info_url = f"https://api.tidal.com/v1/tracks/{id}/?countryCode=US"

        payload = {
            "authorization": f"Bearer {tidal_token}",
        }

        async with httpx.AsyncClient(http2=True) as client:
            track_data = await client.get(url=track_url, headers=payload)
            info_data = await client.get(url=info_url, headers=payload)

            final_data = track_data.json()["manifest"]
            decode_manifest = base64.b64decode(final_data)
            con_json = json.loads(decode_manifest)
            audio_url = con_json.get("urls")[0]
            au_j = {"OriginalTrackUrl": audio_url}
            fetch_info = info_data.json()

            return [fetch_info, track_data.json(), au_j]

    except KeyError:
        raise HTTPException(
            status_code=404,
            detail="Quality not found. check API docs = https://github.com/sachinsenal0x64/Hifi-Tui?tab=readme-ov-file#-api-documentation",
        )
    except httpx.ConnectTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ConnectError:
        raise HTTPException(
            status_code=429,
        )

    except json.JSONDecodeError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.WriteError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadError:
        raise HTTPException(
            status_code=429,
        )


@app.api_route("/lyrics/", methods=["GET"])
async def get_lyrics(id: int):
    try:
        tokz = await refresh()
        tidal_token = tokz
        search_url = f"https://api.tidal.com/v1/tracks/{id}/lyrics?countryCode=US&locale=en_US&deviceType=BROWSER"
        payload = {
            "authorization": f"Bearer {tidal_token}",
        }
        async with httpx.AsyncClient(http2=True) as clinet:
            search_data = await clinet.get(url=search_url, headers=payload)

            return [search_data.json()]

    except KeyError:
        raise HTTPException(
            status_code=404,
            detail="Quality not found. check API docs = https://github.com/sachinsenal0x64/Hifi-Tui?tab=readme-ov-file#-api-documentation",
        )

    except json.JSONDecodeError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ConnectTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ConnectError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.WriteError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadError:
        raise HTTPException(
            status_code=429,
        )


@app.api_route("/song/", methods=["GET"])
async def get_song(q: str, quality: str):
    try:
        tokz = await refresh()
        tidal_token = tokz
        search_url = f"https://api.tidal.com/v1/search/tracks?countryCode=US&query={q}"
        payload = {
            "authorization": f"Bearer {tidal_token}",
        }
        async with httpx.AsyncClient(http2=True) as clinet:
            search_data = await clinet.get(url=search_url, headers=payload)
            try:
                id = search_data.json()["items"][0]["id"]

            except IndexError:
                raise HTTPException(status_code=404)

            track_url = f"https://api.tidal.com/v1/tracks/{id}/playbackinfopostpaywall/v4?audioquality={quality}&playbackmode=STREAM&assetpresentation=FULL"
            track_data = await clinet.get(url=track_url, headers=payload)

            print(track_data.http_version)

            rich.print(track_data.json())
            final_data = track_data.json()["manifest"]
            decode_manifest = base64.b64decode(final_data)
            con_json = json.loads(decode_manifest)
            audio_url = con_json.get("urls")[0]
            au_j = {"OriginalTrackUrl": audio_url}

            return [search_data.json()["items"][0], track_data.json(), au_j]

    except KeyError:
        raise HTTPException(
            status_code=404,
            detail="Quality not found. check API docs = https://github.com/sachinsenal0x64/Hifi-Tui?tab=readme-ov-file#-api-documentation",
        )

    except json.JSONDecodeError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ConnectTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ConnectError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.WriteError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadError:
        raise HTTPException(
            status_code=429,
        )


@app.api_route("/search/", methods=["GET"])
async def search(
    s: Union[str, None] = Query(default=None),
    a: Union[str, None] = Query(default=None),
    al: Union[str, None] = Query(default=None),
    v: Union[str, None] = Query(default=None),
    p: Union[str, None] = Query(default=None),
):
    try:
        tokz = await refresh()
        tidal_token = tokz

        search_url = f"https://api.tidal.com/v1/search/tracks?query={s}&limit=25&offset=0&countryCode=US"

        artist_url = f"https://api.tidal.com/v1/search/top-hits?query={a}&limit=25&offset=0&types=ARTISTS,TRACKS&countryCode=US"

        album_url = f"https://api.tidal.com/v1/search/top-hits?query={al}&limit=25&offset=0&types=ALBUMS&countryCode=US"

        video_url = f"https://api.tidal.com/v1/search/top-hits?query={v}&limit=25&offset=0&types=VIDEOS&countryCode=US"

        playlist_url = f"https://api.tidal.com/v1/search/top-hits?query={p}&limit=25&offset=0&types=PLAYLISTS&countryCode=US"

        header = {"authorization": f"Bearer {tidal_token}"}

        async with httpx.AsyncClient(http2=True) as clinet:
            if s:
                search_data = await clinet.get(url=search_url, headers=header)
                sed = search_data.json()
                return sed

            if a:
                search_data = await clinet.get(url=artist_url, headers=header)
                sed = search_data.json()
                return [sed]

            if al:
                search_data = await clinet.get(url=album_url, headers=header)
                sed = search_data.json()
                return sed

            if v:
                search_data = await clinet.get(url=video_url, headers=header)
                sed = search_data.json()
                return sed

            if p:
                search_data = await clinet.get(url=playlist_url, headers=header)
                sed = search_data.json()
                return sed

    except httpx.ConnectTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ConnectError:
        raise HTTPException(
            status_code=429,
        )

    except json.JSONDecodeError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.WriteError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadError:
        raise HTTPException(
            status_code=429,
        )

    except TypeError:
        raise HTTPException(status_code=404)


@app.api_route("/album/", methods=["GET"])
async def get_album(id: int):
    try:
        tokz = await refresh()
        tidal_token = tokz
        search_url = f"https://api.tidal.com/v1/albums/{id}/?countryCode=US"
        item_url = (
            f"https://api.tidal.com/v1/albums/{id}/items?limit=100&countryCode=US"
        )
        header = {"authorization": f"Bearer {tidal_token}"}
        async with httpx.AsyncClient(http2=True) as clinet:
            album_data = await clinet.get(url=search_url, headers=header)
            album_item = await clinet.get(url=item_url, headers=header)
            sed = album_data.json()
            sed_2 = album_item.json()

            return [sed, sed_2]

    except httpx.ConnectTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ConnectError:
        raise HTTPException(
            status_code=429,
        )

    except json.JSONDecodeError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.WriteError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadError:
        raise HTTPException(
            status_code=429,
        )


@app.api_route("/playlist/", methods=["GET"])
async def get_playlist(id: str):
    try:
        tokz = await refresh()
        tidal_token = tokz
        search_url = f"https://api.tidal.com/v1/playlists/{id}?countryCode=US"
        search_item = (
            f"https://api.tidal.com/v1/playlists/{id}/items?countryCode=US&limit=100"
        )
        header = {"authorization": f"Bearer {tidal_token}"}
        async with httpx.AsyncClient(http2=True) as clinet:
            album_search = await clinet.get(url=search_url, headers=header)
            album_item = await clinet.get(url=search_item, headers=header)
            sed_1 = album_search.json()
            sed_2 = album_item.json()

            return [sed_1, sed_2]

    except httpx.ConnectTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ConnectError:
        raise HTTPException(
            status_code=429,
        )

    except json.JSONDecodeError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.WriteError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadError:
        raise HTTPException(
            status_code=429,
        )


@app.api_route("/artist/", methods=["GET"])
async def get_artist(
    id: Union[int, None] = Query(default=None),
    f: Union[int, None] = Query(default=None),
):
    try:
        tokz = await refresh()
        tidal_token = tokz
        search_url = f"https://api.tidal.com/v1/artists/{id}?countryCode=US"
        header = {"authorization": f"Bearer {tidal_token}"}
        async with httpx.AsyncClient(http2=True) as clinet:
            try:
                if id:
                    global sed_1
                    album_search = await clinet.get(url=search_url, headers=header)
                    sed_1 = album_search.json()
                    artist_ids = []
                    artist_cover = sed_1["picture"].replace("-", "/")
                    artist_ids.append(artist_cover)
                    artist_name = sed_1["name"]
                    artist_ids.append(artist_name)
                    artist_id = sed_1["id"]
                    artist_ids.append(artist_id)

                    json_data = [
                        {
                            "id": artist_ids[i + 2],
                            "name": artist_ids[i + 1],
                            "750": f"https://resources.tidal.com/images/{artist_ids[i]}/750x750.jpg",
                        }
                        for i in range(0, len(artist_ids), 3)
                    ]

                    return [sed_1, json_data]

                try:
                    if f:
                        artist_albums = f"https://api.tidal.com/v1/pages/single-module-page/ae223310-a4c2-4568-a770-ffef70344441/4/a4f964ba-b52e-41e8-b25c-06cd70c1efad/2?artistId={f}&countryCode=US&deviceType=BROWSER"
                        album_data = await clinet.get(url=artist_albums, headers=header)
                        alb = album_data.json()

                        albums_ids = []
                        for album in alb.get("rows")[0]["modules"][0]["pagedList"][
                            "items"
                        ]:
                            albums_ids.append(album["id"])

                        all_tracks = []
                        for album_id in albums_ids:
                            album_endpoint = f"https://api.tidal.com/v1/pages/album?albumId={album_id}&countryCode=US&deviceType=BROWSER"
                            album_info = await clinet.get(
                                url=album_endpoint, headers=header
                            )
                            album_tracks = album_info.json().get("rows")[1]["modules"][
                                0
                            ]["pagedList"]["items"]
                            all_tracks.extend([track["item"] for track in album_tracks])

                        return [alb, all_tracks]
                except Exception:
                    artist_albums = f"https://api.tidal.com/v1/pages/single-module-page/ae223310-a4c2-4568-a770-ffef70344441/4/a4f964ba-b52e-41e8-b25c-06cd70c1efad/2?artistId={f}&countryCode=US&deviceType=BROWSER"
                    album_data = await clinet.get(url=artist_albums, headers=header)
                    alb = album_data.json()
                    return [alb]

            except AttributeError:
                url = f"https://api.tidal.com/v1/pages/artist?artistId={id}&countryCode=US&locale=en_US&deviceType=BROWSER"
                artist_data = await clinet.get(url=url, headers=header)
                art = artist_data.json()
                return [art]

    except httpx.ConnectTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ConnectError:
        raise HTTPException(
            status_code=429,
        )

    except json.JSONDecodeError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.WriteError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadError:
        raise HTTPException(
            status_code=429,
        )


@app.api_route("/cover/", methods=["GET"])
async def get_cover(id: Union[int, None] = None, q: Union[str, None] = None):
    try:
        tokz = await refresh()
        tidal_token = tokz
        if id:
            search_url = f"https://api.tidal.com/v1/tracks/{id}/?countryCode=US"
            header = {"authorization": f"Bearer {tidal_token}"}
            async with httpx.AsyncClient(http2=True) as clinet:
                cover_data = await clinet.get(url=search_url, headers=header)
                tracks = cover_data.json()["album"]

                album_ids = []  # list to store album ids

                album_track_id = tracks["id"]
                album_cover = tracks["cover"].replace("-", "/")
                album_name = tracks["title"]
                album_ids.append(album_cover)
                album_ids.append(album_name)
                album_ids.append(album_track_id)

                json_data = [
                    {
                        "id": album_ids[i + 2],
                        "name": album_ids[i + 1],
                        "1280": f"https://resources.tidal.com/images/{album_ids[i]}/1280x1280.jpg",
                        "640": f"https://resources.tidal.com/images/{album_ids[i]}/640x640.jpg",
                        "80": f"https://resources.tidal.com/images/{album_ids[i]}/80x80.jpg",
                    }
                    for i in range(0, len(album_ids), 3)
                ]

                # Create a list of dictionaries with "cover" and "name" keys
                return json_data

        elif q:
            search_url = (
                f"https://api.tidal.com/v1/search/tracks?countryCode=US&query={q}"
            )
            header = {"authorization": f"Bearer {tidal_token}"}
            async with httpx.AsyncClient() as clinet:
                cover_data = await clinet.get(url=search_url, headers=header)
                tracks = cover_data.json()["items"][:10]

                album_ids = []  # list to store album ids

                for track in tracks:
                    album_track_id = track["id"]
                    print(album_track_id)
                    album_cover = track["album"]["cover"].replace("-", "/")
                    album_name = track["title"]
                    album_ids.append(album_cover)
                    album_ids.append(album_name)
                    album_ids.append(album_track_id)
                json_data = [
                    {
                        "id": album_ids[i + 2],
                        "name": album_ids[i + 1],
                        "1280": f"https://resources.tidal.com/images/{album_ids[i]}/1280x1280.jpg",
                        "640": f"https://resources.tidal.com/images/{album_ids[i]}/640x640.jpg",
                        "80": f"https://resources.tidal.com/images/{album_ids[i]}/80x80.jpg",
                    }
                    for i in range(0, len(album_ids), 3)
                ]

                # Create a list of dictionaries with "cover" and "name" keys
                return json_data

        else:
            raise HTTPException(
                status_code=404,
                detail="Cover not found. check API docs = https://github.com/sachinsenal0x64/Hifi-Tui?tab=readme-ov-file#-api-documentation",
            )

    except KeyError:
        raise HTTPException(
            status_code=404,
            detail="Cover not found. check API docs = https://github.com/sachinsenal0x64/Hifi-Tui?tab=readme-ov-file#-api-documentation",
        )

    except httpx.ConnectTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ConnectError:
        raise HTTPException(
            status_code=429,
        )

    except json.JSONDecodeError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadTimeout:
        raise HTTPException(
            status_code=429,
        )

    except httpx.WriteError:
        raise HTTPException(
            status_code=429,
        )

    except httpx.ReadError:
        raise HTTPException(
            status_code=429,
        )
