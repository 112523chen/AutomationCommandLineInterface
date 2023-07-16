# Automation Command Line Interface

## Description

This is a command line interface for automating tasks. It is written in Golang and the data is stored in a Postgres database.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Usage](#usage)
- [Installation](#installation)
  - [Prerequisites](#prerequisites)
  - [Steps](#steps)

## Features

- [x] Search for solutions to code problems to speed up development
- [x] Search for open roles to speed up job search
- [x] Search for recruiters' linkedin profiles to speed up job search
- [x] Clean up local cache to free up space

## Technologies

- [Golang](https://golang.org/)
- [Postgres](https://www.postgresql.org/)

## Usage

Run the following command to start the CLI

- `~/Executable/cli`

Run the following command to start the CLI with the alias created

- `cli`

## Installation

### Prerequisites

- [Golang](https://golang.org/doc/install)
- [Postgres](https://www.postgresql.org/download/)

### Steps

1. Clone the Executable directory to your local machine

- `git clone https://github.com/112523chen/AutomationCommandLineInterface.git`

2. Create a database in Postgres

- `CREATE DATABASE websites;`
- `CREATE TABLE cli_websites (id SERIAL PRIMARY KEY, url VARCHAR(255)), type VARCHAR(10);`

3. Copy the Exectuable directory to your local machine root directory

- `cp -r Executable/ ~/`

4. Create a `.env` file in the Executable directory

- `touch ~/Executable/.env && open ~/Executable/.env`

5. Add the following to the `.env` file

- `DB_USER=your_username`
- `DB_PASSWORD=your_password`
- `DB_NAME=websites`

6. Run the following command to start the CLI

- `~/Executable/cli`

7. (Optional) Add the following to your `.bashrc` or `.zshrc` file

- `alias cli="~/Executable/cli"`
