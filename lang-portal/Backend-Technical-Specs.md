## Backend Server Technical Specs

## Business Goal: 
A language learning school wants to build a prototype of learning portal which will act as three things:
- Inventory of possible vocabulary that can be learned
- Act as a  Learning record store (LRS), providing correct and wrong score on practice vocabulary
- A unified launchpad to launch different learning apps

You have been tasked with creating the backend API of the application.

## Technical Restrictions:
- Use SQLite3 as the database
- You can use any language or framework 
- Does not require authentication/authorization
- Assume there is a single user
- The backend will be built using Go
- The API will be built using Gin
- The API will always return JSON
- Mage is the task runner for Go

## Database Schema

Our database will be a single sqlite database called 'words.db' that will be in the root of the project folder of 'backend_go'
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
    - activity_name string
    - group_name string
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
#### GET /api/dashboard/last_study_session
##### JSOn Response
    {
      "id": 1,
      "group_id": 1,
      "created_at": "2023-10-01T12:00:00Z",
      "study_activity_id": 1,
      "activity_name": "Vocabulary Practice",
      "group_name": "Basic Vocabulary"
    }
#### GET /api/dashboard/study_progress
    {
      "total_words_studied": 3,
      "total_available_words": 124
    }
#### GET /api/dashboard/quick-stats
    {
      "success_rate": 0.8,
      "total_study_sessions": 4,
      "study_streaks": 4,
      "total_active_groups": 3
    }
#### GET /api/dashboard/mastery_progress
    {
      "total_words": 124,
      "studied_words": 3,
      "mastery_progress": 0.02
    }
#### GET /api/words
        - Pagination with 100 items per page
    {
      "words": [
        {
          "id": 1,
          "japanese": "猫",
          "romaji": "neko",
          "english": "cat",
          "parts": ["noun"]
        },
        // ... up to 100 items
      ],
      "pagination": {
        "current_page": 1,
        "total_pages": 2,
        "total_items": 124,
        "items_per_page": 100
      }
    }
#### GET /api/words/:id
    {
      "id": 1,
      "japanese": "猫",
      "romaji": "neko",
      "english": "cat",
      "parts": ["noun"],
      "stats":{
        "correct_count": 10,
        "wrong_count": 2
      },
      "groups": [
        {
            "id": 1,
            "name": "Basic Greeting"
        }
      ]
    }
#### GET /api/groups
        - Pagination with 100 items per page
    {
      "groups": [
        {
          "id": 1,
          "name": "Basic Vocabulary",
          "word_count": 50
        },
        // ... up to 100 items
      ],
      "pagination": {
        "current_page": 1,
        "total_pages": 1,
        "total_items": 3,
        "items_per_page": 100
      }
    }
#### GET /api/groups/:id
    {
      "id": 1,
      "name": "Basic Vocabulary",
      "word_count": 50
    }
#### GET /api/groups/:id/words
    {
      "words": [
        {
          "id": 1,
          "japanese": "猫",
          "romaji": "neko",
          "english": "cat",
          "parts": ["noun"]
        },
        // ... up to 100 items
      ],
      "pagination": {
        "current_page": 1,
        "total_pages": 1,
        "total_items": 50,
        "items_per_page": 100
      }
    }
#### GET /api/groups/:id/study_sessions
    {
      "study_sessions": [
        {
          "id": 1,
          "group_id": 1,
          "created_at": "2023-10-01T12:00:00Z",
          "study_activity_id": 1,
          "activity_name": "Vocabulary Practice"
        },
        // ... more study sessions
      ],
      "pagination": {
        "current_page": 1,
        "total_pages": 1,
        "total_items": 10,
        "items_per_page": 100
      }
    }
#### GET /api/study_activities/:id
    {
      "id": 1,
      "study_session_id": 1,
      "group_id": 1,
      "created_at": "2023-10-01T12:00:00Z",
      "activity_name": "Vocabulary Practice",
      "group_name": "Basic Vocabulary"
    }
#### GET /api/study_activities/:id/study_sessions
    {
      "study_sessions": [
        {
          "id": 1,
          "group_id": 1,
          "created_at": "2023-10-01T12:00:00Z",
          "study_activity_id": 1,
          "activity_name": "Vocabulary Practice"
        },
        // ... more study sessions
      ],
      "pagination": {
        "current_page": 1,
        "total_pages": 1,
        "total_items": 10,
        "items_per_page": 100
      }
    }
