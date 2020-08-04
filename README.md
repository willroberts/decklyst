# decklyst

Card and Deck API for Duelyst.

Uses a local JSON database of all cards in the game (as of version v1.87.1) to
allow queries for cards by ID. Also decodes the game's internal deck code format
to break down a deck into its component cards.

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
GET http://decklyst.xyz/deck/MTo0MDEsMzo0MDUsMzo0MDcsMzo0MTAsMzo0MTIsMzo0MTMsMzo0MTUsMzoxMDAxMiwzOjEwOTU5LDM6MTA5ODEsMzoxOTAzNywzOjIwMTEzLDM6MjAxMTYsMzoyMDE1Nw==

{
    "faction": "Magmar Aspects",
    "general": "Vaath the Immortal",
    "spiritCost": 2550,
    "averageManaCost": 3.3,
    "manaCurve": {
        "1": 3,
        "2": 12,
        "3": 6,
        "4": 9,
        "5": 6,
        "6": 3
    },
    "cards": [
        {
            "id": 405,
            "name": "Makantor Warbeast",
            "count": 3
        },
        {
            "id": 407,
            "name": "Elucidator",
            "count": 3
        },
        {
            "id": 410,
            "name": "Primordial Gazer",
            "count": 3
        },
        {
            "id": 412,
            "name": "Young Silithar",
            "count": 3
        },
        {
            "id": 413,
            "name": "Veteran Silithar",
            "count": 3
        },
        {
            "id": 415,
            "name": "Spirit Harvester",
            "count": 3
        },
        {
            "id": 10012,
            "name": "Saberspine Tiger",
            "count": 3
        },
        {
            "id": 10959,
            "name": "Dancing Blades",
            "count": 3
        },
        {
            "id": 10981,
            "name": "Healing Mystic",
            "count": 3
        },
        {
            "id": 19037,
            "name": "Primus Fist",
            "count": 3
        },
        {
            "id": 20113,
            "name": "Diretide Frenzy",
            "count": 3
        },
        {
            "id": 20116,
            "name": "Greater Fortitude",
            "count": 3
        },
        {
            "id": 20157,
            "name": "Egg Morph",
            "count": 3
        }
    ]
}
```
