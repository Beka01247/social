# Social - Social Networking Platform

A comprehensive social networking platform built with Go that enables users to create posts, follow other users, and interact through comments. The platform features a REST API with JWT authentication, real-time feeds, and robust user management capabilities.

## 🚀 Features

### Core Functionality

- **User Management**: User registration, authentication, and profile management
- **Posts & Content**: Create, read, update, and delete posts with tagging support
- **Social Features**: Follow/unfollow users, personalized feeds
- **Comments System**: Comment on posts with full CRUD operations
- **User Invitations**: Invite new users with email notifications
- **Role-based Access Control**: Multi-level user permissions and roles

### Technical Features

- **JWT Authentication**: Secure token-based authentication
- **Rate Limiting**: Fixed-window rate limiting to prevent abuse
- **Caching**: Redis integration for improved performance
- **Email Notifications**: SendGrid and Mailtrap integration for user invitations
- **Database Migrations**: Structured database schema versioning
- **Graceful Shutdown**: Proper server shutdown handling
- **Comprehensive Logging**: Structured logging with Zap
- **API Documentation**: Auto-generated Swagger documentation
- **Server Metrics**: Application performance monitoring with expvar
- **Unit Testing**: Partial test coverage for critical components

## 🏗️ Project Structure

```
├── cmd/
│   ├── api/                    # Main API application
│   │   ├── main.go            # Application entry point
│   │   ├── auth.go            # Authentication handlers
│   │   ├── posts.go           # Post management handlers
│   │   ├── users.go           # User management handlers
│   │   ├── feed.go            # User feed handlers
│   │   ├── middleware.go      # HTTP middleware
│   │   └── *_test.go          # Unit tests
│   └── migrate/               # Database migration tools
│       ├── migrations/        # SQL migration files
│       └── seed/              # Database seeding
├── internal/
│   ├── auth/                  # Authentication logic & JWT handling
│   ├── db/                    # Database connection and utilities
│   ├── env/                   # Environment configuration
│   ├── mailer/                # Email service integrations
│   ├── ratelimiter/           # Rate limiting implementation
│   └── store/                 # Data access layer
│       ├── cache/             # Redis caching layer
│       ├── users.go           # User repository
│       ├── posts.go           # Posts repository
│       ├── comments.go        # Comments repository
│       ├── followers.go       # Follow/unfollow logic
│       └── roles.go           # Role management
├── docs/                      # Swagger documentation
├── scripts/                   # Database initialization scripts
└── docker-compose.yml         # Development environment setup
```

## 🛠️ Tech Stack

- **Backend**: Go 1.23
- **Database**: PostgreSQL 16.3
- **Cache**: Redis 6.2
- **Authentication**: JWT tokens
- **Email**: SendGrid / Mailtrap
- **Documentation**: Swagger/OpenAPI
- **Testing**: Go testing framework
- **Containerization**: Docker & Docker Compose

## 📚 API Documentation

The project includes comprehensive API documentation generated with Swagger. Once the server is running, you can access the interactive API docs at:

```
http://localhost:8080/swagger/index.html
```

The API supports the following main endpoints:

- `/v1/users` - User management and profiles
- `/v1/posts` - Post creation and management
- `/v1/users/feed` - Personalized user feeds
- `/v1/health` - Health check endpoint
- `/debug/vars` - Server metrics and statistics

## 🚦 Getting Started

### Prerequisites

- Go 1.23 or later
- Docker and Docker Compose
- Make (optional, for convenience commands)

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/Beka01247/social.git
   cd social
   ```

2. **Start the infrastructure**

   ```bash
   docker-compose up -d
   ```

   This will start:

   - PostgreSQL database on port 5432
   - Redis cache on port 6379
   - Redis Commander (web UI) on port 8081

3. **Set up environment variables**

   ```bash
   cp .envrc.example .envrc
   # Edit .envrc with your configuration
   source .envrc
   ```

4. **Run database migrations**

   ```bash
   make migrate-up
   ```

5. **Seed the database (optional)**

   ```bash
   make seed
   ```

6. **Generate API documentation**

   ```bash
   make gen-docs
   ```

7. **Start the application**
   ```bash
   go run cmd/api/main.go
   ```

The API will be available at `http://localhost:8080`

### Running Tests

```bash
make test
```

## 🐳 Docker Services

The project includes a Docker Compose setup with the following services:

- **PostgreSQL**: Main database with persistent volume
- **Redis**: Caching layer and session storage
- **Redis Commander**: Web-based Redis management interface

## 🔧 Configuration

The application supports configuration through environment variables:

### Server Configuration

- `ADDR`: Server address (default: `:8080`)
- `EXTERNAL_URL`: External API URL for callbacks (default: `localhost:8080`)
- `FRONTEND_URL`: Frontend application URL (default: `localhost:3000`)
- `ENV`: Environment mode (`development`, `production`)

### Database Configuration

- `DB_ADDR`: PostgreSQL connection string
- `DB_MAX_OPEN_CONS`: Maximum open database connections (default: 30)
- `DB_MAX_IDLE_CONS`: Maximum idle database connections (default: 30)
- `DB_MAX_OPEN_TIME`: Maximum connection idle time (default: "15m")

### Authentication

- `JWT_SECRET`: Secret key for JWT token signing

### Email Configuration

- `SENDGRID_API_KEY`: SendGrid API key for email sending
- `FROM_EMAIL`: From email address for notifications
- `MAILTRAP_API_KEY`: Mailtrap API key (for development/testing)

### Redis Configuration

- `REDIS_ENABLED`: Enable Redis caching (default: true)
- `REDIS_ADDR`: Redis connection address (default: `localhost:6379`)
- `REDIS_PW`: Redis password (optional)
- `REDIS_DB`: Redis database number (default: 0)

### Rate Limiting

- `RATELIMITER_ENABLED`: Enable rate limiting (default: true)
- `RATELIMITER_REQUESTS_PER_TIME_FRAME`: Requests per time frame (default: 20)
- `RATELIMITER_TIME_FRAME`: Time frame duration (default: "1m")

## 📊 Monitoring & Metrics

The application exposes runtime metrics and statistics at `/debug/vars`, including:

- Memory usage and garbage collection stats
- Goroutine counts
- Request counters and timing
- Database connection pool statistics

## 🔒 Security Features

- **JWT Authentication**: Secure token-based user authentication
- **Rate Limiting**: Configurable request rate limiting per IP
- **Input Validation**: Comprehensive request payload validation
- **SQL Injection Protection**: Parameterized queries and prepared statements
- **CORS Support**: Configurable cross-origin resource sharing

## 🧪 Testing

The project includes unit tests for critical components:

- API handlers testing
- User management functionality
- Authentication flows
- Database operations

Run tests with: `go test -v ./...`

## 🔄 CI/CD Workflows

The project includes automated GitHub Actions workflows for continuous integration and deployment:

### Audit Workflow (`audit.yaml`)

- **Triggers**: Push to master, Pull Requests
- **Features**:
  - Dependency verification with `go mod verify`
  - Code compilation and build verification
  - Static analysis with `go vet` and `staticcheck`
  - Race condition testing with `go test -race`
  - Runs on Ubuntu 22.04 with Go 1.23.0

### Release Please (`release-please.yaml`)

- **Triggers**: Push to master branch
- **Features**:
  - Automated release management
  - Changelog generation
  - Semantic versioning
  - Pull request creation for releases

### Version Update (`update-api-version.yaml`)

- **Triggers**: Push to master branch
- **Features**:
  - Automatic version extraction from CHANGELOG.md
  - Version synchronization across project files
  - Automated commit and release creation
  - Prevents recursive workflow execution
