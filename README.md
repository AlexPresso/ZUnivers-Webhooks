# ZUnivers-Webhooks

Discord webhooks dispatcher for the ZUnivers card game.

<img src="https://repository-images.githubusercontent.com/419089778/e78f5f5c-49cc-429a-8285-2e043b16fe05" height="320px">

## Disclaimer

This project is not affiliated with the ZUnivers's project. It's a community project.

## Features

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
- ✅ Multiple webhooks dispatching

## Usage

- Install the package with `go install github.com/alexpresso/zunivers-webhooks@latest`
- Open a CLI and run `zunivers-webhooks` to make it create a default config
- Edit `config.json` and add
  your [webhook endpoint(s)](https://support.discord.com/hc/fr/articles/228383668-Utiliser-les-Webhooks). You can add
  multiple webhook URLs to dilute the Discord rate limiting (number of message you can send per second) on multiple
  endpoints. I recommend you add at least 5.
- Run `zunivers-webhooks` again and you're done.

`config.json` example:

```json
{
  "frontBaseUrl": "https://zunivers.zerator.com",
  "cdnBaseUrl": "https://minio-zunivers-prod.prod.poneyy.fr/zunivers-prod",
  "api": {
    "baseUrl": "https://zunivers-api.zerator.com",
    "timeout": 10
  },
  "webhooks": [
    "https://discord.com/api/webhooks/123456789/........."
  ],
  "messages": {
    ...
  }
}
```

