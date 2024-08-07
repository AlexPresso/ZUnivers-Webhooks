# ZUnivers-Webhooks

Discord webhooks dispatcher for the ZUnivers card game.

<img src="https://repository-images.githubusercontent.com/419089778/e78f5f5c-49cc-429a-8285-2e043b16fe05" height="320px">

## Disclaimer

This project is not affiliated with the ZUnivers's project. It's a community project.

## Features
All events/checks can be enabled/disabled based on your needs using the `enabled: true/false` booleans in config file

- ✅ `!daily` reminder
- ✅ Notifies for new patchnotes
- ✅ Notifies for configuration changes
- ✅ Notifies for new webapp versions
- ✅ Notifies for new items/item changes
- ✅ Notifies for new packs/packs changes
- ✅ Notifies for new banners/banners changes
- ✅ Notifies for new "ascension" season
- ✅ Notifies for new event/events changes
- ✅ Notifies for new achievements/achievements changes
- ✅ Notifies for new challenges
- ✅ Notifies for shop changes
- ✅ Notifies for API response model changes (disabled by default)
- ✅ Multiple webhooks dispatching (as a pool to mitigate Discord rate limits)

## Usage

- Install the package with `go install github.com/alexpresso/zunivers-webhooks@latest`
- Open a CLI and run `zunivers-webhooks` to make it create a default config
- Edit `config.json` and add
  your [webhook endpoint(s)](https://support.discord.com/hc/fr/articles/228383668-Utiliser-les-Webhooks). You can add
  multiple webhook URLs to dilute the Discord rate limiting (number of message you can send per second) on multiple
  endpoints. I recommend you add at least 5.
- Run `zunivers-webhooks` again and you're done.

Note: you can also deploy zunivers-webhooks to a kubernetes cluster using the
following [helm chart](https://github.com/AlexPresso/helm.alexpresso.me/tree/main/charts/zunivers-webhooks)

`config.json` example: see [this file](https://github.com/AlexPresso/ZUnivers-Webhooks/blob/main/config.default.json)
