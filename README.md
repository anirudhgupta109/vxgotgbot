# Getting started
## Using the bot as a user
Just add the bot to your group from: https://t.me/twitter2vx_bot
and the bot will auto post on twitter and instagram links

## Running your own instance of the bot
Download the repository and download its dependencies with:

```bash
git clone https://github.com/anirudhgupta109/vxgotgbot
cd vxgotgbot
go mod tidy
```

Set the enviroment variable containing the bots API token:

```bash
export TELEGRAM_APITOKEN=<VALUE HERE>
```

And finally run the bot:
```bash
go run .
```