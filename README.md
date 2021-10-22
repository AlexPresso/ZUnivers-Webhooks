# ZUnivers-Webhooks

Discord webhooks dispatcher for the ZUnivers card game.

![preview](https://i.imgur.com/pAWleMn.png)

## Disclaimer

This is project is not affiliated with the ZUnivers's project. It's a community project.

## Features

- ✅ `!daily` reminder
- ✅ Notify for new patchnotes
- ✅ Notify for configuration changes
- ✅ Notify for new webapp versions
- ✅ Notify for new items/item changes
- ✅ Notify for new packs/packs changes
- ✅ Notify for new "ascension" season
- ✅ Per event multiple webhook dispatching

## Usage

- Install the package with `go install github.com/alexpresso/zunivers-webhooks@latest`
- Copy `config.example.json` to `config.json` + replace all the `https://discordapp/...` URLs by your webhook URLs
- Create one or more [webhook endpoint(s)](https://support.discord.com/hc/fr/articles/228383668-Utiliser-les-Webhooks)
- Copy/paste their URL into the `urls` property of each event you want to listen
- Open a CLI and type `zunivers-webhooks`

