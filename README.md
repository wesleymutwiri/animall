# Animall

Simple Anime helper to ease my anime tracking life

## Things I'd like to automate about anime life

- Get anime release date and get a calendar of anime I'm currently watching
- Update anime list to ensure I can track the current episode I'm currently at
- Schedule downloads for latest anime episodes so that I can remain up to date(not quite sure where to get anime downloads from yet)


## CLI features
- Add new anime and record current episode via CLI
- Update already added anime episode
- Mark an anime season as completed once watched all episodes
- Show current episode and episodes left to finish anime
- Show countdown to new release of episode
- Provide feature for downloading anime episode from the CLI in the background as I do other things(to be done first)

## Web portal Features
- Show a calendar displaying dates when new anime releases are going to premier
- Show list of data from backend of current anime episdoe and status whether finished or currently continuing
- Display countdown to next episode of most anime currently continuing
- Update anime status and be able to add new anime I'm currently following.
- Allow download of anime in the background while I'm doing other things.
- Secure application to be able to run whenever and wherever 
- Distribute through accounts so that other people can also use the site as expected

## Project structure 
- root 
    - cmd
        - root.go (entry point to the CLI application)
    - main.go (runs to ensure application works)
    - api(contains API to be created and used by the web application interface)
        - main.go(for running API instance)
        - dockerfile(easier deployment of application)
    - helpers
        - CRON job runner for getting latest episodes per week from site scrapped
        - Database connection for updating and adding new things