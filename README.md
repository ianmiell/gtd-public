# GTD

My 'Getting Things Done' workflow, in script form. Uses git as a database to track all my tasks.

## Files and directories

- README.md     - this file

- `bin/`        - gtd scripts

- `checklists/` - commonly-used checklists

- `gtd-web/`    - web interface for GTD

- `tasks/`      - all the tasks, one folder per task

- `status/`     - symlinks to statuses of task folders

- `meetings/`   - records of meetings

- `people/`     - records of people

- `priority/`   - symlinks to priority of task folders

## Installation and getting started

Fork the 'core' repository

- Fork the repository: <https://github.com/ianmiell/gtd-public>

- Clone your forked repository to your local machine:
  ```bash
  git clone https://github.com/<your-username>/gtd-public.git
  ```

- Create tasks:
  ```bash
  ./gtd create "My first task"
  ```

- View tasks:
  ```bash
  ./gtd d
  ```

- Close tasks:
  ```bash
  ./gtd close 4
  ```

- There are many more commands. To get help:
  ```bash
  ./gtd help
  ```

