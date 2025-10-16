import asyncio
import json
import webbrowser
from dotenv import load_dotenv
import httpx
import os
import rich


load_dotenv()

client_id = os.getenv("CLIENT_ID")
client_secret = os.getenv("CLIENT_SECRET")
access_token = os.getenv("TIDAL_TOKEN")
refresh_token = os.getenv("TIDAL_REFRESH")
redis_url = os.getenv("REDIS_URL")
redis_port = os.getenv("REDIS_PORT")
redis_password = os.getenv("REDIS_PASSWORD")
user_id = os.getenv("USER_ID")

client_id = client_id
client_secret = client_secret


class Hifi:
    def __init__(self, client_id, scope, url, client_secret):
        self.client_id = client_id
        self.scope = scope
        self.url = url
        self.client_secret = client_secret

    @staticmethod
    def Quality(quality):
        rate = {quality: "HI_RES"}
        return rate[quality]


class Auth(Hifi):
    def __init__(self, client_id, scope, url, client_secret):
        super().__init__(client_id, scope, url, client_secret)
        self.response = None

    async def get_auth_response(self):
        data = {"client_id": self.client_id, "scope": self.scope}
        headers = {
            "User-Agent": "Mozilla/5.0 (Linux; Android 8.0.0; SM-G965F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.109 Mobile Safari/537.36"
        }

        async with httpx.AsyncClient() as client:
            response = await client.post(self.url, data=data, headers=headers)

        self.response = response

    def __str__(self):
        return str(self.response)


async def poll_for_authorization(url, data, auth):
    async with httpx.AsyncClient() as client:
        while True:
            response = await client.post(url, data=data, auth=auth)
            if response.status_code == 200:
                return response.json()
            await asyncio.sleep(5)


async def main():
    authrize = Auth(
        client_id=client_id,
        scope="r_usr+w_usr+w_sub",
        url="https://auth.tidal.com/v1/oauth2/device_authorization",
        client_secret=client_secret,
    )

    await authrize.get_auth_response()
    res = authrize.response.json()

    verifyurl = res["verificationUriComplete"]
    dcode = res["deviceCode"]

    rich.print(verifyurl)
    rich.print(dcode)

    HI_RES = authrize.Quality(quality="True")

    rich.print(HI_RES)

    webbrowser.open(verifyurl)

    url2 = "https://auth.tidal.com/v1/oauth2/token"

    data2 = {
        "client_id": authrize.client_id,
        "scope": authrize.scope,
        "device_code": dcode,
        "grant_type": "urn:ietf:params:oauth:grant-type:device_code",
    }

    basic = (authrize.client_id, authrize.client_secret)

    auth_response = await poll_for_authorization(url2, data2, basic)

    access_token = auth_response["access_token"]
    refresh_token = auth_response["refresh_token"]
    user_id = auth_response["user"]["userId"]
    accs = {
        "access_token": access_token,
        "refresh_token": refresh_token,
        "userID": user_id,
        "client_ID": client_id,
        "client_secret": client_secret,
    }
    with open("token.json", "w") as file:
        json.dump(accs, file, indent=4)

    with open("token.json", "r") as readfile:
        token = json.loads(readfile.read())
        rich.print(token)
        acs_tok = token["access_token"]

    url3 = f"https://api.tidal.com/v1/tracks/286266926/playbackinfopostpaywall?countryCode=en_US&audioquality={HI_RES}&playbackmode=STREAM&assetpresentation=FULL"

    headers = {"authorization": f"Bearer {acs_tok}"}

    async with httpx.AsyncClient() as client:
        res3 = await client.get(url3, headers=headers)

    rich.print(res3.json())

    print("TOKEN IS VALID")


if __name__ == "__main__":
    asyncio.run(main())
