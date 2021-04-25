# User Survey

This simple project allow for the anonymous creation/taking surveys. 
Consist of 2 services and mongo db.
- survey hexagonal service
- cobra interactive cli

Survey service runs on port `:8000` as RESTful service and Cobra client interact with it through API.

### Run the services

`docker-compose build` will build the Dockerfiles of both the services

`docker-compose run client sh` will start the `Cobra cli` in interactive terminal.

### Cli commands

```text
+++++++++++++++++++++++++++++++++++++


Usage:
  cli [command]

Available Commands:
  create      Create a new survey
  help        Help about any command
  show        Show results of a survey
  take        Take a survey

Flags:
  -h, --help   help for cli

Use "cli [command] --help" for more information about a command.
```

### Create a new Survey
`cli create`

```text
+++++++++++++++++++++++++++++++++++++
Title: Music
Question 0:  Do you like Trance?
Question 1:  Do you like Classical?
Question 2:  Do you like American Pop?
Survey result:
 Location: /api/surveys/607d2491ff515ded51cd49a5
 Survey: 607d2491ff515ded51cd49a5
```

### Take a survey
`cli take`

```text
+++++++++++++++++++++++++++++++++++++
🌶 Music
Please fill the Survey.
✔ Yes
✔ Yes
✔ No
Survey result:
 Location: /api/surveys/607d2491ff515ded51cd49a5/results/607d24eaff515ded51cd49a6
 Survey: 607d2491ff515ded51cd49a5
 Result: 607d24eaff515ded51cd49a6
```

### Show results of a survey
`cli show -s 607d2491ff515ded51cd49a5`

```text
+++++++++++++++++++++++++++++++++++++
Use the arrow keys to navigate: ↓ ↑ → ←  and / toggles search
Survey Titles?
  🌶 Music (607d24eaff515ded51cd49a6)
    Music (607d254aff515ded51cd49a7)

--------- Survey ----------
Title:        Music

        0        Do you like Trance?
        true

        1        Do you like Classical?
        true

        2        Do you like American Pop?
        false

```
