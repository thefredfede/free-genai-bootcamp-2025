# #Frontend Technical Specs

## Pages

### Dashboard '/dashboard'
#### Purpose
#### Components
This page contains the following components
- Last study session
    shows last activity used
    shows when last activity used
    summarizes wrong vs correct from last activity
    has a link to the group
- Study progress
    - total words study e.g. 3/124
        - across all study session show the total words studies out of all possible words in our database
    - display a mastery progress e.g 0%
- Quick Stats
    - success rate e.g. 80%
    - total study sessions e.g. 4
    - study streaks e.g. 4 days
    - total active groups e.g. 3
- Start Studying Button
    - goes to study activities page
#### Needed API endpoints
We'll need the following API endpoints to power this page
- GET /api/dashboard/last_study_session
- GET /api/dashboard/study_progress
- GET /api/dashboard/quick-stats

### Study Activities '/study_activities'
#### Purpose
The purpose of this page is to show a collection of study activities with a thmumbnail and its name, to either launch or view the study activity.
#### Components
- Study Activity Card
    - shows a thumbnail of the study activity
    - the name of the study activity
    - a launch button to take us t the lunch page
    - the viewpage to view more information about past study sessions for this study activity
#### Needed API endpoints
- GET /api/study_activities


### Study Activity Show '/study_activities/:id'
#### Purpose
Show details of a study activity and its past study sessions
#### Components
- Name of study activity
- Thumbnail of study activity
- Description of study activity
- Launch button
- Study activities paginated list
    - id
    - activity name
    - group name
    - start time
    - end time (infered by the last word review item submitted)
    - number of review items
#### Needed API endpoints
- GET /api/study_activities/:id
- GET /api/study_activities/:id/study_sessions

### Study Activity Launch '/study_activities/:id/launch'
#### Purpose
The purpose of this page is to launch a study activity
#### Components
- Name of study activity
- Launch form
    - select field for group
    - launch now button
#### Needed API endpoints
- POST /api/study_activities
    - Required params: group_id, study_activity_id
## Bahaviour
After the form is ubmitted, a new tab opens with the study activity based on its URL provided in the database.
After submitting form the page will redirect to the study session show page.

### Words Index '/words'
#### Purpose
The purpose is to show all words in our DB.
#### Components
- Paginated word list
    - Columns
        - Japanese
        - Romaji
        - English
        - Correct Count
        - Wrong Count
    - Pagination with 100 items per page
    - Clicking the Japanese work will take us to the word show page
#### Needed API endpoints
- GET /api/words

### Word Show '/words/:id'
#### Purpose
The purpose is to show information about a specific word
#### Components
- Japanese
- Romaji
- English
- Study Statistics
    - Correct Count
    - Wrong Count
- Word Groups 
    - show a series of pills e.g. tags
    - when grup name is clicked it will take us to the group show page
#### Needed API endpoints
- GET /api/words/:id

### Word Groups Index '/groups'
#### Purpose
To show a list of groups in database
#### Components
- Paginated group lisy
    - Columns
        - Group Name
        - Word Count
    - Clicking the group name wil take us to the group show page
#### Needed API endpoints
- GET /api/groups


### Group Show '/groups/:id'
#### Purpose
To show information about a specific group
#### Components
- Group Name
- Group Statistics
    - Total Word Count
- Words in Group (Paginated list of words)
    - Should use the same components as the words index page
- Study Sessions
    - Should use the same components as the study sessionss index page
#### Needed API endpoints
- GET /api/groups/:id (the name and groups stats)
- GET /api/groups/:id/words
- GET /api/groups/:id/study_sessions 

### Study Sessions Index '/study_sessions'
#### Purpose
The purpose is to show a list of study sessions in our database
#### Components
- Paginated Study Session List
    - Columns
        - Id
        - Activity Name
        - Group Name
        - Start Time
        - End Time
        - Number of Review Items
    - Clicking the study session id will take us to the study session show page
#### Needed API endpoints
- GET /api/study_sessions

### Study Sessions Show '/study_sessions/:id'
#### Purpose
The purpose of the page is to show information about a specific study session.
#### Components
- Study Session Details
    - Activity Name
    - Group Name
    - Start Time
    - End Time
    - Number of Review Items
- Words Review Items (Paginated list of words)
    - Should use the same components as the words index page
#### Needed API endpoints
- GET /api/study_sessions/:id
- GET /api/study_sessions/:id/words

### Settings Page '/settings'
#### Purpose
To make configurations to the study portal
#### Components
- Theme selection eg. Light, Dark, System Default
- Reset History Button
    - This will delete all study sessions and word review items
- Full Reset Button
    - This will drop all tables and re-create with seed data
#### Needed API endpoints
- POST /api/reset_history
- POST /api/full_reset

## Mage Tasks
Mage is a task runner for Go
### Initialize Database

### Migrate database
This task will run a series of migration sql files on the database.
Migrations will live in the 'migratons' folder.
The migration files will be run in order oftheir file name.
### Seed Data
This task will import json files and transform them into target data for our database.

All see files live in the "seeds: folder.
All seed files should be loaded.

In our task we should have DSL to specify each seed file and its expected group word name.
```json
[
    {
        "kanji": "いい",
        "romaji": "ii",
        "english": "good",
        "parts": [
            { "kanji": "い", "romaji": ["i"] },
            { "kanji": "い", "romaji": ["i"] }
    ]
  },
  ...
]
```
