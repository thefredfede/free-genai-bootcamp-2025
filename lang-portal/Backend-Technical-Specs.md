## Backend Server Technical Specs

## Business Goal: 
A language learning school wants to build a prototype of learning portal which will act as three things:
- Inventory of possible vocabulary that can be learned
- Act as a  Learning record store (LRS), providing correct and wrong score on practice vocabulary
- A unified launchpad to launch different learning apps

You have been tasked with creating the backend API of the application.

## Technical Restrictions:
- Use SQLite3 as the database
You can use any language or framework 
- Does not require authentication/authorization
- Assume there is a single user
- The backend will be built using Go
- The API will be built using Gin
- The API will always return JSON

## Database Schema
We have the following tables:
- words - stored vocabulary words
    - id integer
    - japanese string
    - romaji string
    - english string
    - parts json
- word_groups - join table for words and groups many-to-many
    - id integer
    - word_id integer
    - group_id integer
- groups - thematic group of words
    - id integer
    - name string
- study_sessions - records of study sessions grouping word_review_tems
    - id integer
    - group_id integer
    - created_at datetime
    - study_activity_id integer
- study_activities - a specific study activity, linking a study session to group 
    - id integer
    - study_session_id integer
    - group_id integer
    - created_at datetime
- word_review_items - a record of word practise, determining if the word was correct or not
    - word_id integer
    - study_session_id integer
    - correct boolean
    - created_at datetime

### API Endpoints
    - GET /api/dashboard/last_study_session
    - GET /api/dashboard/study_progress
    - GET /api/dashboard/quick-stats
    - GET /api/words
        - Pagination with 100 items per page
    - GET /api/words/:id
    - GET /api/groups
        - Pagination with 100 items per page
    - GET /api/groups/:id
    - GET /api/groups/:id/words
    - GET /api/groups/:id/study_sessions 
    - GET /api/study_activities/:id
    - GET /api/study_activities/:id/study_sessions
    - GET /api/study_sessions
    - GET /api/study_sessions/:id/words
    - POST /api/reset_history
    - POST /api/full_reset
    - POST /api/study_sessions
    - POST /api/study_sessions/:id/review
    - POST /api/study_sessions/:id/words/:word_id/review
        -   required param: correct
    - POST /api/study_activities
        - Required params: group_id, study_activity_id