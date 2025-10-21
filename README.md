> [!IMPORTANT]
This project will be ported to be compatible with Subsonic. Once that’s done, you’ll be able to use any frontend client that supports the Subsonic API.
I can’t provide an ETA right now, but all current issues will be resolved in the next update.
The previous API will be deprecated, yet all its endpoints will continue to work with Subsonic.
[API Docs](https://opensubsonic.netlify.app/docs).

<div align="center">

<h1 align="center"> HIFI </h1>

<h4 align="center">🎵 Opinionated Subsonic-compatible Tidal proxy.<br><br>

</h4>
 
<img width="1568" height="860" alt="hifi tidal" src="https://sachinsenal0x64.github.io/picx-images-hosting/hifi.5fkz01pkwn.webp" />

 HiFi Frontend - [BiniTidal](https://music.binimum.org)  Web UI 



<img width="1568" height="860" alt="hifi tidal" src="https://sachinsenal0x64.github.io/picx-images-hosting/image.3d56fumd2v.webp" />

 Another HiFi Frontend - [Monochrome](https://monochrome.prigoana.com) or [Monochrome](https://monochrome.tf) Web UI 

</div>

<div align="center">
	
[![License: MIT](https://img.shields.io/badge/License-MIT-orange.svg)](https://opensource.org/licenses/MIT)
  

</div>

<br><br>

# 💕 Community

> 🍻 Join the community:  <a href="https://discord.gg/EbfftZ5Dd4">Discord</a>
> [![](https://cdn.statically.io/gh/sachinsenal0x64/picx-images-hosting@master/discord.72y8nlaw5mdc.webp)](https://discord.gg/EbfftZ5Dd4)
 
<br>


> [!IMPORTANT]  
> # F.A.Q / Terms

> We do not encourage piracy. This project is made purely for educational and personal use, intended only for listening, not for downloading.
  If you somehow download copyrighted content, you are solely responsible for complying with the relevant laws in your country.

> The HiFi Project assumes no responsibility for any misuse or legal violations arising from the use of this project.

> I'm currently paying for a Tidal HiFi Plus subscription.

> Community projects may be inspired by or related to the HiFi Project but aren’t officially part of it. These projects are run independently, and the HiFi team isn’t responsible for their content or actions. Each project team should follow any relevant laws and handle its own checks and compliance.

> The HiFi project does not claim ownership of any music or audio content. All rights remain with their respective copyright holders. Users are encouraged to support artists and rights owners by maintaining a valid Tidal subscription. HiFi serves solely as a playback and control interface (e.g., on a Raspberry Pi) for personal, non-commercial use.

> You can access our rest api for free if you want to self-host then need tidal subscription. 

> TUI is plug & play also you can add your own tidal account but by default it has our API so you can listen tidal music for free.

> TL;DR  HIFI API Can Get Any Quality & Codec Which Tidal Offer / Some Qualities & Codecs Need Special Driver / Song / Hardware to get maximum output i always recommend to use `HI_RES` or `LOSSLESS` both are in flac.

> How do I know if this is an MQA or not? - [MQA-CHECKER](https://github.com/purpl3F0x/MQA_identifier)

> File Size - Low (96 kbps) - 3 MB | Low (320 kbps) - 8 MB | High (FLAC, 16-bit, 44.1 kHz) - 30 MB | Max (MQA) - 26 MB | Max (MQA) - 26  | Max (HiRes FLAC, up to 24-bit, 192 kHz) - 30 MB to 140 MB

> AUDIO QUALITY / CODEC  : SONY_360RA | DOLBY ATMOS | MQA 96K  | HI RES FLAC | HI RES LOSSLESS | FLAC LOSSLESS | HIGH | LOW | Up to 24-bit, 192 kHz



<br>


> [!NOTE]  
> API

> [API DOCS](https://github.com/sachinsenal0x64/Hifi-Tui?tab=readme-ov-file#-api-documentation-no-account-required)

> https://hifi.401658.xyz 


<br>



# 🔋 EXTRA 

- [host-hifi-restapi-on-vercel](https://github.com/sachinsenal0x64/host-hifi-restapi-on-vercel)
- [host-hifi-restapi-on-heroku](https://github.com/sachinsenal0x64/host-hifi-restapi-on-heroku)
- [tidal_auth](https://github.com/sachinsenal0x64/hifi-tui/tree/main/tidal_auth)
- [apigateway](https://github.com/sachinsenal0x64/hifi-tui/tree/main/apigateway)


<br>

# 🌱 Community 

- [tidal-dl](https://github.com/Nem-git/tidal) Hifi CLI
- [tidler](https://git.medvidek77.tech/Medvidek77/tidler) Hifi CLI
- [BiniTidal (Web)](https://music.binimum.org) | [SquidWtf (Web)](https://tidal.squid.wtf) | [BiniTidal (Source code)](https://github.com/uimaxbai/tidal-ui) Hifi Frontend
- [Monochrome (Web)](https://monochrome.prigoana.com) | [Monochrome (Web)](https://monochrome.tf) | [Monochrome (Source code)](https://github.com/eduardprigoana/monochrome)  Hifi Frontend
- [Multi-Region Instances](https://github.com/EduardPrigoana/hifi-instances) - Status

<br>


# 💨 Quick Start

## Web UI

- [BiniTidal (Web)](https://music.binimum.org) | [SquidWtf (Web)](https://tidal.squid.wtf) Hifi Frontend
- [Monochrome (Web)](https://monochrome.prigoana.com)  Another Hifi Frontend

<br>

## 🏠 INSTALLATION (Optional)  

> [!TIP]
> You can access [api](https://github.com/sachinsenal0x64/Hifi-Tui?tab=readme-ov-file#-tidal-reverse-api--status) for free without having to self-host it.

> [!NOTE]
> This Required [Tidal](https://tidal.com) subscription / [Redis](https://github.com/redis/redis) & Fill the [.env](https://github.com/sachinsenal0x64/Hifi-Tui/blob/main/api/env-example) file. / Grab Tokens and Ids Using
[tidal_auth.py](https://github.com/sachinsenal0x64/hifi-tui/tree/main/tidal_auth)


<br>

### 🐳 Docker

```sh
# Clone the Repo
git clone https://github.com/sachinsenal0x64/hifi

# Rename .env-example
cd hifi
cp .env.example .env

# Just run
docker compose up

```
<br>



### 🐳 Update

```sh
# Clone the Repo
git clone https://github.com/sachinsenal0x64/hifi

cd hifi
git pull

# Just run
docker compose pull
docker compose down
docker compose up -d --build
docker image prune -f

# or 

cd hifi 

chmod +x update.sh

./update.sh

```

<br>

### 🦄 From Source Code

```sh
git clone https://github.com/sachinsenal0x64/hifi

# Rename .env-example
cd hifi
cp .env.example .env

pip install "fastapi[all]"
pip install -r requirements.txt

python3 -m uvicorn main:app --host 0.0.0.0 --port 8000

```

![fastapi](https://sachinsenal0x64.github.io/picx-images-hosting/300191675-4330ea31-3f15-45b0-962c-ca5a85041f02.5tz3jj54f2ps.webp)


<br>

## 🎞️ DEMO

[![hifi tui](https://img.youtube.com/vi/TfIWc5sQ2M0/0.jpg)](https://www.youtube.com/watch?v=TfIWc5sQ2M0)


<br>

## 📡 API DOCUMENTATION (No account required)

> [!TIP]
> The same API endpoints are used for self-hosted instances — just change your domain.


------------------------------------------------------------------------------------------

<details>

 <summary><code>GET</code>   <code><b>/track/</b></code> </summary>

## Request


<br>

> | Parameter  |   Type    | Description |
> |------------|-----------|-------------|
> | `id`       |  integer  | Track Id = `286266926` |
> | `quality`  |  string   | Song Quality = `HI_RES_LOSSLESS` or `HI_RES` or `LOSSLESS` or `HIGH` or `LOW ` |


<br>

#### Example

>xh

    https GET "https://hifi.401658.xyz/track/?id=286266926&quality=LOSSLESS"
    

![image](https://github.com/sachinsenal0x64/Hifi-Tui/assets/127573781/e586ec03-68eb-4c54-b6ee-251093f4b8a6)

<br>


### Response

  ```json
{
        "albumPeakAmplitude": 1.0,
        "albumReplayGain": -9.18,
        "assetPresentation": "FULL",
        "audioMode": "STEREO",
        "audioQuality": "LOSSLESS",
        "bitDepth": 16,
        "manifest": "base64 manifest",
        "manifestMimeType": "application/vnd.tidal.bts",
        "sampleRate": 44100,
        "trackId": 286266926,
        "trackPeakAmplitude": 0.988482,
        "trackReplayGain": -7.89
    },
    {
        "originalTrack": "aka song track"
    }
```
<br>


### Status Codes

HIFI returns the following status codes in its API:

> | Status Code | Description |
> | :---        | :--- |
> | 200         | `OK` |
> | 422         | `UNPROCESSABLE CONTENT` |
> | 404         | `NOT FOUND` |
> | 500         | `INTERNAL SERVER ERROR` |


</details>

------------------------------------------------------------------------------------------

<details>

 <summary><code>GET</code>   <code><b>/home/</b></code> </summary>

## Request


<br>

> | Parameter  |   Type    | Description | Optional    |
> |------------|-----------|-------------|-------------|
> | `country`  |  string   | Country = `AU`| `true`      |


<br>

#### Example

>xh

    xh GET "https://hifi.401658.xyz/home/?country=AU"
    

<br>


### Response

  ```json
{
  "selfLink": null,
  "id": "eyJwIjoiMDc3ZDJlZmQtNjUwNC00ZWJlLTlkNDktNzNhMTA5YjgyZWVmIiwicFYiOjEyfQ==",
  "title": "Home",
  "rows": [
    {
      "modules": [
        {
          "id": "eyJwIjoiMDc3ZDJlZmQtNjUwNC00ZWJlLTlkNDktNzNhMTA5YjgyZWVmIiwicFYiOjEyLCJtIjoiMDVmNWQ5NDYtYmNmOS00Yzc3LWFmZDEtOTFiMzNjMjEwYmM3IiwibVYiOjEsIm1IIjoiMTQ2NzQxMTMifQ==",
          "type": "PLAYLIST_LIST",
          "width": 100,
          "scroll": "HORIZONTAL",
          "title": "From our editors",
          "description": "",
          "showMore": {
            "title": "View all",
            "apiPath": "pages/single-module-page/077d2efd-6504-4ebe-9d49-73a109b82eef/12/05f5d946-bcf9-4c77-afd1-91b33c210bc7/1"
          },
          "pagedList": {
            "dataApiPath": "pages/data/202c0e67-a076-4f3c-b794-f42429e8b5e7",
            "limit": 15,
            "offset": 0,
            "totalNumberOfItems": 40,
            "items": [
              {
                "uuid": "e4454549-dfb8-46dd-8cbf-32860ac2b318",
                "title": "Usher Essentials",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/e4454549-dfb8-46dd-8cbf-32860ac2b318",
                "image": "f0a04689-5f26-45ac-9054-c06b88b71839",
                "squareImage": "b7e9df64-f3c4-443a-a6d1-1d135e758186",
                "duration": 6200,
                "numberOfTracks": 25,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2024-02-12T06:31:08.733Z",
                "promotedArtists": [
                  {
                    "id": 1518,
                    "name": "USHER",
                    "type": "MAIN",
                    "picture": "1a28d3f8-0ef5-4e90-bac3-330c153e3650",
                    "handle": null,
                    "userId": 0,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "With his relentless work ethic, flawless voice and unparalleled dance moves, Usher has solidified his place among legendary musicians.This playlist recaps Usher's journey to becoming one the best-selling artists of all time. (Photo: Press)"
              },
              {
                "uuid": "1b418bb8-90a7-4f87-901d-707993838346",
                "title": "New Arrivals",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/1b418bb8-90a7-4f87-901d-707993838346",
                "image": "89bf406d-9065-4b41-b566-21a913c8343a",
                "squareImage": "848dfc87-f0d3-4805-aa79-cf264161fa5e",
                "duration": 9910,
                "numberOfTracks": 52,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2025-10-10T10:10:16.916Z",
                "promotedArtists": [
                  {
                    "id": 4916222,
                    "name": "Khalid",
                    "type": "MAIN",
                    "picture": "6064ded8-bd93-4b92-b9d1-2aa28daf86b5",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 3544816,
                    "name": "Demi Lovato",
                    "type": "MAIN",
                    "picture": "8e694491-83bf-4580-b8ea-0e63814a4067",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 23409727,
                    "name": "PinkPantheress",
                    "type": "MAIN",
                    "picture": "224393ce-7f5a-4848-8800-fead1ffa37dc",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 7714725,
                    "name": "Amber Mark",
                    "type": "MAIN",
                    "picture": "e1559ea7-9f31-4f01-b82e-14be843f9032",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "Dig into fresh new releases with this playlist highlighting the most notable songs in a broad range of genres. (Cover: Khalid / Photo: press)          "
              },
              {
                "uuid": "b3aa74f5-4fe3-49ee-9cfa-92b53192053b",
                "title": "New Arrivals: Hip-Hop and R&B",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/b3aa74f5-4fe3-49ee-9cfa-92b53192053b",
                "image": "9ac6271f-ef14-4861-8f04-f6ed334571ad",
                "squareImage": "1c0b7385-6335-4fa6-ab39-05a7d79f73aa",
                "duration": 14177,
                "numberOfTracks": 80,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2025-10-10T04:00:16.074Z",
                "promotedArtists": [
                  {
                    "id": 7272766,
                    "name": "G Herbo",
                    "type": "MAIN",
                    "picture": "6abb8d91-e6fe-473b-a666-1a383c69fe46",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 4973645,
                    "name": "Ty Dolla $ign",
                    "type": "MAIN",
                    "picture": "d8e7fce3-711e-4fc2-80ef-3a7ab6fe9ffb",
                    "handle": null,
                    "userId": 0,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 9644098,
                    "name": "Coi Leray",
                    "type": "MAIN",
                    "picture": "f3a4db12-c23d-4eae-b8f1-950f08131f69",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 9116035,
                    "name": "Shoreline Mafia",
                    "type": "MAIN",
                    "picture": "afb063da-4bac-4dfe-8ba5-8043c1dda15c",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "The number one source for the best new Hip-Hop and R&B. (Cover: G Herbo / Ty Dolla $ign // Photo: Republic Records, Press / Atlantic Records, Press)"
              },
              {
                "uuid": "da0d8a6d-de60-4528-9f1a-1316f054b4c4",
                "title": "New Arrivals: Pop and R&B",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/da0d8a6d-de60-4528-9f1a-1316f054b4c4",
                "image": "c6a96027-32aa-4088-a570-5e82a5859b45",
                "squareImage": "67f70f4e-2fcf-4295-9e0c-5c1b69810702",
                "duration": 16637,
                "numberOfTracks": 87,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2025-10-10T04:00:15.027Z",
                "promotedArtists": [
                  {
                    "id": 7714725,
                    "name": "Amber Mark",
                    "type": "MAIN",
                    "picture": "e1559ea7-9f31-4f01-b82e-14be843f9032",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 5866385,
                    "name": "Anderson .Paak",
                    "type": "MAIN",
                    "picture": "167df412-60db-4821-b0dc-913c89adbda8",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 5439164,
                    "name": "Josh Levi",
                    "type": "MAIN",
                    "picture": "0b6617e6-c54c-4f39-b564-ebbdff628403",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 3851664,
                    "name": "Jason Derulo",
                    "type": "MAIN",
                    "picture": "b1a1cb60-f9a3-4249-9900-652cdde9bbd1",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "This playlist highlights the newest and standout Pop and R&B releases of the week. (Cover: Amber Mark / Josh Levi // Interscope Records, Press / Atlantic Records, Press)"
              },
              {
                "uuid": "27a0e725-0807-4fd4-952b-d79c93da6ec6",
                "title": "New Arrivals: Hip-Hop and Latin",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/27a0e725-0807-4fd4-952b-d79c93da6ec6",
                "image": "2df2c24a-6fba-4023-9df1-2ceaea7a7e49",
                "squareImage": "f10f8c67-c18d-48b8-a56f-773d5d7852b1",
                "duration": 13837,
                "numberOfTracks": 80,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2025-10-10T04:14:08.588Z",
                "promotedArtists": [
                  {
                    "id": 7835837,
                    "name": "BIA",
                    "type": "MAIN",
                    "picture": "61c06cb9-b81e-4a0b-817c-8eed82fa5bc7",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 19760521,
                    "name": "Dei V",
                    "type": "MAIN",
                    "picture": "173bc70a-b30b-49b2-9398-313f941b37ca",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 6389843,
                    "name": "Kodak Black",
                    "type": "MAIN",
                    "picture": "b712dbd5-82d9-4f65-b374-11908041abcc",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 5026732,
                    "name": "Chance the Rapper",
                    "type": "MAIN",
                    "picture": "ac6bb9d7-5278-45dc-b2bc-ac99f9c56a66",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "Hip-hop culture is universal and it's evident in the way we consume music today. With the influx of Latinx stars (Nicky Jam, J Balvin, Bad Bunny) on the charts and collaborating with their fellow mainstream hip-hop MCs (Drake, Cardi B, Meek Mill), it only made sense to combine the worlds. Because at the end of the day, it's all hip-hop, baby. Listen to the best new Hip-Hop & Latin music releases. Playlist previously known as Dímelo Flow. (Cover: BIA / Dei V // Photo: Epic Records, Press / The Orchard, Press)"
              },
              {
                "uuid": "b7882c14-9890-4fbb-b341-eaf88c0939c7",
                "title": "Bruno Mars Essentials",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/b7882c14-9890-4fbb-b341-eaf88c0939c7",
                "image": "335bf674-4859-40d1-86f7-b91cf00dcd71",
                "squareImage": "4e3c461f-6ca3-4e20-96cc-21f804c6919b",
                "duration": 5479,
                "numberOfTracks": 25,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2024-08-14T20:26:00.027Z",
                "promotedArtists": [
                  {
                    "id": 8722,
                    "name": "Mark Ronson",
                    "type": "MAIN",
                    "picture": "0ee81a7f-93e5-47b6-a93f-08b30ba33e81",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 3658521,
                    "name": "Bruno Mars",
                    "type": "MAIN",
                    "picture": "72b86dab-cb66-4cb6-9610-187b94b3363a",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 5866385,
                    "name": "Anderson .Paak",
                    "type": "MAIN",
                    "picture": "167df412-60db-4821-b0dc-913c89adbda8",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "Multi-platinum, four-time Grammy-winner Bruno Mars is no stranger to chart-topping music. With his soulful voice and irresistible retro showmanship, the hit-making machine has captivated audiences worldwide, most recently releasing his third studio album, '24K Magic'. Check out some of Bruno Mars' best in this playlist. (Photo: Atlantic)"
              },
              {
                "uuid": "f8695823-a555-4a65-a6e5-a1bc69377393",
                "title": "Teddy Riley Essentials",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/f8695823-a555-4a65-a6e5-a1bc69377393",
                "image": "581a4e6a-179c-4265-8325-aa79d1543050",
                "squareImage": "c3b5eabf-8f07-46c6-ad85-3c9ea9112fbc",
                "duration": 15299,
                "numberOfTracks": 43,
                "numberOfVideos": 0,
                "lastItemAddedAt": null,
                "promotedArtists": [
                  {
                    "id": 32169,
                    "name": "Guy",
                    "type": "MAIN",
                    "picture": "ccd682e0-fc8c-4baf-89e2-79423e73d635",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 9496,
                    "name": "MC Hammer",
                    "type": "MAIN",
                    "picture": "04bee3e5-6783-41d7-844c-072f9a9b551b",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 6079,
                    "name": "Johnny Kemp",
                    "type": "MAIN",
                    "picture": "c30596ea-d3b9-4a42-86f9-81bdc070a895",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 3952747,
                    "name": "Wreckx-N-Effect",
                    "type": "MAIN",
                    "picture": "95f30045-9bad-4f1a-91cf-6a8df70eeaa8",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "GRAMMY®-winning producer/songwriter/mixer/engineer, Teddy Riley, is the king of new jack swing. Groove along to these Riley contributions. (Photo: Getty)"
              },
              {
                "uuid": "a495aa93-313d-4406-b864-285ad5ea97ff",
                "title": "Toni Braxton Essentials",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/a495aa93-313d-4406-b864-285ad5ea97ff",
                "image": "d309962c-127d-4b82-8b8a-62c2453fb9cc",
                "squareImage": "8a2fc85b-c01c-4ab4-ba49-e60588c7fd43",
                "duration": 6680,
                "numberOfTracks": 25,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2025-10-07T19:54:39.476Z",
                "promotedArtists": [
                  {
                    "id": 1731,
                    "name": "Toni Braxton",
                    "type": "MAIN",
                    "picture": "350f7da4-a4fa-4e35-bfbb-e21559417b56",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 2977,
                    "name": "Kenny G",
                    "type": "MAIN",
                    "picture": "ccf0d5f8-5a81-4064-b1dc-5fdd77a3dea4",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "After selling 10 million copies worldwide of her self-titled album in 1993 (home to the hits \"Breathe Again\" and \"Another Sad Love Song\"), the R&B powerhouse went on to release Billboard Hot 100 No. 1 hits \"Un-break My Heart\" and \"You're Makin' Me High.\" Unbreak your heart with her catalog's finest. (Photo: Rovi)"
              },
              {
                "uuid": "7bb7d887-e048-45ce-8e30-533cfa866638",
                "title": "Sting Essentials",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/7bb7d887-e048-45ce-8e30-533cfa866638",
                "image": "fd733615-10eb-41ab-9ba3-10039fd16634",
                "squareImage": "4a891435-e6f2-45e4-a913-5fe3bd5470dd",
                "duration": 6728,
                "numberOfTracks": 25,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2022-02-11T04:51:35.383Z",
                "promotedArtists": [
                  {
                    "id": 17356,
                    "name": "Sting",
                    "type": "MAIN",
                    "picture": "fb86a997-2ee5-4f0b-8b1f-38cb8aa4a5f5",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "Originally breaking through during the late ´70s as the lead singer, bassist and songwriter of The Police, Sting has created a solid back catalogue throughout his solo career–merging genres ranging from rock, to jazz, reggae and classical. (Photo: Rovi)"
              },
              {
                "uuid": "285d6293-8f77-4dc1-8dab-a262f3d0cb43",
                "title": "JAY-Z Essentials",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/285d6293-8f77-4dc1-8dab-a262f3d0cb43",
                "image": "687f1693-7b14-4260-9553-e1a9f743fb00",
                "squareImage": "6e2c2202-badf-489c-a81c-c779f2ddfdd8",
                "duration": 12079,
                "numberOfTracks": 50,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2023-07-21T05:50:16.973Z",
                "promotedArtists": [
                  {
                    "id": 7804,
                    "name": "JAY Z",
                    "type": "MAIN",
                    "picture": "a437c542-e9b5-4050-aced-c75ca0565c4f",
                    "handle": null,
                    "userId": 0,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "Hear the finest from the Brooklyn legend in this career-spanning collection by Hov. (Photo: Courtesy of Roc Nation, Press)"
              },
              {
                "uuid": "31506ed1-cbe4-4bc4-a3ec-e81e216c7a17",
                "title": "Halsey Essentials",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/31506ed1-cbe4-4bc4-a3ec-e81e216c7a17",
                "image": "fa1da629-a866-4c8d-b551-ba1e20b2f1a0",
                "squareImage": "ba5e6985-3d66-48e1-8089-d3f8c44ab1ee",
                "duration": 5135,
                "numberOfTracks": 25,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2023-09-30T23:11:47.097Z",
                "promotedArtists": [
                  {
                    "id": 5386709,
                    "name": "Halsey",
                    "type": "MAIN",
                    "picture": "e3480b04-578c-4420-90e3-2ea239ad9bae",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "Singer/Songwriter Halsey first stepped into the spotlight in 2015 with her album BADLANDS. Here are Halsey's best singles, features, and deep album album cuts.\n\n(Photo: Capitol Records, press)"
              },
              {
                "uuid": "41f588c8-09ac-434d-b536-66711edd9f58",
                "title": "Autumn Serenades",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/41f588c8-09ac-434d-b536-66711edd9f58",
                "image": "0485e6b1-d466-4964-902e-612fc970aca1",
                "squareImage": "b7cc4373-b61f-42a7-95b5-eb1be281a6f2",
                "duration": 24636,
                "numberOfTracks": 100,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2025-09-21T18:24:38.131Z",
                "promotedArtists": [
                  {
                    "id": 10979,
                    "name": "Mazzy Star",
                    "type": "MAIN",
                    "picture": "7e92d8d6-da16-4329-93b2-7a28bef4f880",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 205,
                    "name": "Neil Young",
                    "type": "MAIN",
                    "picture": "5bd389b5-a9fa-4fae-9a05-7f78ca80f5d0",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 4133018,
                    "name": "Jóhann Jóhannsson",
                    "type": "MAIN",
                    "picture": "1fdd6eed-b5a0-4e06-bfc0-209b8e9bcdeb",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 24899,
                    "name": "Diana Krall",
                    "type": "MAIN",
                    "picture": "d38a6b71-60c3-4e39-b7db-4bc55d67002d",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "With summer inevitably losing its grip and fall season slowly taking over, we welcome the darker, cooler days. This is music for a walk in a foggy park, a cosy evening with a book by the fire, or the soundtrack to rain gently knocking on your windows. (Photo: Unsplash)"
              },
              {
                "uuid": "239d76e6-6b43-440a-addf-1929b99b461b",
                "title": "LHM: Afro-Latinx",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/239d76e6-6b43-440a-addf-1929b99b461b",
                "image": "cfdb246d-1abf-4bc5-b564-f34549675373",
                "squareImage": "a66ace3e-45f3-4bf1-9c9a-5bb64fc9099b",
                "duration": 7912,
                "numberOfTracks": 40,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2025-08-26T15:41:06.943Z",
                "promotedArtists": [
                  {
                    "id": 6072,
                    "name": "Celia Cruz",
                    "type": "MAIN",
                    "picture": "0c333072-af5d-42f2-9a4b-f5c92a7d1a02",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 3584226,
                    "name": "Tego Calderón",
                    "type": "MAIN",
                    "picture": "1510093b-3e36-4c6e-a9e8-61071472f2bc",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 17110,
                    "name": "Don Omar",
                    "type": "MAIN",
                    "picture": "3b9dbc31-85d0-4290-80be-da8425bf5538",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  },
                  {
                    "id": 7301626,
                    "name": "Cardi B",
                    "type": "MAIN",
                    "picture": "c7fcbed0-97f1-4f8a-9d1a-615c29ac39ec",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "A celebration of the voices, rhythms, and stories at the heart of Latin music. From salsa legends like Celia Cruz and Ismael Rivera, to reggaetón pioneers like Tego Calderón and Don Omar, to global stars like Ozuna, Sech, and Cardi B. These songs showcase the richness of Afro-Latinidad across generations. (Photo: Celia Cruz/ Credit: Getty)"
              },
              {
                "uuid": "7023339b-5429-4956-b467-890a07c16499",
                "title": "Stevie Wonder Essentials",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/7023339b-5429-4956-b467-890a07c16499",
                "image": "223e6aab-b383-4293-b24a-a08d99259b35",
                "squareImage": "6faebc79-79e8-4ffb-ad1f-86bf050ab9b1",
                "duration": 6719,
                "numberOfTracks": 25,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2023-09-20T22:02:57.716Z",
                "promotedArtists": [
                  {
                    "id": 239,
                    "name": "Stevie Wonder",
                    "type": "MAIN",
                    "picture": "c9292b86-fcee-48bb-a1e5-84b99e147148",
                    "handle": null,
                    "userId": null,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "Motown legend and revered icon Stevie Wonder has delivered life anthems since he was known as Little Stevie Wonder. This toe-tapping collection of his best demonstrates his envelope-pushing genius as singer, songwriter and musician. (Photo: Republic Records)"
              },
              {
                "uuid": "5f81c527-cf94-46e7-8289-68e4530377a4",
                "title": "Jeezy Essentials",
                "type": "EDITORIAL",
                "url": "https://tidal.com/browse/playlist/5f81c527-cf94-46e7-8289-68e4530377a4",
                "image": "bec38d83-f7dd-4e2d-a2d1-5e5e751c9ea6",
                "squareImage": "b97de96d-a705-4f0c-8dc2-d9582d7ece01",
                "duration": 6067,
                "numberOfTracks": 25,
                "numberOfVideos": 0,
                "lastItemAddedAt": "2025-10-03T20:16:47.777Z",
                "promotedArtists": [
                  {
                    "id": 3724965,
                    "name": "Jeezy",
                    "type": "MAIN",
                    "picture": "03241cb4-6089-4385-8785-13898d96759b",
                    "handle": null,
                    "userId": 0,
                    "contributionLinkUrl": null
                  }
                ],
                "creators": [],
                "description": "Jeezy has built a legendary career over the last decade to become one of gangsta rap's most important voices. Here are the crowd favorites from the Atlanta legend's vast catalog. (Photo: Def Jam, Press)"
              }
            ]
          },
          "supportsPaging": false,
          "layout": null,
          "quickPlay": false,
          "playlistStyle": null,
          "preTitle": null
        }
      ]
    }
  ]
}
```
<br>


### Status Codes

HIFI returns the following status codes in its API:

> | Status Code | Description |
> | :---        | :--- |
> | 200         | `OK` |
> | 422         | `UNPROCESSABLE CONTENT` |
> | 404         | `NOT FOUND` |
> | 500         | `INTERNAL SERVER ERROR` |


</details>

------------------------------------------------------------------------------------------

<details>

 <summary><code>GET</code>   <code><b>/dash/</b></code> </summary>

## Request


<br>

> | Parameter  |   Type    | Description |
> |------------|-----------|-------------|
> | `id`       |  integer  | Track Id = `286266926` |
> | `quality`  |  string   | Song Quality = `HI_RES_LOSSLESS` or `HI_RES` or `LOSSLESS` or `HIGH` or `LOW ` |


<br>

#### Example

>xh

    https GET "https://hifi.401658.xyz/dash/?id=286266926&quality=HI_RES_LOSSLESS"

<br>


### Response

XML-encoded MPEG-DASH manifest. Play on VLC/FFMpeg/Google Shaka Player (see note below).

```xml
<?xml version='1.0' encoding='UTF-8'?>
<MPD
	xmlns="urn:mpeg:dash:schema:mpd:2011"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xmlns:xlink="http://www.w3.org/1999/xlink"
	xmlns:cenc="urn:mpeg:cenc:2013" xsi:schemaLocation="urn:mpeg:dash:schema:mpd:2011 DASH-MPD.xsd" profiles="urn:mpeg:dash:profile:isoff-main:2011" type="static" minBufferTime="PT3.925S" mediaPresentationDuration="PT3M14.607S">
	<Period id="0">
		<AdaptationSet id="0" contentType="audio" mimeType="audio/mp4" segmentAlignment="true">
			<Representation id="0" codecs="flac" bandwidth="1741283" audioSamplingRate="48000">
				<SegmentTemplate timescale="48000" initialization="https://sp-ad-fa.audio.tidal.com/mediatracks/GigIAxIkY2VmNTNmMWRkMDVhMmMxYmU2YTQ1N2ExNmJmNDZmY2YubXA0IiAdAACAQCACKhAKFXRoxO9p3fvNd34d7sokMgUNAACgQQ/0.mp4?token=1759786386~L21lZGlhdHJhY2tzL0dpZ0lBeElrWTJWbU5UTm1NV1JrTURWaE1tTXhZbVUyWVRRMU4yRXhObUptTkRabVkyWXViWEEwSWlBZEFBQ0FRQ0FDS2hBS0ZYUm94TzlwM2Z2TmQzNGQ3c29rTWdVTkFBQ2dRUS8qfmM4NzFkODM0NGYxNThlYmM1Y2U4ODkzYTI3YjgyNTQ2ZTIyYzc4NGU=" media="https://sp-ad-fa.audio.tidal.com/mediatracks/GigIAxIkY2VmNTNmMWRkMDVhMmMxYmU2YTQ1N2ExNmJmNDZmY2YubXA0IiAdAACAQCACKhAKFXRoxO9p3fvNd34d7sokMgUNAACgQQ/$Number$.mp4?token=1759786386~L21lZGlhdHJhY2tzL0dpZ0lBeElrWTJWbU5UTm1NV1JrTURWaE1tTXhZbVUyWVRRMU4yRXhObUptTkRabVkyWXViWEEwSWlBZEFBQ0FRQ0FDS2hBS0ZYUm94TzlwM2Z2TmQzNGQ3c29rTWdVTkFBQ2dRUS8qfmM4NzFkODM0NGYxNThlYmM1Y2U4ODkzYTI3YjgyNTQ2ZTIyYzc4NGU=" startNumber="1">
					<SegmentTimeline>
						<S d="188416" r="48"/>
						<S d="108786"/>
					</SegmentTimeline>
				</SegmentTemplate>
			</Representation>
		</AdaptationSet>
	</Period>
</MPD>
```
#### MPV Player

```sh
 mpv https://hifi.401658.xyz/dash/?id=286266926&quality=HI_RES_LOSSLESS

```

<img width="913" height="218" alt="image" src="https://github.com/user-attachments/assets/c6d0cad4-fc44-412b-8f0e-0ad870cc3ae2" />


#### Google Shaka Player

The API does not support HEAD requests so use this code to disable them:

```js
player.getNetworkingEngine().registerRequestFilter(function(type, request) {
    // Convert any HEAD requests to GET
    if (request.method === 'HEAD') {
        request.method = 'GET';
    }
});
```


### Status Codes

HIFI returns the following status codes in its API:

> | Status Code | Description |
> | :---        | :--- |
> | 200         | `OK` |
> | 422         | `UNPROCESSABLE CONTENT` |
> | 404         | `NOT FOUND` |
> | 500         | `INTERNAL SERVER ERROR` |


</details>

------------------------------------------------------------------------------------------



<details>

 <summary><code>GET</code>   <code><b>/search/</b></code> </summary>

## Request

<br>

> | Parameter | Type | Description |
> |------|--------|-------------|
> | `s`  | string |  Name = `Spaceship`|
> | `a`  | string |  Artist Name = `Kanye West`|
> | `al` | string |  Album Name = `Late Registration`|
> | `v`  | string |  Video Name = `Spaceship`|
> | `p`  | string |  Playlist Name = `Pop Hits`|
> |`li`  | integer | Limit = `125` |
> | `o`  | integer | Offset = `0` |

<br>

#### Example
>xh

    https GET "https://hifi.401658.xyz/search/?s=Consequence"

![2023-11-19_03-05](https://github.com/sachinsenal0x64/Hifi-Tui/assets/127573781/35041774-394c-4b17-9cfd-927e5e113da3)

<br>


### Response

```json

{
  "limit": 1,
  "offset": 0,
  "totalNumberOfItems": 200,
  "items": [
    {
      "id": 82448461,
      "title": "Consequence",
      "duration": 313,
      "replayGain": -9.88,
      "peak": 1,
      "allowStreaming": true,
      "streamReady": true,
      "streamStartDate": "2017-12-05T00:00:00.000+0000",
      "premiumStreamingOnly": false,
      "trackNumber": 10,
      "volumeNumber": 1,
      "version": null,
      "popularity": 6,
      "copyright": "City Slang/big Store",
      "url": "http://www.tidal.com/track/82448461",
      "isrc": "DED620118410",
      "editable": false,
      "explicit": false,
      "audioQuality": "LOSSLESS",
      "audioModes": [
        "STEREO"
      ],
      "artist": {
        "id": 3529689,
        "name": "The Notwist",
        "type": "MAIN"
      },
      "artists": [
        {
          "id": 3529689,
          "name": "The Notwist",
          "type": "MAIN"
        }
      ],
      "album": {
        "id": 82448449,
        "title": "Neon Golden",
        "cover": "ad3ed5f3-37a2-4b27-9002-b83459ab5a0e",
        "videoCover": null
      },
      "mixes": {
        "TRACK_MIX": "001981d70c53d5448599714c407079"
      }
    }
  ]
}

```

<br>

### Status Codes

HIFI returns the following status codes in its API:

> | Status Code | Description |
> | :---        | :--- |
> | 200         | `OK` |
> | 422         | `UNPROCESSABLE CONTENT` |
> | 404         | `NOT FOUND` |
> | 500         | `INTERNAL SERVER ERROR` |


------------------------------------------------------------------------------------------

</details>


------------------------------------------------------------------------------------------

<details>

<summary><code>GET</code> <code><b>/cover/</b></code></summary>


## Request


<br>

> | Parameter  |   Type    | Description |
> |------------|-----------|-------------|
> | `id`       |  integer  | Track Id = `328060990` |
> | `q`     |  string      | Song Name = `Maestro` |

<br>

> #### Sizes = `1280px | 640px | 80px `

<br>

#### Example

>xh

    https GET "https://hifi.401658.xyz/cover/?q=Maestro"
    https GET "https://hifi.401658.xyz/cover/?id=328060990"
    

![image](https://github.com/sachinsenal0x64/Hifi-Tui/assets/127573781/42b43878-00c5-4d35-8210-2cca466bc594)


<br>


### Response

  ```json
[
    {
        "1280": "https://resources.tidal.com/images/6f5c52be/c21c/4fb7/9ce6/0c270f6f1a5a/1280x1280.jpg",
        "640": "https://resources.tidal.com/images/6f5c52be/c21c/4fb7/9ce6/0c270f6f1a5a/640x640.jpg",
        "80": "https://resources.tidal.com/images/6f5c52be/c21c/4fb7/9ce6/0c270f6f1a5a/80x80.jpg",
        "id": 328060988,
        "name": "Maestro: Music by Leonard Bernstein (Original Soundtrack / Dolby Atmos)"
    }
]
```

</details>

------------------------------------------------------------------------------------------


<details>


 <summary><code>GET</code>   <code><b>/song/</b></code> </summary>

## Request


<br>

> | Parameter  |   Type    | Description |
> |------------|-----------|-------------|
> | `q`        |  string   | Song Query   =  `Mine` |
> | `quality`  |  string   | Song Quality = `HI_RES_LOSSLESS` or `HI_RES` or `LOSSLESS` or `HIGH` or `LOW ` |


<br>

#### Example

>xh

    https GET "https://hifi.401658.xyz/song/?q=Mine&quality=HI_RES"
    

![image](https://sachinsenal0x64.github.io/picx-images-hosting/2024-02-07-20:19:04-screenshot.5zw9tsa19wcg.webp)

<br>

### Response

  ```json
{
    "OriginalTrackUrl": "track url",
    "Song Info": {
        "adSupportedStreamReady": true,
        "album": {
            "cover": "22b8ce2a-1912-4fc6-956f-3be5eb4a7f4c",
            "id": 79712262,
            "title": "Mine",
            "vibrantColor": "#a7d9fc",
            "videoCover": null
        },
        "allowStreaming": true,
        "artist": {
            "id": 7384212,
            "name": "Bazzi",
            "picture": "2726f1e5-0435-4c49-a6f7-c2192544638f",
            "type": "MAIN"
        },
        "artists": [
            {
                "id": 7384212,
                "name": "Bazzi",
                "picture": "2726f1e5-0435-4c49-a6f7-c2192544638f",
                "type": "MAIN"
            }
        ],
        "audioModes": [
            "STEREO"
        ],
        "audioQuality": "HI_RES",
        "copyright": "2017",
        "djReady": true,
        "duration": 134,
        "editable": false,
        "explicit": true,
        "id": 79712263,
        "isrc": "USAT21704227",
        "mediaMetadata": {
            "tags": [
                "LOSSLESS",
                "MQA"
            ]
        },
        "mixes": {
            "TRACK_MIX": "0014833cd62b1eecd3b24115e5f8d4"
        },
        "peak": 0.997437,
        "popularity": 64,
        "premiumStreamingOnly": false,
        "replayGain": -10.39,
        "stemReady": false,
        "streamReady": true,
        "streamStartDate": "2017-10-12T00:00:00.000+0000",
        "title": "Mine",
        "trackNumber": 1,
        "url": "http://www.tidal.com/track/79712263",
        "version": null,
        "volumeNumber": 1
    },
    "Track Info": {
        "albumPeakAmplitude": 0.997437,
        "albumReplayGain": -10.39,
        "assetPresentation": "FULL",
        "audioMode": "STEREO",
        "audioQuality": "HI_RES",
        "manifest": "base64 manifest",
        "manifestMimeType": "application/vnd.tidal.bts",
        "trackId": 79712263,
        "trackPeakAmplitude": 0.997437,
        "trackReplayGain": -10.39
    }
}

```

### Status Codes

HIFI returns the following status codes in its API:

> | Status Code | Description |
> | :---        | :--- |
> | 200         | `OK` |
> | 422         | `UNPROCESSABLE CONTENT` |
> | 404         | `NOT FOUND` |
> | 500         | `INTERNAL SERVER ERROR` |


</details>

------------------------------------------------------------------------------------------

<details>

 <summary><code>GET</code>   <code><b>/album/</b></code> </summary>

## Request


<br>

> | Parameter  |   Type    | Description |
> |------------|-----------|-------------|
> | `id`       |  integer  | Album Id = `286266925` |


<br>

#### Example

>xh

    https GET "https://hifi.401658.xyz/album/?id=286266925"
    

![image](https://github.com/sachinsenal0x64/picx-images-hosting/raw/master/2024-02-20-19:33:52-screenshot.99t2w3gelf.webp)

<br>


### Response

  ```json
 {
        "adSupportedStreamReady": true,
        "allowStreaming": true,
        "artist": {
            "id": 7162333,
            "name": "Dua Lipa",
            "picture": "28047130-6ada-4955-b3b9-65bed4508618",
            "type": "MAIN"
        },
        "artists": [
            {
                "id": 7162333,
                "name": "Dua Lipa",
                "picture": "28047130-6ada-4955-b3b9-65bed4508618",
                "type": "MAIN"
            }
        ],
        "audioModes": [
            "SONY_360RA"
        ],
        "audioQuality": "LOW",
        "copyright": "℗ 2017 Dua Lipa Limited under exclusive license to Warner Music UK Limited",
        "cover": "deae7f19-5da7-4d73-97be-ce901911c939",
        "djReady": true,
        "duration": 2456,
        "explicit": false,
        "id": 157117504,
        "mediaMetadata": {
            "tags": [
                "SONY_360RA"
            ]
        },
        "numberOfTracks": 12,
        "numberOfVideos": 0,
        "numberOfVolumes": 1,
        "popularity": 36,
        "premiumStreamingOnly": false,
        "releaseDate": "2020-10-05",
        "stemReady": false,
        "streamReady": true,
        "streamStartDate": "2020-10-05T00:00:00.000+0000",
        "title": "Dua Lipa (360 Reality Audio)",
        "type": "ALBUM",
        "upc": "190295160180",
        "url": "http://www.tidal.com/album/157117504",
        "version": null,
        "vibrantColor": "#6d99c6",
        "videoCover": null
    }
```
<br>


### Status Codes

HIFI returns the following status codes in its API:

> | Status Code | Description |
> | :---        | :--- |
> | 200         | `OK` |
> | 422         | `UNPROCESSABLE CONTENT` |
> | 404         | `NOT FOUND` |
> | 500         | `INTERNAL SERVER ERROR` |


</details>

------------------------------------------------------------------------------------------

<details>

 <summary><code>GET</code>   <code><b>/playlist/</b></code> </summary>

## Request


<br>

> | Parameter  |   Type    | Description |
> |------------|-----------|-------------|
> | `id`       |  string  | Playlist UUID = `910c525f-be8a-41a1-b557-2682af2bcef3` |


<br>

#### Example

>xh

    https GET "https://hifi.401658.xyz/playlist/?id=910c525f-be8a-41a1-b557-2682af2bcef3"
    

![image](https://sachinsenal0x64.github.io/picx-images-hosting/2024-02-20-23:15:31-screenshot.67x6v3b7q9.webp)

<br>


### Response

  ```json
 {
        "created": "2015-04-14T16:32:14.636+0000",
        "creator": {
            "id": 5034071,
            "name": "VIC MENSA",
            "picture": "cdd212a2-dadc-466d-9703-7216a9f66da1",
            "type": null
        },
        "description": "",
        "duration": 2696,
        "image": "c41cfe9b-cda1-4364-b517-f6a706741d24",
        "lastItemAddedAt": null,
        "lastUpdated": "2020-03-24T12:27:23.941+0000",
        "numberOfTracks": 11,
        "numberOfVideos": 0,
        "popularity": 0,
        "promotedArtists": [
            {
                "id": 5034071,
                "name": "VIC MENSA",
                "picture": null,
                "type": "MAIN"
            },
            {
                "id": 25022,
                "name": "Kanye West",
                "picture": null,
                "type": "MAIN"
            },
            {
                "id": 3899583,
                "name": "Theophilus London",
                "picture": null,
                "type": "MAIN"
            },
            {
                "id": 5637986,
                "name": "Allan Kingdom",
                "picture": null,
                "type": "MAIN"
            }
        ],
        "publicPlaylist": false,
        "squareImage": "03750282-401b-481c-bf60-55d6ee9fcc27",
        "title": "My Playlist",
        "type": "ARTIST",
        "url": "http://www.tidal.com/playlist/910c525f-be8a-41a1-b557-2682af2bcef3",
        "uuid": "910c525f-be8a-41a1-b557-2682af2bcef3"
    },
    {
        "items": [
            {
                "cut": null,
                "item": {
                    "adSupportedStreamReady": true,
                    "album": {
                        "cover": "43929b37-df27-4e1a-81b2-70692c058674",
                        "id": 44590541,
                        "releaseDate": "2015-04-16",
                        "title": "U Mad",
                        "vibrantColor": "#FFFFFF",
                        "videoCover": null
                    },
                    "allowStreaming": true,
                    "artist": {
                        "id": 5034071,
                        "name": "VIC MENSA",
                        "picture": "cdd212a2-dadc-466d-9703-7216a9f66da1",
                        "type": "MAIN"
                    },
                    "artists": [
                        {
                            "id": 5034071,
                            "name": "VIC MENSA",
                            "picture": "cdd212a2-dadc-466d-9703-7216a9f66da1",
                            "type": "MAIN"
                        },
                        {
                            "id": 25022,
                            "name": "Kanye West",
                            "picture": "26076dbd-7361-40d3-9335-f944d2c49ea6",
                            "type": "FEATURED"
                        }
                    ],
                    "audioModes": [
                        "STEREO"
                    ],
                    "audioQuality": "LOSSLESS",
                    "copyright": "(C) 2015 Roc Nation Records, LLC",
                    "dateAdded": "2015-04-15T15:03:19.696+0000",
                    "description": null,
                    "djReady": true,
                    "duration": 300,
                    "editable": false,
                    "explicit": true,
                    "id": 44590542,
                    "index": 0,
                    "isrc": "QMJMT1500671",
                    "itemUuid": "90545040-acc7-44c1-9481-7e48f36cefe8",
                    "mediaMetadata": {
                        "tags": [
                            "LOSSLESS"
                        ]
                    },
                    "mixes": {
                        "TRACK_MIX": "00169d5b613bbc32050146c8be21df"
                    },
                    "peak": 0.999359,
                    "popularity": 47,
                    "premiumStreamingOnly": false,
                    "replayGain": -9.38,
                    "stemReady": false,
                    "streamReady": true,
                    "streamStartDate": "2015-04-10T00:00:00.000+0000",
                    "title": "U Mad",
                    "trackNumber": 1,
                    "url": "http://www.tidal.com/track/44590542",
                    "version": null,
                    "volumeNumber": 1
                },
                "type": "track"
            },
```
<br>


### Status Codes

HIFI returns the following status codes in its API:

> | Status Code | Description |
> | :---        | :--- |
> | 200         | `OK` |
> | 422         | `UNPROCESSABLE CONTENT` |
> | 404         | `NOT FOUND` |
> | 500         | `INTERNAL SERVER ERROR` |


</details>

------------------------------------------------------------------------------------------


<details>

 <summary><code>GET</code>   <code><b>/artist/</b></code> </summary>

## Request


<br>

> | Parameter  |   Type    | Description |
> |------------|-----------|-------------|
> | `id`       |  integer  |  Artist ID = `5034071` |
> | `f`        | integer   |  Artist ID = `5034071`, This will retrieve all the albums and tracks.| 


<br>

#### Example

>xh

    https GET "https://hifi.401658.xyz/artist/?id=5034071"
    

![image](https://sachinsenal0x64.github.io/picx-images-hosting/2024-02-21-21:19:27-screenshot.1aoq2k57al.webp)

<br>


### Response

  ```json
 [
    {
        "artistRoles": [
            {
                "category": "Artist",
                "categoryId": -1
            },
            {
                "category": "Songwriter",
                "categoryId": 2
            },
            {
                "category": "Production team",
                "categoryId": 10
            },
            {
                "category": "Producer",
                "categoryId": 1
            },
            {
                "category": "Engineer",
                "categoryId": 3
            },
            {
                "category": "Performer",
                "categoryId": 11
            }
        ],
        "artistTypes": [
            "ARTIST",
            "CONTRIBUTOR"
        ],
        "id": 5034071,
        "mixes": {
            "ARTIST_MIX": "000720bd7d7867c71a4c63b1fe61cf"
        },
        "name": "VIC MENSA",
        "picture": "cdd212a2-dadc-466d-9703-7216a9f66da1",
        "popularity": 66,
        "url": "http://www.tidal.com/artist/5034071"
    },
    [
        {
            "750": "https://resources.tidal.com/images/cdd212a2/dadc/466d/9703/7216a9f66da1/750x750.jpg",
            "id": 5034071,
            "name": "VIC MENSA"
        }
    ]
]

```
<br>


### Status Codes

HIFI returns the following status codes in its API:

> | Status Code | Description |
> | :---        | :--- |
> | 200         | `OK` |
> | 422         | `UNPROCESSABLE CONTENT` |
> | 404         | `NOT FOUND` |
> | 500         | `INTERNAL SERVER ERROR` |


</details>

------------------------------------------------------------------------------------------


<details>

 <summary><code>GET</code>   <code><b>/lyrics/</b></code> </summary>

## Request


<br>

> | Parameter  |   Type    | Description |
> |------------|-----------|-------------|
> | `id`       |  integer  | Track ID = `286266926` |


<br>

#### Example

>xh

    https GET "https://hifi.401658.xyz/lyrics/?id=286266926"
    

![image](https://github.com/sachinsenal0x64/picx-images-hosting/raw/master/2024-04-11-14:31:30-screenshot.3go6nb7skh.webp)

<br>


### Response

  ```json
[
    {
        "isRightToLeft": false,
        "lyrics": "You know I question every motive, everything you say\nThought with you, maybe my heart wasn't meant to break\nCan't believe I let you in, I can't believe I stayed\nAs long as I stayed, yeah\n\nI hope one day someone will take your heart and hold it tight\nMake you feel like you're invincible deep inside\nAnd right when you think that it's perfect, they cross a line\nAnd steal your shine\nLike you did mine\n\nGo ahead and break my heart, that's fine\nSo unkind\nEternal sunshine of the spotless mind\nOh, love is blind\nWhy am I missin' you tonight?\nWas it all a lie?\nSomeone's gonna show you how a heart can be used\nLike you did mine\n\nAnd now, I second-guess my thoughts, every step I take\nI'm losin' hope in love, and I've lost all in faith\nYeah, for a dreamer, I just close my eyes and it's all blank\nI have you to thank, yeah\n\nI hope one day someone will take your heart and hold it tight\nMake you feel like you're invincible deep inside\nAnd right when you think that you'll try again, they cross a line\nAnd steal your shine\nLike you did mine\n\nGo ahead and break my heart, that's fine\nSo unkind\nEternal sunshine of the spotless mind\nOh, love is blind\nWhy am I missin' you tonight?\nWas it all a lie?\nSomeone's gonna show you how a heart can be used\n\nAnd you keep talkin', it'll come back, karma\nIs the truth, I don't make you look bad, you do, darlin'\nSabotage, your choice of art\nWho the hell do you think you are?\n\nGo ahead and break my heart, that's fine\nEternal sunshine of the spotless mind\nWhy am I missin' you tonight?\nSomeone's gonna show you how a heart can be used\nLike you did mine",
        "lyricsProvider": "MUSIXMATCH",
        "providerCommontrackId": "158550594",
        "providerLyricsId": "31545572",
        "subtitles": "[00:00.48] You know I question every motive, everything you say\n[00:04.51] Thought with you, maybe my heart wasn't meant to break\n[00:08.57] Can't believe I let you in, I can't believe I stayed\n[00:13.40] As long as I stayed, yeah\n[00:16.35] I hope one day someone will take your heart and hold it tight\n[00:20.64] Make you feel like you're invincible deep inside\n[00:24.45] And right when you think that it's perfect, they cross a line\n[00:28.97] And steal your shine\n[00:30.77] Like you did mine\n[00:33.04] Go ahead and break my heart, that's fine\n[00:38.97] So unkind\n[00:40.97] Eternal sunshine of the spotless mind\n[00:46.48] Oh, love is blind\n[00:49.13] Why am I missin' you tonight?\n[00:54.53] Was it all a lie?\n[00:57.18] Someone's gonna show you how a heart can be used\n[01:02.86] Like you did mine\n[01:06.11] \n[01:08.32] And now, I second-guess my thoughts, every step I take\n[01:12.67] I'm losin' hope in love, and I've lost all in faith\n[01:16.42] Yeah, for a dreamer, I just close my eyes and it's all blank\n[01:21.56] I have you to thank, yeah\n[01:24.30] I hope one day someone will take your heart and hold it tight\n[01:28.57] Make you feel like you're invincible deep inside\n[01:32.43] And right when you think that you'll try again, they cross a line\n[01:36.96] And steal your shine\n[01:39.06] Like you did mine\n[01:41.11] Go ahead and break my heart, that's fine\n[01:47.09] So unkind\n[01:49.03] Eternal sunshine of the spotless mind\n[01:54.52] Oh, love is blind\n[01:57.17] Why am I missin' you tonight?\n[02:02.52] Was it all a lie?\n[02:04.93] Someone's gonna show you how a heart can be used\n[02:11.65] And you keep talkin', it'll come back, karma\n[02:17.64] Is the truth, I don't make you look bad, you do, darlin'\n[02:23.79] Sabotage, your choice of art\n[02:29.90] Who the hell do you think you are?\n[02:37.59] Go ahead and break my heart, that's fine\n[02:44.06] Eternal sunshine of the spotless mind\n[02:49.03] \n[02:51.56] Why am I missin' you tonight?\n[02:56.72] \n[02:59.23] Someone's gonna show you how a heart can be used\n[03:05.24] Like you did mine\n[03:06.49] ",
        "trackId": 286266926
    }
]


```
<br>


### Status Codes

HIFI returns the following status codes in its API:

> | Status Code | Description |
> | :---        | :--- |
> | 200         | `OK` |
> | 422         | `UNPROCESSABLE CONTENT` |
> | 404         | `NOT FOUND` |
> | 500         | `INTERNAL SERVER ERROR` |


</details>

------------------------------------------------------------------------------------------


<details>

 <summary><code>GET</code>   <code><b>/mix/</b></code> </summary>

## Request


<br>

> | Parameter  |   Type    | Description | Optional|
> |------------|-----------|-------------|-------------|	
> | `id`       |  string   | Mix ID = `286266926000bc267583f06c15474662e5a83ae` | `false`|
> | `country`  |  string   | Country = `AU`| `true`|


<br>

#### Example

>xh

    https GET "https://hifi.401658.xyz/mix/?id=000bc267583f06c15474662e5a83ae&country=us"
    
<br>


### Response

```json
   
 {
  "limit": 100,
  "offset": 0,
  "totalNumberOfItems": 100,
  "items": [
    {
      "item": {
        "id": 3140991,
        "title": "Fireflies",
        "duration": 228,
        "replayGain": -9.11,
        "peak": 0.961304,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-07-28T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 9,
        "volumeNumber": 1,
        "version": null,
        "popularity": 85,
        "copyright": "℗ 2009 Universal Republic Records, a division of UMG Recordings, Inc.",
        "bpm": 90,
        "url": "http://www.tidal.com/track/3140991",
        "isrc": "USUM70972068",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3582596,
          "name": "Owl City",
          "handle": null,
          "type": "MAIN",
          "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
        },
        "artists": [
          {
            "id": 3582596,
            "name": "Owl City",
            "handle": null,
            "type": "MAIN",
            "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
          }
        ],
        "album": {
          "id": 3140982,
          "title": "Ocean Eyes",
          "cover": "76b92beb-399c-4983-9b91-0eef89c796e1",
          "vibrantColor": "#305c86",
          "videoCover": "0812e3f6-bd0a-40b8-ba96-f6f2d21c8ebb"
        },
        "mixes": {
          "TRACK_MIX": "00102fba8ce483de0f908f4c37c889"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 84970907,
        "title": "Hey, Soul Sister",
        "duration": 217,
        "replayGain": -11,
        "peak": 0.997345,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2018-03-02T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 18,
        "volumeNumber": 1,
        "version": null,
        "popularity": 83,
        "copyright": "(P) 2009 Columbia Records, a division of Sony Music Entertainment",
        "bpm": 97,
        "url": "http://www.tidal.com/track/84970907",
        "isrc": "USSM10904113",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 1186,
          "name": "Train",
          "handle": null,
          "type": "MAIN",
          "picture": "4e624013-d349-4d9e-9da7-56031917bb7e"
        },
        "artists": [
          {
            "id": 1186,
            "name": "Train",
            "handle": null,
            "type": "MAIN",
            "picture": "4e624013-d349-4d9e-9da7-56031917bb7e"
          }
        ],
        "album": {
          "id": 84970889,
          "title": "Mother's Day Songs",
          "cover": "86f50e0d-953f-4305-b5d6-a52bf858b77e",
          "vibrantColor": "#fb8485",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00140ca8a6e427d4da22227c12ab95"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3413724,
        "title": "Whatcha Say",
        "duration": 221,
        "replayGain": -11.65,
        "peak": 0.899933,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2010-03-02T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 83,
        "copyright": "℗ 2009 Warner Records Inc.",
        "bpm": 144,
        "url": "http://www.tidal.com/track/3413724",
        "isrc": "USWB10901504",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3851664,
          "name": "Jason Derulo",
          "handle": null,
          "type": "MAIN",
          "picture": "b1a1cb60-f9a3-4249-9900-652cdde9bbd1"
        },
        "artists": [
          {
            "id": 3851664,
            "name": "Jason Derulo",
            "handle": null,
            "type": "MAIN",
            "picture": "b1a1cb60-f9a3-4249-9900-652cdde9bbd1"
          }
        ],
        "album": {
          "id": 3413723,
          "title": "Jason Derulo",
          "cover": "3659764d-fa06-414a-9277-c172d59809d9",
          "vibrantColor": "#99b7d7",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001fb7f887501a46770e2d0b3c364f"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3698265,
        "title": "All The Right Moves",
        "duration": 238,
        "replayGain": -10.92,
        "peak": 0.988647,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-11-17T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 73,
        "copyright": "℗ 2009 UMG Recordings, Inc.",
        "bpm": null,
        "url": "http://www.tidal.com/track/3698265",
        "isrc": "USUM70984099",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 64665,
          "name": "OneRepublic",
          "handle": null,
          "type": "MAIN",
          "picture": "fadac3b8-76b0-4723-bae2-e504bc911e17"
        },
        "artists": [
          {
            "id": 64665,
            "name": "OneRepublic",
            "handle": null,
            "type": "MAIN",
            "picture": "fadac3b8-76b0-4723-bae2-e504bc911e17"
          }
        ],
        "album": {
          "id": 3698263,
          "title": "Waking Up",
          "cover": "c6519e20-1308-472e-94d1-5bbca0246d16",
          "vibrantColor": "#e1d29f",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001ed2ac3ff9a84e3bc6cbe8f47bd5"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35748754,
        "title": "Pork And Beans",
        "duration": 189,
        "replayGain": -9.13,
        "peak": 0.98233,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2008-06-03T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 3,
        "volumeNumber": 1,
        "version": null,
        "popularity": 70,
        "copyright": "℗ 2008 Interscope Records",
        "bpm": 119,
        "url": "http://www.tidal.com/track/35748754",
        "isrc": "USUM70811232",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 30157,
          "name": "Weezer",
          "handle": null,
          "type": "MAIN",
          "picture": "defeed86-e0f2-40f2-8ea4-4a0c42d2642f"
        },
        "artists": [
          {
            "id": 30157,
            "name": "Weezer",
            "handle": null,
            "type": "MAIN",
            "picture": "defeed86-e0f2-40f2-8ea4-4a0c42d2642f"
          }
        ],
        "album": {
          "id": 35748751,
          "title": "Weezer (Red Album)",
          "cover": "cfe2a03c-3b02-4bcc-b43f-60e4c37324bf",
          "vibrantColor": "#f49688",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001105168c337edeacb70cd7370ef5"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3140985,
        "title": "Hello Seattle",
        "duration": 167,
        "replayGain": -9.11,
        "peak": 0.961304,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-07-28T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 3,
        "volumeNumber": 1,
        "version": null,
        "popularity": 61,
        "copyright": "℗ 2007 Owl City Catalog / Federal Distribution",
        "bpm": 119,
        "url": "http://www.tidal.com/track/3140985",
        "isrc": "USZXT0835577",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3582596,
          "name": "Owl City",
          "handle": null,
          "type": "MAIN",
          "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
        },
        "artists": [
          {
            "id": 3582596,
            "name": "Owl City",
            "handle": null,
            "type": "MAIN",
            "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
          }
        ],
        "album": {
          "id": 3140982,
          "title": "Ocean Eyes",
          "cover": "76b92beb-399c-4983-9b91-0eef89c796e1",
          "vibrantColor": "#305c86",
          "videoCover": "0812e3f6-bd0a-40b8-ba96-f6f2d21c8ebb"
        },
        "mixes": {
          "TRACK_MIX": "00143d934edca298bda87926313dbf"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 1311661,
        "title": "Drops of Jupiter (Tell Me)",
        "duration": 260,
        "replayGain": -9.88,
        "peak": 0.998825,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2001-03-27T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 3,
        "volumeNumber": 1,
        "version": null,
        "popularity": 86,
        "copyright": "(P) 2001 Columbia Records, a division of Sony Music Entertainment",
        "bpm": 162,
        "url": "http://www.tidal.com/track/1311661",
        "isrc": "USSM10019751",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS",
            "HIRES_LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 1186,
          "name": "Train",
          "handle": null,
          "type": "MAIN",
          "picture": "4e624013-d349-4d9e-9da7-56031917bb7e"
        },
        "artists": [
          {
            "id": 1186,
            "name": "Train",
            "handle": null,
            "type": "MAIN",
            "picture": "4e624013-d349-4d9e-9da7-56031917bb7e"
          }
        ],
        "album": {
          "id": 1310832,
          "title": "Drops Of Jupiter",
          "cover": "b0471b65-84e3-414c-b559-14c18eda9397",
          "vibrantColor": "#fcdca9",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00198a78a341d09bd061025967b73e"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3069302,
        "title": "Replay",
        "duration": 182,
        "replayGain": -7.95,
        "peak": 0.997314,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2016-11-10T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 83,
        "copyright": "℗ 2009 Reprise Records",
        "bpm": 91,
        "url": "http://www.tidal.com/track/3069302",
        "isrc": "USRE10901161",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3634440,
          "name": "Iyaz",
          "handle": null,
          "type": "MAIN",
          "picture": "dd73e877-9079-4d80-9e62-68aac67f920b"
        },
        "artists": [
          {
            "id": 3634440,
            "name": "Iyaz",
            "handle": null,
            "type": "MAIN",
            "picture": "dd73e877-9079-4d80-9e62-68aac67f920b"
          }
        ],
        "album": {
          "id": 3069301,
          "title": "Replay",
          "cover": "0939aa12-e5e5-4ad6-a99d-e510f201e0a3",
          "vibrantColor": "#dbddb8",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001db6134a14c3e56548b2b5723329"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3413726,
        "title": "In My Head",
        "duration": 199,
        "replayGain": -11.65,
        "peak": 0.792389,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2010-03-02T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 3,
        "volumeNumber": 1,
        "version": null,
        "popularity": 76,
        "copyright": "℗ 2010 Warner Records Inc.",
        "bpm": 110,
        "url": "http://www.tidal.com/track/3413726",
        "isrc": "USWB10904633",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3851664,
          "name": "Jason Derulo",
          "handle": null,
          "type": "MAIN",
          "picture": "b1a1cb60-f9a3-4249-9900-652cdde9bbd1"
        },
        "artists": [
          {
            "id": 3851664,
            "name": "Jason Derulo",
            "handle": null,
            "type": "MAIN",
            "picture": "b1a1cb60-f9a3-4249-9900-652cdde9bbd1"
          }
        ],
        "album": {
          "id": 3413723,
          "title": "Jason Derulo",
          "cover": "3659764d-fa06-414a-9277-c172d59809d9",
          "vibrantColor": "#99b7d7",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001fa96745d078def75c783a4b1c58"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 36039914,
        "title": "You Belong With Me",
        "duration": 231,
        "replayGain": -11.06,
        "peak": 0.99173,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-10-26T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 12,
        "volumeNumber": 1,
        "version": null,
        "popularity": 80,
        "copyright": "℗ 2008 Big Machine Records, LLC",
        "bpm": 130,
        "url": "http://www.tidal.com/track/36039914",
        "isrc": "USCJY0803328",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3557299,
          "name": "Taylor Swift",
          "handle": null,
          "type": "MAIN",
          "picture": "ffa46767-8d63-42bd-9dc7-827881ef460c"
        },
        "artists": [
          {
            "id": 3557299,
            "name": "Taylor Swift",
            "handle": null,
            "type": "MAIN",
            "picture": "ffa46767-8d63-42bd-9dc7-827881ef460c"
          }
        ],
        "album": {
          "id": 36039902,
          "title": "Fearless Platinum Edition",
          "cover": "a59a0f5f-3782-46be-abe5-040c326500a4",
          "vibrantColor": "#e0cda1",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001f4fbe5dcc7f287d242c34391125"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 23849830,
        "title": "Airplanes (feat. Hayley Williams of Paramore)",
        "duration": 180,
        "replayGain": -10.61,
        "peak": 0.999969,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2010-04-27T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 4,
        "volumeNumber": 1,
        "version": null,
        "popularity": 84,
        "copyright": "℗ 2010 Atlantic Recording Corporation for the United States and WEA International Inc. for the world outside of the United States",
        "bpm": 93,
        "url": "http://www.tidal.com/track/23849830",
        "isrc": "USAT21000477",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3667959,
          "name": "B.o.B",
          "handle": null,
          "type": "MAIN",
          "picture": "530f127a-9fd5-4b8c-b2f9-cdde70873130"
        },
        "artists": [
          {
            "id": 3667959,
            "name": "B.o.B",
            "handle": null,
            "type": "MAIN",
            "picture": "530f127a-9fd5-4b8c-b2f9-cdde70873130"
          },
          {
            "id": 3653311,
            "name": "Hayley Williams",
            "handle": null,
            "type": "FEATURED",
            "picture": "c71605cf-3169-4224-bb01-610222b08dd7"
          }
        ],
        "album": {
          "id": 23849826,
          "title": "B.o.B Presents: The Adventures of Bobby Ray",
          "cover": "85b28a9c-c161-4fe5-9c4c-ea8ee8c0da57",
          "vibrantColor": "#e7e452",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00120629d4b164da453ce12fc6eeb9"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3698266,
        "title": "Secrets",
        "duration": 225,
        "replayGain": -10.92,
        "peak": 0.988647,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-11-17T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 3,
        "volumeNumber": 1,
        "version": null,
        "popularity": 81,
        "copyright": "℗ 2009 UMG Recordings, Inc.",
        "bpm": 148,
        "url": "http://www.tidal.com/track/3698266",
        "isrc": "USUM70985644",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 64665,
          "name": "OneRepublic",
          "handle": null,
          "type": "MAIN",
          "picture": "fadac3b8-76b0-4723-bae2-e504bc911e17"
        },
        "artists": [
          {
            "id": 64665,
            "name": "OneRepublic",
            "handle": null,
            "type": "MAIN",
            "picture": "fadac3b8-76b0-4723-bae2-e504bc911e17"
          }
        ],
        "album": {
          "id": 3698263,
          "title": "Waking Up",
          "cover": "c6519e20-1308-472e-94d1-5bbca0246d16",
          "vibrantColor": "#e1d29f",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001bf23488ad3db5c8306b058f89b2"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3140983,
        "title": "Cave In",
        "duration": 242,
        "replayGain": -9.11,
        "peak": 0.961304,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-07-28T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": "Album Version",
        "popularity": 56,
        "copyright": "℗ 2009 Universal Republic Records, a division of UMG Recordings, Inc.",
        "bpm": null,
        "url": "http://www.tidal.com/track/3140983",
        "isrc": "USUM70971816",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3582596,
          "name": "Owl City",
          "handle": null,
          "type": "MAIN",
          "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
        },
        "artists": [
          {
            "id": 3582596,
            "name": "Owl City",
            "handle": null,
            "type": "MAIN",
            "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
          }
        ],
        "album": {
          "id": 3140982,
          "title": "Ocean Eyes",
          "cover": "76b92beb-399c-4983-9b91-0eef89c796e1",
          "vibrantColor": "#305c86",
          "videoCover": "0812e3f6-bd0a-40b8-ba96-f6f2d21c8ebb"
        },
        "mixes": {
          "TRACK_MIX": "0010463bd2d8a8be6ec395ee3e85db"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3922234,
        "title": "Dynamite",
        "duration": 204,
        "replayGain": -10.57,
        "peak": 0.998962,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2010-06-01T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 85,
        "copyright": "℗ 2010 Universal Island Records Ltd. A Universal Music Company.",
        "bpm": 118,
        "url": "http://www.tidal.com/track/3922234",
        "isrc": "GBUM71003721",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3530775,
          "name": "Taio Cruz",
          "handle": null,
          "type": "MAIN",
          "picture": "45b6ac44-966e-4924-bc83-8fd4576b9884"
        },
        "artists": [
          {
            "id": 3530775,
            "name": "Taio Cruz",
            "handle": null,
            "type": "MAIN",
            "picture": "45b6ac44-966e-4924-bc83-8fd4576b9884"
          }
        ],
        "album": {
          "id": 3922233,
          "title": "Rokstarr",
          "cover": "3d33bae4-901d-47d7-9db6-1f2974b95902",
          "vibrantColor": "#FFFFFF",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0011789c0c289c661dc736ef20d8ab"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35213010,
        "title": "Drive By",
        "duration": 196,
        "replayGain": -12.22,
        "peak": 0.997375,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2012-04-17T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 80,
        "copyright": "(P) 2012 Columbia Records, a Division of Sony Music Entertainment",
        "bpm": 122,
        "url": "http://www.tidal.com/track/35213010",
        "isrc": "USSM11106876",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 1186,
          "name": "Train",
          "handle": null,
          "type": "MAIN",
          "picture": "4e624013-d349-4d9e-9da7-56031917bb7e"
        },
        "artists": [
          {
            "id": 1186,
            "name": "Train",
            "handle": null,
            "type": "MAIN",
            "picture": "4e624013-d349-4d9e-9da7-56031917bb7e"
          }
        ],
        "album": {
          "id": 35213008,
          "title": "California 37",
          "cover": "2c8954c6-de8c-4ac0-bdb7-e6c524d18154",
          "vibrantColor": "#45889f",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001b612d26501dbcd57c9fcbacb61d"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 2334623,
        "title": "My Life Would Suck Without You",
        "duration": 211,
        "replayGain": -11.87,
        "peak": 0.997375,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-03-10T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 75,
        "copyright": "(P) 2009 19 Recordings Limited",
        "bpm": 145,
        "url": "http://www.tidal.com/track/2334623",
        "isrc": "GBCTA0800348",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 1772,
          "name": "Kelly Clarkson",
          "handle": null,
          "type": "MAIN",
          "picture": "88c4b579-d555-4b5b-b63f-e8e23cb9a073"
        },
        "artists": [
          {
            "id": 1772,
            "name": "Kelly Clarkson",
            "handle": null,
            "type": "MAIN",
            "picture": "88c4b579-d555-4b5b-b63f-e8e23cb9a073"
          }
        ],
        "album": {
          "id": 2334497,
          "title": "All I Ever Wanted",
          "cover": "76b90d25-7888-4e8b-9b9b-2b858ec328bd",
          "vibrantColor": "#f1d4a1",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001aebb0ecc9f25d4ff4d2dd6828ab"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35358101,
        "title": "I Knew You Were Trouble.",
        "duration": 218,
        "replayGain": -10.03,
        "peak": 0.987396,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2012-10-22T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 4,
        "volumeNumber": 1,
        "version": null,
        "popularity": 81,
        "copyright": "℗ 2012 Big Machine Records, LLC.",
        "bpm": 154,
        "url": "http://www.tidal.com/track/35358101",
        "isrc": "USCJY1231039",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3557299,
          "name": "Taylor Swift",
          "handle": null,
          "type": "MAIN",
          "picture": "acce3554-c0dd-428b-b39a-18a06897a7c3"
        },
        "artists": [
          {
            "id": 3557299,
            "name": "Taylor Swift",
            "handle": null,
            "type": "MAIN",
            "picture": "acce3554-c0dd-428b-b39a-18a06897a7c3"
          }
        ],
        "album": {
          "id": 35358097,
          "title": "Red",
          "cover": "fa796c92-e85c-4fe5-b267-e4c0e2b74159",
          "vibrantColor": "#b01e48",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001fbddf4d1ce637528a2c64fa084b"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3413725,
        "title": "Ridin' Solo",
        "duration": 216,
        "replayGain": -11.65,
        "peak": 0.831787,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2010-03-02T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 77,
        "copyright": "℗ 2010 Warner Records Inc.",
        "bpm": 90,
        "url": "http://www.tidal.com/track/3413725",
        "isrc": "USWB10905329",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3851664,
          "name": "Jason Derulo",
          "handle": null,
          "type": "MAIN",
          "picture": "b1a1cb60-f9a3-4249-9900-652cdde9bbd1"
        },
        "artists": [
          {
            "id": 3851664,
            "name": "Jason Derulo",
            "handle": null,
            "type": "MAIN",
            "picture": "b1a1cb60-f9a3-4249-9900-652cdde9bbd1"
          }
        ],
        "album": {
          "id": 3413723,
          "title": "Jason Derulo",
          "cover": "3659764d-fa06-414a-9277-c172d59809d9",
          "vibrantColor": "#99b7d7",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00119f23f47ed01fe82f634d44d35e"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35128374,
        "title": "Hey There Delilah",
        "duration": 233,
        "replayGain": -12.13,
        "peak": 0.999969,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2007-02-20T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 13,
        "volumeNumber": 1,
        "version": "Bonus Track",
        "popularity": 83,
        "copyright": "℗ 2006 Hollywood Records, Inc.",
        "bpm": 104,
        "url": "http://www.tidal.com/track/35128374",
        "isrc": "USHR10622325",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 10371,
          "name": "Plain White T's",
          "handle": null,
          "type": "MAIN",
          "picture": "a1d8ebb1-9d20-4925-85d8-f2a0872408d7"
        },
        "artists": [
          {
            "id": 10371,
            "name": "Plain White T's",
            "handle": null,
            "type": "MAIN",
            "picture": "a1d8ebb1-9d20-4925-85d8-f2a0872408d7"
          }
        ],
        "album": {
          "id": 35128361,
          "title": "Every Second Counts",
          "cover": "75734f61-46d4-46fe-a91d-5b3b0d1d3b3c",
          "vibrantColor": "#2da1c6",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001e14d525aca5f525a31ac0e88606"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 23849834,
        "title": "Magic (feat. Rivers Cuomo)",
        "duration": 196,
        "replayGain": -10.61,
        "peak": 0.998444,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2010-04-27T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 8,
        "volumeNumber": 1,
        "version": null,
        "popularity": 76,
        "copyright": "℗ 2010 Atlantic Recording Corporation for the United States and WEA International Inc. for the world outside of the United States",
        "bpm": 83,
        "url": "http://www.tidal.com/track/23849834",
        "isrc": "USAT21000545",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3667959,
          "name": "B.o.B",
          "handle": null,
          "type": "MAIN",
          "picture": "530f127a-9fd5-4b8c-b2f9-cdde70873130"
        },
        "artists": [
          {
            "id": 3667959,
            "name": "B.o.B",
            "handle": null,
            "type": "MAIN",
            "picture": "530f127a-9fd5-4b8c-b2f9-cdde70873130"
          },
          {
            "id": 3500116,
            "name": "Rivers Cuomo",
            "handle": null,
            "type": "FEATURED",
            "picture": "07c7fc48-7311-45ce-b75c-4b5fe92ba1a2"
          }
        ],
        "album": {
          "id": 23849826,
          "title": "B.o.B Presents: The Adventures of Bobby Ray",
          "cover": "85b28a9c-c161-4fe5-9c4c-ea8ee8c0da57",
          "vibrantColor": "#e7e452",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001552c2303d61e56ab21eba6ec4f2"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3698269,
        "title": "Good Life",
        "duration": 253,
        "replayGain": -10.92,
        "peak": 0.988647,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-11-17T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 6,
        "volumeNumber": 1,
        "version": null,
        "popularity": 79,
        "copyright": "℗ 2009 Mosley Music/Interscope Records",
        "bpm": 95,
        "url": "http://www.tidal.com/track/3698269",
        "isrc": "USUM70999110",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 64665,
          "name": "OneRepublic",
          "handle": null,
          "type": "MAIN",
          "picture": "fadac3b8-76b0-4723-bae2-e504bc911e17"
        },
        "artists": [
          {
            "id": 64665,
            "name": "OneRepublic",
            "handle": null,
            "type": "MAIN",
            "picture": "fadac3b8-76b0-4723-bae2-e504bc911e17"
          }
        ],
        "album": {
          "id": 3698263,
          "title": "Waking Up",
          "cover": "c6519e20-1308-472e-94d1-5bbca0246d16",
          "vibrantColor": "#e1d29f",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0011396aa0c42f5ecc33ead0475d81"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 102998487,
        "title": "We're All To Blame",
        "duration": 219,
        "replayGain": -12.12,
        "peak": 0.805115,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2019-02-01T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 3,
        "volumeNumber": 1,
        "version": null,
        "popularity": 61,
        "copyright": "℗ 2004 Island Records, a division of UMG Recordings, Inc.",
        "bpm": null,
        "url": "http://www.tidal.com/track/102998487",
        "isrc": "USIR20400704",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 30124,
          "name": "Sum 41",
          "handle": null,
          "type": "MAIN",
          "picture": "c86ecbb3-027e-4e69-b697-501599bdb55e"
        },
        "artists": [
          {
            "id": 30124,
            "name": "Sum 41",
            "handle": null,
            "type": "MAIN",
            "picture": "c86ecbb3-027e-4e69-b697-501599bdb55e"
          }
        ],
        "album": {
          "id": 102998484,
          "title": "Chuck",
          "cover": "06f16e4b-1810-4e24-94b8-43405136dbc6",
          "vibrantColor": "#ccba29",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001ba09c684b027d7187dcb9add874"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 13420417,
        "title": "Break Your Heart",
        "duration": 185,
        "replayGain": -10.57,
        "peak": 0.998962,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2014-08-18T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 79,
        "copyright": "℗ 2010 Universal Island Records Ltd. A Universal Music Company.",
        "bpm": 122,
        "url": "http://www.tidal.com/track/13420417",
        "isrc": "GBUM71003723",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3530775,
          "name": "Taio Cruz",
          "handle": null,
          "type": "MAIN",
          "picture": "45b6ac44-966e-4924-bc83-8fd4576b9884"
        },
        "artists": [
          {
            "id": 3530775,
            "name": "Taio Cruz",
            "handle": null,
            "type": "MAIN",
            "picture": "45b6ac44-966e-4924-bc83-8fd4576b9884"
          },
          {
            "id": 24907,
            "name": "Ludacris",
            "handle": null,
            "type": "FEATURED",
            "picture": "e4a99d8e-6ab3-4470-a727-f41006d666cd"
          }
        ],
        "album": {
          "id": 13420415,
          "title": "Rokstarr (Special Edition)",
          "cover": "9b9d0a44-60f6-4f29-8abe-85d1e1411ba2",
          "vibrantColor": "#FFFFFF",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00103439766cd901c4990082abbfbd"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3140984,
        "title": "The Bird And The Worm",
        "duration": 208,
        "replayGain": -9.11,
        "peak": 0.961304,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-07-28T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 55,
        "copyright": "℗ 2009 Universal Republic Records, a division of UMG Recordings, Inc.",
        "bpm": null,
        "url": "http://www.tidal.com/track/3140984",
        "isrc": "USUM70972062",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3582596,
          "name": "Owl City",
          "handle": null,
          "type": "MAIN",
          "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
        },
        "artists": [
          {
            "id": 3582596,
            "name": "Owl City",
            "handle": null,
            "type": "MAIN",
            "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
          }
        ],
        "album": {
          "id": 3140982,
          "title": "Ocean Eyes",
          "cover": "76b92beb-399c-4983-9b91-0eef89c796e1",
          "vibrantColor": "#305c86",
          "videoCover": "0812e3f6-bd0a-40b8-ba96-f6f2d21c8ebb"
        },
        "mixes": {
          "TRACK_MIX": "001c335e6ac71c7526d177f29a1e2b"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3701010,
        "title": "Down",
        "duration": 212,
        "replayGain": -9.52,
        "peak": 1,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2021-04-14T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 80,
        "copyright": "℗ 2010 Cash Money Records Inc.",
        "bpm": 132,
        "url": "http://www.tidal.com/track/3701010",
        "isrc": "USCM50900094",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3512185,
          "name": "Jay Sean",
          "handle": null,
          "type": "MAIN",
          "picture": "12f19ac8-2642-49dc-b194-a181023f2fe0"
        },
        "artists": [
          {
            "id": 3512185,
            "name": "Jay Sean",
            "handle": null,
            "type": "MAIN",
            "picture": "12f19ac8-2642-49dc-b194-a181023f2fe0"
          },
          {
            "id": 27518,
            "name": "Lil Wayne",
            "handle": null,
            "type": "FEATURED",
            "picture": "5b9d2f85-f9ea-45e5-aa50-2db1651f3237"
          }
        ],
        "album": {
          "id": 3701009,
          "title": "Down",
          "cover": "1a212a66-7807-492e-b597-3ce4b35454d1",
          "vibrantColor": "#f2656a",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0019b0f0968311724ccc6ef68db7b7"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 33810622,
        "title": "TiK ToK",
        "duration": 200,
        "replayGain": -11.47,
        "peak": 0.997345,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2010-05-10T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 84,
        "copyright": "(P) 2009 RCA Records, a division of Sony Music Entertainment",
        "bpm": 120,
        "url": "http://www.tidal.com/track/33810622",
        "isrc": "USRC10900433",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3652939,
          "name": "Kesha",
          "handle": null,
          "type": "MAIN",
          "picture": "2d915c0c-8ad0-4b00-a6ec-d9d931e30c9d"
        },
        "artists": [
          {
            "id": 3652939,
            "name": "Kesha",
            "handle": null,
            "type": "MAIN",
            "picture": "2d915c0c-8ad0-4b00-a6ec-d9d931e30c9d"
          }
        ],
        "album": {
          "id": 33810620,
          "title": "Animal (Expanded Edition)",
          "cover": "39fc0d06-b7b1-4606-859a-667e471a7155",
          "vibrantColor": "#fe81d9",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001743fabd59f6148bc2975f2b5dfd"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35736704,
        "title": "Clumsy",
        "duration": 240,
        "replayGain": -9.82,
        "peak": 0.988556,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2006-09-19T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 76,
        "copyright": "℗ 2006 UMG Recordings, Inc.",
        "bpm": 92,
        "url": "http://www.tidal.com/track/35736704",
        "isrc": "USUM70609116",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 10664,
          "name": "Fergie",
          "handle": null,
          "type": "MAIN",
          "picture": "09a83abe-3916-4940-807b-d6389d931358"
        },
        "artists": [
          {
            "id": 10664,
            "name": "Fergie",
            "handle": null,
            "type": "MAIN",
            "picture": "09a83abe-3916-4940-807b-d6389d931358"
          }
        ],
        "album": {
          "id": 35736702,
          "title": "The Dutchess",
          "cover": "1ea9157c-2931-4655-b3d2-69cfba7b71fc",
          "vibrantColor": "#e4887e",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001e714e9770d10ae0e3e6dece01b1"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35285118,
        "title": "Counting Stars",
        "duration": 258,
        "replayGain": -10.45,
        "peak": 0.999176,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2014-04-14T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 87,
        "copyright": "℗ 2013 Mosley Music/Interscope Records",
        "bpm": 122,
        "url": "http://www.tidal.com/track/35285118",
        "isrc": "USUM71301306",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 64665,
          "name": "OneRepublic",
          "handle": null,
          "type": "MAIN",
          "picture": "fadac3b8-76b0-4723-bae2-e504bc911e17"
        },
        "artists": [
          {
            "id": 64665,
            "name": "OneRepublic",
            "handle": null,
            "type": "MAIN",
            "picture": "fadac3b8-76b0-4723-bae2-e504bc911e17"
          }
        ],
        "album": {
          "id": 35285117,
          "title": "Native",
          "cover": "6cf20079-b352-4bad-b14a-561c08191440",
          "vibrantColor": "#f5e6ba",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001934c43dfe7b5f70809285bdc883"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 30924579,
        "title": "No Brains",
        "duration": 166,
        "replayGain": -11.3,
        "peak": 0.999146,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2002-11-26T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 6,
        "volumeNumber": 1,
        "version": "Album Version (Edited)",
        "popularity": 46,
        "copyright": "℗ 2002 The Island Def Jam Music Group",
        "bpm": null,
        "url": "http://www.tidal.com/track/30924579",
        "isrc": "USIR20200962",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 30124,
          "name": "Sum 41",
          "handle": null,
          "type": "MAIN",
          "picture": "c86ecbb3-027e-4e69-b697-501599bdb55e"
        },
        "artists": [
          {
            "id": 30124,
            "name": "Sum 41",
            "handle": null,
            "type": "MAIN",
            "picture": "c86ecbb3-027e-4e69-b697-501599bdb55e"
          }
        ],
        "album": {
          "id": 30924573,
          "title": "Does This Look Infected?",
          "cover": "11dd08a7-e018-4c19-b0a6-4fd4aa4a77ea",
          "vibrantColor": "#f4a95a",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001bda240d1aaff50102aeee995d94"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 122351390,
        "title": "Before You Go",
        "duration": 216,
        "replayGain": -9.31,
        "peak": 0.889557,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2019-11-19T12:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 82,
        "copyright": "℗ 2019 Universal Music GmbH",
        "bpm": 112,
        "url": "http://www.tidal.com/track/122351390",
        "isrc": "DEUM71905868",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS",
            "HIRES_LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 8583514,
          "name": "Lewis Capaldi",
          "handle": null,
          "type": "MAIN",
          "picture": "bee09376-f217-46ba-b1e2-4b99405d0d13"
        },
        "artists": [
          {
            "id": 8583514,
            "name": "Lewis Capaldi",
            "handle": null,
            "type": "MAIN",
            "picture": "bee09376-f217-46ba-b1e2-4b99405d0d13"
          }
        ],
        "album": {
          "id": 122351387,
          "title": "Before You Go",
          "cover": "68723f1c-f9a4-4f0f-9068-92dd6b140422",
          "vibrantColor": "#f7b43c",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00178e1b55160ae09a97e00078b657"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 2350373,
        "title": "Rainbow Veins",
        "duration": 281,
        "replayGain": -6.31,
        "peak": 1,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2008-12-16T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 47,
        "copyright": "℗ 2008 Owl City Catalog / Federal Distribution",
        "bpm": null,
        "url": "http://www.tidal.com/track/2350373",
        "isrc": "USZXT0835560",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3582596,
          "name": "Owl City",
          "handle": null,
          "type": "MAIN",
          "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
        },
        "artists": [
          {
            "id": 3582596,
            "name": "Owl City",
            "handle": null,
            "type": "MAIN",
            "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
          }
        ],
        "album": {
          "id": 2350371,
          "title": "Maybe I'm Dreaming",
          "cover": "0ea2b28d-8548-4c47-a60e-c346b03182d0",
          "vibrantColor": "#74b5f6",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001fdc0e577451d9390765f35760ca"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 121131955,
        "title": "Pick Up The Phone",
        "duration": 278,
        "replayGain": -11.28,
        "peak": 0.988617,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2011-07-26T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 6,
        "volumeNumber": 1,
        "version": null,
        "popularity": 64,
        "copyright": "Epitaph",
        "bpm": null,
        "url": "http://www.tidal.com/track/121131955",
        "isrc": "USEP41118006",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3988097,
          "name": "Falling In Reverse",
          "handle": null,
          "type": "MAIN",
          "picture": "0240a6bb-08ff-4a3e-9d98-17a8fd37ffb4"
        },
        "artists": [
          {
            "id": 3988097,
            "name": "Falling In Reverse",
            "handle": null,
            "type": "MAIN",
            "picture": "0240a6bb-08ff-4a3e-9d98-17a8fd37ffb4"
          }
        ],
        "album": {
          "id": 121131949,
          "title": "The Drug In Me Is You",
          "cover": "27d09b40-154c-4bc5-89ba-cc5ec44bf0f6",
          "vibrantColor": "#9e142a",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001825e131be4f00060531a0a724f0"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 4044469,
        "title": "Your Love Is My Drug",
        "duration": 187,
        "replayGain": -10.09,
        "peak": 0.997375,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2010-06-07T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 78,
        "copyright": "(P) 2009 RCA Records, a division of Sony Music Entertainment",
        "bpm": 119,
        "url": "http://www.tidal.com/track/4044469",
        "isrc": "USRC10900735",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3652939,
          "name": "Kesha",
          "handle": null,
          "type": "MAIN",
          "picture": "2d915c0c-8ad0-4b00-a6ec-d9d931e30c9d"
        },
        "artists": [
          {
            "id": 3652939,
            "name": "Kesha",
            "handle": null,
            "type": "MAIN",
            "picture": "2d915c0c-8ad0-4b00-a6ec-d9d931e30c9d"
          }
        ],
        "album": {
          "id": 4044413,
          "title": "Your Love Is My Drug",
          "cover": "495b62d4-727f-4668-b341-1b25148004ec",
          "vibrantColor": "#ebe7bc",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0016293b56fb0b44b35db9ad51fdc1"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 4104027,
        "title": "Not Afraid",
        "duration": 248,
        "replayGain": -13.03,
        "peak": 0.988647,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2013-12-06T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 7,
        "volumeNumber": 1,
        "version": null,
        "popularity": 84,
        "copyright": "℗ 2010 Aftermath Records",
        "bpm": 86,
        "url": "http://www.tidal.com/track/4104027",
        "isrc": "USUM71011769",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 17275,
          "name": "Eminem",
          "handle": null,
          "type": "MAIN",
          "picture": "695e593e-451a-4dcc-ae3f-84e3023bbd1f"
        },
        "artists": [
          {
            "id": 17275,
            "name": "Eminem",
            "handle": null,
            "type": "MAIN",
            "picture": "695e593e-451a-4dcc-ae3f-84e3023bbd1f"
          }
        ],
        "album": {
          "id": 4104020,
          "title": "Recovery",
          "cover": "7d7f91c7-6b9d-4b09-8465-25715535ec40",
          "vibrantColor": "#c03446",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001baecd7075d7621d7dba7cefcd6b"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35736712,
        "title": "Big Girls Don't Cry (Personal)",
        "duration": 268,
        "replayGain": -9.82,
        "peak": 0.988556,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2006-09-19T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 10,
        "volumeNumber": 1,
        "version": null,
        "popularity": 81,
        "copyright": "℗ 2006 UMG Recordings, Inc.",
        "bpm": 113,
        "url": "http://www.tidal.com/track/35736712",
        "isrc": "USUM70609115",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 10664,
          "name": "Fergie",
          "handle": null,
          "type": "MAIN",
          "picture": "09a83abe-3916-4940-807b-d6389d931358"
        },
        "artists": [
          {
            "id": 10664,
            "name": "Fergie",
            "handle": null,
            "type": "MAIN",
            "picture": "09a83abe-3916-4940-807b-d6389d931358"
          }
        ],
        "album": {
          "id": 35736702,
          "title": "The Dutchess",
          "cover": "1ea9157c-2931-4655-b3d2-69cfba7b71fc",
          "vibrantColor": "#e4887e",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0013248dda4adbf940079b45391704"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 209256604,
        "title": "Boom Boom Pow",
        "duration": 251,
        "replayGain": -8.85,
        "peak": 0.988556,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-06-09T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 76,
        "copyright": "℗ 2009 Interscope Records",
        "bpm": 130,
        "url": "http://www.tidal.com/track/209256604",
        "isrc": "USUM70955624",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 7806,
          "name": "Black Eyed Peas",
          "handle": null,
          "type": "MAIN",
          "picture": "1b99c28b-9e06-4b6c-8074-9418764427ce"
        },
        "artists": [
          {
            "id": 7806,
            "name": "Black Eyed Peas",
            "handle": null,
            "type": "MAIN",
            "picture": "1b99c28b-9e06-4b6c-8074-9418764427ce"
          }
        ],
        "album": {
          "id": 209256603,
          "title": "THE E.N.D. (THE ENERGY NEVER DIES)",
          "cover": "54599f8d-55a2-42c6-93c4-e3ab4446ebaf",
          "vibrantColor": "#309f49",
          "videoCover": "1b8a88ea-2025-4a67-af31-ac6f0a4035ab"
        },
        "mixes": {
          "TRACK_MIX": "0015a711d9b41d82b10057f60a8037"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35285123,
        "title": "I Lived",
        "duration": 235,
        "replayGain": -10.45,
        "peak": 0.999115,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2014-04-14T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 6,
        "volumeNumber": 1,
        "version": null,
        "popularity": 75,
        "copyright": "℗ 2013 Mosley Music/Interscope Records",
        "bpm": 120,
        "url": "http://www.tidal.com/track/35285123",
        "isrc": "USUM71301307",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 64665,
          "name": "OneRepublic",
          "handle": null,
          "type": "MAIN",
          "picture": "fadac3b8-76b0-4723-bae2-e504bc911e17"
        },
        "artists": [
          {
            "id": 64665,
            "name": "OneRepublic",
            "handle": null,
            "type": "MAIN",
            "picture": "fadac3b8-76b0-4723-bae2-e504bc911e17"
          }
        ],
        "album": {
          "id": 35285117,
          "title": "Native",
          "cover": "6cf20079-b352-4bad-b14a-561c08191440",
          "vibrantColor": "#f5e6ba",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0017dc662f9f690ea1b116406a581d"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 14165831,
        "title": "Teenage Dream",
        "duration": 228,
        "replayGain": -10.94,
        "peak": 0.988647,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2012-03-23T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 84,
        "copyright": "℗ 2010 Capitol Records, LLC",
        "bpm": 120,
        "url": "http://www.tidal.com/track/14165831",
        "isrc": "USCA21001255",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3531205,
          "name": "Katy Perry",
          "handle": null,
          "type": "MAIN",
          "picture": "1208c7d7-f423-41f5-9124-d9b0a60c6207"
        },
        "artists": [
          {
            "id": 3531205,
            "name": "Katy Perry",
            "handle": null,
            "type": "MAIN",
            "picture": "1208c7d7-f423-41f5-9124-d9b0a60c6207"
          }
        ],
        "album": {
          "id": 14165830,
          "title": "Teenage Dream: The Complete Confection",
          "cover": "9a78b0d6-405d-4686-92cf-8442c5857cca",
          "vibrantColor": "#f9d460",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001f3e5c77d030e58eda65e13b5e5d"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 121044447,
        "title": "Alone",
        "duration": 279,
        "replayGain": -10.55,
        "peak": 0.891249,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2013-06-18T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 5,
        "volumeNumber": 1,
        "version": null,
        "popularity": 60,
        "copyright": "Epitaph",
        "bpm": null,
        "url": "http://www.tidal.com/track/121044447",
        "isrc": "USEP41310021",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS",
            "HIRES_LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3988097,
          "name": "Falling In Reverse",
          "handle": null,
          "type": "MAIN",
          "picture": "0240a6bb-08ff-4a3e-9d98-17a8fd37ffb4"
        },
        "artists": [
          {
            "id": 3988097,
            "name": "Falling In Reverse",
            "handle": null,
            "type": "MAIN",
            "picture": "0240a6bb-08ff-4a3e-9d98-17a8fd37ffb4"
          }
        ],
        "album": {
          "id": 121044442,
          "title": "Fashionably Late (Deluxe Edition)",
          "cover": "04cbabc2-b414-4ce5-ade1-d62673a921b6",
          "vibrantColor": "#ed85aa",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00101d79ebe3b7bff253bd2181db11"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 16604848,
        "title": "Shooting Star",
        "duration": 247,
        "replayGain": -10.33,
        "peak": 0.977112,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2012-08-21T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 61,
        "copyright": "℗ 2012 Universal Republic Records, a division of UMG Recordings, Inc.",
        "bpm": 124,
        "url": "http://www.tidal.com/track/16604848",
        "isrc": "USUM71204753",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3582596,
          "name": "Owl City",
          "handle": null,
          "type": "MAIN",
          "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
        },
        "artists": [
          {
            "id": 3582596,
            "name": "Owl City",
            "handle": null,
            "type": "MAIN",
            "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
          }
        ],
        "album": {
          "id": 16604846,
          "title": "The Midsummer Station",
          "cover": "11739bd5-770a-4176-8e12-86d9cbf86398",
          "vibrantColor": "#ead7b0",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001d44fd95c477295494e5a6c06dae"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 19291985,
        "title": "Shadow Moses",
        "duration": 243,
        "replayGain": -10.87,
        "peak": 0.86099,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2021-04-01T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 6,
        "volumeNumber": 1,
        "version": null,
        "popularity": 68,
        "copyright": "(P) 2013 Sony Music Entertainment UK Limited",
        "bpm": null,
        "url": "http://www.tidal.com/track/19291985",
        "isrc": "GBARL1202153",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS",
            "HIRES_LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 19137,
          "name": "Bring Me The Horizon",
          "handle": null,
          "type": "MAIN",
          "picture": "d7bab9a5-ef15-4c83-a6dc-78fb120c6861"
        },
        "artists": [
          {
            "id": 19137,
            "name": "Bring Me The Horizon",
            "handle": null,
            "type": "MAIN",
            "picture": "d7bab9a5-ef15-4c83-a6dc-78fb120c6861"
          }
        ],
        "album": {
          "id": 19291979,
          "title": "Sempiternal (Expanded Edition)",
          "cover": "f49be36d-4032-4c61-b1aa-85b9b0833e18",
          "vibrantColor": "#ae8637",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0019d4e9d5e21d75593737775144c2"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 33810623,
        "title": "Take It Off",
        "duration": 215,
        "replayGain": -11.47,
        "peak": 0.997345,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2010-05-10T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 3,
        "volumeNumber": 1,
        "version": null,
        "popularity": 78,
        "copyright": "(P) 2009 RCA Records, a division of Sony Music Entertainment",
        "bpm": 125,
        "url": "http://www.tidal.com/track/33810623",
        "isrc": "USRC10900736",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3652939,
          "name": "Kesha",
          "handle": null,
          "type": "MAIN",
          "picture": "2d915c0c-8ad0-4b00-a6ec-d9d931e30c9d"
        },
        "artists": [
          {
            "id": 3652939,
            "name": "Kesha",
            "handle": null,
            "type": "MAIN",
            "picture": "2d915c0c-8ad0-4b00-a6ec-d9d931e30c9d"
          }
        ],
        "album": {
          "id": 33810620,
          "title": "Animal (Expanded Edition)",
          "cover": "39fc0d06-b7b1-4606-859a-667e471a7155",
          "vibrantColor": "#fe81d9",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001e77cf824bd982b6b15432bdc802"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 4104035,
        "title": "Love The Way You Lie",
        "duration": 263,
        "replayGain": -13.03,
        "peak": 0.999786,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2017-11-21T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 15,
        "volumeNumber": 1,
        "version": null,
        "popularity": 87,
        "copyright": "℗ 2010 Aftermath Records",
        "bpm": 87,
        "url": "http://www.tidal.com/track/4104035",
        "isrc": "USUM71015397",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 17275,
          "name": "Eminem",
          "handle": null,
          "type": "MAIN",
          "picture": "695e593e-451a-4dcc-ae3f-84e3023bbd1f"
        },
        "artists": [
          {
            "id": 17275,
            "name": "Eminem",
            "handle": null,
            "type": "MAIN",
            "picture": "695e593e-451a-4dcc-ae3f-84e3023bbd1f"
          }
        ],
        "album": {
          "id": 4104020,
          "title": "Recovery",
          "cover": "7d7f91c7-6b9d-4b09-8465-25715535ec40",
          "vibrantColor": "#c03446",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001ed2045ded3395cf082a6109ff31"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 77707061,
        "title": "Bad Romance",
        "duration": 295,
        "replayGain": -12.82,
        "peak": 0.998849,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2017-02-03T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 82,
        "copyright": "℗ 2009 UMG Recordings, Inc.",
        "bpm": 119,
        "url": "http://www.tidal.com/track/77707061",
        "isrc": "USUM70918596",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS",
            "HIRES_LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3534754,
          "name": "Lady Gaga",
          "handle": null,
          "type": "MAIN",
          "picture": "1594f44b-f4ed-4bec-80aa-276e0e759063"
        },
        "artists": [
          {
            "id": 3534754,
            "name": "Lady Gaga",
            "handle": null,
            "type": "MAIN",
            "picture": "1594f44b-f4ed-4bec-80aa-276e0e759063"
          }
        ],
        "album": {
          "id": 77707060,
          "title": "The Fame Monster",
          "cover": "bbf2a946-214c-4ac6-b25f-7c40bc19bd0d",
          "vibrantColor": "#FFFFFF",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00139ffbed25d0ba6ba70837ee1659"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 1783276,
        "title": "Viva La Vida",
        "duration": 242,
        "replayGain": -8.85,
        "peak": 0.997253,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2008-06-16T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 7,
        "volumeNumber": 1,
        "version": null,
        "popularity": 88,
        "copyright": "℗ 2008 Parlophone Records Limited",
        "bpm": 138,
        "url": "http://www.tidal.com/track/1783276",
        "isrc": "GBAYE0800265",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 8812,
          "name": "Coldplay",
          "handle": null,
          "type": "MAIN",
          "picture": "b4579672-5b91-4679-a27a-288f097a4da5"
        },
        "artists": [
          {
            "id": 8812,
            "name": "Coldplay",
            "handle": null,
            "type": "MAIN",
            "picture": "b4579672-5b91-4679-a27a-288f097a4da5"
          }
        ],
        "album": {
          "id": 1783269,
          "title": "Viva La Vida or Death and All His Friends",
          "cover": "e5be9793-54bf-4580-853e-c93c372a9d07",
          "vibrantColor": "#9b382c",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001a2269068d0dbcd7137f00f4c7f0"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 1214593,
        "title": "Apologize",
        "duration": 208,
        "replayGain": -10.57,
        "peak": 0.988556,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2007-11-20T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 4,
        "volumeNumber": 1,
        "version": null,
        "popularity": 79,
        "copyright": "℗ 2007 UMG Recordings, Inc.",
        "bpm": 118,
        "url": "http://www.tidal.com/track/1214593",
        "isrc": "USUM70757102",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 64665,
          "name": "OneRepublic",
          "handle": null,
          "type": "MAIN",
          "picture": "fadac3b8-76b0-4723-bae2-e504bc911e17"
        },
        "artists": [
          {
            "id": 64665,
            "name": "OneRepublic",
            "handle": null,
            "type": "MAIN",
            "picture": "fadac3b8-76b0-4723-bae2-e504bc911e17"
          }
        ],
        "album": {
          "id": 1214589,
          "title": "Dreaming Out Loud",
          "cover": "1bb1e623-cffb-4bb4-a5de-0b360e51e45d",
          "vibrantColor": "#6788ba",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001b9d17ae7a37c948ed727dbd912e"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35358572,
        "title": "Hot N Cold",
        "duration": 220,
        "replayGain": -11.84,
        "peak": 0.888916,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2008-06-17T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 7,
        "volumeNumber": 1,
        "version": null,
        "popularity": 84,
        "copyright": "℗ 2008 Capitol Records, LLC",
        "bpm": 132,
        "url": "http://www.tidal.com/track/35358572",
        "isrc": "USCA20802544",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3531205,
          "name": "Katy Perry",
          "handle": null,
          "type": "MAIN",
          "picture": "1208c7d7-f423-41f5-9124-d9b0a60c6207"
        },
        "artists": [
          {
            "id": 3531205,
            "name": "Katy Perry",
            "handle": null,
            "type": "MAIN",
            "picture": "1208c7d7-f423-41f5-9124-d9b0a60c6207"
          }
        ],
        "album": {
          "id": 35358565,
          "title": "One Of The Boys",
          "cover": "2b96c653-f448-477b-accb-9175c717ab60",
          "vibrantColor": "#dc5764",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001b6302ab304f33326643b042c5e1"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 1897824,
        "title": "Breakeven",
        "duration": 261,
        "replayGain": -10.49,
        "peak": 0.997345,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-12-08T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 5,
        "volumeNumber": 1,
        "version": null,
        "popularity": 83,
        "copyright": "(P) 2008 Sony Music Entertainment UK Limited",
        "bpm": 94,
        "url": "http://www.tidal.com/track/1897824",
        "isrc": "GBARL0800147",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3530277,
          "name": "The Script",
          "handle": null,
          "type": "MAIN",
          "picture": "aef4d418-71a9-469e-8887-c65e621468a5"
        },
        "artists": [
          {
            "id": 3530277,
            "name": "The Script",
            "handle": null,
            "type": "MAIN",
            "picture": "aef4d418-71a9-469e-8887-c65e621468a5"
          }
        ],
        "album": {
          "id": 1897819,
          "title": "The Script",
          "cover": "cb5923e4-98ff-474a-a239-c2eedb8293b1",
          "vibrantColor": "#a0c5de",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001f4990038e90a0416396fbe81228"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 270283999,
        "title": "Watch The World Burn",
        "duration": 204,
        "replayGain": -8.79,
        "peak": 0.971021,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2023-01-31T10:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 70,
        "copyright": "Epitaph",
        "bpm": 162,
        "url": "http://www.tidal.com/track/270283999",
        "isrc": "USEP42231004",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS",
            "HIRES_LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3988097,
          "name": "Falling In Reverse",
          "handle": null,
          "type": "MAIN",
          "picture": "0240a6bb-08ff-4a3e-9d98-17a8fd37ffb4"
        },
        "artists": [
          {
            "id": 3988097,
            "name": "Falling In Reverse",
            "handle": null,
            "type": "MAIN",
            "picture": "0240a6bb-08ff-4a3e-9d98-17a8fd37ffb4"
          }
        ],
        "album": {
          "id": 270283998,
          "title": "Watch The World Burn",
          "cover": "1877250d-6844-4e86-962e-b2a35dccdf10",
          "vibrantColor": "#a5b6d7",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00111ffa4c563c50528a6eed37ad78"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 105083424,
        "title": "To the Sky",
        "duration": 219,
        "replayGain": -8.86,
        "peak": 0.982697,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2025-05-07T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 55,
        "copyright": "℗ 2010 Universal Republic Records",
        "bpm": null,
        "url": "http://www.tidal.com/track/105083424",
        "isrc": "USNLR1000372",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3582596,
          "name": "Owl City",
          "handle": null,
          "type": "MAIN",
          "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
        },
        "artists": [
          {
            "id": 3582596,
            "name": "Owl City",
            "handle": null,
            "type": "MAIN",
            "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
          }
        ],
        "album": {
          "id": 105083423,
          "title": "Legend of the Guardians: The Owls of Ga'Hoole (Original Motion Picture Soundtrack)",
          "cover": "b128261f-36e1-4c54-808a-5a9f82690e2d",
          "vibrantColor": "#ead8a5",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00174d434bf4f2f829707e5e3f0456"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 23496115,
        "title": "Pompeii",
        "duration": 214,
        "replayGain": -10.98,
        "peak": 0.989166,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2014-01-14T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 84,
        "copyright": "℗ 2013 Virgin Records Limited",
        "bpm": 127,
        "url": "http://www.tidal.com/track/23496115",
        "isrc": "GBAAA1200795",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 4526830,
          "name": "Bastille",
          "handle": null,
          "type": "MAIN",
          "picture": "1fc92fb1-4e36-441e-9fbd-4ab1b701726c"
        },
        "artists": [
          {
            "id": 4526830,
            "name": "Bastille",
            "handle": null,
            "type": "MAIN",
            "picture": "1fc92fb1-4e36-441e-9fbd-4ab1b701726c"
          }
        ],
        "album": {
          "id": 23496114,
          "title": "All This Bad Blood",
          "cover": "641894a5-dd48-4829-9c5e-c6827e22913d",
          "vibrantColor": "#f2eddf",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00133eb8c860c2ecccf50bd491c61e"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 77707062,
        "title": "Alejandro",
        "duration": 275,
        "replayGain": -12.82,
        "peak": 0.998849,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2017-02-03T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 76,
        "copyright": "℗ 2009 UMG Recordings, Inc.",
        "bpm": 99,
        "url": "http://www.tidal.com/track/77707062",
        "isrc": "USUM70905526",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS",
            "HIRES_LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3534754,
          "name": "Lady Gaga",
          "handle": null,
          "type": "MAIN",
          "picture": "1594f44b-f4ed-4bec-80aa-276e0e759063"
        },
        "artists": [
          {
            "id": 3534754,
            "name": "Lady Gaga",
            "handle": null,
            "type": "MAIN",
            "picture": "1594f44b-f4ed-4bec-80aa-276e0e759063"
          }
        ],
        "album": {
          "id": 77707060,
          "title": "The Fame Monster",
          "cover": "bbf2a946-214c-4ac6-b25f-7c40bc19bd0d",
          "vibrantColor": "#FFFFFF",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0014435d6e49eaf58650cfd511f44c"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 331558,
        "title": "This Love",
        "duration": 206,
        "replayGain": -10.37,
        "peak": 0.988556,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2007-04-24T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 85,
        "copyright": "℗ 2003 UMG Recordings, Inc.",
        "bpm": 95,
        "url": "http://www.tidal.com/track/331558",
        "isrc": "USJAY0300080",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 1565,
          "name": "Maroon 5",
          "handle": null,
          "type": "MAIN",
          "picture": "dc5be69b-84d4-42d9-be13-b9154ae4e159"
        },
        "artists": [
          {
            "id": 1565,
            "name": "Maroon 5",
            "handle": null,
            "type": "MAIN",
            "picture": "dc5be69b-84d4-42d9-be13-b9154ae4e159"
          }
        ],
        "album": {
          "id": 331556,
          "title": "Songs About Jane",
          "cover": "4b3436f5-c2b8-4613-8c9b-492b05f07580",
          "vibrantColor": "#ed8c82",
          "videoCover": "38a6f93a-086d-49d2-a6f4-5d57dc5fb230"
        },
        "mixes": {
          "TRACK_MIX": "001c82b5735d11cd9974475c2b4d05"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 4935186,
        "title": "Just the Way You Are",
        "duration": 221,
        "replayGain": -10.74,
        "peak": 1,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2013-02-21T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 88,
        "copyright": "℗ 2010 Elektra Entertainment Group Inc.",
        "bpm": 109,
        "url": "http://www.tidal.com/track/4935186",
        "isrc": "USAT21001269",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS",
            "HIRES_LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3658521,
          "name": "Bruno Mars",
          "handle": null,
          "type": "MAIN",
          "picture": "72b86dab-cb66-4cb6-9610-187b94b3363a"
        },
        "artists": [
          {
            "id": 3658521,
            "name": "Bruno Mars",
            "handle": null,
            "type": "MAIN",
            "picture": "72b86dab-cb66-4cb6-9610-187b94b3363a"
          }
        ],
        "album": {
          "id": 4935184,
          "title": "Doo-Wops & Hooligans",
          "cover": "d1cb1172-edf2-4793-979f-7700ecad8432",
          "vibrantColor": "#fbf0c2",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001581b2e95fc6b9478f86eb5d7d32"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 33599342,
        "title": "Superheroes",
        "duration": 245,
        "replayGain": -10.79,
        "peak": 0.997375,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2014-09-30T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 74,
        "copyright": "(P) 2014 Sony Music Entertainment UK Limited",
        "bpm": 166,
        "url": "http://www.tidal.com/track/33599342",
        "isrc": "GBARL1400978",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3530277,
          "name": "The Script",
          "handle": null,
          "type": "MAIN",
          "picture": "aef4d418-71a9-469e-8887-c65e621468a5"
        },
        "artists": [
          {
            "id": 3530277,
            "name": "The Script",
            "handle": null,
            "type": "MAIN",
            "picture": "aef4d418-71a9-469e-8887-c65e621468a5"
          }
        ],
        "album": {
          "id": 33599340,
          "title": "No Sound Without Silence",
          "cover": "4d8c9e2d-a922-4061-9678-2b5be71d8d9c",
          "vibrantColor": "#77accf",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001e633972a9d2432a19215adc90f2"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 33747024,
        "title": "Pocketful of Sunshine",
        "duration": 203,
        "replayGain": -10.94,
        "peak": 0.997345,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2008-01-22T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 81,
        "copyright": "(P) 2008 Sony Music Entertainment UK Limited",
        "bpm": 110,
        "url": "http://www.tidal.com/track/33747024",
        "isrc": "GBARL0701374",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 504,
          "name": "Natasha Bedingfield",
          "handle": null,
          "type": "MAIN",
          "picture": "19454c7b-6a02-4c2c-bfc0-a4b9219aee48"
        },
        "artists": [
          {
            "id": 504,
            "name": "Natasha Bedingfield",
            "handle": null,
            "type": "MAIN",
            "picture": "19454c7b-6a02-4c2c-bfc0-a4b9219aee48"
          }
        ],
        "album": {
          "id": 33747022,
          "title": "Pocketful Of Sunshine",
          "cover": "29b58694-9b43-4433-b3c9-b7c48da482c3",
          "vibrantColor": "#f6d994",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001f67fe571c066e8fb105b420327b"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35421134,
        "title": "Deer In The Headlights",
        "duration": 181,
        "replayGain": -11.23,
        "peak": 0.999969,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2011-06-14T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": "Album Version",
        "popularity": 53,
        "copyright": "℗ 2011 Universal Republic Records, a division of UMG Recordings, Inc.",
        "bpm": null,
        "url": "http://www.tidal.com/track/35421134",
        "isrc": "USUM71104669",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3582596,
          "name": "Owl City",
          "handle": null,
          "type": "MAIN",
          "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
        },
        "artists": [
          {
            "id": 3582596,
            "name": "Owl City",
            "handle": null,
            "type": "MAIN",
            "picture": "eb4132ca-f4a2-4ea0-9e69-16648a159a4d"
          }
        ],
        "album": {
          "id": 35421132,
          "title": "All Things Bright And Beautiful",
          "cover": "c4b50a32-e7e4-4d5f-acf0-0630f6d81a0f",
          "vibrantColor": "#f0e6d3",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00167c6cc8a3b8f29a858c562d2344"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35456858,
        "title": "Somebody That I Used To Know",
        "duration": 245,
        "replayGain": -11.49,
        "peak": 1,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2012-01-20T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 3,
        "volumeNumber": 1,
        "version": null,
        "popularity": 90,
        "copyright": "℗ 2011 Samples 'n' Seconds Records, under exclusive license to Universal Republic Records, a Division of UMG Recordings, Inc.",
        "bpm": 129,
        "url": "http://www.tidal.com/track/35456858",
        "isrc": "AUZS21100040",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3544059,
          "name": "Gotye",
          "handle": null,
          "type": "MAIN",
          "picture": "2a2bde69-4a30-42b8-91c9-ee7d0ff65150"
        },
        "artists": [
          {
            "id": 3544059,
            "name": "Gotye",
            "handle": null,
            "type": "MAIN",
            "picture": "2a2bde69-4a30-42b8-91c9-ee7d0ff65150"
          },
          {
            "id": 4301985,
            "name": "Kimbra",
            "handle": null,
            "type": "FEATURED",
            "picture": "475140b8-51e8-4f55-9275-96fec5274036"
          }
        ],
        "album": {
          "id": 35456855,
          "title": "Making Mirrors",
          "cover": "0e0856fd-b8de-40e9-b380-d6f3f2544444",
          "vibrantColor": "#f3ebd7",
          "videoCover": "6afd28f0-2f13-4fc8-bb58-eaa8d338566d"
        },
        "mixes": {
          "TRACK_MIX": "001bf8fb37844207d35b953a4adb0c"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 331560,
        "title": "She Will Be Loved",
        "duration": 257,
        "replayGain": -10.37,
        "peak": 0.988556,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2007-04-24T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 4,
        "volumeNumber": 1,
        "version": null,
        "popularity": 86,
        "copyright": "℗ 2003 A&M/Octone Records",
        "bpm": 102,
        "url": "http://www.tidal.com/track/331560",
        "isrc": "USJAY0300082",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 1565,
          "name": "Maroon 5",
          "handle": null,
          "type": "MAIN",
          "picture": "dc5be69b-84d4-42d9-be13-b9154ae4e159"
        },
        "artists": [
          {
            "id": 1565,
            "name": "Maroon 5",
            "handle": null,
            "type": "MAIN",
            "picture": "dc5be69b-84d4-42d9-be13-b9154ae4e159"
          }
        ],
        "album": {
          "id": 331556,
          "title": "Songs About Jane",
          "cover": "4b3436f5-c2b8-4613-8c9b-492b05f07580",
          "vibrantColor": "#ed8c82",
          "videoCover": "38a6f93a-086d-49d2-a6f4-5d57dc5fb230"
        },
        "mixes": {
          "TRACK_MIX": "001d727688495dd4cf15c258bcc625"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 16156372,
        "title": "A Thousand Miles",
        "duration": 237,
        "replayGain": -10.55,
        "peak": 1,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2011-01-04T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 82,
        "copyright": "℗ 2002 UMG Recordings, Inc.",
        "bpm": 95,
        "url": "http://www.tidal.com/track/16156372",
        "isrc": "USIR10210955",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 27565,
          "name": "Vanessa Carlton",
          "handle": null,
          "type": "MAIN",
          "picture": "237a702f-d1a4-4d69-aeaf-21a162ee5f02"
        },
        "artists": [
          {
            "id": 27565,
            "name": "Vanessa Carlton",
            "handle": null,
            "type": "MAIN",
            "picture": "237a702f-d1a4-4d69-aeaf-21a162ee5f02"
          }
        ],
        "album": {
          "id": 16156371,
          "title": "Best Of",
          "cover": "4313cfbb-f09c-4454-98cb-b4b673da01db",
          "vibrantColor": "#957e43",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00187512c33bbaff91ead95b81f465"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 4935185,
        "title": "Grenade",
        "duration": 222,
        "replayGain": -10.74,
        "peak": 1,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2013-02-21T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 86,
        "copyright": "℗ 2010 Elektra Entertainment Group Inc.",
        "bpm": 111,
        "url": "http://www.tidal.com/track/4935185",
        "isrc": "USAT21001883",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS",
            "HIRES_LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3658521,
          "name": "Bruno Mars",
          "handle": null,
          "type": "MAIN",
          "picture": "72b86dab-cb66-4cb6-9610-187b94b3363a"
        },
        "artists": [
          {
            "id": 3658521,
            "name": "Bruno Mars",
            "handle": null,
            "type": "MAIN",
            "picture": "72b86dab-cb66-4cb6-9610-187b94b3363a"
          }
        ],
        "album": {
          "id": 4935184,
          "title": "Doo-Wops & Hooligans",
          "cover": "d1cb1172-edf2-4793-979f-7700ecad8432",
          "vibrantColor": "#fbf0c2",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0016a1612069ae505c7b5b1063237f"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 14492425,
        "title": "Everybody Talks",
        "duration": 177,
        "replayGain": -11.7,
        "peak": 0.999969,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2012-04-17T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 3,
        "volumeNumber": 1,
        "version": null,
        "popularity": 82,
        "copyright": "℗ 2011 UMG Recordings, Inc.",
        "bpm": 155,
        "url": "http://www.tidal.com/track/14492425",
        "isrc": "USUM71119189",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3665225,
          "name": "Neon Trees",
          "handle": null,
          "type": "MAIN",
          "picture": "e6f17398-759e-45a0-9673-6ded6811e199"
        },
        "artists": [
          {
            "id": 3665225,
            "name": "Neon Trees",
            "handle": null,
            "type": "MAIN",
            "picture": "e6f17398-759e-45a0-9673-6ded6811e199"
          }
        ],
        "album": {
          "id": 14492422,
          "title": "Picture Show",
          "cover": "1c2d7c90-034e-485a-be1f-24a669c7e6ee",
          "vibrantColor": "#f8af88",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0019768c833a193c29829e5bf473fc"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3182077,
        "title": "One Time",
        "duration": 216,
        "replayGain": -11.09,
        "peak": 0.998962,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-11-17T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 76,
        "copyright": "℗ 2009 The Island Def Jam Music Group",
        "bpm": 146,
        "url": "http://www.tidal.com/track/3182077",
        "isrc": "USUM70964274",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3639248,
          "name": "Justin Bieber",
          "handle": null,
          "type": "MAIN",
          "picture": "32097e23-45f4-4020-896d-ab6194448529"
        },
        "artists": [
          {
            "id": 3639248,
            "name": "Justin Bieber",
            "handle": null,
            "type": "MAIN",
            "picture": "32097e23-45f4-4020-896d-ab6194448529"
          }
        ],
        "album": {
          "id": 3182076,
          "title": "My World",
          "cover": "0e25de79-fdd5-401e-a5df-17d801eacf99",
          "vibrantColor": null,
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0014f4622570ef4dfc1a2648949375"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 2796001,
        "title": "21 Guns",
        "duration": 321,
        "replayGain": -11.68,
        "peak": 0.997375,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-06-24T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 16,
        "volumeNumber": 1,
        "version": null,
        "popularity": 81,
        "copyright": "℗ 2009 Reprise Records",
        "bpm": 164,
        "url": "http://www.tidal.com/track/2796001",
        "isrc": "USRE10900679",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 13866,
          "name": "Green Day",
          "handle": null,
          "type": "MAIN",
          "picture": "2d0d9bcc-ab61-4942-bbfe-5e79b08c9a41"
        },
        "artists": [
          {
            "id": 13866,
            "name": "Green Day",
            "handle": null,
            "type": "MAIN",
            "picture": "2d0d9bcc-ab61-4942-bbfe-5e79b08c9a41"
          }
        ],
        "album": {
          "id": 2795985,
          "title": "21st Century Breakdown",
          "cover": "d833a658-0168-416c-92c4-f132da8086ff",
          "vibrantColor": "#d9c24f",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001465bfc14e445fe926d854fc25cd"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35422915,
        "title": "Misery",
        "duration": 216,
        "replayGain": -11.61,
        "peak": 0.966248,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2011-07-12T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 73,
        "copyright": "℗ 2010 Interscope Records",
        "bpm": 103,
        "url": "http://www.tidal.com/track/35422915",
        "isrc": "USUM71015280",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 1565,
          "name": "Maroon 5",
          "handle": null,
          "type": "MAIN",
          "picture": "0db7eb1b-6132-4929-a700-7d3f3cc098b7"
        },
        "artists": [
          {
            "id": 1565,
            "name": "Maroon 5",
            "handle": null,
            "type": "MAIN",
            "picture": "0db7eb1b-6132-4929-a700-7d3f3cc098b7"
          }
        ],
        "album": {
          "id": 35422914,
          "title": "Hands All Over (Deluxe)",
          "cover": "11cf508f-05e8-4cdf-ab87-d36621e3507a",
          "vibrantColor": "#e9676f",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00178d8e9954d256afe4043f849f38"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3514582,
        "title": "Billionaire (feat. Bruno Mars)",
        "duration": 211,
        "replayGain": -8.13,
        "peak": 1,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2013-08-22T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 79,
        "copyright": "℗ 2009 Fueled By Ramen, LLC. for the United States and WEA International Inc. for the world outside of the United States. All rights reserved.",
        "bpm": 172,
        "url": "http://www.tidal.com/track/3514582",
        "isrc": "USAT21000257",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3667957,
          "name": "Travie McCoy",
          "handle": null,
          "type": "MAIN",
          "picture": "a132b3b1-fe9a-48d3-83c4-7d36c6d90da8"
        },
        "artists": [
          {
            "id": 3667957,
            "name": "Travie McCoy",
            "handle": null,
            "type": "MAIN",
            "picture": "a132b3b1-fe9a-48d3-83c4-7d36c6d90da8"
          },
          {
            "id": 3658521,
            "name": "Bruno Mars",
            "handle": null,
            "type": "FEATURED",
            "picture": "72b86dab-cb66-4cb6-9610-187b94b3363a"
          }
        ],
        "album": {
          "id": 3514581,
          "title": "Billionaire (feat. Bruno Mars)",
          "cover": "e9f3e6d2-47a4-48f4-b303-21ca79a348d9",
          "vibrantColor": "#acd8de",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0011fe99bcfd1e337bdf104b4e1c47"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 4935189,
        "title": "The Lazy Song",
        "duration": 189,
        "replayGain": -10.74,
        "peak": 0.92256,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2013-02-21T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 5,
        "volumeNumber": 1,
        "version": null,
        "popularity": 83,
        "copyright": "℗ 2010 Elektra Entertainment Group Inc. for the United States and WEA International Inc. for the world outside of the United States. All Rights Reserved.",
        "bpm": 172,
        "url": "http://www.tidal.com/track/4935189",
        "isrc": "USAT21001886",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS",
            "HIRES_LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3658521,
          "name": "Bruno Mars",
          "handle": null,
          "type": "MAIN",
          "picture": "72b86dab-cb66-4cb6-9610-187b94b3363a"
        },
        "artists": [
          {
            "id": 3658521,
            "name": "Bruno Mars",
            "handle": null,
            "type": "MAIN",
            "picture": "72b86dab-cb66-4cb6-9610-187b94b3363a"
          }
        ],
        "album": {
          "id": 4935184,
          "title": "Doo-Wops & Hooligans",
          "cover": "d1cb1172-edf2-4793-979f-7700ecad8432",
          "vibrantColor": "#fbf0c2",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0011fcc6b6b8b8bb546be966202328"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 36771441,
        "title": "Naturally",
        "duration": 203,
        "replayGain": -11.88,
        "peak": 0.999969,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-09-29T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 6,
        "volumeNumber": 1,
        "version": null,
        "popularity": 70,
        "copyright": "℗ 2009 Hollywood Records, Inc.",
        "bpm": 133,
        "url": "http://www.tidal.com/track/36771441",
        "isrc": "USHR10924574",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3638248,
          "name": "Selena Gomez",
          "handle": null,
          "type": "MAIN",
          "picture": "adaad98e-0083-49dd-b400-a30ca256dc7c"
        },
        "artists": [
          {
            "id": 3638248,
            "name": "Selena Gomez",
            "handle": null,
            "type": "MAIN",
            "picture": "adaad98e-0083-49dd-b400-a30ca256dc7c"
          },
          {
            "id": 24092879,
            "name": "The Scene",
            "handle": null,
            "type": "MAIN",
            "picture": "b9292544-6f0a-4c48-8ba2-97d90de24207"
          }
        ],
        "album": {
          "id": 36771435,
          "title": "Kiss & Tell",
          "cover": "e2e7827c-fec6-454d-9f56-ac6adc188e48",
          "vibrantColor": "#fde8d9",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001f51dc79397956d8950f5e3676bc"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3524426,
        "title": "Baby",
        "duration": 214,
        "replayGain": -10.25,
        "peak": 0.998962,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2010-03-23T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 85,
        "copyright": "℗ 2010 The Island Def Jam Music Group",
        "bpm": 130,
        "url": "http://www.tidal.com/track/3524426",
        "isrc": "USUM70919263",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3639248,
          "name": "Justin Bieber",
          "handle": null,
          "type": "MAIN",
          "picture": "32097e23-45f4-4020-896d-ab6194448529"
        },
        "artists": [
          {
            "id": 3639248,
            "name": "Justin Bieber",
            "handle": null,
            "type": "MAIN",
            "picture": "32097e23-45f4-4020-896d-ab6194448529"
          },
          {
            "id": 24907,
            "name": "Ludacris",
            "handle": null,
            "type": "FEATURED",
            "picture": "e4a99d8e-6ab3-4470-a727-f41006d666cd"
          }
        ],
        "album": {
          "id": 3524425,
          "title": "My World 2.0",
          "cover": "5bbdd64a-1ec9-47e3-b6b1-42443eaa863a",
          "vibrantColor": "#5e3087",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001c978486bd376eb20c4595507715"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 318716604,
        "title": "This Is War",
        "duration": 218,
        "replayGain": -10.46,
        "peak": 1,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2023-10-03T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 5,
        "volumeNumber": 1,
        "version": null,
        "popularity": 57,
        "copyright": "2022 Thriller Records",
        "bpm": 172,
        "url": "http://www.tidal.com/track/318716604",
        "isrc": "ZZOPM2110442",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 9880493,
          "name": "GAWNE",
          "handle": null,
          "type": "MAIN",
          "picture": "75faad0a-19ee-4c05-ab33-53b9149f69a4"
        },
        "artists": [
          {
            "id": 9880493,
            "name": "GAWNE",
            "handle": null,
            "type": "MAIN",
            "picture": "75faad0a-19ee-4c05-ab33-53b9149f69a4"
          },
          {
            "id": 6170699,
            "name": "ATLUS",
            "handle": null,
            "type": "MAIN",
            "picture": "a6b71409-ea6d-4516-9ad0-b7a35774ef4c"
          },
          {
            "id": 3615033,
            "name": "Tech N9ne",
            "handle": null,
            "type": "MAIN",
            "picture": "bc30a183-0a9c-44a4-8f93-5dc11f70b688"
          }
        ],
        "album": {
          "id": 318716582,
          "title": "Waves",
          "cover": "eedda811-cc1a-44b8-9637-bfc53ea9b4f0",
          "vibrantColor": "#5ac1e5",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001e7dc48ab046863262f5ed4e865d"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 70891471,
        "title": "Perfect",
        "duration": 263,
        "replayGain": -11.18,
        "peak": 0.997314,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2017-03-03T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 5,
        "volumeNumber": 1,
        "version": null,
        "popularity": 89,
        "copyright": "℗ 2017 Warner Music UK Limited",
        "bpm": 95,
        "url": "http://www.tidal.com/track/70891471",
        "isrc": "GBAHS1700024",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3995478,
          "name": "Ed Sheeran",
          "handle": null,
          "type": "MAIN",
          "picture": "05d72ae4-319f-4237-821f-1d7af9ec8acf"
        },
        "artists": [
          {
            "id": 3995478,
            "name": "Ed Sheeran",
            "handle": null,
            "type": "MAIN",
            "picture": "05d72ae4-319f-4237-821f-1d7af9ec8acf"
          }
        ],
        "album": {
          "id": 70891466,
          "title": "÷ (Deluxe)",
          "cover": "d76ab40a-551b-43e8-8bbc-3bf25ce33aaf",
          "vibrantColor": "#58c4e4",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001fa523eb8a9c2dfc6bcad49e866e"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 31873694,
        "title": "If We Ever Meet Again (Featuring Katy Perry)",
        "duration": 292,
        "replayGain": -9.9,
        "peak": 0.999969,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-12-08T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 9,
        "volumeNumber": 1,
        "version": null,
        "popularity": 69,
        "copyright": "℗ 2009 Blackground Records/Interscope Records",
        "bpm": 126,
        "url": "http://www.tidal.com/track/31873694",
        "isrc": "USUM70913828",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 15708,
          "name": "Timbaland",
          "handle": null,
          "type": "MAIN",
          "picture": "61acd1c9-525f-4e7a-ab44-3c3fe1685a3c"
        },
        "artists": [
          {
            "id": 15708,
            "name": "Timbaland",
            "handle": null,
            "type": "MAIN",
            "picture": "61acd1c9-525f-4e7a-ab44-3c3fe1685a3c"
          },
          {
            "id": 3531205,
            "name": "Katy Perry",
            "handle": null,
            "type": "FEATURED",
            "picture": "1208c7d7-f423-41f5-9124-d9b0a60c6207"
          }
        ],
        "album": {
          "id": 31873685,
          "title": "Shock Value II",
          "cover": "20041b8a-a907-4965-9568-7a6aec699960",
          "vibrantColor": "#18d2e7",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001a193c6cd03cae56d54a1296b1d8"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 33713333,
        "title": "Sk8er Boi",
        "duration": 204,
        "replayGain": -10.4,
        "peak": 0.997375,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2007-04-11T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 3,
        "volumeNumber": 1,
        "version": null,
        "popularity": 84,
        "copyright": "(P) 2002 Arista Records, LLC.",
        "bpm": 150,
        "url": "http://www.tidal.com/track/33713333",
        "isrc": "USAR10200229",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 1567,
          "name": "Avril Lavigne",
          "handle": null,
          "type": "MAIN",
          "picture": "eef765c3-8017-4817-9911-2c59729ce465"
        },
        "artists": [
          {
            "id": 1567,
            "name": "Avril Lavigne",
            "handle": null,
            "type": "MAIN",
            "picture": "eef765c3-8017-4817-9911-2c59729ce465"
          }
        ],
        "album": {
          "id": 33713330,
          "title": "Let Go",
          "cover": "88858072-3555-4190-90cd-073dff3ff6a8",
          "vibrantColor": "#823f3c",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0017549a9855640d29fa8863188e90"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 7689577,
        "title": "Glad You Came",
        "duration": 198,
        "replayGain": -11.95,
        "peak": 0.997375,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2022-01-13T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 74,
        "copyright": "℗ 2011 Global Talent Records Limited, under exclusive licence to Universal Island Records, a division of Universal Music Operations Limited",
        "bpm": 128,
        "url": "http://www.tidal.com/track/7689577",
        "isrc": "GBUM71104495",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3942582,
          "name": "The Wanted",
          "handle": null,
          "type": "MAIN",
          "picture": "d72a8dda-f42c-45e0-b682-31d555846387"
        },
        "artists": [
          {
            "id": 3942582,
            "name": "The Wanted",
            "handle": null,
            "type": "MAIN",
            "picture": "d72a8dda-f42c-45e0-b682-31d555846387"
          }
        ],
        "album": {
          "id": 7689576,
          "title": "Glad You Came",
          "cover": "743d67ea-851a-4794-b966-d2d61b50ee68",
          "vibrantColor": "#e4f0f2",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001739d4bbc3f39e38cc452983122c"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 292734,
        "title": "Bad Day",
        "duration": 234,
        "replayGain": -9.88,
        "peak": 0.997284,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2013-10-17T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 3,
        "volumeNumber": 1,
        "version": null,
        "popularity": 80,
        "copyright": "℗ 2004 Warner Records Inc.",
        "bpm": 140,
        "url": "http://www.tidal.com/track/292734",
        "isrc": "USWB10401971",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 13876,
          "name": "Daniel Powter",
          "handle": null,
          "type": "MAIN",
          "picture": "3da28118-2cbf-4403-91da-e273feb2a5aa"
        },
        "artists": [
          {
            "id": 13876,
            "name": "Daniel Powter",
            "handle": null,
            "type": "MAIN",
            "picture": "3da28118-2cbf-4403-91da-e273feb2a5aa"
          }
        ],
        "album": {
          "id": 292731,
          "title": "Daniel Powter",
          "cover": "6711a757-1b2c-4c0e-b40b-b1ab6e6127fa",
          "vibrantColor": "#ecd991",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001fb7d47ce0df3862af6a93e8d29e"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 10614512,
        "title": "Tonight Tonight",
        "duration": 200,
        "replayGain": -9.57,
        "peak": 0.997314,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2011-11-29T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 79,
        "copyright": "(P) 2011 RCA Records, a division of Sony Music Entertainment",
        "bpm": 100,
        "url": "http://www.tidal.com/track/10614512",
        "isrc": "USJI11100019",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3639230,
          "name": "Hot Chelle Rae",
          "handle": null,
          "type": "MAIN",
          "picture": "6a18c717-c688-4edd-a277-4446b24cfc5f"
        },
        "artists": [
          {
            "id": 3639230,
            "name": "Hot Chelle Rae",
            "handle": null,
            "type": "MAIN",
            "picture": "6a18c717-c688-4edd-a277-4446b24cfc5f"
          }
        ],
        "album": {
          "id": 10614373,
          "title": "Whatever",
          "cover": "0621add6-9ff3-4db9-8d2d-1d7ebeae929f",
          "vibrantColor": "#5dbbc5",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001f9d6e033d226ca2335f5f314e98"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 14989955,
        "title": "Some Nights",
        "duration": 277,
        "replayGain": -9.75,
        "peak": 0.997284,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2012-02-21T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 83,
        "copyright": "℗ 2011 Fueled By Ramen LLC",
        "bpm": 108,
        "url": "http://www.tidal.com/track/14989955",
        "isrc": "USAT21104050",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3762073,
          "name": "fun.",
          "handle": null,
          "type": "MAIN",
          "picture": "17b912d0-ca47-4bd8-afa3-2295cb7ba2b0"
        },
        "artists": [
          {
            "id": 3762073,
            "name": "fun.",
            "handle": null,
            "type": "MAIN",
            "picture": "17b912d0-ca47-4bd8-afa3-2295cb7ba2b0"
          }
        ],
        "album": {
          "id": 14989953,
          "title": "Some Nights",
          "cover": "3705a2ff-7b04-4df0-9a6f-c6f326c0232e",
          "vibrantColor": "#83423c",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001e282a3dea08268217c4b4a19168"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 2509410,
        "title": "Right Round (feat. Ke$ha)",
        "duration": 208,
        "replayGain": -9.23,
        "peak": 0.997375,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2025-04-04T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 53,
        "copyright": "℗ 2009 Atlantic Recording Corporation for the United States and WEA International for the world outside of the United States",
        "bpm": 125,
        "url": "http://www.tidal.com/track/2509410",
        "isrc": "USAT20900316",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3500259,
          "name": "Flo Rida",
          "handle": null,
          "type": "MAIN",
          "picture": "2ab3695b-31bd-41bb-b107-f4b79b8c4d23"
        },
        "artists": [
          {
            "id": 3500259,
            "name": "Flo Rida",
            "handle": null,
            "type": "MAIN",
            "picture": "2ab3695b-31bd-41bb-b107-f4b79b8c4d23"
          },
          {
            "id": 3652939,
            "name": "Kesha",
            "handle": null,
            "type": "FEATURED",
            "picture": "2d915c0c-8ad0-4b00-a6ec-d9d931e30c9d"
          }
        ],
        "album": {
          "id": 2509409,
          "title": "Right Round (feat. Ke$ha)",
          "cover": "219ab19f-1f2b-4708-9832-a47d15c3e527",
          "vibrantColor": "#e5d6a3",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001bf7dedd5bf4e9266e6b80017e0c"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 80236586,
        "title": "That's All",
        "duration": 210,
        "replayGain": -8.96,
        "peak": 1,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2017-10-23T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 61,
        "copyright": "2017 Evan Kijek",
        "bpm": null,
        "url": "http://www.tidal.com/track/80236586",
        "isrc": "USHM91726727",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 6996473,
          "name": "Skitz Kraven",
          "handle": null,
          "type": "MAIN",
          "picture": "3c581c47-0058-429e-971a-de37887505c7"
        },
        "artists": [
          {
            "id": 6996473,
            "name": "Skitz Kraven",
            "handle": null,
            "type": "MAIN",
            "picture": "3c581c47-0058-429e-971a-de37887505c7"
          }
        ],
        "album": {
          "id": 80236585,
          "title": "That's All",
          "cover": "1e03f751-485d-455b-936f-092beefe7672",
          "vibrantColor": "#FFFFFF",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001af78791332f7218ee70c111e174"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3006321,
        "title": "Bulletproof",
        "duration": 206,
        "replayGain": -12.6,
        "peak": 0.999237,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2021-03-08T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 4,
        "volumeNumber": 1,
        "version": null,
        "popularity": 81,
        "copyright": "℗ 2009 Polydor Ltd. (UK)",
        "bpm": 123,
        "url": "http://www.tidal.com/track/3006321",
        "isrc": "GBUM70903779",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3630789,
          "name": "La Roux",
          "handle": null,
          "type": "MAIN",
          "picture": "7ba03612-2002-4e7a-823c-bd5e44ef9fc6"
        },
        "artists": [
          {
            "id": 3630789,
            "name": "La Roux",
            "handle": null,
            "type": "MAIN",
            "picture": "7ba03612-2002-4e7a-823c-bd5e44ef9fc6"
          }
        ],
        "album": {
          "id": 3006317,
          "title": "La Roux",
          "cover": "0ed96bab-61bf-4a9c-826a-9f90f8184bfc",
          "vibrantColor": "#eab04d",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0015ed29029202affcede6543bf18c"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 2256444,
        "title": "You Found Me",
        "duration": 242,
        "replayGain": -9.53,
        "peak": 0.997345,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-02-03T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 3,
        "volumeNumber": 1,
        "version": null,
        "popularity": 82,
        "copyright": "(P) 2009 Epic Records, a division of Sony Music Entertainment",
        "bpm": null,
        "url": "http://www.tidal.com/track/2256444",
        "isrc": "USSM10803009",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 6895,
          "name": "The Fray",
          "handle": null,
          "type": "MAIN",
          "picture": "097cd296-0879-4ce7-a801-59433b4d43fd"
        },
        "artists": [
          {
            "id": 6895,
            "name": "The Fray",
            "handle": null,
            "type": "MAIN",
            "picture": "097cd296-0879-4ce7-a801-59433b4d43fd"
          }
        ],
        "album": {
          "id": 2256322,
          "title": "The Fray",
          "cover": "f8afe9c3-4a03-4eee-bd62-f8df7b9808a3",
          "vibrantColor": "#e4d398",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001d819d55a28e867a9367d82af5d6"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 5888036,
        "title": "Whataya Want from Me",
        "duration": 226,
        "replayGain": -10.62,
        "peak": 0.997375,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2011-03-15T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 11,
        "volumeNumber": 1,
        "version": null,
        "popularity": 70,
        "copyright": "(P) 2009 19 Recordings Ltd. under exclusive license to RCA Records, a division of Sony Music Entertainment",
        "bpm": 93,
        "url": "http://www.tidal.com/track/5888036",
        "isrc": "GBCTA0900344",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3620693,
          "name": "Adam Lambert",
          "handle": null,
          "type": "MAIN",
          "picture": "9764d344-b292-4cf8-9bad-f4eb0078f20c"
        },
        "artists": [
          {
            "id": 3620693,
            "name": "Adam Lambert",
            "handle": null,
            "type": "MAIN",
            "picture": "9764d344-b292-4cf8-9bad-f4eb0078f20c"
          }
        ],
        "album": {
          "id": 5887999,
          "title": "10th Anniversary - The Hits - Volume 1",
          "cover": "aa625470-7cc7-428a-b0bf-94823ec82636",
          "vibrantColor": "#a5d6e4",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001260421e891ad5d85a99e9a3d4d7"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 121131485,
        "title": "Home Is For The Heartless",
        "duration": 248,
        "replayGain": -10.26,
        "peak": 0.988556,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2010-06-29T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 10,
        "volumeNumber": 1,
        "version": null,
        "popularity": 54,
        "copyright": "Resist Records, Under Exclusive License To Epitaph",
        "bpm": null,
        "url": "http://www.tidal.com/track/121131485",
        "isrc": "USEP41014010",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 17808,
          "name": "Parkway Drive",
          "handle": null,
          "type": "MAIN",
          "picture": "29c93abb-cf68-4e77-8581-b363826e593e"
        },
        "artists": [
          {
            "id": 17808,
            "name": "Parkway Drive",
            "handle": null,
            "type": "MAIN",
            "picture": "29c93abb-cf68-4e77-8581-b363826e593e"
          }
        ],
        "album": {
          "id": 121131475,
          "title": "Deep Blue",
          "cover": "e6ea5a23-8d7a-4960-aa11-beeed78e560a",
          "vibrantColor": "#a5cbd9",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0012ea57e84bce578fdfc9a1c1f772"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 166711390,
        "title": "Dear X, You Don't Own Me",
        "duration": 216,
        "replayGain": -12.11,
        "peak": 1,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2021-01-04T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 50,
        "copyright": "(P) 2010 INO Records",
        "bpm": null,
        "url": "http://www.tidal.com/track/166711390",
        "isrc": "USM2C1000601",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3531137,
          "name": "Disciple",
          "handle": null,
          "type": "MAIN",
          "picture": "ccaeb98a-1c3f-44d5-8c52-13303fe7a0c1"
        },
        "artists": [
          {
            "id": 3531137,
            "name": "Disciple",
            "handle": null,
            "type": "MAIN",
            "picture": "ccaeb98a-1c3f-44d5-8c52-13303fe7a0c1"
          }
        ],
        "album": {
          "id": 166711389,
          "title": "The Best of Disciple",
          "cover": "12554338-747d-4c75-a456-786b489d733d",
          "vibrantColor": "#bd6c62",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001b7dbafae4daf6c0e1c69ceed962"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 5119488,
        "title": "How to Save a Life",
        "duration": 263,
        "replayGain": -11.72,
        "peak": 0.997314,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2005-09-13T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 3,
        "volumeNumber": 1,
        "version": null,
        "popularity": 86,
        "copyright": "(P) 2005 Epic Records, a division of Sony Music Entertainment",
        "bpm": 122,
        "url": "http://www.tidal.com/track/5119488",
        "isrc": "USSM10601178",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 6895,
          "name": "The Fray",
          "handle": null,
          "type": "MAIN",
          "picture": "097cd296-0879-4ce7-a801-59433b4d43fd"
        },
        "artists": [
          {
            "id": 6895,
            "name": "The Fray",
            "handle": null,
            "type": "MAIN",
            "picture": "097cd296-0879-4ce7-a801-59433b4d43fd"
          }
        ],
        "album": {
          "id": 5119485,
          "title": "How To Save A Life",
          "cover": "9825549e-70c5-4962-8819-281f745d2678",
          "vibrantColor": "#f8ea8e",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0013b358ecab0c8fc99dfe9f62da98"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 1317019,
        "title": "Kids",
        "duration": 303,
        "replayGain": -12.03,
        "peak": 0.997314,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-11-24T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 5,
        "volumeNumber": 1,
        "version": null,
        "popularity": 87,
        "copyright": "(P) 2007 Columbia Records, a division of Sony Music Entertainment",
        "bpm": 123,
        "url": "http://www.tidal.com/track/1317019",
        "isrc": "USSM10702135",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3504370,
          "name": "MGMT",
          "handle": null,
          "type": "MAIN",
          "picture": "3cab46b4-9eef-44a8-b8e0-a2036c8f9a34"
        },
        "artists": [
          {
            "id": 3504370,
            "name": "MGMT",
            "handle": null,
            "type": "MAIN",
            "picture": "3cab46b4-9eef-44a8-b8e0-a2036c8f9a34"
          }
        ],
        "album": {
          "id": 1317014,
          "title": "Oracular Spectacular",
          "cover": "5da65637-df7f-4332-b521-261b3e103a40",
          "vibrantColor": "#bb1c34",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0010118e749467d9c9e4d62b0363a7"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 111301453,
        "title": "Impossible",
        "duration": 227,
        "replayGain": -12.34,
        "peak": 1,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2019-06-21T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 23,
        "volumeNumber": 1,
        "version": null,
        "popularity": 63,
        "copyright": "℗ 2010 Universal Records, a Division of UMG Recordings, Inc. and SRC Records, Inc.",
        "bpm": 90,
        "url": "http://www.tidal.com/track/111301453",
        "isrc": "USUM71002550",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3557975,
          "name": "Shontelle",
          "handle": null,
          "type": "MAIN",
          "picture": "9cc63ca0-cef8-4720-ac00-5be2aacdf231"
        },
        "artists": [
          {
            "id": 3557975,
            "name": "Shontelle",
            "handle": null,
            "type": "MAIN",
            "picture": "9cc63ca0-cef8-4720-ac00-5be2aacdf231"
          }
        ],
        "album": {
          "id": 111301426,
          "title": "Breakup Songs",
          "cover": "40a1e734-a78f-4ebf-86f7-39a67f88395a",
          "vibrantColor": "#e34d3d",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001e0ba31b43364f18927ae0156f31"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35454220,
        "title": "Price Tag",
        "duration": 223,
        "replayGain": -11.01,
        "peak": 0.998993,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2011-11-21T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 76,
        "copyright": "℗ 2010 Universal Republic Records, a division of UMG Recordings, Inc. & Lava Music, LLC",
        "bpm": 172,
        "url": "http://www.tidal.com/track/35454220",
        "isrc": "USUM71029357",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3828746,
          "name": "Jessie J",
          "handle": null,
          "type": "MAIN",
          "picture": "f4090c0a-95f7-4580-b4df-cc9ca5cdd2ec"
        },
        "artists": [
          {
            "id": 3828746,
            "name": "Jessie J",
            "handle": null,
            "type": "MAIN",
            "picture": "f4090c0a-95f7-4580-b4df-cc9ca5cdd2ec"
          },
          {
            "id": 3667959,
            "name": "B.o.B",
            "handle": null,
            "type": "FEATURED",
            "picture": "530f127a-9fd5-4b8c-b2f9-cdde70873130"
          }
        ],
        "album": {
          "id": 35454219,
          "title": "Who You Are (Platinum Edition)",
          "cover": "1bca6f42-286e-4e14-ba08-0edb3871c1be",
          "vibrantColor": "#e7e0a5",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001df33f0a5d4ac00f21d40165c351"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 3383653,
        "title": "Fight For This Love",
        "duration": 223,
        "replayGain": -9.72,
        "peak": 0.988525,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2015-07-10T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 4,
        "volumeNumber": 1,
        "version": null,
        "popularity": 58,
        "copyright": "℗ 2009 Polydor Ltd. (UK)",
        "bpm": 123,
        "url": "http://www.tidal.com/track/3383653",
        "isrc": "GBUM70911955",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 4610414,
          "name": "Cheryl",
          "handle": null,
          "type": "MAIN",
          "picture": "c83ede6e-d47a-4834-981b-8ca2cc0a4041"
        },
        "artists": [
          {
            "id": 4610414,
            "name": "Cheryl",
            "handle": null,
            "type": "MAIN",
            "picture": "c83ede6e-d47a-4834-981b-8ca2cc0a4041"
          }
        ],
        "album": {
          "id": 3383649,
          "title": "3 Words",
          "cover": "1ee454c8-a1c6-40b5-bf38-6fea4c044dc2",
          "vibrantColor": "#FFFFFF",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001b447035ef9624763974e904e397"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35082899,
        "title": "I'm Yours",
        "duration": 243,
        "replayGain": -10.61,
        "peak": 0.999176,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-07-14T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 79,
        "copyright": "℗ 2008 WEA International Inc",
        "bpm": 151,
        "url": "http://www.tidal.com/track/35082899",
        "isrc": "USEE10800667",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 14721,
          "name": "Jason Mraz",
          "handle": null,
          "type": "MAIN",
          "picture": "7e0f9138-6770-4076-acf5-2451d4161773"
        },
        "artists": [
          {
            "id": 14721,
            "name": "Jason Mraz",
            "handle": null,
            "type": "MAIN",
            "picture": "7e0f9138-6770-4076-acf5-2451d4161773"
          }
        ],
        "album": {
          "id": 35082898,
          "title": "Yours Truly: The I'm Yours Collection",
          "cover": "b3e90d4b-7e95-4db5-8ba0-792491d25425",
          "vibrantColor": "#e1d097",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0018beb32b272ffc6537cb2deef0e5"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 4787050,
        "title": "Survive",
        "duration": 192,
        "replayGain": -11.88,
        "peak": 0.986114,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2010-10-08T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 7,
        "volumeNumber": 1,
        "version": null,
        "popularity": 47,
        "copyright": "A Virgin Records America Release; ℗ 2009 Capitol Records, LLC",
        "bpm": null,
        "url": "http://www.tidal.com/track/4787050",
        "isrc": "USVI20900229",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 10640,
          "name": "Sick Puppies",
          "handle": null,
          "type": "MAIN",
          "picture": "311ef149-5e8a-4016-b024-9b67c668eff0"
        },
        "artists": [
          {
            "id": 10640,
            "name": "Sick Puppies",
            "handle": null,
            "type": "MAIN",
            "picture": "311ef149-5e8a-4016-b024-9b67c668eff0"
          }
        ],
        "album": {
          "id": 4787043,
          "title": "Tri-polar",
          "cover": "7f6b2984-6167-4a24-90ec-31233917f247",
          "vibrantColor": "#FFFFFF",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0010e63bf79782ceff2618da825b2f"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 90320810,
        "title": "So What",
        "duration": 215,
        "replayGain": -13.61,
        "peak": 0.997345,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2018-06-22T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 21,
        "volumeNumber": 1,
        "version": null,
        "popularity": 79,
        "copyright": "(P) 2008 Arista Records LLC",
        "bpm": 126,
        "url": "http://www.tidal.com/track/90320810",
        "isrc": "USLF20800067",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 1900,
          "name": "P!nk",
          "handle": null,
          "type": "MAIN",
          "picture": "1df06ed4-fd75-45f4-ab93-492c3e3acba7"
        },
        "artists": [
          {
            "id": 1900,
            "name": "P!nk",
            "handle": null,
            "type": "MAIN",
            "picture": "1df06ed4-fd75-45f4-ab93-492c3e3acba7"
          }
        ],
        "album": {
          "id": 90320789,
          "title": "Orgullo Gay (Streaming Only)",
          "cover": "816f6564-9a95-4c44-a256-b1f55b8c145b",
          "vibrantColor": "#f4d266",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0012f441f2406354feb3422cad6bab"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 40913191,
        "title": "Never Giving Up",
        "duration": 220,
        "replayGain": -10.69,
        "peak": 0.988403,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2015-02-24T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 14,
        "volumeNumber": 1,
        "version": null,
        "popularity": 60,
        "copyright": "℗ 2015 Rise Records",
        "bpm": null,
        "url": "http://www.tidal.com/track/40913191",
        "isrc": "USEK71527014",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3877856,
          "name": "Of Mice & Men",
          "handle": null,
          "type": "MAIN",
          "picture": "7c6698e6-14de-44b2-8dc7-a7040a209d44"
        },
        "artists": [
          {
            "id": 3877856,
            "name": "Of Mice & Men",
            "handle": null,
            "type": "MAIN",
            "picture": "7c6698e6-14de-44b2-8dc7-a7040a209d44"
          }
        ],
        "album": {
          "id": 40913177,
          "title": "Restoring Force: Full Circle",
          "cover": "2ae98e8c-b03d-4355-9f45-3b977c72f87f",
          "vibrantColor": "#FFFFFF",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001be3cff385641e7db04faf826c5e"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 7094875,
        "title": "Beautiful Soul",
        "duration": 214,
        "replayGain": -10.42,
        "peak": 0.997742,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2004-09-28T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 2,
        "volumeNumber": 1,
        "version": null,
        "popularity": 80,
        "copyright": "℗ 2004 Hollywood Records, Inc.",
        "bpm": 90,
        "url": "http://www.tidal.com/track/7094875",
        "isrc": "USHR10421324",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 9688,
          "name": "Jesse McCartney",
          "handle": null,
          "type": "MAIN",
          "picture": "7c992209-e62b-4631-ba83-c4105ad166e0"
        },
        "artists": [
          {
            "id": 9688,
            "name": "Jesse McCartney",
            "handle": null,
            "type": "MAIN",
            "picture": "7c992209-e62b-4631-ba83-c4105ad166e0"
          }
        ],
        "album": {
          "id": 7094873,
          "title": "Beautiful Soul",
          "cover": "eacdb6d5-fe75-4b91-8a8b-535562b5c5a3",
          "vibrantColor": "#c6df9d",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001147a3187223bce1143bb0506c68"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 76773937,
        "title": "Chasing Cars",
        "duration": 268,
        "replayGain": -10.89,
        "peak": 0.810486,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2017-08-04T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 11,
        "volumeNumber": 1,
        "version": null,
        "popularity": 79,
        "copyright": "℗ 2006 Polydor Ltd. (UK)",
        "bpm": 104,
        "url": "http://www.tidal.com/track/76773937",
        "isrc": "GBUM70600345",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 16998,
          "name": "Snow Patrol",
          "handle": null,
          "type": "MAIN",
          "picture": "164be9c0-a0b9-407c-b175-36cbc7f4afed"
        },
        "artists": [
          {
            "id": 16998,
            "name": "Snow Patrol",
            "handle": null,
            "type": "MAIN",
            "picture": "164be9c0-a0b9-407c-b175-36cbc7f4afed"
          }
        ],
        "album": {
          "id": 76773926,
          "title": "00s Mixtape",
          "cover": "4201f5ce-fd8e-4813-a8b5-b93bb4bb74ce",
          "vibrantColor": "#d9a09c",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001850eb2ac3a0a0c6e606f8388460"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 4833306,
        "title": "Raise Your Glass",
        "duration": 203,
        "replayGain": -10.03,
        "peak": 0.997375,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2010-11-01T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 81,
        "copyright": "(P) 2010 Arista Records LLC",
        "bpm": 122,
        "url": "http://www.tidal.com/track/4833306",
        "isrc": "USLF21000090",
        "editable": false,
        "explicit": true,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 1900,
          "name": "P!nk",
          "handle": null,
          "type": "MAIN",
          "picture": "1df06ed4-fd75-45f4-ab93-492c3e3acba7"
        },
        "artists": [
          {
            "id": 1900,
            "name": "P!nk",
            "handle": null,
            "type": "MAIN",
            "picture": "1df06ed4-fd75-45f4-ab93-492c3e3acba7"
          }
        ],
        "album": {
          "id": 4833129,
          "title": "Raise Your Glass",
          "cover": "1a31ff0b-3de6-49e4-bdf8-6abf07998d83",
          "vibrantColor": "#ca578a",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0012d214851a2447de0efa339d75eb"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 1906287,
        "title": "Use Somebody",
        "duration": 231,
        "replayGain": -10.15,
        "peak": 0.997314,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2008-09-23T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 4,
        "volumeNumber": 1,
        "version": null,
        "popularity": 85,
        "copyright": "(P) 2008 RCA Records, a division of Sony Music Entertainment",
        "bpm": 136,
        "url": "http://www.tidal.com/track/1906287",
        "isrc": "USRC10800301",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 1553,
          "name": "Kings Of Leon",
          "handle": null,
          "type": "MAIN",
          "picture": "31f31abf-7fc4-479c-b8af-0c42eb53d59e"
        },
        "artists": [
          {
            "id": 1553,
            "name": "Kings Of Leon",
            "handle": null,
            "type": "MAIN",
            "picture": "31f31abf-7fc4-479c-b8af-0c42eb53d59e"
          }
        ],
        "album": {
          "id": 1906011,
          "title": "Only By The Night",
          "cover": "808c5d3c-5961-4a98-9573-b8cdc8c8ef40",
          "vibrantColor": "#e5ada9",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "0017a2e2a80ddbc9cfa9f2e1c6dbd8"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 35754165,
        "title": "We Are Golden",
        "duration": 238,
        "replayGain": -10.22,
        "peak": 0.999969,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2009-09-22T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 1,
        "volumeNumber": 1,
        "version": null,
        "popularity": 51,
        "copyright": "℗ 2009 Casablanca Music, LLC",
        "bpm": 123,
        "url": "http://www.tidal.com/track/35754165",
        "isrc": "USC7R0900125",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 16926,
          "name": "MIKA",
          "handle": null,
          "type": "MAIN",
          "picture": "ea5fc0d1-27a9-4c92-a36c-42fc37150e86"
        },
        "artists": [
          {
            "id": 16926,
            "name": "MIKA",
            "handle": null,
            "type": "MAIN",
            "picture": "ea5fc0d1-27a9-4c92-a36c-42fc37150e86"
          }
        ],
        "album": {
          "id": 35754164,
          "title": "The Boy Who Knew Too Much",
          "cover": "becc487e-ec4a-412b-a121-8bce28791e24",
          "vibrantColor": "#e6de9e",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00146e7326a598cc61d2cbb928c123"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 23273354,
        "title": "Human",
        "duration": 245,
        "replayGain": -12,
        "peak": 0.984375,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2013-11-11T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 8,
        "volumeNumber": 1,
        "version": null,
        "popularity": 76,
        "copyright": "℗ 2008 UMG Recordings, Inc.",
        "bpm": 135,
        "url": "http://www.tidal.com/track/23273354",
        "isrc": "USUM70837367",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 15686,
          "name": "The Killers",
          "handle": null,
          "type": "MAIN",
          "picture": "cea18ed3-1dde-4858-b3df-b3564be7e360"
        },
        "artists": [
          {
            "id": 15686,
            "name": "The Killers",
            "handle": null,
            "type": "MAIN",
            "picture": "cea18ed3-1dde-4858-b3df-b3564be7e360"
          }
        ],
        "album": {
          "id": 23273346,
          "title": "Direct Hits",
          "cover": "b942b2c7-1827-4238-bf9a-44787ec11ccf",
          "vibrantColor": "#FFFFFF",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "001eb691cf866c3f777fb3b4697d33"
        }
      },
      "type": "track"
    },
    {
      "item": {
        "id": 1706786,
        "title": "Shake It",
        "duration": 180,
        "replayGain": -9.89,
        "peak": 0.997375,
        "allowStreaming": true,
        "streamReady": true,
        "payToStream": false,
        "adSupportedStreamReady": true,
        "djReady": true,
        "stemReady": false,
        "streamStartDate": "2007-09-18T00:00:00.000+0000",
        "premiumStreamingOnly": false,
        "trackNumber": 4,
        "volumeNumber": 1,
        "version": null,
        "popularity": 80,
        "copyright": "(P) 2007 Sony Music Entertainment",
        "bpm": 150,
        "url": "http://www.tidal.com/track/1706786",
        "isrc": "USSM10702537",
        "editable": false,
        "explicit": false,
        "audioQuality": "LOSSLESS",
        "audioModes": [
          "STEREO"
        ],
        "mediaMetadata": {
          "tags": [
            "LOSSLESS"
          ]
        },
        "upload": false,
        "accessType": "PUBLIC",
        "spotlighted": false,
        "artist": {
          "id": 3529763,
          "name": "Metro Station",
          "handle": null,
          "type": "MAIN",
          "picture": "b263a4d6-d9ef-4182-908f-c4cfe738097f"
        },
        "artists": [
          {
            "id": 3529763,
            "name": "Metro Station",
            "handle": null,
            "type": "MAIN",
            "picture": "b263a4d6-d9ef-4182-908f-c4cfe738097f"
          }
        ],
        "album": {
          "id": 1706782,
          "title": "Metro Station",
          "cover": "aee09d46-9316-41a9-9812-532688daa829",
          "vibrantColor": "#dad7a4",
          "videoCover": null
        },
        "mixes": {
          "TRACK_MIX": "00194c630df5bb6f50396bc302df43"
        }
      },
      "type": "track"
    }
  ]
}

```
<br>


### Status Codes

HIFI returns the following status codes in its API:

> | Status Code | Description |
> | :---        | :--- |
> | 200         | `OK` |
> | 422         | `UNPROCESSABLE CONTENT` |
> | 404         | `NOT FOUND` |
> | 500         | `INTERNAL SERVER ERROR` |


</details>

------------------------------------------------------------------------------------------

<br>

## 🫂 Contributing
Please refer to [CONTRIBUTING.md](./CONTRIBUTING.md).


<br>

## 🔐 Security Policy
Please refer to [SECURITY.md](./SECURITY.md).

<br>

## 👩‍⚖️ License

This project is licensed under the terms of the MIT license.