#### GET /api/study_sessions
    {
      "study_sessions": [
        {
          "id": 1,
          "activity_name": "Vocabulary Practice",
          "group_name": "Basic Vocabulary",
          "start_time": "2023-10-01T12:00:00Z",
          "end_time": "2023-10-01T12:30:00Z",
          "number_of_review_items": 20
        },
        // ... more study sessions
      ],
      "pagination": {
        "current_page": 1,
        "total_pages": 1,
        "total_items": 10,
        "items_per_page": 100
      }
    }
#### GET /api/study_sessions/:id
    {
      "id": 1,
      "activity_name": "Vocabulary Practice",
      "group_name": "Basic Vocabulary",
      "start_time": "2023-10-01T12:00:00Z",
      "end_time": "2023-10-01T12:30:00Z",
      "number_of_review_items": 20,
      "words": [
        {
          "id": 1,
          "japanese": "猫",
          "romaji": "neko",
          "english": "cat",
          "correct": true,
          "created_at": "2023-10-01T12:05:00Z"
        },
        // ... more words
      ]
    }
#### GET /api/study_sessions/:id/words
    {
      "words": [
        {
          "id": 1,
          "japanese": "猫",
          "romaji": "neko",
          "english": "cat",
          "correct": true,
          "created_at": "2023-10-01T12:05:00Z"
        },
        // ... more words
      ],
      "pagination": {
        "current_page": 1,
        "total_pages": 1,
        "total_items": 20,
        "items_per_page": 100
      }
    }
#### POST /api/reset_history
    {
      "message": "History reset successfully"
    }
#### POST /api/full_reset
    {
      "message": "Full reset successfully"
    }
#### POST /api/study_sessions
    {
      "group_id": 1,
      "study_activity_id": 1
    }
#### POST /api/study_sessions/:id/review
##### Request Payload
    {
      "correct": true
    }
##### JSON Response
    {
      "message": "Review recorded successfully"
    }
#### POST /api/study_sessions/:id/words/:word_id/review
##### Request params:
    - id (study_session_id) integer
    - word_id integer
    - correct boolean
##### Request Payload
    {
      "correct": true
    }
##### JSON Response
    {
      "message": "Word review recorded successfully"
    }
#### POST /api/study_activities
##### Request params: 
    - group_id integer
    - study_activity_id integer
##### Request Payload:
    {
      "group_id": 1,
      "study_activity_id": 1
    }
##### JSON Response
    {
      "id": 1,
      "study_session_id": 1,
      "group_id": 1,
      "created_at": "2023-10-01T12:00:00Z"
    }

### Directory Structure
lang-portal-backend
└── lang-portal-backend
|   ├── cmd
|       └── main.go
|   ├── internal
|       ├── handlers
|           ├── dashboard.go
|           ├── groups.go
|           ├── study_activities.go
|           ├── study_sessions.go
|           ├── words.go
|           └── reset.go
|       ├── models
|           ├── group.go
|           ├── study_activity.go
|           ├── study_session.go
|           ├── word.go
|           └── word_review_item.go
|       ├── repositories
|           ├── group_repository.go
|           ├── study_activity_repository.go
|           ├── study_session_repository.go
|           ├── word_repository.go
|           └── word_review_item_repository.go
|       ├── services
|           ├── group_service.go
|           ├── study_activity_service.go
|           ├── study_session_service.go
|           ├── word_service.go
|           └── word_review_item_service.go
|       └── db
|           ├── migrations
|               └── 001_create_tables.sql
|           └── seeds
|               └── seed_data.json
|   ├── words.db
|   ├── go.mod
|   ├── go.sum
|   └── README.md

# Navigate to the backend_go directory
cd /Users/securingus/Downloads/BootCamp/free-genai-bootcamp-2025/lang-portal/backend_go

# Initialize a new Go module
go mod init lang-portal-backend

# Create the directory structure
mkdir -p cmd internal/handlers internal/models internal/repositories internal/services internal/db/migrations internal/db/seeds

# Create the initial files
touch cmd/main.go internal/handlers/dashboard.go internal/handlers/groups.go internal/handlers/study_activities.go internal/handlers/study_sessions.go internal/handlers/words.go internal/handlers/reset.go internal/models/group.go internal/models/study_activity.go internal/models/study_session.go internal/models/word.go internal/models/word_review_item.go internal/repositories/group_repository.go internal/repositories/study_activity_repository.go internal/repositories/study_session_repository.go internal/repositories/word_repository.go internal/repositories/word_review_item_repository.go internal/services/group_service.go internal/services/study_activity_service.go internal/services/study_session_service.go internal/services/word_service.go internal/services/word_review_item_service.go internal/db/migrations/001_create_tables.sql internal/db/seeds/seed_data.json words.db README.md

# Add necessary dependencies
go get github.com/gin-gonic/gin
go get github.com/mattn/go-sqlite3
go get github.com/magefile/mage