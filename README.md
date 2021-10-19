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
- ❌ Notify for new items/item changes (SoonTM)
- ❌ Notify for new packs (SoonTM)
- ✅ Per event multiple webhook dispatching

## Usage

- Install the package with `go install github.com/alexpresso/zunivers-webhooks@latest`
- Copy `config.example.json` to `config.json`
- Create one or more [webhook endpoint(s)](https://support.discord.com/hc/fr/articles/228383668-Utiliser-les-Webhooks)
- Copy/paste their URL into the `urls` property of each event you want to listen
- Open a CLI and type `zunivers-webhooks`

`config.json` example:

```json
{
  "frontBaseUrl": "https://zunivers.zerator.com",
  "api": {
    "baseUrl": "https://zunivers-api.zerator.com",
    "timeout": 10
  },
  "webhooks": {
    "config_changed": {
      "urls": [
        "https://discord.com/api/webhooks/123456789123456111/.............................."
      ],
      "message": "Un paramètre de configuration a changé !"
    },
    "new_patchnote": {
      "urls": [
        "https://discord.com/api/webhooks/123456789123456111/.............................."
      ],
      "message": "Une nouvelle patchnote a été publiée !"
    },
    "status_changed": {
      "urls": [
        "https://discord.com/api/webhooks/123456789123456111/.............................."
      ],
      "message": "Le statut de l'application a changé."
    },
    "new_day": {
      "urls": [
        "https://discord.com/api/webhooks/123456789123456111/.............................."
      ],
      "message": "C'est l'heure du `!daily` ! (<#808432657838768168>)"
    }
  }
}
```
