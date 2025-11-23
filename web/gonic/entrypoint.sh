#!/bin/sh
set -e

# create directories INSIDE fly volumes
mkdir -p /data /cache /music /podcasts /playlists

exec gonic

