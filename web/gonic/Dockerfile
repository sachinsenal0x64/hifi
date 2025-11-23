FROM sentriz/gonic:latest

EXPOSE 4747

ENV GONIC_LISTEN_ADDR=:4747 \
    GONIC_DB_PATH=/data/gonic.db \
    GONIC_MUSIC_PATH=/music \
    GONIC_PODCAST_PATH=/podcasts \
    GONIC_CACHE_PATH=/cache \
    GONIC_PLAYLISTS_PATH=/playlists

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

