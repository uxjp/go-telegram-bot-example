# go-telegram-bot-example

This repo serves as a starting point for developing Telegram Bots with golang.  
This `README.md` contains a comprehensive step-by-step guide that I used for setting up the bot.

## Security

I personally consider using `.env` files a bad practice, so I don't use them to store my Token.  
Why don’t I like it? Because I've accidentally committed secrets too many times. If that wasn’t reason enough, I also think it takes too much time to manually set up these `.env` files in numerous folders when handling multiple environments.
a
The scope of this repo is local development, and for this purpose, I'm storing my secrets with a free tool called  
`doppler.com`. It allows me to log in locally and have my secrets injected into my app while I execute the command to run it, like:

```bash
doppler run -- go run bot.go
```