# decklyst

Card and Deck API for Duelyst

## Sample API Responses

```
GET http://decklyst.xyz/card/501

{
    "id": 501,
    "name": "Faie Bloodwing",
    "factionId": 6,
    "faction": "Vanar Kindred",
    "setName": "Core",
    "rarityId": 0,
    "rarity": "Basic",
    "mana": 0,
    "attack": 2,
    "hp": 25,
    "category": "unit",
    "type": "General",
    "isGeneral": true,
    "description": "<b>Bloodborn Spell:</b> Deal 2 damage to all enemies in the enemy General's Column.",
    "searchableContent": "Faie Bloodwing <b>Bloodborn Spell:</b> Deal 2 damage to all enemies in the enemy General's Column. Basic minion unit general Core Set Core core 2atk 25hp ",
    "frame": "f6_general_idle_",
    "plist": "https://assets-counterplaygames.netdna-ssl.com/production/resources/units/f6_general.plist",
    "sprite": "https://assets-counterplaygames.netdna-ssl.com/production/resources/units/f6_general.png"
}
``` 

```
GET http://decklyst.xyz/deck/MTo1MDEsMzo1MTcsMjo1MzgsMzo1NDAsMzoxMDMwMiwyOjEwMzAzLDM6MTAzMDUsMjoxMTA5NCwzOjExMDk3LDM6MjAxMzQsMzoyMDEzOSwyOjIwMTQ0LDI6MjAxNDcsMjoyMDIwNywzOjIwMjM3LDM6MjAyNjE=

{
    "Faction": "Vanar Kindred",
    "General": "Faie Bloodwing",
    "Cards": {
        "Aspect of the Fox": 2,
        "Blue Conjurer": 2,
        "Chromatic Cold": 3,
        "Circulus": 3,
        "Flash Freeze": 3,
        "Frigid Corona": 3,
        "Frostburn": 2,
        "Grandmaster Embla": 2,
        "Gravity Well": 2,
        "Hearth-Sister": 3,
        "Mana Deathgrip": 3,
        "Manaforger": 3,
        "Owlbeast Sage": 3,
        "Prismatic Illusionist": 2,
        "Trinity Wing": 3
    }
}
```

## To Do

1. Fix deploy to check for existing assets.
2. Remove console.log() statements from GetCardJSON().
3. GraphQL?
