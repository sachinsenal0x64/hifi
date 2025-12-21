<div align="center">

<h1 align="center"> Hifi </h1>

<h4 align="center">🎵 Tidal Music integration for Subsonic/Jellyfin/Plexamp.<br><br>

</h4>

<img width="1568" height="860" alt="hifi tidal" src="https://sachinsenal0x64.github.io/picx-images-hosting/foamshot-2025-11-19-02-30-23-(Edit)-1.8l0imacwd1.webp" />

<img width="1568" height="860" alt="hifi tidal" src="https://sachinsenal0x64.github.io/picx-images-hosting/castafiore-1.9o0985zqqu.webp" />

HiFi running on - [Castafiore](https://github.com/sawyerf/Castafiore)

<br>

<img width="1568" height="860" alt="hifi tidal" src="https://sachinsenal0x64.github.io/picx-images-hosting/download-1.64eau537wz.webp" />

<img width="1568" height="860" alt="hifi tidal" src="https://sachinsenal0x64.github.io/picx-images-hosting/Untitled.9rjuho5l5w.webp" />

HiFi running on - [Feishin](https://feishin.vercel.app) Web UI

</div>

<div align="center">

</div>

<br><br>

# 💕 Community

> 🍻 Join the community: <a href="https://discord.gg/EbfftZ5Dd4">Discord</a> [![](https://cdn.statically.io/gh/sachinsenal0x64/picx-images-hosting@master/discord.72y8nlaw5mdc.webp)](https://discord.gg/EbfftZ5Dd4)

<br>

# Quickstart

<img width="1568" height="860" alt="hifi tidal" src="https://sachinsenal0x64.github.io/picx-images-hosting/hifi.b9c0j0foq.webp" />

> [!NOTE]
>
> #### Create your HiFi account here, and then use the account with supported Subsonic client.

> [!IMPORTANT]
>
> #### Use the latest stable release of the supported clients.

- [hifi.401658.xyz](https://hifi.401658.xyz)

<br>

## Jellyfin Music Clients Support (WIP)

> [!IMPORTANT]
>
> #### ETA — It is developed in my free time, and the next major update is Jellyfin support nothing else, so there is no actual ETA for it.

HiFi Subsonic compatibility will only work with Feishin, SubStreamer, and Ultrasonic. All future updates will be exclusive to those apps. If HiFi does not work with the Subsonic client, it will not be implemented.

I’m going to support https://www.plex.tv/plexamp. It’s not a new thing I just forgot to add it. I forgot about Plex itself when I moved to Jellyfin, so I won’t support any other third-party Plex apps. It will be exclusive to the official Plexamp app.

The following jellyfin music clients will continue to be compatible with HiFi from the day one release and will remain fully functional unless the app itself is broken.

All other clients should work as well, since they all use Jellyfin’s API. If they don’t, I will eventually make them work. If you want to request support for a specific jellyfin music app, open an [issue](https://github.com/sachinsenal0x64/hifi/issues). Before doing so, please check the HiFi proposal guide below. Anything else that comes without a hifi proposal check will be rejected.

- https://www.fintunes.app
- https://github.com/jmshrv/finamp
- https://streamyfin.app
- https://github.com/jeffvli/feishin
- https://symfonium.app
- https://github.com/LLukas22/Jellyswarrm
- https://github.com/Devioxic/Foxy-Desktop
- https://github.com/Stannnnn/jelly-app
- https://github.com/Taiko2k/Tauon
- https://github.com/pathetic/bloodin
- https://github.com/saltpi/iPlay
- https://github.com/dweymouth/supersonic
- https://github.com/Jellify-Music/App

<br>

> [!IMPORTANT]
>
> # Project Terms

- We do not encourage piracy. This project is made purely for educational and personal use, intended only for listening, not for downloading.
  If you somehow download copyrighted content, you are solely responsible for complying with the relevant laws in your country.

- The HiFi Project assumes no responsibility for any misuse or legal violations arising from the use of this project.

- HiFi project requires a valid Tidal subscription. We encourage users to maintain an active subscription.

- The subsonic client use the maximum quality available from the HiFi API.

- Community projects may be inspired by or related to the HiFi Project but aren’t officially part of it. These projects are run independently, and the HiFi team isn’t responsible for their content or actions. Each project team should follow any relevant laws and handle its own checks and compliance.

- The HiFi project does not claim ownership of any music or audio content. All rights remain with their respective copyright holders. Users are encouraged to support artists and rights owners by maintaining a valid Tidal subscription. HiFi serves solely as a playback and control interface (e.g., on a Raspberry Pi) for personal, non-commercial use.

<br>

> [!NOTE]
>
> ## Hifi project is for

- Tidal subscribers who don’t like the official app and want more client options.

- Anyone who want to try Tidal music before making a commitment.

- Anyone who wants to listen to hifi-res lossless audio.

- Anyone who doesn’t want to store terabytes of music.

- Anyone who want to access over 110 million songs through their Subsonic-based client.

- Anyone who doesn’t need to download and manually add tracks.

- Developers who want to create their own Subsonic-based client using HiFi as a tidal proxy.

<br>

> [!NOTE]
>
> ## HiFi proposal check

<img width="1568" height="860" alt="hifi tidal" src="https://sachinsenal0x64.github.io/picx-images-hosting/HiFi-proposal-check.26lwrtfomq.png" />

<br>

## API

[API Docs](https://opensubsonic.netlify.app/docs) |
[Hosted API](https://hifi.401658.xyz)

## How can I create my own app based on HiFi?

First, check which [routes](https://github.com/sachinsenal0x64/hifi/tree/main/routes) are supported by Hifi. Hifi uses OpenSubsonic [API](https://opensubsonic.netlify.app), but that doesn’t mean it supports every endpoint that OpenSubsonic provides. Hifi only supports the necessary endpoints required to play music. For example, You can connect the Hifi server to [Feishin](https://github.com/jeffvli/feishin) to observe how it works in practice.

<br>

## Installation (Self-host Hifi)

> [!NOTE]
>
> #### Tidal subscription required.

Pending...

Reason: The current installation is a little bit complex, so I will make this as simple as possible.

<br>

## Contributing

Please refer to [CONTRIBUTING.md](./CONTRIBUTING.md).

## Security Policy

Please refer to [SECURITY.md](./SECURITY.md).

## License

Hifi is released under the [MIT License](https://opensource.org/license/MIT).
