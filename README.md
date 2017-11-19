# IMT6281 submission 3

## Dependencies

Dependencies given are required to make sense out of the following code. Which was also worked on during this submission as a contribution.

- [s1kx/unison](https://github.com/s1kx/unison)
- [andersfylling/concurrencyparser](https://github.com/andersfylling/concurrencyparser)

Please review `concurrencyparser` as this is a main module used to parse the message queries in the server.

## About

The bot is hosted on a Synology server. The reason Heroku was not used is due to the free Heroku version requiring that given environment variable `PORT` is bound to by a active web service. If this web service does not have any incoming requests at least once an hour, the Heroku dyno will go idle and the Bot will die.

A Docker solution was created in stead to support cross platform hosting. It requires golang 1.8 and glide. Prebuild version can be found on [Docker hub](https://hub.docker.com/r/andersfylling/imt2681bot/). Note. See `Setting up from scratch` for required environment variables.

## Add to your Discord server

Clicking [this link](https://discordapp.com/api/oauth2/authorize?scope=bot&permissions=0&client_id=376052576761937920) will add the bot to your server. The bot instantly replies to any new message events and have no commands. The only requirements required is to read messages and send messages.

If you want this bot to only work on one channel, simply deactivate message permissions for `@everyone` and give the bot message permissions on desired channel.

## Setting up from scratch

Required environment variables, regardless of hosting solution, are:

- IMT_BOT_TOKEN
- IMT_BOT_COMMAND_PREFIX

`IMT_BOT_TOKEN` is required, and the bot will not run without it. This is the Discord token found in the developer sites.

`IMT_BOT_COMMAND_PREFIX` is optional, and honestly not needed. It defines what the command invoker prefix is. By default mention is used to invoke commands, but this bot does not have any commands yet.
