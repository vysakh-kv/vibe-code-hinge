# Cursor Development Guidelines for Vibe Project

## Overview

This document provides guidelines for developers using Cursor AI when working on the Vibe dating app project. Following these guidelines will help maintain consistency and quality across the codebase.

## Core Principles

### 1. Keep It Simple

- Favor simple solutions over complex ones
- Don't overengineer features or introduce unnecessary abstractions
- Keep code readable and maintainable for the entire team

### 2. Reuse Existing Components

- Before creating new components or utilities, check if existing ones can be reused or extended
- Maintain consistency by leveraging established patterns in the codebase
- Don't duplicate functionality that already exists

### 3. Preserve Existing Functionality

- **IMPORTANT**: Do not remove or significantly alter existing functionality without explicit confirmation from users
- When modifying existing features, ensure backward compatibility when possible
- Test thoroughly to ensure changes don't break existing features

### 4. Maintain Project Structure

- Follow the established project architecture and folder structure
- Keep frontend components in appropriate directories based on their purpose
- Organize backend code according to clean architecture principles

### 5. Follow Coding Standards

- Adhere to Vue's composition API patterns for frontend
- Follow Go's idiomatic practices for backend
- Maintain consistent naming conventions across the codebase

### 6. Keep Context Documentation Updated

- **CRITICAL**: Update `llm_context.txt` when making significant changes to the project
- This file provides essential context for AI tools and new developers
- Document new features, endpoints, tech stack changes, and other important developments
- Accurate context documentation leads to better assistance and more efficient development

## Practical Guidelines for Cursor Use

### When Generating Code

1. **Start with Context**: Always understand the surrounding code before making changes
2. **Incremental Changes**: Generate smaller, focused changes rather than large overhauls
3. **Test Generated Code**: Verify that generated code works as expected before committing

### When Refactoring

1. **Confirm First**: Always confirm with users before significant refactoring
2. **Document Why**: Add comments explaining the purpose of complex refactoring
3. **Preserve Behavior**: Ensure refactored code maintains the same behavior as the original

### When Adding Features

1. **Check Requirements**: Ensure you understand the feature requirements completely
2. **Leverage Existing Code**: Use or extend existing components and utilities where possible
3. **Consider Edge Cases**: Account for error states, loading states, and edge cases
4. **Update Context**: Add new features to the `llm_context.txt` file

### When Fixing Bugs

1. **Understand Root Cause**: Take time to understand the root cause before fixing
2. **Minimal Changes**: Make the smallest change necessary to fix the issue
3. **Add Tests**: When possible, add tests to prevent regression

### When To Update llm_context.txt

Update the `llm_context.txt` file when:

1. **Adding new features or components**
2. **Changing or extending APIs**
3. **Introducing new technology or libraries**
4. **Changing project architecture or structure**
5. **Modifying development workflow or processes**
6. **Updating database schema**
7. **Adding or changing configuration requirements**

This keeps AI agents and new developers informed about the current state of the project.

## Project-Specific Guidelines

- Frontend components should follow the existing naming and styling conventions
- Backend endpoints should follow RESTful design principles
- Mobile-first approach for all UI development
- Document any new environment variables or configuration options
- Keep sensitive information out of the codebase
- Utilize the Render PostgreSQL database with the provided credentials 

### 7. External Asset Attribution

- When using external assets (libraries, images, code snippets, etc.), add appropriate attribution in code comments
- Include basic information like asset name, creator, license, and source when available
- These credits should NOT be displayed in the UI or included in API responses
- Attribution comments are for development purposes only
- This is a guideline for Cursor to keep in mind while developing, not a strict implementation requirement 