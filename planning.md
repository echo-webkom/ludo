### Backlog

Flytte til todo

Struktur

```
tittel
laget av
knapp for å flytte til todo
```

### Todo

Har all info om todoen

Trykker på 'velg':

- flyttes til doing
- setter ditt navn på den
- lager lokal branch
- pusher til repo
- linker todo med branchen i repo
  - bruker github API for å få info om branchen

struktur:
tittel
laget av
description
knapp for å velge (flytte til doing)

doing
se status på en branch - antall commits - klar for merge

    når pr åpnes:
        - flyttes til review
        - notify person som skal review

review
når pr er merged flyttes automatisk til done
knapp for å manuellt flytte til done ('skip review')

done (archives et annet sted)
egen liste i en annen tab

flere ideer:
slack bot: - kan legge til i backlog med command - pinge folk som trenger review

github API ting vi trenger:
fetche info om en spesifik branch
om det er en pr åpen
om den er merged
hvis branchen er deleted så antar vi merged (fjerne todo)

stack:
Go
HTMX
GitHub API
API first

htmx app -> ludo API -> turso og github
slack -^

/ludo
/board - web app
/api - API
/server - api endpoint og server
/board - service for todo list (over db)
/database - database service
/github - github service
