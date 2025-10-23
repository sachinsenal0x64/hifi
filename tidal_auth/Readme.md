### Get tokens

```sh
git clone https://github.com/sachinsenal0x64/hifi
cd hifi/tidal_auth

# Rename .env-example
cp .env.example .env

# Edit the .env file
CLIENT_ID=4N3n6Q1x95LL5K7p
CLIENT_SECRET=oKOXfJW371cX6xaZ0PyhgGNBdNLlBZd4AKKYougMjik=

pip install -r requirements.txt

python tidal_auth.py

```
