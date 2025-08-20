# Claude Code Best Practices Guide

A comprehensive guide for effectively using Claude Code to maximize productivity and code quality.

## 📋 Table of Contents

1. [Getting Started](#getting-started)
2. [Keyboard Shortcuts & Interface](#keyboard-shortcuts--interface)
3. [Project Setup Best Practices](#project-setup-best-practices)
4. [Effective Communication](#effective-communication)
5. [Code Development Workflow](#code-development-workflow)
6. [Advanced Workflows & Use Cases](#advanced-workflows--use-cases)
7. [Subagents & Specialized Tasks](#subagents--specialized-tasks)
8. [Background Processing](#background-processing)
9. [File Management](#file-management)
10. [Testing and Quality Assurance](#testing-and-quality-assurance)
11. [Git and Version Control](#git-and-version-control)
12. [Advanced Features](#advanced-features)
13. [Troubleshooting](#troubleshooting)
14. [Common Patterns](#common-patterns)

---

## 🚀 Getting Started

### Essential Setup

1. **Create CLAUDE.md**: Always include project-specific instructions
```markdown
# Project Name

## Context
Brief description of what this project does

## Development Setup
- How to install dependencies
- How to run the project
- Testing commands

## Important Notes
- Any special considerations
- Coding conventions
- Architecture decisions
```

2. **Initialize Git**: Ensure your project is version controlled
```bash
git init
git add .
git commit -m "Initial commit"
```

---

## ⌨️ Keyboard Shortcuts & Interface

### Essential Shortcuts in Interactive Mode

| Shortcut | Action | Description |
|----------|--------|-------------|
| `Ctrl+C` | Cancel | Stop current operation |
| `Ctrl+D` | Exit | Exit Claude Code |
| `↑` / `↓` | History | Navigate command history |
| `Tab` | Complete | Auto-complete commands |
| `Ctrl+L` | Clear | Clear screen |
| `Ctrl+R` | Search | Search command history |

### Chat Interface Commands

```bash
# Start new conversation
/new

# Resume previous conversation  
/resume

# Show help
/help

# Show conversation memory
/memory

# Clear current conversation
/clear

# Exit Claude Code
/exit or /quit
```

### File Navigation Shortcuts

```bash
# Quick file operations
"Open file explorer in src/"           # Navigate directories
"Show recent files"                     # View recently modified files
"Find all TypeScript files"            # Quick file search
"Show git status"                       # Check repository status
```

### Productivity Shortcuts

```bash
# Quick actions
"Run tests"                            # Execute test suite
"Build project"                        # Build application
"Start dev server"                     # Launch development server
"Lint code"                           # Check code style
"Format code"                         # Auto-format files
```

---

## 🏗️ Project Setup Best Practices

### Directory Structure
```
project/
├── CLAUDE.md          # Project instructions for Claude
├── README.md          # Project documentation
├── src/               # Source code
├── tests/             # Test files
├── docs/              # Documentation
└── .gitignore         # Git ignore rules
```

### CLAUDE.md Best Practices
- **Be Specific**: Include exact commands for building, testing, linting
- **Set Context**: Explain the project's purpose and architecture
- **Define Conventions**: Coding style, naming patterns, file organization
- **Include Examples**: Show preferred code patterns

Example CLAUDE.md:
```markdown
# My Project

Web application built with React + TypeScript + Node.js

## Commands
- Build: `npm run build`
- Test: `npm test`
- Lint: `npm run lint`
- Dev: `npm run dev`

## Conventions
- Use arrow functions for React components
- Prefer TypeScript strict mode
- Follow Airbnb ESLint rules
```

---

## 💬 Effective Communication

### Clear Requests
**❌ Vague:**
```
"Fix the bug in the app"
```

**✅ Specific:**
```
"The user login form throws a validation error when email contains '+' character. 
Fix the email validation regex in src/components/LoginForm.tsx"
```

### Providing Context
**Always include:**
- What you're trying to achieve
- Current behavior vs expected behavior
- Relevant file paths
- Error messages (if any)

### Using Examples
```
"Create a React component similar to UserCard.tsx but for displaying products. 
It should show product name, price, and image."
```

---

## ⚡ Code Development Workflow

### 1. Planning Phase
Use Claude's planning capabilities:
```
"I need to add user authentication. Can you help me plan the implementation?"
```

Claude will break down the task into manageable steps.

### 2. Implementation Phase
Work incrementally:
```
"Let's start with the authentication middleware first"
```

### 3. Review Phase
Ask for code review:
```
"Please review the authentication implementation for security best practices"
```

### Iterative Development Pattern
1. **Describe** the feature/change needed
2. **Review** Claude's plan
3. **Implement** step by step
4. **Test** each component
5. **Refactor** if needed

---

## 🔄 Advanced Workflows & Use Cases

### Complex Multi-Step Development

#### 1. Feature Development with Testing
```bash
# Step 1: Analysis & Planning
"Analyze the current user management system and plan adding role-based permissions"

# Step 2: Database Schema
"Design and implement the database migration for user roles"

# Step 3: Backend Implementation
"Create the role-based authentication middleware"

# Step 4: API Endpoints
"Add CRUD endpoints for role management with proper validation"

# Step 5: Frontend Integration
"Update the React components to support role-based UI rendering"

# Step 6: Testing Suite
"Create comprehensive tests covering all role scenarios"

# Step 7: Documentation
"Update API docs and create user guide for role management"
```

#### 2. Performance Optimization Workflow
```bash
# Step 1: Profiling
"Run performance analysis on the application and identify bottlenecks"

# Step 2: Database Optimization
"Optimize slow database queries and add proper indexing"

# Step 3: Frontend Optimization
"Implement code splitting and lazy loading for React components"

# Step 4: Caching Strategy
"Add Redis caching for frequently accessed data"

# Step 5: Bundle Optimization
"Analyze and optimize webpack bundle size"

# Step 6: Monitoring Setup
"Set up performance monitoring and alerting"
```

### Large-Scale Refactoring

#### Architecture Migration
```bash
# Example: Moving from REST to GraphQL
"Plan the migration from REST API to GraphQL step by step"
"Create GraphQL schema based on existing REST endpoints"
"Implement GraphQL resolvers with existing business logic"
"Create GraphQL client-side queries replacing REST calls"
"Add comprehensive testing for GraphQL endpoints"
"Update documentation and deployment scripts"
```

#### Legacy Code Modernization
```bash
# Example: JavaScript to TypeScript migration
"Analyze the JavaScript codebase and create TypeScript migration plan"
"Convert core utility functions to TypeScript first"
"Add type definitions for existing data models"
"Update React components with proper TypeScript interfaces"
"Configure build pipeline for TypeScript compilation"
"Update all import/export statements and dependencies"
```

### Security Implementation

#### Authentication System
```bash
"Design secure JWT-based authentication system"
"Implement password hashing with bcrypt and salt"
"Add rate limiting for login attempts"
"Create secure password reset functionality"
"Implement 2FA with TOTP"
"Add security headers and CORS configuration"
"Create security audit logging"
```

#### Data Protection
```bash
"Implement data encryption for sensitive user information"
"Add input validation and sanitization for all endpoints"
"Create secure file upload with virus scanning"
"Implement API key management system"
"Add database query protection against SQL injection"
"Set up vulnerability scanning and monitoring"
```

---

## 🤖 Subagents & Specialized Tasks

Claude Code uses specialized subagents for different types of tasks. Understanding when and how to use them maximizes efficiency.

### General-Purpose Agent
**Best for:** Complex searches, multi-step tasks, research

```bash
# When to use:
"Search the entire codebase for authentication patterns and analyze them"
"Find all database connection configurations across the project"
"Research how error handling is implemented throughout the application"

# The general-purpose agent excels at:
- Searching through large codebases
- Finding patterns across multiple files
- Analyzing complex relationships
- Research tasks requiring multiple tool uses
```

### Output Style Setup Agent
**Best for:** Configuring Claude Code's output formatting

```bash
# When to use:
"Configure Claude Code to show more detailed error messages"
"Set up custom output formatting for code reviews"
"Create a custom style for API documentation generation"

# This agent handles:
- Output formatting preferences
- Code display styles
- Error message formatting
- Custom report templates
```

### Status Line Setup Agent
**Best for:** Configuring the Claude Code status line

```bash
# When to use:
"Configure the status line to show current git branch and test status"
"Set up custom status indicators for build status"
"Add project-specific information to the status line"

# This agent manages:
- Status line configuration
- Custom status indicators
- Project information display
- Real-time status updates
```

### Requesting Specific Subagents

```bash
# Explicit subagent requests
"Use the general-purpose agent to find all React components that handle user authentication"
"Have the output-style agent create a custom format for displaying test results"
"Ask the statusline agent to add database connection status to the status bar"
```

### Subagent Best Practices

#### 1. Let Claude Choose Automatically
```bash
# Good: Let Claude decide which agent to use
"Search for all API endpoints and analyze their security implementations"

# Claude will automatically select the best subagent for the task
```

#### 2. Be Specific About Complex Tasks
```bash
# Good: Specific multi-step request
"Find all authentication-related code, analyze the security patterns, and create a comprehensive security audit report"

# This automatically triggers the general-purpose agent for complex analysis
```

#### 3. Use for Configuration Tasks
```bash
# Configuration requests automatically route to appropriate agents
"Set up Claude Code to display more verbose output for debugging"
"Configure the interface to show additional project information"
```

---

## ⚡ Background Processing

Claude Code can handle long-running tasks in the background, allowing you to continue working while tasks execute.

### Long-Running Commands

#### Build and Test Operations
```bash
# Start background build
"Run the full test suite in the background while I continue working"

# Monitor progress
"Check the status of the background test run"

# Background build processes
"Start the production build in the background"
"Run the linting process in the background"
"Execute database migrations in the background"
```

#### Data Processing Tasks
```bash
# Large file operations
"Process all log files in the /logs directory in the background"
"Run data migration script in the background"
"Generate API documentation in the background"

# Database operations
"Run database backup in the background"
"Execute data cleanup scripts in the background"
"Perform database indexing in the background"
```

### Monitoring Background Tasks

```bash
# Check running processes
"Show all background processes"
"What's the status of the test run?"
"Is the build still running?"

# Get output from background tasks
"Show me the output from the background test run"
"Get the latest logs from the build process"
```

### Background Task Patterns

#### 1. Development Workflow
```bash
# Start long-running tasks in background
"Start the development server in the background"
"Run file watcher for automatic compilation in the background"
"Start database seeding in the background"

# Continue with other work
"While that's running, let me review the user authentication code"
"Meanwhile, update the API documentation"
```

#### 2. Testing and Quality Assurance
```bash
# Comprehensive testing in background
"Run the entire test suite in the background"
"Execute performance tests in the background"
"Run security scanning in the background"

# Parallel development
"While tests are running, implement the new user profile feature"
```

#### 3. Deployment Preparation
```bash
# Build and package in background
"Build the production bundle in the background"
"Generate deployment assets in the background"
"Run pre-deployment checks in the background"

# Concurrent documentation
"While building, update the deployment documentation"
```

### Managing Background Processes

```bash
# Process control
"Kill the background build process"
"Restart the development server in the background"
"Stop all background processes"

# Output management
"Filter the background test output to show only failures"
"Save the background build logs to a file"
"Send me notifications when background tasks complete"
```

---

## 📁 File Management

### Reading Files Efficiently
**Multiple file analysis:**
```
"Analyze the authentication flow by reading LoginForm.tsx, authMiddleware.js, 
and UserService.ts"
```

### Creating vs Editing
- **Always prefer editing** existing files over creating new ones
- **Ask Claude to search** for existing similar functionality first

### File Organization
```
"Organize the utility functions in src/utils/ following the existing pattern"
```

---

## 🧪 Testing and Quality Assurance

### Test-Driven Development
```
"Create unit tests for the UserService class, then implement the class to pass the tests"
```

### Quality Checks
Always run after changes:
```bash
npm run lint      # Code style
npm run test      # Unit tests
npm run build     # Build verification
npm run typecheck # TypeScript checks
```

### Claude Integration
```
"Run the test suite and fix any failing tests"
```

Claude will:
1. Execute tests
2. Analyze failures
3. Fix the issues
4. Re-run tests to verify

---

## 🔧 Git and Version Control

### Commit Best Practices
```
"Create a commit for the authentication feature with a descriptive message"
```

Claude follows conventional commit format:
```
feat: add user authentication with JWT tokens

- Implement login/logout endpoints
- Add JWT middleware for protected routes
- Create user session management

🤖 Generated with Claude Code
Co-Authored-By: Claude <noreply@anthropic.com>
```

### Branch Management
```
"Create a new branch for the payment integration feature"
```

### Pull Requests
```
"Create a pull request for the authentication feature"
```

Claude will:
1. Analyze all changes
2. Write comprehensive PR description
3. Create test plan
4. Submit PR via GitHub CLI

---

## 🚀 Advanced Features

### Multi-file Operations
```
"Refactor the API endpoints to use async/await instead of Promises across all files"
```

### Code Generation Patterns
```
"Generate CRUD operations for the User model following the existing Product model pattern"
```

### Architecture Decisions
```
"Help me choose between Redux and Context API for state management in this React app"
```

### Performance Optimization
```
"Analyze the application for performance bottlenecks and suggest optimizations"
```

---

## 🔍 Troubleshooting

### Debug Mode
```
"The app crashes with this error: [paste error]. Help me debug step by step."
```

### Dependency Issues
```
"I'm getting a module not found error for 'axios'. Help me resolve this."
```

### Environment Problems
```
"The development server won't start. Here's the error output: [paste output]"
```

### Performance Issues
```
"The page loads slowly. Can you analyze the code and identify bottlenecks?"
```

---

## 🎯 Common Patterns

### Feature Implementation
```
1. "I need to add user profile editing functionality"
2. Claude creates todo list with steps
3. "Let's start with the API endpoint"
4. Implement backend
5. "Now create the React component"
6. Implement frontend
7. "Add tests for this feature"
8. Create comprehensive tests
9. "Run the full test suite"
10. Verify everything works
```

### Bug Fixing
```
1. "There's a bug where users can't upload files larger than 1MB"
2. Claude analyzes the codebase
3. Identifies the issue in upload middleware
4. Provides fix with explanation
5. Tests the fix
6. Commits the change
```

### Code Review
```
1. "Please review this pull request for security issues"
2. Claude analyzes all changed files
3. Provides detailed feedback
4. Suggests improvements
5. Helps implement fixes
```

### Refactoring
```
1. "The UserController.js file is getting too large"
2. Claude analyzes the file structure
3. Proposes refactoring plan
4. Breaks down into smaller modules
5. Maintains functionality
6. Updates imports and tests
```

---

## 📚 Pro Tips

### 1. Be Conversational
Claude works best with natural language:
```
"I'm trying to implement user authentication, but I'm not sure about the best approach for this Node.js app"
```

### 2. Use Context
Reference previous work:
```
"Using the same pattern as the Product API, create a User API"
```

### 3. Ask for Explanations
```
"Explain why you chose this approach over alternatives"
```

### 4. Request Documentation
```
"Add comprehensive JSDoc comments to these functions"
```

### 5. Validate Assumptions
```
"Before implementing, confirm that we should use MongoDB for the database"
```

### 6. Batch Operations
```
"Update all React components to use TypeScript and add proper prop types"
```

### 7. Security Focus
```
"Review this authentication code for security vulnerabilities"
```

---

## ⚠️ Common Mistakes to Avoid

### ❌ Don't:
- Give vague instructions
- Skip providing context
- Forget to test changes
- Ignore linting errors
- Skip code review
- Commit without running tests

### ✅ Do:
- Be specific and detailed
- Provide relevant file paths
- Test incrementally
- Fix all quality issues
- Ask for code review
- Verify before committing

---

## 🎨 Example Workflows

### New Feature Development
```bash
# 1. Planning
"I need to add user notifications. Help me plan this feature."

# 2. Backend Implementation
"Create the notification API endpoints following RESTful principles"

# 3. Frontend Implementation  
"Create React components for displaying and managing notifications"

# 4. Testing
"Add comprehensive tests for the notification system"

# 5. Integration
"Integrate notifications with the existing user dashboard"

# 6. Quality Check
"Run all tests and fix any issues"

# 7. Documentation
"Update the API documentation with the new notification endpoints"

# 8. Deployment Prep
"Create a migration script for the notification database schema"
```

### Bug Investigation
```bash
# 1. Report Issue
"Users report that the search function returns duplicate results"

# 2. Analysis
"Analyze the search implementation in src/services/SearchService.js"

# 3. Root Cause
"The issue is in the database query - it's not using DISTINCT"

# 4. Fix
"Update the query to remove duplicates and add proper indexing"

# 5. Test
"Create test cases to prevent this regression"

# 6. Verify
"Run the search functionality tests to confirm the fix"
```

---

## 🏆 Success Metrics

### Code Quality
- ✅ All tests passing
- ✅ Linting rules followed
- ✅ TypeScript compilation clean
- ✅ Security best practices applied

### Development Velocity
- ✅ Features delivered incrementally
- ✅ Bugs fixed promptly
- ✅ Code reviews completed
- ✅ Documentation updated

### Team Collaboration
- ✅ Clear commit messages
- ✅ Comprehensive PR descriptions
- ✅ Shared coding standards
- ✅ Knowledge documentation

---

## 🎯 Real-World Use Case Scenarios

### Scenario 1: Joining a New Project

#### Day 1: Understanding the Codebase
```bash
# Get project overview
"Analyze this codebase and give me a comprehensive overview of the architecture"

# Understand dependencies
"Show me all the main dependencies and explain what each one is used for"

# Find key files
"Identify the main entry points and most important files I should understand first"

# Development setup
"Help me set up the development environment following the project's conventions"
```

#### Week 1: First Contributions
```bash
# Find easy wins
"Find simple bugs or improvements I can work on to get familiar with the codebase"

# Understand testing
"Show me how testing is set up and help me write my first test"

# Code style learning
"Analyze the code style and conventions used in this project"
```

### Scenario 2: Emergency Bug Fix

#### Production Issue Response
```bash
# Rapid diagnosis
"There's a critical bug in production. The error is [paste error]. Help me diagnose this quickly"

# Hot fix development
"Create a minimal fix for this issue that can be deployed immediately"

# Testing the fix
"Create tests that verify this bug is fixed and won't regress"

# Emergency deployment
"Prepare this fix for emergency deployment with proper documentation"
```

### Scenario 3: Feature Development Sprint

#### Sprint Planning
```bash
# Requirements analysis
"Help me break down this user story into technical tasks: [paste requirements]"

# Architecture planning
"Design the architecture for this feature considering our existing codebase"

# Estimation
"Estimate the complexity and time required for each task"
```

#### Sprint Execution
```bash
# Daily development
"Let's implement the user authentication API endpoints first"
"Now create the frontend components for the login flow"
"Add comprehensive tests for all the authentication scenarios"

# Code review preparation
"Review my code changes and identify any issues before I create the PR"
```

### Scenario 4: Legacy System Maintenance

#### Technical Debt Management
```bash
# Debt identification
"Analyze this codebase and identify the main areas of technical debt"

# Prioritization
"Help me prioritize which technical debt to tackle first based on impact and effort"

# Incremental improvements
"Create a plan to gradually refactor the user management system"
```

#### Documentation and Knowledge Transfer
```bash
# Code documentation
"Generate comprehensive documentation for this undocumented legacy module"

# Knowledge capture
"Create a guide explaining how the payment processing system works"

# Migration planning
"Plan the migration from this legacy framework to the modern stack"
```

### Scenario 5: Performance Crisis

#### Performance Investigation
```bash
# Bottleneck identification
"The application is running slowly. Help me identify performance bottlenecks"

# Database optimization
"Analyze these slow database queries and optimize them"

# Frontend performance
"The React app is slow. Help me identify and fix performance issues"
```

#### Optimization Implementation
```bash
# Backend optimization
"Implement caching for these frequently accessed endpoints"

# Database tuning
"Add proper indexing and optimize these database queries"

# Frontend optimization
"Implement code splitting and lazy loading for better performance"
```

### Scenario 6: Security Audit

#### Security Assessment
```bash
# Vulnerability scan
"Perform a security audit of this authentication system"

# Code review for security
"Review this payment processing code for security vulnerabilities"

# Dependencies check
"Check all dependencies for known security vulnerabilities"
```

#### Security Hardening
```bash
# Input validation
"Add comprehensive input validation and sanitization to all endpoints"

# Authentication hardening
"Implement proper session management and CSRF protection"

# Data protection
"Add encryption for sensitive data storage and transmission"
```

### Scenario 7: API Development

#### REST API Creation
```bash
# API design
"Design a RESTful API for the user management system"

# Implementation
"Implement CRUD operations for users with proper validation"

# Documentation
"Generate comprehensive API documentation with examples"

# Testing
"Create integration tests for all API endpoints"
```

#### GraphQL Migration
```bash
# Schema design
"Convert this REST API to GraphQL, starting with the schema design"

# Resolver implementation
"Implement GraphQL resolvers for user and product queries"

# Client updates
"Update the frontend to use GraphQL queries instead of REST calls"
```

### Scenario 8: DevOps and Deployment

#### CI/CD Pipeline Setup
```bash
# Pipeline configuration
"Set up a GitHub Actions workflow for automated testing and deployment"

# Docker setup
"Create Docker configuration for this application"

# Infrastructure as code
"Create Terraform configuration for AWS deployment"
```

#### Monitoring and Logging
```bash
# Logging implementation
"Add comprehensive logging to this application"

# Monitoring setup
"Set up application performance monitoring and alerting"

# Health checks
"Implement health check endpoints for deployment verification"
```

---

## 🔧 Advanced Integration Patterns

### IDE Integration

#### VS Code Extension Usage
```bash
# Leverage IDE features
"Use the diagnostics to identify and fix TypeScript errors"
"Execute this Python code in the Jupyter kernel to test the algorithm"
"Get language server diagnostics for this Go file"
```

#### Multi-File Refactoring
```bash
# Cross-file operations
"Rename this interface across all TypeScript files in the project"
"Update all imports after moving this module to a new location"
"Refactor this shared utility function used in multiple components"
```

### Git Integration Workflows

#### Branch Management
```bash
# Feature branch workflow
"Create a feature branch for the user profile implementation"
"Merge the authentication feature branch with proper conflict resolution"
"Rebase this feature branch onto the latest main branch"
```

#### Advanced Git Operations
```bash
# Complex git operations
"Cherry-pick the bug fix commits from the hotfix branch"
"Create a patch file for these specific changes"
"Split this large commit into smaller, focused commits"
```

### Continuous Integration Patterns

#### Test Automation
```bash
# Automated testing workflows
"Set up pre-commit hooks for code quality checks"
"Create a test matrix for multiple Node.js versions"
"Implement visual regression testing for the UI components"
```

#### Deployment Automation
```bash
# Deployment workflows
"Create a blue-green deployment strategy for zero-downtime updates"
"Set up automatic rollback on deployment failure"
"Implement feature flag deployment for gradual rollouts"
```

---

## 📖 Additional Resources

- [Claude Code Documentation](https://docs.anthropic.com/claude-code)
- [GitHub CLI Reference](https://cli.github.com/manual/)
- [Conventional Commits](https://www.conventionalcommits.org/)
- [Git Best Practices](https://git-scm.com/book)

### Quick Reference Cards

#### Essential Commands
```bash
# Project initialization
claude --init                  # Initialize Claude Code in project
claude --resume                # Resume previous conversation

# File operations
"Read file.js"                 # Quick file read
"Edit function in file.js"     # Targeted editing
"Create new component"         # File creation

# Testing and quality
"Run tests"                    # Execute test suite
"Lint code"                    # Code quality check
"Build project"               # Build verification

# Git operations
"Git status"                   # Repository status
"Create commit"                # Commit changes
"Create PR"                    # Pull request creation
```

#### Common Patterns
```bash
# Development cycle
"Analyze → Plan → Implement → Test → Review → Deploy"

# Bug fixing cycle
"Reproduce → Diagnose → Fix → Test → Document"

# Feature development cycle
"Research → Design → Implement → Test → Document → Deploy"
```

---

*This guide is living documentation. Update it as you discover new patterns and best practices while working with Claude Code.*