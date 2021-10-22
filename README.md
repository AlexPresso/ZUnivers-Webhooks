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
- Open a CLI and run `zunivers-webhooks` to make it create a default config
- Edit `config.json` and add
  your [webhook endpoint(s)](https://support.discord.com/hc/fr/articles/228383668-Utiliser-les-Webhooks) URLs under the
  needed events.
- Run `zunivers-webhooks` again and you're done.

`config.json` example for the `config_changed` event:

```json
{
  "config_changed": {
    "urls": [
      "https://discord.com/api/webhooks/123456789/........."
    ],
    "message": "Un paramètre de configuration a changé !"
  }
}
```

